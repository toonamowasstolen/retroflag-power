package events

import "testing"

func TestLifecycleEventTypes(t *testing.T) {
	tests := map[Type]string{
		TypeDaemonStarting:         "daemon.starting",
		TypeDaemonReady:            "daemon.ready",
		TypeShutdownSignalReceived: "shutdown.signal_received",
		TypeDaemonStopped:          "daemon.stopped",
	}

	for eventType, want := range tests {
		if got := string(eventType); got != want {
			t.Errorf("event type = %q, want %q", got, want)
		}
	}
}
