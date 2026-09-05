package assert

import (
	"reflect"
)

// DeepEqual asserts got is [reflect.DeepEqual] to want.
//
// Prefer to use [TB.SlicesEqual] if possible.
func (be TB) DeepEqual[T any](got, want T) TB {
	be.Helper()
	// Pass as pointers to get around the nil-interface problem
	if !reflect.DeepEqual(&got, &want) {
		be.fatalf("reflect.DeepEqual(%#v, %#v) == false", got, want)
	}
	return be
}
