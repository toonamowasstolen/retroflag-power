package actions

import "testing"

func TestNewDryRunNoop(t *testing.T) {
	want := Action{
		Type:    TypeNoop,
		Message: "no operation planned",
		DryRun:  true,
	}

	if got := NewDryRunNoop(want.Message); got != want {
		t.Fatalf("NewDryRunNoop() = %#v, want %#v", got, want)
	}
}
