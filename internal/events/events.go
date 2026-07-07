package events

type Type string

const (
	TypeDaemonStarting         Type = "daemon.starting"
	TypeDaemonReady            Type = "daemon.ready"
	TypeShutdownSignalReceived Type = "shutdown.signal_received"
	TypeDaemonStopped          Type = "daemon.stopped"
	TypePowerIntentReceived    Type = "power.intent_received"
	TypeDryRunPlanPrepared     Type = "power.dry_run_plan_prepared"
	TypeNoopExecutionCompleted Type = "power.noop_execution_completed"
)

type Event struct {
	Type    Type
	Message string
}
