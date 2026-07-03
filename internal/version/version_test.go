package version

import "testing"

func TestString(t *testing.T) {
	const want = "retroflag-powerd 0.1.0-dev"

	if got := String(); got != want {
		t.Fatalf("String() = %q, want %q", got, want)
	}
}
