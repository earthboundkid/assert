package assert

import "reflect"

// Zero asserts value is not [Truthy].
func (be TB) Zero[T any](value T) TB {
	be.Helper()
	if Truthy(value) {
		be.fatalf("got: %v", value)
	}
	return be
}

// Truthy asserts value is [Truthy].
func (be TB) Truthy[T any](value T) TB {
	be.Helper()
	if !Truthy(value) {
		be.fatalf("got: %v", value)
	}
	return be
}

// Truthy returns
//
//   - !v.IsZero(), for types with an IsZero() method.
//   - len(v) != 0, for slices and maps.
//   - v != the zero value of T, for all other types.
func Truthy[T any](v T) bool {
	switch m := any(v).(type) {
	case interface{ IsZero() bool }:
		return !m.IsZero()
	}
	return reflectValue(&v)
}

func reflectValue(vp any) bool {
	switch rv := reflect.ValueOf(vp).Elem(); rv.Kind() {
	case reflect.Map, reflect.Slice:
		return rv.Len() != 0
	default:
		return !rv.IsZero()
	}
}
