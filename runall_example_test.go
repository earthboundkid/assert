package assert_test

import (
	"strings"
	"testing"

	"github.com/earthboundkid/assert"
)

func ExampleRun() {
	// TestCapitalize
	_ = func(t *testing.T) {
		type testcase struct {
			in, want string
		}
		assert.Run(t, map[string]testcase{
			"blank":            {in: "", want: ""},
			"a":                {in: "a", want: "A"},
			"already upper":    {in: "A", want: "A"},
			"multi character":  {in: "Abc", want: "ABC"},
			"other characters": {in: " a,.c", want: " A,.C"},
		}, func(be assert.Tester, tc testcase) {
			be.Equal(strings.ToUpper(tc.in), tc.want)
		})
	}
}
