package assert

func (be Tester) Type[T any](v any) T {
	be.tb.Helper()
	vv, ok := v.(T)
	if !ok {
		be.fatalf("could not assert (%v).(%T)", v, *new(T))
	}
	return vv
}
