package assert

import (
	"reflect"
	"regexp"
)

// Match asserts got matches the [regexp] pattern.
//
// The pattern must compile.
func (be Tester) Match[byteseq ~string | ~[]byte](got byteseq, pattern string) Tester {
	be.tb.Helper()
	reg := regexp.MustCompile(pattern)
	if !match(reg, got) {
		be.fatalf("missing match: /%s/ !~ %q", pattern, got)
	}
	return be
}

// NotMatch asserts got does not matches the [regexp] pattern.
//
// The pattern must compile.
func (be Tester) NotMatch[byteseq ~string | ~[]byte](got byteseq, pattern string) Tester {
	be.tb.Helper()
	reg := regexp.MustCompile(pattern)
	if match(reg, got) {
		be.fatalf("unexpected match: /%s/ =~ %q", pattern, got)
	}
	return be
}

func match[byteseq ~string | ~[]byte](reg *regexp.Regexp, got byteseq) bool {
	switch rv := reflect.ValueOf(got); rv.Kind() {
	case reflect.String:
		return reg.MatchString(rv.String())
	case reflect.Slice:
		return reg.Match(rv.Bytes())
	}
	panic("unreachable")
}
