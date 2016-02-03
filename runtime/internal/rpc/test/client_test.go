// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package test

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"testing"
	"time"

	"v.io/v23"
	"v.io/v23/context"
	"v.io/v23/flow"
	"v.io/v23/flow/message"
	"v.io/v23/naming"
	"v.io/v23/options"
	"v.io/v23/rpc"
	"v.io/v23/security"
	"v.io/v23/vdlroot/signature"
	"v.io/v23/verror"
	"v.io/x/lib/gosh"
	"v.io/x/ref"
	"v.io/x/ref/internal/logger"
	"v.io/x/ref/lib/signals"
	_ "v.io/x/ref/runtime/factories/generic"
	"v.io/x/ref/runtime/internal/lib/tcputil"
	inaming "v.io/x/ref/runtime/internal/naming"
	irpc "v.io/x/ref/runtime/internal/rpc"
	"v.io/x/ref/runtime/protocols/debug"
	"v.io/x/ref/services/mounttable/mounttablelib"
	"v.io/x/ref/test"
	"v.io/x/ref/test/testutil"
	"v.io/x/ref/test/v23test"
)

func TestMain(m *testing.M) {
	v23test.TestMain(m)
}

// testInit creates a new v23test.Shell, starts a root mount table, and
// optionally starts a simple server.
func testInit(t *testing.T, startServer bool) (sh *v23test.Shell, ctx *context.T, name string, cleanup func()) {
	sh = v23test.NewShell(t, v23test.Opts{})
	ctx = sh.Ctx
	startRootMT(t, sh, false)
	var cleanupServer func()
	name, cleanupServer = startSimpleServer(t, ctx)
	cleanup = func() {
		cleanupServer()
		sh.Cleanup()
	}
	return
}

////////////////////////////////////////
// Root mount table

// TODO(sadovsky): Switch to using v23test.Shell.StartRootMountTable.
var rootMT = gosh.RegisterFunc("rootMT", func(deb bool) error {
	ctx, shutdown := v23.Init()
	defer shutdown()
	if deb {
		p, a := debug.WrapAddress("tcp", "127.0.0.1:0")
		ctx = v23.WithListenSpec(ctx, rpc.ListenSpec{
			Addrs: rpc.ListenAddrs{{Protocol: p, Address: a}},
		})
	}
	mt, err := mounttablelib.NewMountTableDispatcher(ctx, "", "", "mounttable")
	if err != nil {
		return fmt.Errorf("mounttablelib.NewMountTableDispatcher failed: %s", err)
	}
	ctx, server, err := v23.WithNewDispatchingServer(ctx, "", mt, options.ServesMountTable(true))
	if err != nil {
		return fmt.Errorf("root failed: %v", err)
	}
	fmt.Printf("PID=%d\n", os.Getpid())
	for _, ep := range server.Status().Endpoints {
		fmt.Printf("MT_NAME=%s\n", ep.Name())
	}
	<-signals.ShutdownOnSignals(ctx)
	return nil
})

func startRootMT(t *testing.T, sh *v23test.Shell, deb bool) {
	cmd := sh.FuncCmd(rootMT, deb)
	cmd.Start()
	cmd.S.ExpectVar("PID")
	rootName := cmd.S.ExpectVar("MT_NAME")
	if len(rootName) == 0 {
		sh.HandleError(errors.New("no MT_NAME"))
		return
	}
	sh.Vars[ref.EnvNamespacePrefix] = rootName
	if err := v23.GetNamespace(sh.Ctx).SetRoots(rootName); err != nil {
		sh.HandleError(err)
		return
	}
}

////////////////////////////////////////
// Echo server

type treeDispatcher struct{ id string }

func (d treeDispatcher) Lookup(_ *context.T, suffix string) (interface{}, security.Authorizer, error) {
	return &echoServerObject{d.id, suffix}, nil, nil
}

type echoServerObject struct {
	id, suffix string
}

func (es *echoServerObject) Echo(_ *context.T, _ rpc.ServerCall, m string) (string, error) {
	if len(es.suffix) > 0 {
		return fmt.Sprintf("%s.%s: %s\n", es.id, es.suffix, m), nil
	}
	return fmt.Sprintf("%s: %s\n", es.id, m), nil
}

func (es *echoServerObject) Sleep(_ *context.T, _ rpc.ServerCall, d string) error {
	duration, err := time.ParseDuration(d)
	if err != nil {
		return err
	}
	time.Sleep(duration)
	return nil
}

var echoServer = gosh.RegisterFunc("echoServer", func(id, mp, addr string) error {
	ctx, shutdown := v23.Init()
	defer shutdown()

	if addr == "" {
		addr = "127.0.0.1:0"
	}
	ctx = v23.WithListenSpec(ctx, rpc.ListenSpec{Addrs: rpc.ListenAddrs{{Protocol: "tcp", Address: addr}}})

	disp := &treeDispatcher{id: id}
	ctx, server, err := v23.WithNewDispatchingServer(ctx, mp, disp)
	if err != nil {
		return err
	}
	fmt.Printf("PID=%d\n", os.Getpid())
	for _, ep := range server.Status().Endpoints {
		fmt.Printf("NAME=%s\n", ep.Name())
	}
	<-signals.ShutdownOnSignals(ctx)
	return nil
})

// Returns server Cmd and name.
func startEchoServer(t *testing.T, sh *v23test.Shell, id, mp, addr string) (*v23test.Cmd, string) {
	cmd := sh.FuncCmd(echoServer, id, mp, addr)
	cmd.Start()
	cmd.S.ExpectVar("PID")
	name := cmd.S.ExpectVar("NAME")
	return cmd, name
}

////////////////////////////////////////
// Echo client

var echoClient = gosh.RegisterFunc("echoClient", func(name string, args ...string) error {
	ctx, shutdown := v23.Init()
	defer shutdown()

	client := v23.GetClient(ctx)
	for _, a := range args {
		var r string
		if err := client.Call(ctx, name, "Echo", []interface{}{a}, []interface{}{&r}); err != nil {
			return err
		}
		fmt.Printf(r)
	}
	return nil
})

func runEchoClient(t *testing.T, sh *v23test.Shell) {
	cmd := sh.FuncCmd(echoClient, "echoServer", "a message")
	cmd.Start()
	cmd.S.Expect("echoServer: a message")
	cmd.Wait()
}

////////////////////////////////////////
// Misc helpers

func numServers(t *testing.T, ctx *context.T, name string, expected int) int {
	for {
		me, err := v23.GetNamespace(ctx).Resolve(ctx, name)
		if err == nil && len(me.Servers) == expected {
			return expected
		}
		time.Sleep(10 * time.Millisecond)
	}
}

////////////////////////////////////////
// Tests

// TODO(cnicolaou): figure out how to test and see what the internals
// of tryCall are doing - e.g. using stats counters.
func TestMultipleEndpoints(t *testing.T) {
	sh, ctx, _, cleanup := testInit(t, false)
	defer cleanup()

	cmd, _ := startEchoServer(t, sh, "echoServer", "echoServer", "")

	// Verify that there is 1 entry for echoServer in the mount table.
	if got, want := numServers(t, ctx, "echoServer", 1), 1; got != want {
		t.Fatalf("got: %d, want: %d", got, want)
	}

	runEchoClient(t, sh)

	// Create a fake set of 100 entries in the mount table.
	for i := 0; i < 100; i++ {
		// 203.0.113.0 is TEST-NET-3 from RFC5737
		ep := naming.FormatEndpoint("tcp", fmt.Sprintf("203.0.113.%d:443", i))
		n := naming.JoinAddressName(ep, "")
		if err := v23.GetNamespace(ctx).Mount(ctx, "echoServer", n, time.Hour); err != nil {
			t.Fatalf("unexpected error: %s", err)
		}
	}

	// Verify that there are 101 entries for echoServer in the mount table.
	if got, want := numServers(t, ctx, "echoServer", 101), 101; got != want {
		t.Fatalf("got: %q, want: %q", got, want)
	}

	// TODO(cnicolaou): ok, so it works, but I'm not sure how
	// long it should take or if the parallel connection code
	// really works. Use counters to inspect it for example.
	runEchoClient(t, sh)

	cmd.Terminate(os.Interrupt)

	// Verify that there are 100 entries for echoServer in the mount table.
	if got, want := numServers(t, ctx, "echoServer", 100), 100; got != want {
		t.Fatalf("got: %d, want: %d", got, want)
	}
}

func TestTimeout(t *testing.T) {
	t.Skip("https://github.com/vanadium/issues/issues/650")

	_, ctx, _, cleanup := testInit(t, false)
	defer cleanup()

	client := v23.GetClient(ctx)
	ctx, _ = context.WithTimeout(ctx, 100*time.Millisecond)
	name := naming.JoinAddressName(naming.FormatEndpoint("tcp", "203.0.113.10:443"), "")
	_, err := client.StartCall(ctx, name, "echo", []interface{}{"args don't matter"})
	t.Log(err)
	if verror.ErrorID(err) != verror.ErrTimeout.ID {
		t.Fatalf("wrong error: %s", err)
	}
}

func logErrors(t *testing.T, msg string, logerr, logstack, debugString bool, err error) {
	_, file, line, _ := runtime.Caller(2)
	loc := fmt.Sprintf("%s:%d", filepath.Base(file), line)
	if logerr {
		t.Logf("%s: %s: %v", loc, msg, err)
	}
	if logstack {
		t.Logf("%s: %s: %v", loc, msg, verror.Stack(err).String())
	}
	if debugString {
		t.Logf("%s: %s: %v", loc, msg, verror.DebugString(err))
	}
}

func TestStartCallErrors(t *testing.T) {
	// TODO(suharshs,mattr): This test will be enabled once the new client is written and solves the flakiness.
	t.Skip("TODO(suharshs,mattr): This test is very flaky and is temporarily disabled.")

	sh, ctx, _, cleanup := testInit(t, false)
	defer cleanup()

	client := v23.GetClient(ctx)
	ns := v23.GetNamespace(ctx)

	logErr := func(msg string, err error) {
		logErrors(t, msg, true, false, false, err)
	}

	emptyCtx := &context.T{}
	emptyCtx = context.WithLogger(emptyCtx, logger.Global())
	_, err := client.StartCall(emptyCtx, "noname", "nomethod", nil)
	if verror.ErrorID(err) != verror.ErrBadArg.ID {
		t.Fatalf("wrong error: %s", err)
	}
	logErr("no context", err)

	// TODO(mattr): This test doesn't pass because xclient doesn't look for this
	// error.
	// p1 := security.PublicKeyAuthorizer(testutil.NewPrincipal().PublicKey())
	// p2 := security.PublicKeyAuthorizer(testutil.NewPrincipal().PublicKey())
	// _, err = client.StartCall(ctx, "noname", "nomethod", nil, options.ServerAuthorizer{p1}, options.ServerAuthorizer{p2})
	// if verror.ErrorID(err) != verror.ErrBadArg.ID {
	// 	t.Fatalf("wrong error: %s", err)
	// }
	// logErr("too many public keys", err)

	// This will fail with NoServers, but because there is no mount table
	// to communicate with. The error message should include a
	// 'connection refused' string.
	ns.SetRoots("/127.0.0.1:8101")
	_, err = client.StartCall(ctx, "noname", "nomethod", nil, options.NoRetry{})
	if verror.ErrorID(err) != verror.ErrNoServers.ID {
		t.Fatalf("wrong error: %s", err)
	}
	if want := "connection refused"; !strings.Contains(verror.DebugString(err), want) {
		t.Fatalf("wrong error: %s - doesn't contain %q", err, want)
	}
	logErr("no mount table", err)

	// This will fail with NoServers, but because there really is no
	// name registered with the mount table.
	startRootMT(t, sh, false)
	_, err = client.StartCall(ctx, "noname", "nomethod", nil, options.NoRetry{})
	if verror.ErrorID(err) != verror.ErrNoServers.ID {
		t.Fatalf("wrong error: %s", err)
	}
	if unwanted := "connection refused"; strings.Contains(err.Error(), unwanted) {
		t.Fatalf("wrong error: %s - does contain %q", err, unwanted)
	}
	logErr("no name registered", err)

	// The following tests will fail with NoServers, but because there are
	// no protocols that the client and servers (mount table, and "name") share.
	nctx, nclient, err := v23.WithNewClient(ctx, irpc.PreferredProtocols([]string{"wsh"}))

	addr := naming.FormatEndpoint("nope", "127.0.0.1:1081")
	if err := ns.Mount(ctx, "name", addr, time.Minute); err != nil {
		t.Fatal(err)
	}

	// This will fail in its attempt to call ResolveStep to the mount table
	// because we are using both the new context and the new client.
	_, err = nclient.StartCall(nctx, "name", "nomethod", nil, options.NoRetry{})
	if verror.ErrorID(err) != verror.ErrNoServers.ID {
		t.Fatalf("wrong error: %s", err)
	}
	if want := "ResolveStep"; !strings.Contains(err.Error(), want) {
		t.Fatalf("wrong error: %s - doesn't contain %q", err, want)
	}
	logErr("mismatched protocols", err)

	// This will fail in its attempt to invoke the actual RPC because
	// we are using the old context (which supplies the context for the calls
	// to ResolveStep) and the new client.
	_, err = nclient.StartCall(ctx, "name", "nomethod", nil, options.NoRetry{})
	if verror.ErrorID(err) != verror.ErrNoServers.ID {
		t.Fatalf("wrong error: %s", err)
	}
	if want := "nope"; !strings.Contains(err.Error(), want) {
		t.Fatalf("wrong error: %s - doesn't contain %q", err, want)
	}
	if unwanted := "ResolveStep"; strings.Contains(err.Error(), unwanted) {
		t.Fatalf("wrong error: %s - does contain %q", err, unwanted)

	}
	logErr("mismatched protocols", err)

	// The following two tests will fail due to a timeout.
	ns.SetRoots("/203.0.113.10:8101")
	nctx, _ = context.WithTimeout(ctx, 100*time.Millisecond)
	// This call will timeout talking to the mount table.
	call, err := client.StartCall(nctx, "name", "noname", nil, options.NoRetry{})
	if verror.ErrorID(err) != verror.ErrTimeout.ID {
		t.Fatalf("wrong error: %s", err)
	}
	if call != nil {
		t.Fatalf("expected call to be nil")
	}
	logErr("timeout to mount table", err)

	// TODO(mattr): This test is temporarily disabled.  It's very flaky
	// and we're about to move to a new version of the client.  We will
	// fix it in the new client.
	// This, second test, will fail due a timeout contacting the server itself.
	// addr = naming.FormatEndpoint("tcp", "203.0.113.10:8101")

	// nctx, _ = context.WithTimeout(ctx, 100*time.Millisecond)
	// new_name := naming.JoinAddressName(addr, "")
	// call, err = client.StartCall(nctx, new_name, "noname", nil, options.NoRetry{})
	// if verror.ErrorID(err) != verror.ErrTimeout.ID {
	// 	t.Fatalf("wrong error: %s", err)
	// }
	// if call != nil {
	// 	t.Fatalf("expected call to be nil")
	// }
	// logErr("timeout to server", err)
}

type closeConn struct {
	ctx *context.T
	flow.Conn
	wg *sync.WaitGroup
}

func (c *closeConn) ReadMsg() ([]byte, error) {
	buf, err := c.Conn.ReadMsg()
	if err == nil {
		if m, err := message.Read(c.ctx, buf); err == nil {
			if _, ok := m.(*message.Data); ok {
				c.Conn.Close()
				c.wg.Done()
				return nil, io.EOF
			}
		}
	}
	return buf, err
}

func TestStartCallBadProtocol(t *testing.T) {
	sh := v23test.NewShell(t, v23test.Opts{})
	ctx := sh.Ctx
	defer sh.Cleanup()
	startRootMT(t, sh, true)

	client := v23.GetClient(ctx)

	logErr := func(msg string, err error) {
		logErrors(t, msg, true, false, false, err)
	}

	wg := &sync.WaitGroup{}
	nctx := debug.WithFilter(ctx, func(c flow.Conn) flow.Conn {
		wg.Add(1)
		return &closeConn{ctx: ctx, Conn: c, wg: wg}
	})
	call, err := client.StartCall(nctx, "name", "noname", nil, options.NoRetry{})
	if verror.ErrorID(err) != verror.ErrNoServers.ID {
		t.Errorf("wrong error: %s", verror.DebugString(err))
	}
	if call != nil {
		t.Errorf("expected call to be nil")
	}
	logErr("broken connection", err)

	// Make sure we failed because we really did close the connection
	// with our filter
	wg.Wait()
}

func TestStartCallSecurity(t *testing.T) {
	_, ctx, name, cleanup := testInit(t, true)
	defer cleanup()

	client := v23.GetClient(ctx)

	logErr := func(msg string, err error) {
		logErrors(t, msg, true, false, false, err)
	}

	// Create a context with a new principal that doesn't match the server,
	// so that the client will not trust the server.
	ctx1, err := v23.WithPrincipal(ctx, testutil.NewPrincipal("test-blessing"))
	if err != nil {
		t.Fatal(err)
	}
	call, err := client.StartCall(ctx1, name, "noname", nil, options.NoRetry{})
	if verror.ErrorID(err) != verror.ErrNotTrusted.ID {
		t.Fatalf("wrong error: %s", err)
	}
	if call != nil {
		t.Fatalf("expected call to be nil")
	}
	logErr("client does not trust server", err)
}

var childPing = gosh.RegisterFunc("childPing", func(name string) error {
	ctx, shutdown := v23.Init()
	defer shutdown()

	got := ""
	if err := v23.GetClient(ctx).Call(ctx, name, "Ping", nil, []interface{}{&got}); err != nil {
		return fmt.Errorf("unexpected error: %s", err)
	}
	fmt.Printf("RESULT=%s\n", got)
	return nil
})

func TestTimeoutResponse(t *testing.T) {
	_, ctx, name, cleanup := testInit(t, true)
	defer cleanup()

	ctx, cancel := context.WithTimeout(ctx, time.Millisecond)
	err := v23.GetClient(ctx).Call(ctx, name, "Sleep", nil, nil)
	if got, want := verror.ErrorID(err), verror.ErrTimeout.ID; got != want {
		t.Fatalf("got %v, want %v", verror.DebugString(err), want)
	}
	cancel()
}

func TestArgsAndResponses(t *testing.T) {
	_, ctx, name, cleanup := testInit(t, true)
	defer cleanup()

	call, err := v23.GetClient(ctx).StartCall(ctx, name, "Sleep", []interface{}{"too many args"})
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	err = call.Finish()
	if got, want := verror.ErrorID(err), verror.ErrBadProtocol.ID; got != want {
		t.Fatalf("got %v, want %v", got, want)
	}

	call, err = v23.GetClient(ctx).StartCall(ctx, name, "Ping", nil)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	pong := ""
	dummy := ""
	err = call.Finish(&pong, &dummy)
	if got, want := verror.ErrorID(err), verror.ErrBadProtocol.ID; got != want {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestAccessDenied(t *testing.T) {
	_, ctx, name, cleanup := testInit(t, true)
	defer cleanup()

	ctx1, err := v23.WithPrincipal(ctx, testutil.NewPrincipal("test-blessing"))
	// Client must recognize the server, otherwise it won't even send the request.
	security.AddToRoots(v23.GetPrincipal(ctx1), v23.GetPrincipal(ctx).BlessingStore().Default())
	if err != nil {
		t.Fatal(err)
	}
	call, err := v23.GetClient(ctx1).StartCall(ctx1, name, "Sleep", nil)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	err = call.Finish()
	if got, want := verror.ErrorID(err), verror.ErrNoAccess.ID; got != want {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestCanceledBeforeFinish(t *testing.T) {
	_, ctx, name, cleanup := testInit(t, true)
	defer cleanup()

	ctx, cancel := context.WithCancel(ctx)
	call, err := v23.GetClient(ctx).StartCall(ctx, name, "Sleep", nil)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	// Cancel before we call finish.
	cancel()
	err = call.Finish()
	if got, want := verror.ErrorID(err), verror.ErrCanceled.ID; got != want {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestCanceledDuringFinish(t *testing.T) {
	_, ctx, name, cleanup := testInit(t, true)
	defer cleanup()

	ctx, cancel := context.WithCancel(ctx)
	call, err := v23.GetClient(ctx).StartCall(ctx, name, "Sleep", nil)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	// Cancel whilst the RPC is running.
	go func() {
		time.Sleep(100 * time.Millisecond)
		cancel()
	}()
	err = call.Finish()
	if got, want := verror.ErrorID(err), verror.ErrCanceled.ID; got != want {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestRendezvous(t *testing.T) {
	sh, ctx, _, cleanup := testInit(t, false)
	defer cleanup()

	// We start the client before we start the server, StartCall will reresolve
	// the name until it finds an entry or times out after an exponential
	// backoff of some minutes.
	name := "echoServer"
	go func() {
		time.Sleep(100 * time.Millisecond)
		startEchoServer(t, sh, "message", name, "")
	}()

	call, err := v23.GetClient(ctx).StartCall(ctx, name, "Echo", []interface{}{"hello"})
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	response := ""
	if err := call.Finish(&response); err != nil {
		if got, want := verror.ErrorID(err), verror.ErrCanceled.ID; got != want {
			t.Fatalf("got %v, want %v", got, want)
		}
	}
	if got, want := response, "message: hello\n"; got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestCallback(t *testing.T) {
	sh, _, name, cleanup := testInit(t, true)
	defer cleanup()

	cmd := sh.FuncCmd(childPing, name)
	cmd.Start()
	if got, want := cmd.S.ExpectVar("RESULT"), "pong"; got != want {
		t.Errorf("got %q, want %q", got, want)
	}
	cmd.Wait()
}

func TestStreamTimeout(t *testing.T) {
	_, ctx, name, cleanup := testInit(t, true)
	defer cleanup()

	want := 10
	ctx, _ = context.WithTimeout(ctx, 300*time.Millisecond)
	call, err := v23.GetClient(ctx).StartCall(ctx, name, "Source", []interface{}{want})
	if err != nil {
		if verror.ErrorID(err) != verror.ErrTimeout.ID {
			t.Fatalf("verror should be a timeout not %s: stack %s",
				err, verror.Stack(err))
		}
		return
	}

	for {
		got := 0
		err := call.Recv(&got)
		if err == nil {
			if got != want {
				t.Fatalf("got %d, want %d", got, want)
			}
			want++
			continue
		}
		if got, want := verror.ErrorID(err), verror.ErrTimeout.ID; got != want {
			t.Fatalf("got %v, want %v", got, want)
		}
		break
	}
	err = call.Finish()
	if got, want := verror.ErrorID(err), verror.ErrTimeout.ID; got != want {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestStreamAbort(t *testing.T) {
	_, ctx, name, cleanup := testInit(t, true)
	defer cleanup()

	call, err := v23.GetClient(ctx).StartCall(ctx, name, "Sink", nil)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	want := 10
	for i := 0; i <= want; i++ {
		if err := call.Send(i); err != nil {
			t.Fatalf("unexpected error: %s", err)
		}
	}
	call.CloseSend()
	err = call.Send(100)
	if got, want := verror.ErrorID(err), verror.ErrAborted.ID; got != want {
		t.Fatalf("got %v, want %v", got, want)
	}

	result := 0
	err = call.Finish(&result)
	if err != nil {
		t.Errorf("unexpected error: %#v", err)
	}
	if got := result; got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestNoServersAvailable(t *testing.T) {
	_, ctx, _, cleanup := testInit(t, false)
	defer cleanup()

	name := "noservers"
	call, err := v23.GetClient(ctx).StartCall(ctx, name, "Sleep", nil, options.NoRetry{})
	if err != nil {
		if got, want := verror.ErrorID(err), verror.ErrNoServers.ID; got != want {
			t.Fatalf("got %v, want %v", got, want)
		}
		return
	}
	err = call.Finish()
	if got, want := verror.ErrorID(err), verror.ErrNoServers.ID; got != want {
		t.Fatalf("got %v, want %v", got, want)
	}

}

func TestNoMountTable(t *testing.T) {
	_, ctx, _, cleanup := testInit(t, false)
	defer cleanup()

	v23.GetNamespace(ctx).SetRoots()
	name := "a_mount_table_entry"

	// If there is no mount table, then we'll get a NoServers error message.
	ctx, _ = context.WithTimeout(ctx, 300*time.Millisecond)
	_, err := v23.GetClient(ctx).StartCall(ctx, name, "Sleep", nil)
	if got, want := verror.ErrorID(err), verror.ErrNoServers.ID; got != want {
		t.Fatalf("got %v, want %v", got, want)
	}
}

// TestReconnect verifies that the client transparently re-establishes the
// connection to the server if the server dies and comes back (on the same
// endpoint).
func TestReconnect(t *testing.T) {
	t.Skip()

	sh, ctx, _, cleanup := testInit(t, false)
	defer cleanup()

	cmd, serverName := startEchoServer(t, sh, "mymessage", "", "")
	serverEP, _ := naming.SplitAddressName(serverName)
	ep, _ := inaming.NewEndpoint(serverEP)

	makeCall := func(ctx *context.T, opts ...rpc.CallOpt) (string, error) {
		ctx, _ = context.WithDeadline(ctx, time.Now().Add(10*time.Second))
		call, err := v23.GetClient(ctx).StartCall(ctx, serverName, "Echo", []interface{}{"bratman"}, opts...)
		if err != nil {
			return "", fmt.Errorf("START: %s", err)
		}
		var result string
		if err := call.Finish(&result); err != nil {
			return "", err
		}
		return result, nil
	}

	expected := "mymessage: bratman\n"
	if result, err := makeCall(ctx); err != nil || result != expected {
		t.Errorf("Got (%q, %v) want (%q, nil)", result, err, expected)
	}

	// Kill the server, verify client can't talk to it anymore.
	cmd.Terminate(os.Interrupt)
	if _, err := makeCall(ctx, options.NoRetry{}); err == nil || (!strings.HasPrefix(err.Error(), "START") && !strings.Contains(err.Error(), "EOF")) {
		t.Fatalf(`Got (%v) want ("START: <err>" or "EOF") as server is down`, err)
	}

	// Resurrect the server with the same address, verify client
	// re-establishes the connection. This is racy if another
	// process grabs the port.
	cmd, _ = startEchoServer(t, sh, "mymessage again", "", ep.Address)
	defer cmd.Terminate(os.Interrupt)
	expected = "mymessage again: bratman\n"
	if result, err := makeCall(ctx); err != nil || result != expected {
		t.Errorf("Got (%q, %v) want (%q, nil)", result, err, expected)
	}
}

func TestMethodErrors(t *testing.T) {
	_, ctx, name, cleanup := testInit(t, true)
	defer cleanup()

	var (
		i, j int
		s    string
	)

	testCases := []struct {
		testName, objectName, method string
		args, results                []interface{}
		wantID                       verror.ID
		wantMessage                  string
	}{
		{
			testName:   "unknown method",
			objectName: name,
			method:     "NoMethod",
			wantID:     verror.ErrUnknownMethod.ID,
		},
		{
			testName:   "unknown suffix",
			objectName: name + "/NoSuffix",
			method:     "Ping",
			wantID:     verror.ErrUnknownSuffix.ID,
		},
		{
			testName:    "too many args",
			objectName:  name,
			method:      "Ping",
			args:        []interface{}{1, 2},
			results:     []interface{}{&i},
			wantID:      verror.ErrBadProtocol.ID,
			wantMessage: "wrong number of input arguments",
		},
		{
			testName:    "wrong number of results",
			objectName:  name,
			method:      "Ping",
			results:     []interface{}{&i, &j},
			wantID:      verror.ErrBadProtocol.ID,
			wantMessage: "results, but want",
		},
		{
			testName:    "wrong number of results",
			objectName:  name,
			method:      "Ping",
			results:     []interface{}{&i, &j},
			wantID:      verror.ErrBadProtocol.ID,
			wantMessage: "results, but want",
		},
		{
			testName:    "mismatched arg types",
			objectName:  name,
			method:      "Echo",
			args:        []interface{}{1},
			results:     []interface{}{&s},
			wantID:      verror.ErrBadProtocol.ID,
			wantMessage: "aren't compatible",
		},
		{
			testName:    "mismatched result types",
			objectName:  name,
			method:      "Ping",
			results:     []interface{}{&i},
			wantID:      verror.ErrBadProtocol.ID,
			wantMessage: "aren't compatible",
		},
	}

	clt := v23.GetClient(ctx)
	for _, test := range testCases {
		testPrefix := fmt.Sprintf("test(%s) failed", test.testName)
		call, err := clt.StartCall(ctx, test.objectName, test.method, test.args)
		if err != nil {
			t.Fatalf("%s: %v", testPrefix, err)
		}
		verr := call.Finish(test.results...)
		if verror.ErrorID(verr) != test.wantID {
			t.Errorf("%s: wrong error: %v", testPrefix, verr)
		} else if got, want := verr.Error(), test.wantMessage; !strings.Contains(got, want) {
			t.Errorf("%s: want %q to contain %q", testPrefix, got, want)
		}
		logErrors(t, test.testName, false, false, false, verr)
	}
}

func TestReservedMethodErrors(t *testing.T) {
	_, ctx, name, cleanup := testInit(t, true)
	defer cleanup()

	logErr := func(msg string, err error) {
		logErrors(t, msg, true, false, false, err)
	}

	// This call will fail because the __xx suffix is not supported by
	// the dispatcher implementing Signature.
	clt := v23.GetClient(ctx)
	call, err := clt.StartCall(ctx, name+"/__xx", "__Signature", nil)
	if err != nil {
		t.Fatal(err)
	}
	sig := []signature.Interface{}
	verr := call.Finish(&sig)
	if verror.ErrorID(verr) != verror.ErrUnknownSuffix.ID {
		t.Fatalf("wrong error: %s", verr)
	}
	logErr("unknown suffix", verr)

	// This call will fail for the same reason, but with a different error,
	// saying that MethodSignature is an unknown method.
	call, err = clt.StartCall(ctx, name+"/__xx", "__MethodSignature", []interface{}{"dummy"})
	if err != nil {
		t.Fatal(err)
	}
	verr = call.Finish(&sig)
	if verror.ErrorID(verr) != verror.ErrUnknownMethod.ID {
		t.Fatalf("wrong error: %s", verr)
	}
	logErr("unknown method", verr)
}

type changeProtocol struct {
	tcp tcputil.TCP

	mu      sync.Mutex
	count   int
	address string
}

func (c *changeProtocol) Dial(ctx *context.T, protocol, address string, timeout time.Duration) (flow.Conn, error) {
	return c.tcp.Dial(ctx, "tcp", c.address, timeout)
}
func (c *changeProtocol) Resolve(ctx *context.T, proctocol, address string) (string, []string, error) {
	defer c.mu.Unlock()
	c.mu.Lock()
	// Return a new resolved address everytime to disallow cache hits based on resolved address.
	c.count++
	return "tcp", []string{fmt.Sprintf("resolved%d", c.count)}, nil
}
func (c *changeProtocol) Listen(ctx *context.T, protocol, address string) (flow.Listener, error) {
	protocol = "tcp"
	address = "127.0.0.1:0"
	l, err := c.tcp.Listen(ctx, protocol, address)
	if err != nil {
		return nil, err
	}
	c.mu.Lock()
	c.address = l.Addr().String()
	c.mu.Unlock()
	return l, nil
}

func TestDNSResolutionChange(t *testing.T) {
	// Scenario: A DNS Resolution changes the ip address a hostname resolves to. It
	// is important that we make sure that the flow and type flow dialed are still on
	// the same connection or the rpc will never succeed.
	flow.RegisterProtocol("change", &changeProtocol{})
	ctx, shutdown := test.V23Init()
	defer shutdown()

	ctx, cancel := context.WithCancel(ctx)
	ctx = v23.WithListenSpec(ctx, rpc.ListenSpec{
		Addrs: rpc.ListenAddrs{{Protocol: "change", Address: "hostname"}},
	})
	_, server, err := v23.WithNewServer(ctx, "", &testServer{}, nil)
	if err != nil {
		t.Fatal(err)
	}
	defer func() { <-server.Closed() }()
	defer cancel()

	// The call should succeed.
	if err := v23.GetClient(ctx).Call(ctx, "/@6@change@hostname@@00000000000000000000000000000000@s@@@", "Closure", nil, nil); err != nil {
		t.Error(err)
	}
}
