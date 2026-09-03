package assert_test

import (
	"testing"

	"github.com/earthboundkid/assert"
)

func TestLen(t *testing.T) {
	be := assert.Strict(t)
	// Make sure integers aren't treated as rangeable
	be.Nonzero(assert.Panicked(func() {
		be.EqualLength(0, 0)
	}))
}

func TestMatch(t *testing.T) {
	be := assert.Strict(t)
	// Make sure bad regexp patterns panic
	pval := assert.Panicked(func() {
		be.Match("", `\`)
	})
	be.Nonzero(pval)
	s, ok := pval.(string)
	be.
		True(ok).
		Match(s, `^regexp: Compile\(`)
}
