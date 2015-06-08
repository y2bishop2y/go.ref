// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main_test

import (
	"bytes"
	"fmt"
	"io"
	"math/rand"
	"runtime"
	"strings"
	"sync"
	"testing"
	"time"

	"v.io/v23/context"
	"v.io/v23/naming"
	"v.io/v23/services/device"

	"v.io/x/lib/cmdline"
	"v.io/x/ref/test"

	cmd_device "v.io/x/ref/services/device/device"
)

func simplePrintHandler(entry cmd_device.GlobResult, _ *context.T, stdout, _ io.Writer) error {
	fmt.Fprintf(stdout, "%v\n", entry)
	return nil
}

// newEnforceFullParallelismHandler returns a handler that should be invoked in
// parallel n times.
func newEnforceFullParallelismHandler(t *testing.T, n int) cmd_device.GlobHandler {
	// The WaitGroup is used to verify parallel execution: each run of the
	// handler decrements the counter, and then waits for the counter to
	// reach zero.  If any of the runs of the handler were sequential, the
	// counter would never reach zero since a deadlock would ensue.  A
	// timeout protects against the deadlock to ensure the test fails fast.
	var wg sync.WaitGroup
	wg.Add(n)
	return func(entry cmd_device.GlobResult, ctx *context.T, stdout, stderr io.Writer) error {
		wg.Done()
		simplePrintHandler(entry, ctx, stdout, stderr)
		waitDoneCh := make(chan struct{})
		go func() {
			wg.Wait()
			close(waitDoneCh)
		}()
		select {
		case <-waitDoneCh:
		case <-time.After(5 * time.Second):
			t.Errorf("Timed out waiting for WaitGroup.  Potential parallelism issue.")
		}
		return nil
	}
}

// maybeGosched flips a coin to decide where to call Gosched.  It's used to
// shuffle up the order of handler execution a bit (to prevent the goroutines
// spawned by the glob library from always executing in a fixed order, e.g. the
// order in which they're created).
func maybeGosched() {
	if rand.Intn(2) == 0 {
		runtime.Gosched()
	}
}

// newEnforceKindParallelismHandler returns a handler that should be invoked
// nInstallations times in parallel for installations, then nInstances times in
// parallel for instances, then nDevices times in parallel for device service
// objects.
func newEnforceKindParallelismHandler(t *testing.T, nInstallations, nInstances, nDevices int) cmd_device.GlobHandler {
	// Each of these handlers ensures parallelism within each kind
	// (installations, instances, device objects).
	hInstallations := newEnforceFullParallelismHandler(t, nInstallations)
	hInstances := newEnforceFullParallelismHandler(t, nInstances)
	hDevices := newEnforceFullParallelismHandler(t, nDevices)

	// These channels are used to verify that all installation handlers must
	// execute before all instance handlers, which in turn must execute
	// before all device handlers.
	instancesCh := make(chan struct{}, nInstances)
	devicesCh := make(chan struct{}, nDevices)
	return func(entry cmd_device.GlobResult, ctx *context.T, stdout, stderr io.Writer) error {
		maybeGosched()
		switch entry.Kind {
		case cmd_device.ApplicationInstallationObject:
			select {
			case <-instancesCh:
				t.Errorf("Instance before installation")
			case <-devicesCh:
				t.Errorf("Device before installation")
			default:
			}
			return hInstallations(entry, ctx, stdout, stderr)
		case cmd_device.ApplicationInstanceObject:
			select {
			case <-devicesCh:
				t.Errorf("Device before instance")
			default:
			}
			instancesCh <- struct{}{}
			return hInstances(entry, ctx, stdout, stderr)
		case cmd_device.DeviceServiceObject:
			devicesCh <- struct{}{}
			return hDevices(entry, ctx, stdout, stderr)
		}
		t.Errorf("Unknown entry: %v", entry.Kind)
		return nil
	}
}

// newEnforceNoParallelismHandler returns a handler meant to be invoked sequentially
// for each of the suffixes contained in the expected slice.
func newEnforceNoParallelismHandler(t *testing.T, n int, expected []string) cmd_device.GlobHandler {
	if n != len(expected) {
		t.Errorf("Test invariant broken: %d != %d", n, len(expected))
	}
	orderedSuffixes := make(chan string, n)
	for _, e := range expected {
		orderedSuffixes <- e
	}
	return func(entry cmd_device.GlobResult, ctx *context.T, stdout, stderr io.Writer) error {
		maybeGosched()
		_, suffix := naming.SplitAddressName(entry.Name)
		expect := <-orderedSuffixes
		if suffix != expect {
			t.Errorf("Expected %s, got %s", expect, suffix)
		}
		simplePrintHandler(entry, ctx, stdout, stderr)
		return nil
	}
}

// TestGlob tests the internals of the globbing support for the device tool.
func TestGlob(t *testing.T) {
	ctx, shutdown := test.V23Init()
	defer shutdown()
	tapes := newTapeMap()
	rootTape := tapes.forSuffix("")
	server, endpoint, err := startServer(t, ctx, tapes)
	if err != nil {
		return
	}
	appName := naming.JoinAddressName(endpoint.String(), "app")
	defer stopServer(t, server)

	allGlobArgs := []string{"glob1", "glob2"}
	allGlobResponses := [][]string{
		[]string{"app/3", "app/4", "app/6", "app/5", "app/9", "app/7"},
		[]string{"app/2", "app/1", "app/8"},
	}
	allStatusResponses := map[string][]interface{}{
		"app/1": []interface{}{instanceRunning},
		"app/2": []interface{}{installationUninstalled},
		"app/3": []interface{}{instanceUpdating},
		"app/4": []interface{}{installationActive},
		"app/5": []interface{}{instanceNotRunning},
		"app/6": []interface{}{deviceService},
		"app/7": []interface{}{installationActive},
		"app/8": []interface{}{deviceUpdating},
		"app/9": []interface{}{instanceUpdating},
	}
	outLine := func(suffix string, s device.Status) string {
		r, err := cmd_device.NewGlobResult(appName+"/"+suffix, s)
		if err != nil {
			t.Errorf("NewGlobResult failed: %v", err)
			return ""
		}
		return fmt.Sprintf("%v", *r)
	}
	var (
		app1Out = outLine("1", instanceRunning)
		app2Out = outLine("2", installationUninstalled)
		app3Out = outLine("3", instanceUpdating)
		app4Out = outLine("4", installationActive)
		app5Out = outLine("5", instanceNotRunning)
		app6Out = outLine("6", deviceService)
		app7Out = outLine("7", installationActive)
		app8Out = outLine("8", deviceUpdating)
		app9Out = outLine("9", instanceUpdating)
	)

	noParallelismHandler := newEnforceNoParallelismHandler(t, len(allStatusResponses), []string{"app/2", "app/4", "app/7", "app/1", "app/3", "app/5", "app/9", "app/6", "app/8"})

	for _, c := range []struct {
		handler         cmd_device.GlobHandler
		globResponses   [][]string
		statusResponses map[string][]interface{}
		gs              cmd_device.GlobSettings
		globPatterns    []string
		expected        string
	}{
		// Verifies output is correct and in the expected order (first
		// installations, then instances, then device services).
		{
			simplePrintHandler,
			allGlobResponses,
			allStatusResponses,
			cmd_device.GlobSettings{},
			allGlobArgs,
			joinLines(app2Out, app4Out, app7Out, app1Out, app3Out, app5Out, app9Out, app6Out, app8Out),
		},
		// Verifies that full parallelism runs all the handlers
		// simultaneously.
		{
			newEnforceFullParallelismHandler(t, len(allStatusResponses)),
			allGlobResponses,
			allStatusResponses,
			cmd_device.GlobSettings{HandlerParallelism: cmd_device.FullParallelism},
			allGlobArgs,
			joinLines(app2Out, app4Out, app7Out, app1Out, app3Out, app5Out, app9Out, app6Out, app8Out),
		},
		// Verifies that "by kind" parallelism runs all installation
		// handlers in parallel, then all instance handlers in parallel,
		// then all device service handlers in parallel.
		{
			newEnforceKindParallelismHandler(t, 3, 4, 2),
			allGlobResponses,
			allStatusResponses,
			cmd_device.GlobSettings{HandlerParallelism: cmd_device.KindParallelism},
			allGlobArgs,
			joinLines(app2Out, app4Out, app7Out, app1Out, app3Out, app5Out, app9Out, app6Out, app8Out),
		},
		// Verifies that the no parallelism option runs all handlers
		// sequentially.
		{
			noParallelismHandler,
			allGlobResponses,
			allStatusResponses,
			cmd_device.GlobSettings{HandlerParallelism: cmd_device.NoParallelism},
			allGlobArgs,
			joinLines(app2Out, app4Out, app7Out, app1Out, app3Out, app5Out, app9Out, app6Out, app8Out),
		},
		// TODO(caprita): Test the following cases:
		// Various filters.
		// No glob arguments.
		// No glob results.
		// Error in glob.
		// Error in status.
		// Error in handler.
	} {
		tapes.rewind()
		var rootTapeResponses []interface{}
		for _, r := range c.globResponses {
			rootTapeResponses = append(rootTapeResponses, GlobResponse{r})
		}
		rootTape.SetResponses(rootTapeResponses...)
		for n, r := range c.statusResponses {
			tapes.forSuffix(n).SetResponses(r...)
		}
		var stdout, stderr bytes.Buffer
		env := &cmdline.Env{Stdout: &stdout, Stderr: &stderr}
		var args []string
		for _, p := range c.globPatterns {
			args = append(args, naming.JoinAddressName(endpoint.String(), p))
		}
		err := cmd_device.Run(ctx, env, args, c.handler, c.gs)
		if err != nil {
			t.Errorf("%v", err)
		}
		if expected, got := c.expected, strings.TrimSpace(stdout.String()); got != expected {
			t.Errorf("Unexpected output. Got:\n%v\nExpected:\n%v\n", got, expected)
		}
	}
}