package assert_test

import (
	"strings"
	"testing"

	"github.com/earthboundkid/assert"
)

func TestRunAll(t *testing.T) {
	// TestCapitalize
	// Keep in-sync with run_example_test.go
	type testcase struct {
		in, want string
	}
	assert.RunAll(t, map[string]testcase{
		"blank":            {in: "", want: ""},
		"a":                {in: "a", want: "A"},
		"already upper":    {in: "A", want: "A"},
		"multi character":  {in: "Abc", want: "ABC"},
		"other characters": {in: " a,.c", want: " A,.C"},
	}, func(be assert.TB, tc testcase) {
		be.Equal(strings.ToUpper(tc.in), tc.want)
	})
}

func TestRun(t *testing.T) {
	// Keep in-sync with run_example_test.go
	assert.
		Run(t, "empty case", func(be assert.TB) {
			be.Equal(strings.ToUpper(""), "")
		}).
		Run("mixed case", func(be assert.TB) {
			be.Equal(strings.ToUpper("aBc"), "ABC")
		})
}

func TestTB_Run(t *testing.T) {
	// Test sub-sub-tests
	assert.FailsNow(t).
		Run("test", func(be assert.TB) {
			assert.FailsNow(be).Run("subtest", func(be assert.TB) {
				be.Run("sub-subtest", func(be assert.TB) {
					be.Equal(1, 1)
				})
			})
		})
}

func BenchmarkTB_Run(b *testing.B) {
	assert.FailsNow(b).
		Run("test", func(be assert.TB) {
			for be.TB.(*testing.B).Loop() {
				be.Equal(1, 1)
			}
		})
}
