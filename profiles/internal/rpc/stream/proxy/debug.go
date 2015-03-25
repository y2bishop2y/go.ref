// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package proxy

import (
	"bytes"
	"fmt"
)

// DebugString dumps out the routing table at the proxy in text format.
// The format is meant for debugging purposes and may change without notice.
func (p *Proxy) debugString() string {
	var buf bytes.Buffer
	servers := p.servers.List()
	p.mu.RLock()
	defer p.mu.RUnlock()
	fmt.Fprintf(&buf, "Proxy with endpoint: %q. #Processes:%d #Servers:%d\n", p.endpoint(), len(p.processes), len(servers))
	fmt.Fprintf(&buf, "=========\n")
	fmt.Fprintf(&buf, "PROCESSES\n")
	fmt.Fprintf(&buf, "=========\n")
	index := 1
	for process, _ := range p.processes {
		fmt.Fprintf(&buf, "(%d) - %v", index, process)
		index++
		process.mu.RLock()
		fmt.Fprintf(&buf, " NextVCI:%d #Severs:%d\n", process.nextVCI, len(process.servers))
		for vci, d := range process.routingTable {
			fmt.Fprintf(&buf, "    VCI %4d --> VCI %4d @ %s\n", vci, d.VCI, d.Process)
		}
		process.mu.RUnlock()
	}
	fmt.Fprintf(&buf, "=======\n")
	fmt.Fprintf(&buf, "SERVERS\n")
	fmt.Fprintf(&buf, "=======\n")
	for ix, is := range servers {
		fmt.Fprintf(&buf, "(%d) %v\n", ix+1, is)
	}
	return buf.String()
}
