package want

import "testing"

func Eq[T comparable](t *testing.T, a, b T) {
	if a != b {
		fatal_want_got(t, a, b)
	}
}
