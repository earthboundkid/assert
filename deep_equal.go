package assert

import (
	"reflect"
)

// DeepEqual calls t.Fatalf if want and got are different according to reflect.DeepEqual.
func (be Tester) DeepEqual[T any](got, want T) Tester {
	be.tb.Helper()
	// Pass as pointers to get around the nil-interface problem
	if !reflect.DeepEqual(&want, &got) {
		be.fatalf("reflect.DeepEqual(%#v, %#v) == false", want, got)
	}
	return be
}
