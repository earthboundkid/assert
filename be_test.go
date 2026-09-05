package assert_test

import (
	"errors"
	"maps"
	"strings"
	"testing"
	"time"

	"github.com/earthboundkid/assert"
)

func Test(t *testing.T) {
	okayTests := []func(be assert.Tester){
		func(be assert.Tester) { be.Zero(time.Time{}.Local()) },
		func(be assert.Tester) { be.Zero([]string(nil)) },
		func(be assert.Tester) { be.NotZero([]string{""}) },
		func(be assert.Tester) { _ = be.OK(func() (int, error) { return 1, nil }()) },
		func(be assert.Tester) { be.True(true) },
		func(be assert.Tester) { be.False(false) },
		func(be assert.Tester) { be.EqualLength(map[int]int{}, 0) },
		func(be assert.Tester) { be.EqualLength(map[int]int{1: 1}, 1) },
		func(be assert.Tester) {
			ch := make(chan int, 1)
			be.EqualLength(ch, 0)
		},
		func(be assert.Tester) {
			ch := make(chan int, 1)
			ch <- 1
			be.EqualLength(ch, 1)
		},
		func(be assert.Tester) {
			seq2 := maps.All(map[int]int{1: 1})
			be.EqualLength(seq2, 1)
		},
	}

	for _, test := range okayTests {
		var buf strings.Builder
		m := &mockingT{w: &buf}
		test(assert.Continue(m))
		if m.Failed() {
			t.Fatal("failed too soon")
		}
		if buf.String() != "" {
			t.Fatal("wrote too much")
		}
	}

	badTests := []func(be assert.Tester){
		func(be assert.Tester) { be.SlicesEqual([]string{}, []string{""}) },
		func(be assert.Tester) { be.NotZero(time.Time{}.Local()) },
		func(be assert.Tester) { be.Zero([]string{""}) },
		func(be assert.Tester) { be.NotZero([]string(nil)) },
		func(be assert.Tester) {
			be.OK(func() (int, error) { return 0, errors.New("") }())
		},
		func(be assert.Tester) { be.True(false) },
		func(be assert.Tester) { be.False(true) },
		func(be assert.Tester) {
			seq2 := maps.All(map[int]int{1: 1})
			be.EqualLength(seq2, 0)
		},
		func(be assert.Tester) {
			ch := make(chan int, 1)
			be.EqualLength(ch, 1)
		},
		func(be assert.Tester) {
			ch := make(chan int, 1)
			ch <- 1
			be.EqualLength(ch, 0)
		},
		func(be assert.Tester) {
			ch := make(chan int, 1)
			close(ch)
			be.EqualLength(ch, 1)
		},
		func(be assert.Tester) {
			be.Panicked(func() {})
		},
	}

	for _, test := range badTests {
		var buf strings.Builder
		m := &mockingT{w: &buf}
		test(assert.FailNow(m))
		if !m.hasFailedNow {
			t.Fatal("did not fail")
		}
		if buf.String() == "" {
			t.Fatal("wrote too little")
		}
	}
}

func TestTester_TB(t *testing.T) {
	var buf strings.Builder
	m := &mockingT{w: &buf}
	be := assert.Continue(m)
	be.TB().Logf("hi")
	assert.FailNow(t).Equal(buf.String(), "hi\n")
}
