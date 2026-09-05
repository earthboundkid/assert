# Assert [![Go Reference](https://pkg.go.dev/badge/github.com/earthboundkid/assert.svg)](https://pkg.go.dev/github.com/earthboundkid/assert) [![Coverage Status](https://coveralls.io/repos/github/earthboundkid/assert/badge.svg)](https://coveralls.io/github/earthboundkid/assert)

Package assert is the minimalist testing helper for Go.

Inspired by [Mat Ryer](https://github.com/matryer/is) and [Alex Edwards](https://www.alexedwards.net/blog/easy-test-assertions-with-go-generics).

## Features

- Simple and readable test assertions using generics, chainable off a `Tester`
- Built-in helpers for common cases like `Tester.OK` and `Tester.Equal`
- Choose `Fail` or `FailNow` semantics depending on how you construct your `Tester` (`assert.Continue(t)` or `assert.FailNow(t)`)
- Helpers for testing against golden files with the testfile subpackage
- No sub-dependencies: just uses standard library

## Example usage

Test for simple equality using generics:

```go
// First create a testing helper
be := assert.Continue(t)

// Test two unequal strings
be.Equal("hello", "world")    // bad
// t.Fail(); t.Log("want: world; got: hello")
// Test two equal strings
be.Equal("goodbye", "goodbye") // good
// Test equal integers, etc.
be.Equal(resp.StatusCode, 200)
be.Equal(gotPtr, tc.wantPtr)

// Test for inequality
be.NotEqual("hello", "world")    // good
be.NotEqual("goodbye", "goodbye") // bad
// t.Fail(); t.Log("got: goodbye")
```

Chain related tests:

```go
be.
	Equal(x, 1).
	Equal(y, 2)
```

Test for equality of slices:

```go
s := []int{1, 2, 3}
be.SlicesEqual([]int{1, 2, 3}, s) // good
be.SlicesEqual([]int{3, 2, 1}, s) // bad
// t.Fail(); t.Log("got: [3 2 1]; want: [1 2 3]")
```

Handle errors:

```go
f := be.OK(os.Open("nosuchfile")) // bad, and also returns nil *os.File
be.Zero(f)                        // good

var err error
be.Zero(err)                      // good
be.ErrorIs(nil, err)              // good
be.NotZero(err)                   // bad
be.ErrorIs(err, os.ErrPermission) // bad

err = errors.New("(O_o)")
be.ErrorAsType[*os.PathError](err) // bad
be.NotZero(err)                    // good

```

Check for regexp matching:

```go
be.Match(mystring, `world`)               // good
be.Match(mystring, `World`)               // bad
// t.Fail(); t.Log(`missing match: /World/ !~ "hello, world"`)
be.Match([]byte("\a\b\x00\r\t"), `^\W*$`)    // good
be.NotMatch([]byte("\a\b\x00\r\t"), `^\W*$`) // bad
```

Check how long something rangeable is:

```go
seq := strings.FieldsSeq("1 2 3 4")
be.EqualLength(seq, 4)     // good
be.EqualLength(seq, 1)     // bad
be.AtLeastLength(seq, 1)   // good
be.AtLeastLength(seq, 5)   // bad
be.AtLeastLength("123", 3) // good
be.AtLeastLength("123", 4) // bad
```

Test anything else:

```go
be.True(o.IsValid())
```

Test using goldenfiles:

```go
// Start a sub-test for each .txt file
testfile.Run(t, "testdata/*.txt", func(t *testing.T, path string) {
	// Read the file
	input := testfile.Read(t, path)

	// Do some conversion on it
	type myStruct struct{ Whatever string }
	got := myStruct{strings.ToUpper(input)}

	// See if the struct is equivalent to a .json file
	wantFile := testfile.Ext(path, ".json")
	testfile.EqualJSON(t, wantFile, got)

	// If it's not equivalent,
	// the got struct will be dumped
	// to a file named testdata/-failed-test-name.json
})
```

## Philosophy
Tests usually should not fail. When they do fail, the failure should be repeatable. Therefore, it doesn't make sense to spend a lot of time writing good test messages. (This is unlike error messages, which should happen fairly often, and in production, irrepeatably.) Package assert is designed to simply fail a test quickly and quietly if a condition is not met with a reference to the line number of the failing test. If the reason for having the test is not immediately clear from context, you can write a comment, just like in normal code. If you do need more extensive reporting to figure out why a test is failing, use `testing.TB.Log` to capture more information.

The assertions in package assert are methods of the `Tester` type, which captures a `testing.TB` and can either call `testing.TB.Fail` or `testing.TB.FailNow` on failure depending on how you want the assertions to work.

Most tests just need simple equality testing, which is handled by `Tester.Equal` (for comparable types), `Tester.SlicesEqual` (for slices of comparable types), and `Tester.DeepEqual` (which relies on `reflect.DeepEqual`). Another common test is that a string or byte slice should contain or not some substring, which is handled by `Tester.Match` and `Tester.NotMatch`. Rather than package assert providing every possible test helper, you are encouraged to write your own advanced helpers for use with `Tester.True`, while package assert takes away the drudgery of writing yet another simple `func nilErr(t *testing.T, err) { ... }`.

The `github.com/earthboundkid/assert/testfile` subpackage has functions that make it easy to write file-based tests that ensure that the output of some transformation matches a [golden file](https://softwareengineering.stackexchange.com/questions/358786/what-are-golden-files). Subtests can automatically be run for all files matching a glob pattern, such as `testfile.Run(t, "testdata/*/input.txt", ...)`. If the test fails, the failure output will be written to a file, such as "testdata/basic-test/-failed-output.txt", and then the output can be examined via diff testing with standard tools. Set the environmental variable `TESTFILE_UPDATE` to update the golden file.
