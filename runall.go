package assert

import (
	"maps"
	"slices"
	"testing"
)

func RunAll[T any](t *testing.T, m map[string]T, fn func(Tester, T)) {
	t.Helper()

	names := slices.Sorted(maps.Keys(m))
	for _, name := range names {
		t.Run(name, func(t *testing.T) { fn(FailNow(t), m[name]) })
	}
}
