package maps_test

import (
	"testing"

	"github.com/neutrinocorp/nolan/collection/maps"
)

func TestHashMap(t *testing.T) {
	// var mp maps.Map[string, int] = maps.HashMap[string, int]{}
	var mp maps.Map[string, int] = maps.NewLinkedHashMap[string, int]()
	entries := []maps.Entry[string, int]{
		{
			Key:   "foo",
			Value: 1,
		},
		{
			Key:   "bar",
			Value: 2,
		},
		{
			Key:   "baz",
			Value: 3,
		},
		{
			Key:   "foobar",
			Value: 4,
		},
	}
	mp.PutAllEntries(entries...)
	mp.ForEach(func(a string, b int) bool {
		t.Logf("key: %s, val: %d", a, b)
		return false
	})
	t.Log(mp.Len())
	mp.Remove("bar")
	mp.Remove("baz")
	mp.ForEach(func(a string, b int) bool {
		t.Logf("key: %s, val: %d", a, b)
		return false
	})
	t.Log(mp.Len())
}
