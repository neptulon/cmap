package cmap_test

import (
	"log"

	"github.com/nbusy/cmap"
)

// Example demonstrating the use concurrent-map.
func Example() {
	m := cmap.New()
	m.Set("foo", "bar")

	if val, ok := m.Get("foo"); ok {
		bar := val.(string)
		log.Println(bar)
	}

	m.Delete("foo")
}
