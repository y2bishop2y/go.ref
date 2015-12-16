// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package server

// This file defines methods for consistently destroying unreferenced database
// stores by using a garbage collection log keeping dbInfos for inactive (not
// fully created or deleted) databases.

import (
	"fmt"

	"v.io/v23/context"
	"v.io/v23/verror"
	"v.io/v23/vom"
	"v.io/x/lib/vlog"
	"v.io/x/ref/services/syncbase/server/util"
	"v.io/x/ref/services/syncbase/store"
)

func dbGCKey(path string) string {
	return util.JoinKeyParts(util.DbGCPrefix, path)
}

// putDbGCEntry puts a dbInfo into the garbage collection log. It should be
// used before creating a new database and when marking a database for
// destruction (deleting the dbInfo from active databases).
func putDbGCEntry(ctx *context.T, stw store.StoreWriter, dbInfo *DbInfo) error {
	return util.Put(ctx, stw, dbGCKey(dbInfo.RootDir), dbInfo)
}

// delDbGCEntry removes a dbInfo from the garbage collection log. It should
// be used after successfully destroying the database and when finalizing
// database creation (putting the dbInfo into active databases).
func delDbGCEntry(ctx *context.T, stw store.StoreWriter, dbInfo *DbInfo) error {
	return util.Delete(ctx, stw, dbGCKey(dbInfo.RootDir))
}

// deleteDatabaseEntry marks a database for destruction by moving its dbInfo
// record from active databases into the garbage collection log. It returns
// the dbInfo that can be passed to finalizeDatabaseDestroy.
func deleteDatabaseEntry(ctx *context.T, tx store.Transaction, a *app, dbName string) (*DbInfo, error) {
	dbInfo, err := a.getDbInfo(ctx, tx, dbName)
	if err != nil {
		return nil, err
	}
	if err := putDbGCEntry(ctx, tx, dbInfo); err != nil {
		return nil, err
	}
	if err := a.delDbInfo(ctx, tx, dbName); err != nil {
		return nil, err
	}
	return dbInfo, nil
}

// finalizeDatabaseDestroy attempts to close (if stRef is not nil) and destroy
// the database store. If successful, it removes the dbInfo record from the
// garbage collection log.
func finalizeDatabaseDestroy(ctx *context.T, stw store.StoreWriter, dbInfo *DbInfo, stRef store.Store) error {
	vlog.VI(2).Infof("server: app: destroying store at %q for database %s (closing: %v)", dbInfo.RootDir, dbInfo.Name, stRef != nil)
	if stRef != nil {
		// TODO(ivanpi): Safer to crash syncbased on Close() failure? Otherwise,
		// already running calls might continue using the database. Alternatively,
		// explicitly cancel running calls in Destroy.
		if err := stRef.Close(); err != nil {
			return wrapGCError(dbInfo, err)
		}
	}
	if err := util.DestroyStore(dbInfo.Engine, dbInfo.RootDir); err != nil {
		return wrapGCError(dbInfo, err)
	}
	if err := delDbGCEntry(ctx, stw, dbInfo); err != nil {
		return wrapGCError(dbInfo, err)
	}
	return nil
}

// runGCInactiveDatabases iterates over the garbage collection log, attempting
// to destroy each database store, removing dbInfo records from the log when
// destruction is successful.
// NOTE: runGCInactiveDatabases should not be run in parallel with database
// creation or deletion since it can cause spurious failures (e.g. garbage
// collecting a database that is being created).
// TODO(ivanpi): Consider adding mutex to allow running GC concurrently.
func runGCInactiveDatabases(ctx *context.T, st store.Store) error {
	vlog.VI(2).Infof("server: app: starting garbage collection of inactive databases")
	total := 0
	deleted := 0
	gcIt := st.Scan(util.ScanPrefixArgs(util.DbGCPrefix, ""))
	var diBytes []byte
	for gcIt.Advance() {
		diBytes = gcIt.Value(diBytes)
		var dbInfo DbInfo
		if err := vom.Decode(diBytes, &dbInfo); err != nil {
			verror.New(verror.ErrInternal, ctx, err)
		}
		total++
		if err := finalizeDatabaseDestroy(ctx, st, &dbInfo, nil); err != nil {
			vlog.Error(err)
		} else {
			deleted++
		}
	}
	if err := gcIt.Err(); err != nil {
		return verror.New(verror.ErrInternal, ctx, err)
	}
	vlog.VI(2).Infof("server: app: garbage collected %d out of %d inactive databases", deleted, total)
	return nil
}

func wrapGCError(dbInfo *DbInfo, err error) error {
	return fmt.Errorf("failed to destroy store at %q for database %s: %v", dbInfo.RootDir, dbInfo.Name, err)
}