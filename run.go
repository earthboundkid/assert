package assert

import (
	"maps"
	"slices"
	"testing"
)

// Run runs all the test cases in the map with [*testing.Run] using map keys as sub-test names.
// The [TB] associated with the sub-test is [FailsNow] by default.
func Run[Testcase any](t *testing.T, m map[string]Testcase, f func(be TB, tc Testcase)) {
	t.Helper()

	names := slices.Sorted(maps.Keys(m))
	for _, name := range names {
		t.Run(name, func(t *testing.T) { f(FailsNow(t), m[name]) })
	}
}
