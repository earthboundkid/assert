package assert

import (
	"reflect"
	"slices"
)

// EqualLength asserts that seq has a length that is exactly want.
//
// The type of seq must be array, array pointer, slice, map, string, channel, [iter.Seq], or [iter.Seq2].
// For channels and iterators, the values are consumed to get the sequence length.
func (be TB) EqualLength(seq any, want int) TB {
	be.Helper()
	getLen(be, seq, want, true)
	return be
}

func getLen(be TB, seq any, want int, exactly bool) {
	rv := reflect.ValueOf(seq)
	rt := rv.Type()
	kind := rt.Kind()
	got, atLeast := 0, "at least "

	switch {
	case slices.Contains([]reflect.Kind{
		reflect.Array, reflect.Slice, reflect.Map, reflect.String,
	}, kind) ||
		(kind == reflect.Pointer && rt.Elem().Kind() == reflect.Array):
		atLeast = ""
		got = rv.Len()
	case kind == reflect.Chan:
		for {
			chosen, _, recvOK := reflect.Select([]reflect.SelectCase{
				{Dir: reflect.SelectRecv, Chan: rv},
				{Dir: reflect.SelectDefault},
			})

			if chosen == 1 || !recvOK { // default case or channel closed
				break
			}
			if got++; got > want {
				break
			}
		}
	case kind == reflect.Func && rt.CanSeq():
		for range rv.Seq() {
			if got++; got > want {
				break
			}
		}
	case kind == reflect.Func && rt.CanSeq2():
		for range rv.Seq2() {
			if got++; got > want {
				break
			}
		}
	default:
		panic("seq must be a non-integer rangeable type")
	}

	switch {
	case !exactly && want > got:
		be.fatalf("want len(seq) >= %d; got %d", want, got)
	case exactly && want > got:
		be.fatalf("want len(seq) == %d; got %d", want, got)
	case exactly && want < got:
		be.fatalf("want len(seq) == %d; got %s%d", want, atLeast, got)
	}
}

// AtLeastLength asserts that seq has a length that is at least want.
//
// The type of seq must be array, array pointer, slice, map, string, channel, [iter.Seq], or [iter.Seq2].
// For channels and iterators, the values are consumed to get the sequence length.
func (be TB) AtLeastLength(seq any, want int) TB {
	be.Helper()
	getLen(be, seq, want, false)
	return be
}
