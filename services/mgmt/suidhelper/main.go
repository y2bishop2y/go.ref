package main

// suidhelper should be installed setuid root. Having done this, it will
// run the provided command as the specified user identity.
// suidhelper deliberately attempts to be as simple as possible to
// simplify reviewing it for security concerns.

import (
	"flag"
	"os"

	"veyron.io/veyron/veyron/services/mgmt/suidhelper/impl"
)

func main() {
	flag.Parse()
	if err := impl.Run(os.Environ()); err != nil {
		flag.Usage()
	}
}
