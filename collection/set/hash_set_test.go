package set_test

import (
	"testing"

	"github.com/neutrinocorp/nolan/collection/set"
)

func TestHashSet(t *testing.T) {
	st := set.HashSet[int]{}
	st.AddSlice(1, 2, 3, 4, 5)
	t.Log(st.Contains(1))
	t.Log(st.Contains(6))
	t.Log(st.ContainsSlice(1, 2, 3, 4, 5))
	t.Log(st.ContainsSlice(1, 2, 3, 4, 5, 6))
	st.Add(6)
	t.Log(st.ContainsSlice(1, 2, 3, 4, 5, 6))
}
