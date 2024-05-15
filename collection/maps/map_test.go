package maps_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/neutrinocorp/nolan/collection/maps"
)

func TestMap_Get(t *testing.T) {
	linkedPopulated := maps.NewLinkedHashMap[string, int]()
	linkedPopulated.Put("foo", 10)
	tests := []struct {
		name  string
		in    maps.Map[string, int]
		exp   int
		expOk bool
	}{
		{
			name:  "hash empty",
			in:    maps.HashMap[string, int]{},
			exp:   0,
			expOk: false,
		},
		{
			name: "hash value",
			in: maps.HashMap[string, int]{
				"foo": 10,
			},
			exp:   10,
			expOk: true,
		},
		{
			name:  "linked empty",
			in:    maps.NewLinkedHashMap[string, int](),
			exp:   0,
			expOk: false,
		},
		{
			name:  "linked value",
			in:    linkedPopulated,
			exp:   10,
			expOk: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out, ok := tt.in.Get("foo")
			assert.Equal(t, tt.exp, out)
			assert.Equal(t, tt.expOk, ok)
		})
	}
}

func TestMap_GetWithFallback(t *testing.T) {
	linkedPopulated := maps.NewLinkedHashMap[string, int]()
	linkedPopulated.Put("foo", 10)
	tests := []struct {
		name string
		in   maps.Map[string, int]
		exp  int
	}{
		{
			name: "hash empty",
			in:   maps.HashMap[string, int]{},
			exp:  -1,
		},
		{
			name: "hash value",
			in: maps.HashMap[string, int]{
				"foo": 10,
			},
			exp: 10,
		},
		{
			name: "linked empty",
			in:   maps.NewLinkedHashMap[string, int](),
			exp:  -1,
		},
		{
			name: "linked value",
			in:   linkedPopulated,
			exp:  10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := tt.in.GetWithFallback("foo", -1)
			assert.Equal(t, tt.exp, out)
		})
	}
}

func TestMap_Put(t *testing.T) {
	linkedPopulated := maps.NewLinkedHashMap[string, int]()
	linkedPopulated.Put("foo", 10)
	tests := []struct {
		name  string
		src   maps.Map[string, int]
		in    maps.Entry[string, int]
		exp   int
		expOk bool
	}{
		{
			name:  "hash empty entry",
			src:   maps.HashMap[string, int]{},
			in:    maps.Entry[string, int]{},
			exp:   0,
			expOk: true,
		},
		{
			name: "hash empty",
			src:  maps.HashMap[string, int]{},
			in: maps.Entry[string, int]{
				Key:   "foo",
				Value: 1,
			},
			exp:   1,
			expOk: true,
		},
		{
			name: "hash value",
			src: maps.HashMap[string, int]{
				"foo": 10,
			},
			in: maps.Entry[string, int]{
				Key:   "foo",
				Value: 1,
			},
			exp:   1,
			expOk: false,
		},

		{
			name:  "linked empty entry",
			src:   maps.NewLinkedHashMap[string, int](),
			in:    maps.Entry[string, int]{},
			exp:   0,
			expOk: true,
		},
		{
			name: "linked empty",
			src:  maps.NewLinkedHashMap[string, int](),
			in: maps.Entry[string, int]{
				Key:   "foo",
				Value: 1,
			},
			exp:   1,
			expOk: true,
		},
		{
			name: "linked value",
			src:  linkedPopulated,
			in: maps.Entry[string, int]{
				Key:   "foo",
				Value: 1,
			},
			exp:   1,
			expOk: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.src.Put(tt.in.Key, tt.in.Value)
			out, _ := tt.src.Get(tt.in.Key)
			assert.Equal(t, tt.exp, out)
		})
	}
}

func TestMap_PutIfAbsent(t *testing.T) {
	linkedPopulated := maps.NewLinkedHashMap[string, int]()
	linkedPopulated.Put("foo", 10)
	tests := []struct {
		name  string
		src   maps.Map[string, int]
		in    maps.Entry[string, int]
		exp   int
		expOk bool
	}{
		{
			name:  "hash empty entry",
			src:   maps.HashMap[string, int]{},
			in:    maps.Entry[string, int]{},
			exp:   0,
			expOk: true,
		},
		{
			name: "hash empty",
			src:  maps.HashMap[string, int]{},
			in: maps.Entry[string, int]{
				Key:   "foo",
				Value: 1,
			},
			exp:   1,
			expOk: true,
		},
		{
			name: "hash value",
			src: maps.HashMap[string, int]{
				"foo": 10,
			},
			in: maps.Entry[string, int]{
				Key:   "foo",
				Value: 1,
			},
			exp:   10,
			expOk: false,
		},

		{
			name:  "linked empty entry",
			src:   maps.NewLinkedHashMap[string, int](),
			in:    maps.Entry[string, int]{},
			exp:   0,
			expOk: true,
		},
		{
			name: "linked empty",
			src:  maps.NewLinkedHashMap[string, int](),
			in: maps.Entry[string, int]{
				Key:   "foo",
				Value: 1,
			},
			exp:   1,
			expOk: true,
		},
		{
			name: "linked value",
			src:  linkedPopulated,
			in: maps.Entry[string, int]{
				Key:   "foo",
				Value: 1,
			},
			exp:   10,
			expOk: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ok := tt.src.PutIfAbsent(tt.in.Key, tt.in.Value)
			assert.Equal(t, tt.expOk, ok)
			out, _ := tt.src.Get(tt.in.Key)
			assert.Equal(t, tt.exp, out)
		})
	}
}

func TestMap_PutAllEntries(t *testing.T) {
	linkedPopulated := maps.NewLinkedHashMap[string, int]()
	linkedPopulated.Put("foo", 10)
	tests := []struct {
		name string
		src  maps.Map[string, int]
		in   []maps.Entry[string, int]
		exp  int
	}{
		{
			name: "hash empty entry",
			src:  maps.HashMap[string, int]{},
			in:   []maps.Entry[string, int]{},
			exp:  0,
		},
		{
			name: "hash empty",
			src:  maps.HashMap[string, int]{},
			in: []maps.Entry[string, int]{
				{
					Key:   "foo",
					Value: 1,
				},
			},
			exp: 1,
		},
		{
			name: "hash value",
			src: maps.HashMap[string, int]{
				"foo": 10,
			},
			in: []maps.Entry[string, int]{
				{
					Key:   "foo",
					Value: 1,
				},
			},
			exp: 1,
		},
		{
			name: "linked empty entry",
			src:  maps.NewLinkedHashMap[string, int](),
			in:   []maps.Entry[string, int]{},
			exp:  0,
		},
		{
			name: "linked empty",
			src:  maps.NewLinkedHashMap[string, int](),
			in: []maps.Entry[string, int]{
				{
					Key:   "foo",
					Value: 1,
				},
			},
			exp: 1,
		},
		{
			name: "linked value",
			src:  linkedPopulated,
			in: []maps.Entry[string, int]{
				{
					Key:   "foo",
					Value: 1,
				},
			},
			exp: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.src.PutAllEntries(tt.in...)
			for _, entry := range tt.in {
				val, _ := tt.src.Get(entry.Key)
				assert.Equal(t, tt.exp, val)
			}
		})
	}
}

func TestLinkedHashMap_Ordering(t *testing.T) {
	var mp maps.Map[string, int] = maps.NewLinkedHashMap[string, int]()
	keys := []string{"foo", "bar", "baz", "foobar"}
	values := []int{1, 2, 3, 4}

	require.Equal(t, len(keys), len(values))
	entries := make([]maps.Entry[string, int], 0, len(keys))
	for i, key := range keys {
		entries = append(entries, maps.Entry[string, int]{
			Key:   key,
			Value: values[i],
		})
	}
	mp.PutAllEntries(entries...)
	iterCount := 0
	mp.ForEach(func(a string, b int) bool {
		defer func() {
			iterCount++
		}()
		assert.Equal(t, keys[iterCount], a)
		assert.Equal(t, values[iterCount], b)
		return false
	})
	assert.Equal(t, len(keys), iterCount)
	assert.Equal(t, keys, mp.KeysSlice())
	assert.Equal(t, values, mp.ValuesSlice())
	assert.Equal(t, len(keys), mp.Len())
	mp.Remove("bar")
	mp.Remove("baz")
	assert.Equal(t, len(keys)-2, mp.Len())
	keys = []string{"foo", "foobar"}
	values = []int{1, 4}
	assert.Equal(t, keys, mp.KeysSlice())
	assert.Equal(t, values, mp.ValuesSlice())
}
