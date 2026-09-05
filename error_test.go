package assert_test

import (
	"errors"
	"testing"

	"github.com/earthboundkid/assert"
)

func TestNotOK(t *testing.T) {
	type testcase struct {
		v   int
		err error
		ok  bool
	}
	assert.Run(t, map[string]testcase{
		"zero+nil":  {v: 0, err: nil, ok: false},
		"value+nil": {v: 1, err: nil, ok: false},
		"value+err": {v: 1, err: errors.New(""), ok: false},
		"zero+err":  {v: 0, err: errors.New(""), ok: true},
	}, func(be assert.TB, tc testcase) {
		m := &mockingT{}
		assert.FailsNow(m).NotOK(tc.v, tc.err)
		be.Equal(m.hasFailed, !tc.ok)
	})
}
