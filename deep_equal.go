package assert

import (
	"reflect"
)

// DeepEqual asserts got is [reflect.DeepEqual] to want.
//
// Prefer to use [SlicesEqual] if possible.
func (be Tester) DeepEqual[T any](got, want T) Tester {
	be.tb.Helper()
	// Pass as pointers to get around the nil-interface problem
	if !reflect.DeepEqual(&got, &want) {
		be.fatalf("reflect.DeepEqual(%#v, %#v) == false", got, want)
	}
	return be
}
