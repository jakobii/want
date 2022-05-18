package want

import "testing"

func Fatal(t *testing.T, want, got interface{}) {
	t.Fatalf("\nwant: %+v\ngot %+v\n", want, got)
}
