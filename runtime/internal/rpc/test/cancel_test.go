// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package test

import (
	"io"
	"testing"

	"v.io/v23"
	"v.io/v23/context"
	"v.io/v23/rpc"
	"v.io/v23/security"
	"v.io/v23/verror"
	"v.io/x/ref"
	"v.io/x/ref/runtime/internal/flow/conn"
	"v.io/x/ref/runtime/internal/rpc/stream/vc"
	"v.io/x/ref/test"
)

type canceld struct {
	name     string
	child    string
	started  chan struct{}
	canceled chan struct{}
}

func (c *canceld) Run(ctx *context.T, _ rpc.ServerCall) error {
	close(c.started)
	client := v23.GetClient(ctx)
	var done chan struct{}
	if c.child != "" {
		done = make(chan struct{})
		go func() {
			client.Call(ctx, c.child, "Run", nil, nil)
			close(done)
		}()
	}
	<-ctx.Done()
	if done != nil {
		<-done
	}
	close(c.canceled)
	return nil
}

func makeCanceld(ctx *context.T, name, child string) (*canceld, error) {
	c := &canceld{
		name:     name,
		child:    child,
		started:  make(chan struct{}, 0),
		canceled: make(chan struct{}, 0),
	}
	_, _, err := v23.WithNewServer(ctx, name, c, security.AllowEveryone())
	if err != nil {
		return nil, err
	}
	return c, nil
}

// TestCancellationPropagation tests that cancellation propogates along an
// RPC call chain without user intervention.
func TestCancellationPropagation(t *testing.T) {
	ctx, shutdown := test.V23Init()
	defer shutdown()

	c1, err := makeCanceld(ctx, "c1", "c2")
	if err != nil {
		t.Fatalf("Can't start server:", err, verror.DebugString(err))
	}
	c2, err := makeCanceld(ctx, "c2", "")
	if err != nil {
		t.Fatalf("Can't start server:", err)
	}

	ctx, cancel := context.WithCancel(ctx)
	done := make(chan struct{})
	go func() {
		v23.GetClient(ctx).Call(ctx, "c1", "Run", nil, nil)
		close(done)
	}()

	<-c1.started
	<-c2.started
	cancel()
	<-c1.canceled
	<-c2.canceled
	<-done
}

type cancelTestServer struct {
	started   chan struct{}
	cancelled chan struct{}
	t         *testing.T
}

func newCancelTestServer(t *testing.T) *cancelTestServer {
	return &cancelTestServer{
		started:   make(chan struct{}),
		cancelled: make(chan struct{}),
		t:         t,
	}
}

func (s *cancelTestServer) CancelStreamReader(ctx *context.T, call rpc.StreamServerCall) error {
	close(s.started)
	var b []byte
	if err := call.Recv(&b); err != io.EOF {
		s.t.Errorf("Got error %v, want io.EOF", err)
	}
	<-ctx.Done()
	close(s.cancelled)
	return nil
}

// CancelStreamIgnorer doesn't read from it's input stream so all it's
// buffers fill.  The intention is to show that call.Done() is closed
// even when the stream is stalled.
func (s *cancelTestServer) CancelStreamIgnorer(ctx *context.T, _ rpc.StreamServerCall) error {
	close(s.started)
	<-ctx.Done()
	close(s.cancelled)
	return nil
}

func waitForCancel(t *testing.T, ts *cancelTestServer, cancel context.CancelFunc) {
	<-ts.started
	cancel()
	<-ts.cancelled
}

// TestCancel tests cancellation while the server is reading from a stream.
func TestCancel(t *testing.T) {
	ctx, shutdown := test.V23Init()
	defer shutdown()
	var (
		sctx = withPrincipal(t, ctx, "server")
		cctx = withPrincipal(t, ctx, "client")
		ts   = newCancelTestServer(t)
	)
	_, _, err := v23.WithNewServer(sctx, "cancel", ts, security.AllowEveryone())
	if err != nil {
		t.Fatal(err)
	}
	cctx, cancel := context.WithCancel(cctx)
	done := make(chan struct{})
	go func() {
		v23.GetClient(cctx).Call(cctx, "cancel", "CancelStreamReader", nil, nil)
		close(done)
	}()
	waitForCancel(t, ts, cancel)
	<-done
}

// TestCancelWithFullBuffers tests that even if the writer has filled the buffers and
// the server is not reading that the cancel message gets through.
func TestCancelWithFullBuffers(t *testing.T) {
	ctx, shutdown := test.V23Init()
	defer shutdown()
	var (
		sctx = withPrincipal(t, ctx, "server")
		cctx = withPrincipal(t, ctx, "client")
		ts   = newCancelTestServer(t)
	)
	_, _, err := v23.WithNewServer(sctx, "cancel", ts, security.AllowEveryone())
	if err != nil {
		t.Fatal(err)
	}
	cctx, cancel := context.WithCancel(cctx)
	call, err := v23.GetClient(cctx).StartCall(cctx, "cancel", "CancelStreamIgnorer", nil)
	if err != nil {
		t.Fatalf("Start call failed: %v", err)
	}

	// Fill up all the write buffers to ensure that cancelling works even when the stream
	// is blocked.
	if ref.RPCTransitionState() >= ref.XServers {
		call.Send(conn.DefaultBytesBufferedPerFlow)
	} else {
		call.Send(make([]byte, vc.MaxSharedBytes))
		call.Send(make([]byte, vc.DefaultBytesBufferedPerFlow))
	}
	done := make(chan struct{})
	go func() {
		call.Finish()
		close(done)
	}()

	waitForCancel(t, ts, cancel)
	<-done
}
