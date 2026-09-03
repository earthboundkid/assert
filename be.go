package assert

import (
	"reflect"
	"slices"
	"testing"
)

type Tester struct {
	tb      testing.TB
	relaxed bool
}

func Strict(t testing.TB) Tester {
	return Tester{t, false}
}

func Relaxed(t testing.TB) Tester {
	return Tester{t, true}
}

func (be Tester) fatalf(format string, args ...any) {
	be.tb.Helper()
	be.tb.Logf(format, args...)
	if be.relaxed {
		be.tb.Fail()
	} else {
		be.tb.FailNow()
	}
}

// Equal calls t.Fatalf if want != got.
func (be Tester) Equal[T comparable](got, want T) Tester {
	be.tb.Helper()
	if want != got {
		be.fatalf("want: %v; got: %v", want, got)
	}
	return be
}

// Unequal calls t.Fatalf if got == bad.
func (be Tester) Unequal[T comparable](got, bad T) Tester {
	be.tb.Helper()
	if got == bad {
		be.fatalf("got: %v", got)
	}
	return be
}

// AllEqual calls t.Fatalf if want != got.
func (be Tester) AllEqual[T comparable](got, want []T) Tester {
	be.tb.Helper()
	if !slices.Equal(got, want) {
		be.fatalf("got: %v; want: %v", got, want)
	}
	return be
}

// Zero calls t.Fatalf if value != the zero value for T.
func (be Tester) Zero[T any](value T) Tester {
	be.tb.Helper()
	if truthy(value) {
		be.fatalf("got: %v", value)
	}
	return be
}

// Nonzero calls t.Fatalf if value == the zero value for T.
func (be Tester) Nonzero[T any](value T) Tester {
	be.tb.Helper()
	if !truthy(value) {
		be.fatalf("got: %v", value)
	}
	return be
}

func truthy[T any](v T) bool {
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

func (be Tester) OK[T any](value T, err error) T {
	be.tb.Helper()
	if err != nil {
		be.fatalf("err != nil: %v", err)
		return *new(T)
	}
	return value
}

func (be Tester) True(value bool) Tester {
	be.tb.Helper()
	if !value {
		be.fatalf("got: false")
	}
	return be
}

func (be Tester) False(value bool) Tester {
	be.tb.Helper()
	if value {
		be.fatalf("got: true")
	}
	return be
}
