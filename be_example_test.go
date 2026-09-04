package assert_test

import (
	"errors"
	"os"
	"strings"

	"github.com/earthboundkid/assert"
)

func Example() {
	be := assert.Relaxed(&mockingT{})

	be.
		Equal("hello", "world").    // bad
		Equal("goodbye", "goodbye") // good
	be.
		Unequal("hello", "world").    // good
		Unequal("goodbye", "goodbye") // bad

	s := []int{1, 2, 3}
	be.
		AllEqual([]int{1, 2, 3}, s). // good
		AllEqual([]int{3, 2, 1}, s)  // bad

	var err error
	be.
		Zero(err).                     // good
		ErrorIs(nil, err).             // good
		Nonzero(err).                  // bad
		ErrorIs(err, os.ErrPermission) // bad

	err = errors.New("(O_o)")
	be.ErrorAsType[*os.PathError](err) // bad
	be.Nonzero(err)                    // good

	f := be.OK(os.Open("nosuchfile")) // bad
	be.Zero(f)                        // good

	type mytype string
	var mystring mytype = "hello, world"
	be.Match(mystring, `world`)                 // good
	be.Match(mystring, `World`)                 // bad
	be.Match([]byte("\a\b\x00\r\t"), `^\W*$`)   // good
	be.NoMatch([]byte("\a\b\x00\r\t"), `^\W*$`) // bad

	seq := strings.FieldsSeq("1 2 3 4")
	be.
		EqualLength(seq, 4).     // good
		EqualLength(seq, 1).     // bad
		AtLeastLength(seq, 1).   // good
		AtLeastLength(seq, 5).   // bad
		AtLeastLength("123", 3). // good
		AtLeastLength("123", 4)  // bad

	// Output:
	// want: world; got: hello
	// got: goodbye
	// got: [3 2 1]; want: [1 2 3]
	// got: <nil>
	// got errors.Is(<nil>, permission denied) == false
	// got errors.AsType[*fs.PathError]((O_o)) == false
	// err != nil: open nosuchfile: no such file or directory
	// missing match: /World/ !~ "hello, world"
	// unexpected match: /^\W*$/ =~ "\a\b\x00\r\t"
	// want len(seq) == 1; got at least 2
	// want len(seq) >= 5; got 4
	// want len(seq) >= 4; got 3
}
