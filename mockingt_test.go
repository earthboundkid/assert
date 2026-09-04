package assert_test

import (
	"cmp"
	"fmt"
	"io"
	"os"
	"sync"
	"testing"
)

type mockingT struct {
	testing.T
	m            sync.Mutex
	hasFailed    bool
	hasFailedNow bool
	w            io.Writer
}

func (m *mockingT) FailNow() {
	m.m.Lock()
	defer m.m.Unlock()
	m.hasFailed = true
	m.hasFailedNow = true
}

func (m *mockingT) Logf(format string, args ...any) {
	out := cmp.Or(m.w, io.Writer(os.Stdout))
	fmt.Fprintf(out, format+"\n", args...)
}

func (*mockingT) Helper() {}

func (m *mockingT) Fatalf(format string, args ...any) {
	m.Errorf(format, args...)
	m.FailNow()
}

func (m *mockingT) Fail() {
	m.m.Lock()
	defer m.m.Unlock()
	m.hasFailed = true
}

func (m *mockingT) Errorf(format string, args ...any) {
	m.Logf(format, args...)
	m.Fail()
}

func (m *mockingT) Failed() bool {
	m.m.Lock()
	defer m.m.Unlock()
	return m.hasFailed
}
