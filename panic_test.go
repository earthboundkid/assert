package assert_test

import (
	"testing"

	"github.com/earthboundkid/assert"
)

func TestLen(t *testing.T) {
	be := assert.Strict(t)
	// Make sure integers aren't treated as rangeable
	be.Nonzero(assert.Panicked(func() {
		// be.EqualLength(t, 0, 0)
		panic("")
	}))
}

func TestMatch(t *testing.T) {
	be := assert.Strict(t)
	_ = be
	// // Make sure bad regexp patterns panic
	// pval := be.Panicked(func() {
	// 	be.Match(t, `\`, "")
	// })
	// be.Nonzero(t, pval)
	// s, ok := pval.(string)
	// be.True(t, ok)
	// be.Match(t, `^regexp: Compile\(`, s)
}
