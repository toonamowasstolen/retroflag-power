package app

import (
	"bytes"
	"context"
	"log"
	"testing"

	"github.com/toonamowasstolen/retroflag-power/internal/config"
)

func TestRunLogsLifecycle(t *testing.T) {
	var output bytes.Buffer
	logger := log.New(&output, "", 0)
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	New(logger, config.Default()).Run(ctx)

	const want = `retroflag-powerd 0.1.0-dev starting dry_run=true
retroflag-powerd ready
shutdown signal received
retroflag-powerd stopped
`
	if got := output.String(); got != want {
		t.Fatalf("Run() logs:\n%q\nwant:\n%q", got, want)
	}
}
