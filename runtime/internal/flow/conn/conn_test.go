// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package conn

import (
	"bytes"
	"crypto/rand"
	"io"
	"testing"
	"time"

	"v.io/v23"
	_ "v.io/x/ref/runtime/factories/fake"
	"v.io/x/ref/test"
	"v.io/x/ref/test/goroutines"
)

const leakWaitTime = 100 * time.Millisecond

var randData []byte

func init() {
	test.Init()
	randData = make([]byte, 2*defaultBufferSize)
	if _, err := rand.Read(randData); err != nil {
		panic("Could not read random data.")
	}
}

func TestLargeWrite(t *testing.T) {
	defer goroutines.NoLeaks(t, leakWaitTime)()

	ctx, shutdown := v23.Init()
	df, flows, cl := setupFlow(t, ctx, ctx, true)
	defer cl()
	defer shutdown()

	want := randData
	finished := make(chan struct{})
	go func(x []byte) {
		mid := len(x) / 2
		wrote, err := df.WriteMsgAndClose(x[:mid], x[mid:])
		if err != nil {
			t.Fatalf("Unexpected error for write: %v", err)
		}
		if wrote != len(x) {
			t.Errorf("got %d want %d", wrote, len(x))
		}
		close(finished)
	}(want)

	af := <-flows
	read := 0
	for len(want) > 0 {
		got, err := af.ReadMsg()
		if err != nil && err != io.EOF {
			t.Fatalf("Unexpected error: %v", err)
		}
		if !bytes.Equal(got, want[:len(got)]) {
			pl := len(got)
			if pl > 100 {
				pl = 100
			}
			pg, pw := got[:pl], want[:pl]
			t.Fatalf("On read %d got: %v want %v", read, pg, pw)
		}
		want = want[len(got):]
		read++
	}
	if len(want) != 0 {
		t.Errorf("got %d leftover bytes, expected 0.", len(want))
	}
	<-finished
	<-df.Closed()
	<-af.Closed()
}
