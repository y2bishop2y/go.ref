// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package raft

import (
	"fmt"
	"testing"
	"time"

	"v.io/v23"
	"v.io/x/lib/vlog"
	_ "v.io/x/ref/runtime/factories/generic"
)

// waitForLogAgreement makes sure all working servers agree on the log.
func waitForLogAgreement(rs []*raft, timeout time.Duration) bool {
	start := time.Now()
	for {
		indexMap := make(map[Index]string)
		termMap := make(map[Term]string)
		for _, r := range rs {
			id := r.Id()
			indexMap[r.p.LastIndex()] = id
			termMap[r.p.LastTerm()] = id
		}
		if len(indexMap) == 1 && len(termMap) == 1 {
			vlog.Infof("tada, all logs agree at %v@%v", indexMap, termMap)
			return true
		}
		if time.Now().Sub(start) > timeout {
			vlog.Errorf("oops, logs disagree %v@%v", indexMap, termMap)
			return false
		}
		time.Sleep(10 * time.Millisecond)
	}
}

// waitForAppliedAgreement makes sure all working servers have applied all logged commands.
func waitForAppliedAgreement(rs []*raft, cs []*client, timeout time.Duration) bool {
	start := time.Now()
	for {
		notyet := false
		indexMap := make(map[Index]string)
		termMap := make(map[Term]string)
		for _, r := range rs {
			id, role, _ := r.Status()
			if role == RoleStopped {
				vlog.Infof("ignoring %s", id)
				continue
			}
			if r.p.LastIndex() != r.lastApplied() {
				notyet = true
				break
			}
			indexMap[r.p.LastIndex()] = id
			termMap[r.p.LastTerm()] = id
		}
		if len(indexMap) == 1 && len(termMap) == 1 && !notyet {
			vlog.Infof("tada, all applys agree at %v@%v", indexMap, termMap)
			return true
		}
		if time.Now().Sub(start) > timeout {
			vlog.Errorf("oops, applys disagree %v@%v", indexMap, termMap)
			return false
		}
		time.Sleep(10 * time.Millisecond)
	}
}

// waitForAllRunning waits until all servers are in leader or follower state.
func waitForAllRunning(rs []*raft, timeout time.Duration) bool {
	n := len(rs)
	start := time.Now()
	for {
		i := 0
		for _, r := range rs {
			_, role, _ := r.Status()
			if role == RoleLeader || role == RoleFollower {
				i++
			}
		}
		if i == n {
			return true
		}
		if time.Now().Sub(start) > timeout {
			return false
		}
		time.Sleep(500 * time.Millisecond)
	}
}

func TestAppend(t *testing.T) {
	vlog.Infof("TestAppend")
	ctx, shutdown := v23.Init()
	defer shutdown()

	rs, cs := buildRafts(t, ctx, 3, nil)
	defer cleanUp(rs)
	thb := rs[0].heartbeat

	// One of the raft members should time out not hearing a leader and start an election.
	r1 := waitForElection(t, rs, 5*thb)
	if r1 == nil {
		t.Fatalf("too long to find a leader")
	}
	time.Sleep(time.Millisecond)
	if !waitForLeaderAgreement(rs, thb) {
		t.Fatalf("no leader agreement")
	}

	// Append some entries.
	start := time.Now()
	for i := 0; i < 100; i++ {
		cmd := fmt.Sprintf("the rain in spain %d", i)
		if apperr, err := r1.Append(ctx, []byte(cmd)); apperr != nil || err != nil {
			t.Fatalf("append %s failed with %s", cmd, err)
		}
	}
	if !waitForAppliedAgreement(rs, cs, 2*thb) {
		t.Fatalf("no log agreement")
	}
	vlog.Infof("appends took %s", time.Now().Sub(start))

	// Kill off one instance and see if we keep committing.
	r1.Stop()
	r2 := waitForElection(t, rs, 5*thb)
	if r2 == nil {
		t.Fatalf("too long to find a leader")
	}
	if !waitForLeaderAgreement(rs, thb) {
		t.Fatalf("no leader agreement")
	}

	// Append some entries.
	for i := 0; i < 10; i++ {
		cmd := fmt.Sprintf("the rain in spain %d", i+10)
		if apperr, err := r2.Append(ctx, []byte(cmd)); apperr != nil || err != nil {
			t.Fatalf("append %s failed with %s", cmd, err)
		}
	}
	if !waitForAppliedAgreement(rs, cs, thb) {
		t.Fatalf("no log agreement")
	}

	// Restart the stopped server and wait for it to become a follower.
	restart(t, ctx, rs, cs, r1)
	if !waitForAllRunning(rs, 3*thb) {
		t.Fatalf("server didn't joing after restart")
	}

	if !waitForAppliedAgreement(rs, cs, 3*thb) {
		t.Fatalf("no log agreement")
	}

	vlog.Infof("TestAppend passed")
}
