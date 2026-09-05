package assert

import (
	"maps"
	"slices"
	"testing"
)

// Run runs all the test cases in the map with t.Run using the key as the name.
// The Tester associated with the sub-test is FailNow by default.
func Run[T any](t *testing.T, m map[string]T, f func(Tester, T)) {
	t.Helper()

	names := slices.Sorted(maps.Keys(m))
	for _, name := range names {
		t.Run(name, func(t *testing.T) { f(FailNow(t), m[name]) })
	}
}
