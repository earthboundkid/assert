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
	okayTests := []func(be assert.TB){
		func(be assert.TB) { be.Falsey(time.Time{}.Local()) },
		func(be assert.TB) { be.Falsey([]string(nil)) },
		func(be assert.TB) { be.Truthy([]string{""}) },
		func(be assert.TB) { _ = be.OK(func() (int, error) { return 1, nil }()) },
		func(be assert.TB) { _, _ = be.OK2(func() (int, string, error) { return 1, "x", nil }()) },
		func(be assert.TB) { be.True(true) },
		func(be assert.TB) { be.False(false) },
		func(be assert.TB) { be.EqualLength(map[int]int{}, 0) },
		func(be assert.TB) { be.EqualLength(map[int]int{1: 1}, 1) },
		func(be assert.TB) {
			ch := make(chan int, 1)
			be.EqualLength(ch, 0)
		},
		func(be assert.TB) {
			ch := make(chan int, 1)
			ch <- 1
			be.EqualLength(ch, 1)
		},
		func(be assert.TB) {
			seq2 := maps.All(map[int]int{1: 1})
			be.EqualLength(seq2, 1)
		},
	}

	for _, test := range okayTests {
		var buf strings.Builder
		m := &mockingT{w: &buf}
		test(assert.Continues(m))
		if m.Failed() {
			t.Fatal("failed too soon")
		}
		if buf.String() != "" {
			t.Fatal("wrote too much")
		}
	}

	badTests := []func(be assert.TB){
		func(be assert.TB) { be.SlicesEqual([]string{}, []string{""}) },
		func(be assert.TB) { be.Truthy(time.Time{}.Local()) },
		func(be assert.TB) { be.Falsey([]string{""}) },
		func(be assert.TB) { be.Truthy([]string(nil)) },
		func(be assert.TB) {
			be.OK(func() (int, error) { return 0, errors.New("") }())
		},
		func(be assert.TB) {
			be.OK2(func() (int, string, error) { return 0, "", errors.New("") }())
		},
		func(be assert.TB) { be.True(false) },
		func(be assert.TB) { be.False(true) },
		func(be assert.TB) {
			seq2 := maps.All(map[int]int{1: 1})
			be.EqualLength(seq2, 0)
		},
		func(be assert.TB) {
			ch := make(chan int, 1)
			be.EqualLength(ch, 1)
		},
		func(be assert.TB) {
			ch := make(chan int, 1)
			ch <- 1
			be.EqualLength(ch, 0)
		},
		func(be assert.TB) {
			ch := make(chan int, 1)
			close(ch)
			be.EqualLength(ch, 1)
		},
		func(be assert.TB) {
			be.Panicked(func() {})
		},
	}

	for _, test := range badTests {
		var buf strings.Builder
		m := &mockingT{w: &buf}
		test(assert.FailsNow(m))
		if !m.hasFailedNow {
			t.Fatal("did not fail")
		}
		if buf.String() == "" {
			t.Fatal("wrote too little")
		}
	}
}

func TestTB_Logf(t *testing.T) {
	var buf strings.Builder
	m := &mockingT{w: &buf}
	be := assert.Continues(m)
	be.Logf("hi")
	assert.FailsNow(t).Equal(buf.String(), "hi\n")
}

func TestContinues(t *testing.T) {
	// Make sure Continues and FailNows work as both functions and methods
	{
		m := &mockingT{}
		be := assert.Continues(m)
		be.Equal(1, 0)
		assert.Continues(t).
			True(m.hasFailed).
			False(m.hasFailedNow)
	}
	{
		m := &mockingT{}
		be := assert.Continues(m).FailsNow()
		be.Equal(1, 0)
		assert.Continues(t).
			True(m.hasFailed).
			True(m.hasFailedNow)
	}
	{
		m := &mockingT{}
		be := assert.FailsNow(m)
		be.Equal(1, 0)
		assert.Continues(t).
			True(m.hasFailed).
			True(m.hasFailedNow)
	}
	{
		m := &mockingT{}
		be := assert.FailsNow(m).Continues()
		be.Equal(1, 0)
		assert.Continues(t).
			True(m.hasFailed).
			False(m.hasFailedNow)
	}
}
