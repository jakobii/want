package want

import "testing"

func Fatal(t *testing.T, want, got interface{}) {
	t.Fatalf("\nwant: %+v\n got %+v\n", want, got)
}
