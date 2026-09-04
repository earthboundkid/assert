package assert_test

import (
	"github.com/earthboundkid/assert"
)

func ExampleTester_DeepEqual() {
	be := assert.Continue(&mockingT{})

	// good
	m1 := map[int]bool{1: true, 2: false}
	m2 := map[int]bool{1: true, 2: false}
	be.DeepEqual(m1, m2)

	// bad
	var s1 []int
	s2 := []int{}
	be.DeepEqual(s1, s2) // DeepEqual is picky about nil vs. len 0

	// Output:
	// reflect.DeepEqual([]int(nil), []int{}) == false
}
