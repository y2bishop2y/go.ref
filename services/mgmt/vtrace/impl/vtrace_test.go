package impl_test

import (
	"io"
	"testing"

	"v.io/core/veyron2"
	"v.io/core/veyron2/ipc"
	"v.io/core/veyron2/naming"
	"v.io/core/veyron2/rt"
	service "v.io/core/veyron2/services/mgmt/vtrace"
	"v.io/core/veyron2/vtrace"

	"v.io/core/veyron/profiles"
	"v.io/core/veyron/services/mgmt/vtrace/impl"
)

func setup(t *testing.T) (string, ipc.Server, veyron2.Runtime) {
	runtime, err := rt.New()
	if err != nil {
		t.Fatalf("Could not create runtime: %s", err)
	}

	server, err := runtime.NewServer()
	if err != nil {
		t.Fatalf("Could not create server: %s", err)
	}
	endpoints, err := server.Listen(profiles.LocalListenSpec)
	if err != nil {
		t.Fatalf("Listen failed: %s", err)
	}
	if err := server.Serve("", impl.NewVtraceService(), nil); err != nil {
		t.Fatalf("Serve failed: %s", err)
	}
	return endpoints[0].String(), server, runtime
}

func TestVtraceServer(t *testing.T) {
	endpoint, server, runtime := setup(t)
	defer server.Stop()

	sctx := runtime.NewContext()
	sctx, span := vtrace.SetNewSpan(sctx, "The Span")
	vtrace.ForceCollect(sctx)
	span.Finish()
	id := span.Trace()

	client := service.StoreClient(naming.JoinAddressName(endpoint, ""))

	trace, err := client.Trace(runtime.NewContext(), id)
	if err != nil {
		t.Fatalf("Unexpected error getting trace: %s", err)
	}
	if len(trace.Spans) != 1 {
		t.Errorf("Returned trace should have 1 span, found %#v", trace)
	}
	if trace.Spans[0].Name != "The Span" {
		t.Errorf("Returned span has wrong name: %#v", trace)
	}

	call, err := client.AllTraces(runtime.NewContext())
	if err != nil {
		t.Fatalf("Unexpected error getting traces: %s", err)
	}
	ntraces := 0
	stream := call.RecvStream()
	var tr *vtrace.TraceRecord
	for stream.Advance() {
		trace := stream.Value()
		if trace.ID == id {
			tr = &trace
		}
		ntraces++
	}
	if err = stream.Err(); err != nil && err != io.EOF {
		t.Fatalf("Unexpected error reading trace stream: %s", err)
	}
	if ntraces != 1 {
		t.Fatalf("Expected 1 trace, got %#v", ntraces)
	}
	if tr == nil {
		t.Fatalf("Desired trace %x not found.", id)
	}
	if len(tr.Spans) != 1 {
		t.Errorf("Returned trace should have 1 span, found %#v", tr)
	}
	if tr.Spans[0].Name != "The Span" {
		t.Fatalf("Returned span has wrong name: %#v", tr)
	}
}
