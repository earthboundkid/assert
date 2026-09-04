package assert

import (
	"reflect"
)

// DeepEqual calls t.Fatalf if want and got are different according to reflect.DeepEqual.
func (be Tester) DeepEqual[T any](got, want T) Tester {
	be.tb.Helper()
	// Pass as pointers to get around the nil-interface problem
	if !reflect.DeepEqual(&got, &want) {
		be.fatalf("reflect.DeepEqual(%#v, %#v) == false", got, want)
	}
	return be
}
