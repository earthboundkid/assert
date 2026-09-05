package assert

// Catch runs the callback and returns the recovered panic, if any.
func Catch(fn func()) (r any) {
	defer func() {
		r = recover()
	}()
	fn()
	return
}

// Panicked asserts that fn panics when run.
func (be TB) Panicked(fn func()) TB {
	be.Helper()
	if pval := Catch(fn); pval == nil {
		be.fatalf("expected panic; got nil")
	}
	return be
}
