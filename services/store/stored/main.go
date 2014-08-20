// stored is a storage server.
//
// Usage:
//
//     stored [--name=<mount>] [--db=<dbName>]
//
//     - <name> is the Veyron mount point name, default /global/vstore/<hostname>/<username>.
//     - <dbName> is the filename in which to store the data.
//
// The store service has Object name, <name>/.store.
// The raw store service has Object name, <name>/.store.raw.
// Individual values with path <path> have name <name>/<path>.
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/user"

	vflag "veyron/security/flag"
	"veyron/services/store/server"
	"veyron/services/store/viewer"

	"veyron2/rt"
	"veyron2/storage/vstore"

	_ "veyron/services/store/typeregistryhack"
)

var (
	mountName string

	// TODO(rthellend): Remove the protocol and address flags when the config
	// manager is working.
	protocol = flag.String("protocol", "tcp", "protocol to listen on")
	address  = flag.String("address", ":0", "address to listen on")

	dbName = flag.String("db", "/var/tmp/veyron_store.db",
		"Metadata database")
	viewerPort = flag.Int("viewerPort", 5000,
		"IPV4 port to serve viewer from, or 0 to disable viewer")
)

func init() {
	username := "unknown"
	if u, err := user.Current(); err == nil {
		username = u.Username
	}
	hostname := "unknown"
	if h, err := os.Hostname(); err == nil {
		hostname = h
	}
	dir := "global/vstore/" + hostname + "/" + username
	flag.StringVar(&mountName, "name", dir, "Mount point for media")
}

// main starts the store service, taking args from command line flags.
func main() {
	r := rt.Init()

	// Create a new server instance.
	s, err := r.NewServer()
	if err != nil {
		log.Fatal("r.NewServer() failed: ", err)
	}

	// Create a new StoreService.
	storeService, err := server.New(
		server.ServerConfig{Admin: r.Identity().PublicID(), DBName: *dbName})
	if err != nil {
		log.Fatal("server.New() failed: ", err)
	}
	defer storeService.Close()

	// Create the authorizer.
	auth := vflag.NewAuthorizerOrDie()

	// Register the services.
	storeDisp := server.NewStoreDispatcher(storeService, auth)
	// Create an endpoint and start listening.
	ep, err := s.Listen(*protocol, *address)
	if err != nil {
		log.Fatal("s.Listen() failed: ", err)
	}
	// Publish the service in the mount table.
	log.Printf("Mounting store on %s, endpoint /%s", mountName, ep)
	if err := s.Serve(mountName, storeDisp); err != nil {
		log.Fatal("s.Serve() failed: ", err)
	}

	// Run viewer if requested.
	if *viewerPort > 0 {
		go viewer.ListenAndServe(fmt.Sprintf(":%d", *viewerPort), mountName, vstore.New())
	}

	// Wait forever.
	done := make(chan struct{})
	<-done
}
