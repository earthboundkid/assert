package assert_test

import (
	"github.com/earthboundkid/assert"
)

func ExampleTester_Panicked() {
	be := assert.Continue(&mockingT{})

	divide := func(num, denom int) int {
		return num / denom
	}

	// Test that division by zero panics
	be.Panicked(func() {
		divide(1, 0)
	})

	// Because a panic fails a test by default,
	// testing that an operation does not panic is less necessary,
	// but may be helpful in a table test.
	for _, testcase := range []struct {
		num, denom, want int
		shouldPanic      bool
	}{
		{0, 1, 0, false},
		{1, 1, 1, false},
		{1, 0, 0xbadc0ffee, true},
		{0, 0, 0xbadc0ffee, true},
	} {
		got := 0xbadc0ffee
		panicVal := assert.Catch(func() {
			got = divide(testcase.num, testcase.denom)
		})
		be.Equal(got, testcase.want)
		be.Equal(panicVal != nil, testcase.shouldPanic)
	}
	// Output:
}
