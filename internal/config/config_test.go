package config

import "testing"

func TestDefault(t *testing.T) {
	want := Config{
		AppName: "retroflag-powerd",
		Version: "0.1.0-dev",
		DryRun:  true,
	}

	if got := Default(); got != want {
		t.Fatalf("Default() = %#v, want %#v", got, want)
	}
}
