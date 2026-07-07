package events

import "testing"

func TestLifecycleEventTypes(t *testing.T) {
	tests := map[Type]string{
		TypeDaemonStarting:         "daemon.starting",
		TypeDaemonReady:            "daemon.ready",
		TypeShutdownSignalReceived: "shutdown.signal_received",
		TypeDaemonStopped:          "daemon.stopped",
		TypePowerIntentReceived:    "power.intent_received",
		TypeDryRunPlanPrepared:     "power.dry_run_plan_prepared",
		TypeNoopExecutionCompleted: "power.noop_execution_completed",
	}

	for eventType, want := range tests {
		if got := string(eventType); got != want {
			t.Errorf("event type = %q, want %q", got, want)
		}
	}
}
