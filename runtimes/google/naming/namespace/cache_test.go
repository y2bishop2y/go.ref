package namespace

import (
	"fmt"
	"testing"
	"time"

	"v.io/v23/naming"
)

func compatible(server string, servers []naming.MountedServer) bool {
	if len(servers) == 0 {
		return server == ""
	}
	return servers[0].Server == server
}

func future(secs uint32) time.Time {
	return time.Now().Add(time.Duration(secs) * time.Second)
}

// TestCache tests the cache directly rather than via the namespace methods.
func TestCache(t *testing.T) {
	preload := []struct {
		name   string
		suffix string
		server string
	}{
		{"/h1//a/b/c/d", "c/d", "/h2"},
		{"/h2//c/d", "d", "/h3"},
		{"/h3//d", "", "/h4:1234"},
	}
	c := newTTLCache()
	for _, p := range preload {
		e := &naming.MountEntry{Name: p.suffix, Servers: []naming.MountedServer{naming.MountedServer{Server: p.server, Expires: future(30)}}}
		c.remember(p.name, e)
	}

	tests := []struct {
		name    string
		suffix  string
		server  string
		succeed bool
	}{
		{"/h1//a/b/c/d", "c/d", "/h2", true},
		{"/h2//c/d", "d", "/h3", true},
		{"/h3//d", "", "/h4:1234", true},
		{"/notintcache", "", "", false},
		{"/h1//a/b/f//g", "f/g", "/h2", true},
		{"/h3//d//e", "e", "/h4:1234", true},
	}
	for _, p := range tests {
		e, err := c.lookup(p.name)
		if (err == nil) != p.succeed {
			t.Errorf("%s: lookup failed", p.name)
		}
		if e.Name != p.suffix || !compatible(p.server, e.Servers) {
			t.Errorf("%s: got %v, %s not %s, %s", p.name, e.Name, e.Servers, p.suffix, p.server)
		}
	}
}

func TestCacheLimit(t *testing.T) {
	c := newTTLCache().(*ttlCache)
	e := &naming.MountEntry{Servers: []naming.MountedServer{naming.MountedServer{Server: "the rain in spain", Expires: future(3000)}}}
	for i := 0; i < maxCacheEntries; i++ {
		c.remember(fmt.Sprintf("%d", i), e)
		if len(c.entries) > maxCacheEntries {
			t.Errorf("unexpected cache size: got %d not %d", len(c.entries), maxCacheEntries)
		}
	}
	// Adding one more element should reduce us to 3/4 full.
	c.remember(fmt.Sprintf("%d", maxCacheEntries), e)
	if len(c.entries) != cacheHisteresisSize {
		t.Errorf("cache shrunk wrong amount: got %d not %d", len(c.entries), cacheHisteresisSize)
	}
}

func TestCacheTTL(t *testing.T) {
	before := time.Now()
	c := newTTLCache().(*ttlCache)
	// Fill cache.
	e := &naming.MountEntry{Servers: []naming.MountedServer{naming.MountedServer{Server: "the rain in spain", Expires: future(3000)}}}
	for i := 0; i < maxCacheEntries; i++ {
		c.remember(fmt.Sprintf("%d", i), e)
	}
	// Time out half the entries.
	i := len(c.entries) / 2
	for k := range c.entries {
		c.entries[k].Servers[0].Expires = before
		if i == 0 {
			break
		}
		i--
	}
	// Add an entry and make sure we now have room.
	c.remember(fmt.Sprintf("%d", maxCacheEntries+2), e)
	if len(c.entries) > cacheHisteresisSize {
		t.Errorf("entries did not timeout: got %d not %d", len(c.entries), cacheHisteresisSize)
	}
}

func TestFlushCacheEntry(t *testing.T) {
	preload := []struct {
		name   string
		server string
	}{
		{"/h1//a/b", "/h2//c"},
		{"/h2//c", "/h3"},
		{"/h3//d", "/h4:1234"},
	}
	ns, _ := New()
	c := ns.resolutionCache.(*ttlCache)
	for _, p := range preload {
		e := &naming.MountEntry{Servers: []naming.MountedServer{naming.MountedServer{Server: "p.server", Expires: future(3000)}}}
		c.remember(p.name, e)
	}
	toflush := "/h1/xyzzy"
	if ns.FlushCacheEntry(toflush) {
		t.Errorf("%s should not have caused anything to flush", toflush)
	}
	toflush = "/h1/a/b/d/e"
	if !ns.FlushCacheEntry(toflush) {
		t.Errorf("%s should have caused something to flush", toflush)
	}
	name := preload[2].name
	if _, err := c.lookup(name); err != nil {
		t.Errorf("%s should not have been flushed", name)
	}
	if len(c.entries) != 2 {
		t.Errorf("%s flushed too many entries", toflush)
	}
	toflush = preload[1].name
	if !ns.FlushCacheEntry(toflush) {
		t.Errorf("%s should have caused something to flush", toflush)
	}
	if _, ok := c.entries[toflush]; ok {
		t.Errorf("%s should have been flushed", name)
	}
	if len(c.entries) != 1 {
		t.Errorf("%s flushed too many entries", toflush)
	}
}

func disabled(ctls []naming.CacheCtl) bool {
	for _, c := range ctls {
		if v, ok := c.(naming.DisableCache); ok && bool(v) {
			return true
		}
	}
	return false
}

func TestCacheDisableEnable(t *testing.T) {
	ns, _ := New()

	// Default should be working resolution cache.
	name := "/h1//a"
	serverName := "/h2//"
	c := ns.resolutionCache.(*ttlCache)
	e := &naming.MountEntry{Servers: []naming.MountedServer{naming.MountedServer{Server: serverName, Expires: future(3000)}}}
	c.remember(name, e)
	if ne, err := c.lookup(name); err != nil || ne.Servers[0].Server != serverName {
		t.Errorf("should have found the server in the cache")
	}

	// Turn off the resolution cache.
	ctls := ns.CacheCtl(naming.DisableCache(true))
	if !disabled(ctls) {
		t.Errorf("caching not disabled")
	}
	nc := ns.resolutionCache.(nullCache)
	nc.remember(name, e)
	if _, err := nc.lookup(name); err == nil {
		t.Errorf("should not have found the server in the cache")
	}

	// Turn on the resolution cache.
	ctls = ns.CacheCtl(naming.DisableCache(false))
	if disabled(ctls) {
		t.Errorf("caching disabled")
	}
	c = ns.resolutionCache.(*ttlCache)
	c.remember(name, e)
	if ne, err := c.lookup(name); err != nil || ne.Servers[0].Server != serverName {
		t.Errorf("should have found the server in the cache")
	}
}
