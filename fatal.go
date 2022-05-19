package want

import "testing"

func Fatal(t *testing.T, want, got interface{}) {
	fatal_want_got(t, want, got)
}

func fatal_want_got(t *testing.T, want, got interface{}) {
	t.Fatalf("\nwant: %+v\ngot: %+v\n", want, got)
}
