package want

import "testing"

func Eq[T comparable](t *testing.T, a, b T) {
	if a != b {
		Fatal(t, a, b)
	}
}
