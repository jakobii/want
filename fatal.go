package want

import "testing"

func Fatal(t *testing.T, want, got interface{}) {
	t.Fatalf("want: %v\n got %v\n", want, got)
}
