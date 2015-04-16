// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package statslib implements the Stats interface from
// v.io/v23/services/stats.
package statslib

import (
	"reflect"
	"time"

	"v.io/v23/context"
	"v.io/v23/naming"
	"v.io/v23/rpc"
	"v.io/v23/services/stats"
	"v.io/v23/services/watch"
	"v.io/v23/vdl"
	"v.io/v23/verror"
	"v.io/x/lib/vlog"
	libstats "v.io/x/ref/lib/stats"
)

type statsService struct {
	suffix    string
	watchFreq time.Duration
}

const pkgPath = "v.io/x/ref/services/internal/statslib"

var (
	errOperationFailed = verror.Register(pkgPath+".errOperationFailed", verror.NoRetry, "{1:}{2:} operation failed{:_}")
)

// NewStatsService returns a stats server implementation. The value of watchFreq
// is used to specify the time between WatchGlob updates.
func NewStatsService(suffix string, watchFreq time.Duration) interface{} {
	return stats.StatsServer(&statsService{suffix, watchFreq})
}

// Glob__ returns the name of all objects that match pattern.
func (i *statsService) Glob__(ctx *context.T, call rpc.ServerCall, pattern string) (<-chan naming.GlobReply, error) {
	vlog.VI(1).Infof("%v.Glob__(%q)", i.suffix, pattern)

	ch := make(chan naming.GlobReply)
	go func() {
		defer close(ch)
		it := libstats.Glob(i.suffix, pattern, time.Time{}, false)
		for it.Advance() {
			ch <- naming.GlobReplyEntry{naming.MountEntry{Name: it.Value().Key}}
		}
		if err := it.Err(); err != nil {
			vlog.VI(1).Infof("libstats.Glob(%q, %q) failed: %v", i.suffix, pattern, err)
		}
	}()
	return ch, nil
}

// WatchGlob returns the name and value of the objects that match the request,
// followed by periodic updates when values change.
func (i *statsService) WatchGlob(ctx *context.T, call watch.GlobWatcherWatchGlobServerCall, req watch.GlobRequest) error {
	vlog.VI(1).Infof("%v.WatchGlob(%+v)", i.suffix, req)

	var t time.Time
Loop:
	for {
		prevTime := t
		t = time.Now()
		it := libstats.Glob(i.suffix, req.Pattern, prevTime, true)
		changes := []watch.Change{}
		for it.Advance() {
			v := it.Value()
			c := watch.Change{
				Name:  v.Key,
				State: watch.Exists,
				Value: vdl.ValueOf(v.Value),
			}
			changes = append(changes, c)
		}
		if err := it.Err(); err != nil {
			if verror.ErrorID(err) == verror.ErrNoExist.ID {
				return verror.New(verror.ErrNoExist, ctx, i.suffix)
			}
			return verror.New(errOperationFailed, ctx, i.suffix)
		}
		for _, change := range changes {
			if err := call.SendStream().Send(change); err != nil {
				return err
			}
		}
		select {
		case <-ctx.Done():
			break Loop
		case <-time.After(i.watchFreq):
		}
	}
	return nil
}

// Value returns the value of the receiver object.
func (i *statsService) Value(ctx *context.T, _ rpc.ServerCall) (*vdl.Value, error) {
	vlog.VI(1).Infof("%v.Value()", i.suffix)

	rv, err := libstats.Value(i.suffix)
	switch {
	case verror.ErrorID(err) == verror.ErrNoExist.ID:
		return nil, verror.New(verror.ErrNoExist, ctx, i.suffix)
	case verror.ErrorID(err) == stats.ErrNoValue.ID:
		return nil, stats.NewErrNoValue(ctx, i.suffix)
	case err != nil:
		return nil, verror.New(errOperationFailed, ctx, i.suffix)
	}
	vv, err := vdl.ValueFromReflect(reflect.ValueOf(rv))
	if err != nil {
		return nil, verror.New(verror.ErrInternal, ctx, i.suffix, err)
	}
	return vv, nil
}