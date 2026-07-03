package version

const (
	Name    = "retroflag-powerd"
	Version = "0.1.0-dev"
)

func String() string {
	return Name + " " + Version
}
