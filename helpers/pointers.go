package helpers

func PtrTo[T any](t T) *T {
	return &t
}
