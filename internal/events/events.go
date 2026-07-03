package events

type Type string

const (
	TypeDaemonStarting         Type = "daemon.starting"
	TypeDaemonReady            Type = "daemon.ready"
	TypeShutdownSignalReceived Type = "shutdown.signal_received"
	TypeDaemonStopped          Type = "daemon.stopped"
)

type Event struct {
	Type    Type
	Message string
}
