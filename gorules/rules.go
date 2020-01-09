package gorules

import "github.com/quasilyte/go-ruleguard/dsl/fluent"

func _(m fluent.Matcher) {
	m.Import("github.com/bradfitz/gomemcache/memcache")
	m.Match("memcache.New").Report("usage of raw memcached client")
}
