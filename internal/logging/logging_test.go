package logging

import (
	"bytes"
	"strings"
	"testing"
)

func TestNewWritesToGivenWriter(t *testing.T) {
	var buf bytes.Buffer
	logger := New(&buf)

	logger.Print("hello")

	if got := buf.String(); !strings.Contains(got, "hello") {
		t.Fatalf("logger output = %q, want it to contain %q", got, "hello")
	}
}

func TestBannerContainsAppNameAndVersion(t *testing.T) {
	got := Banner("retroflag-powerd", "0.1.0-dev")

	if !strings.Contains(got, "retroflag-powerd 0.1.0-dev") {
		t.Fatalf("Banner() = %q, want it to contain %q", got, "retroflag-powerd 0.1.0-dev")
	}
}

func TestBannerHasNoTrailingReadyLine(t *testing.T) {
	// The " ready\n" substring is load-bearing elsewhere (main.go's
	// readySignalWriter watches for it to detect daemon readiness) — the
	// banner must never accidentally contain it.
	got := Banner("retroflag-powerd", "0.1.0-dev")

	if strings.Contains(got, " ready\n") {
		t.Fatalf("Banner() = %q, must not contain %q", got, " ready\n")
	}
}

func TestWriteBannerWritesToGivenWriter(t *testing.T) {
	var buf bytes.Buffer

	WriteBanner(&buf, "retroflag-powerd", "0.1.0-dev")

	if got := buf.String(); got != Banner("retroflag-powerd", "0.1.0-dev") {
		t.Fatalf("WriteBanner output = %q, want %q", got, Banner("retroflag-powerd", "0.1.0-dev"))
	}
}
