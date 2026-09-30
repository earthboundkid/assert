package assert_test

import (
	"strings"
	"testing"

	"github.com/earthboundkid/assert"
)

func ExampleRunAll() {
	// TestCapitalize
	_ = func(t *testing.T) {
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
}

func ExampleTB_Run() {
	// TestCapitalize
	_ = func(t *testing.T) {
		assert.FailsNow(t).
			Run("empty case", func(be assert.TB) {
				be.Equal(strings.ToUpper(""), "")
			}).
			Run("mixed case", func(be assert.TB) {
				be.Equal(strings.ToUpper("aBc"), "ABC")
			})
	}
}
