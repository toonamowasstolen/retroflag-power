package diagnostics

import (
	"bytes"
	"testing"
)

func TestRenderTextPrintsStub(t *testing.T) {
	var got bytes.Buffer

	RenderText(&got)

	const want = "retroflag-powerd diagnostics\n" +
		"Local diagnostics are planned but not implemented yet.\n" +
		"This command is local-only and read-only in this build.\n" +
		"No GPIO, shutdown, systemd, SafeShutdown, file, telemetry, or network action was performed.\n"
	if got.String() != want {
		t.Fatalf("RenderText() = %q, want %q", got.String(), want)
	}
}

func TestRenderJSONPrintsDeterministicStub(t *testing.T) {
	var got bytes.Buffer

	RenderJSON(&got)

	const want = "{\n" +
		"  \"command\": \"retroflag-powerd diagnostics\",\n" +
		"  \"implemented\": false,\n" +
		"  \"local_only\": true,\n" +
		"  \"read_only\": true,\n" +
		"  \"message\": \"Local diagnostics are planned but not implemented yet.\",\n" +
		"  \"actions_performed\": []\n" +
		"}\n"
	if got.String() != want {
		t.Fatalf("RenderJSON() = %q, want %q", got.String(), want)
	}
}

func TestRenderUsagePrintsUsageText(t *testing.T) {
	var got bytes.Buffer

	RenderUsage(&got)

	if got.String() != UsageText() {
		t.Fatalf("RenderUsage() = %q, want %q", got.String(), UsageText())
	}
}
