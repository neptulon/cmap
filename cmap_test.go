package cmap

import "testing"

func TestMap(t *testing.T) {
	m := New()
	if m == nil {
		t.Fatal("newly created map is nil")
	}

	if m.Len() != 0 {
		t.Fatal("newly created map is not empty")
	}

	m.Set("foo", "bar")

	if val, ok := m.Get("foo"); ok {
		if bar := val.(string); bar != "bar" {
			t.Fatal("retrieved value is not 'bar'")
		}
	} else {
		t.Fatal("cannot value using key 'foo' in map")
	}

	m.Delete("foo")

	if m.Len() != 0 {
		t.Fatal("failed to delte 'foo' from map")
	}
}
