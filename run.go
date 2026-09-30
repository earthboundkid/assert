package assert

import (
	"fmt"
	"maps"
	"slices"
	"testing"
)

// Run runs all the test cases in the map with [TB.Run] using map keys as sub-test names.
//
// The [TB] associated with the sub-test is [FailsNow] by default.
//
// Deprecated: Use [RunAll].
//
//go:fix inline
func Run[Testcase any](tb testing.TB, m map[string]Testcase, f func(be TB, tc Testcase)) TB {
	return RunAll(tb, m, f)
}

// RunAll runs all the test cases in the map with [TB.RunAll] using map keys as sub-test names.
//
// The [TB] associated with the sub-test is [FailsNow] by default.
func RunAll[Testcase any](tb testing.TB, m map[string]Testcase, f func(be TB, tc Testcase)) TB {
	tb.Helper()

	be, ok := tb.(TB)
	if !ok {
		be = FailsNow(tb)
	}

	return be.RunAll(m, f)
}

// RunAll runs all the test cases in the map with [TB.Run] using map keys as sub-test names.
func (be TB) RunAll[Testcase any](m map[string]Testcase, f func(be TB, tc Testcase)) TB {
	be.Helper()

	names := slices.Sorted(maps.Keys(m))
	for _, name := range names {
		be.Run(name, func(be TB) { f(be, m[name]) })
	}
	return be
}

// Run calls [*testing.T.Run] or [*testing.B.Run] on the underlying [testing.TB].
// Run panics if the underlying testing.TB is a [*testing.F].
func (be TB) Run(name string, f func(be TB)) TB {
	be.Helper()

	switch tb := be.TB.(type) {
	// Unwrap
	case TB:
		tb.Run(name, f)
	case *testing.T:
		tb.Run(name, func(t *testing.T) {
			f(TB{TB: t, relaxed: be.relaxed})
		})
	case *testing.B:
		tb.Run(name, func(b *testing.B) {
			f(TB{TB: b, relaxed: be.relaxed})
		})
	default:
		panic(fmt.Sprintf("bad underlying TB type: %T", be.TB))
	}
	return be
}
