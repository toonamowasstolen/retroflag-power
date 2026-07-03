package actions

type Type string

const TypeNoop Type = "noop"

type Action struct {
	Type    Type
	Message string
	DryRun  bool
}

func NewDryRunNoop(message string) Action {
	return Action{
		Type:    TypeNoop,
		Message: message,
		DryRun:  true,
	}
}
