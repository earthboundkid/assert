package assert_test

import (
	"fmt"

	"github.com/earthboundkid/assert"
)

func ExampleTruthy() {
	fmt.Printf("%#v is %t-ish\n", 0, assert.Truthy(0))
	fmt.Printf("%#v is %t-ish\n", 1, assert.Truthy(1))
	fmt.Printf("%#v is %t-ish\n", "", assert.Truthy(""))
	fmt.Printf("%#v is %t-ish\n", "hi", assert.Truthy("hi"))
	fmt.Printf("%#v is %t-ish\n", error(nil), assert.Truthy[error](nil))
	// Output:
	// 0 is false-ish
	// 1 is true-ish
	// "" is false-ish
	// "hi" is true-ish
	// <nil> is false-ish
}
