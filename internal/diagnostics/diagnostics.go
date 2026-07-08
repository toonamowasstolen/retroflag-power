package diagnostics

import (
	"fmt"
	"io"
)

const textStub = "retroflag-powerd diagnostics\n" +
	"Local diagnostics are planned but not implemented yet.\n" +
	"This command is local-only and read-only in this build.\n" +
	"No GPIO, shutdown, systemd, SafeShutdown, file, telemetry, or network action was performed.\n"

const usageText = "retroflag-powerd diagnostics\n" +
	"\n" +
	"Usage:\n" +
	"  retroflag-powerd diagnostics [--format text|json]\n" +
	"  retroflag-powerd diagnostics --help\n" +
	"  retroflag-powerd diagnostics -h\n" +
	"\n" +
	"Status:\n" +
	"  Diagnostics collection is planned but not implemented yet.\n" +
	"  This command is local-only and read-only in this build.\n" +
	"  No diagnostics bundle is generated, and no GPIO, OS, display, audio, process, SafeShutdown, file, telemetry, or network state is inspected.\n" +
	"\n" +
	"Supported formats:\n" +
	"  text\n" +
	"  json\n" +
	"\n" +
	"Examples:\n" +
	"  retroflag-powerd diagnostics\n" +
	"  retroflag-powerd diagnostics --format text\n" +
	"  retroflag-powerd diagnostics --format json\n"

const jsonStub = "{\n" +
	"  \"command\": \"retroflag-powerd diagnostics\",\n" +
	"  \"implemented\": false,\n" +
	"  \"local_only\": true,\n" +
	"  \"read_only\": true,\n" +
	"  \"message\": \"Local diagnostics are planned but not implemented yet.\",\n" +
	"  \"actions_performed\": []\n" +
	"}\n"

func RenderText(w io.Writer) {
	fmt.Fprint(w, textStub)
}

func RenderUsage(w io.Writer) {
	fmt.Fprint(w, usageText)
}

func RenderJSON(w io.Writer) {
	fmt.Fprint(w, jsonStub)
}

func UsageText() string {
	return usageText
}
