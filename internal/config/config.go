package config

import "github.com/toonamowasstolen/retroflag-power/internal/version"

type Config struct {
	AppName string
	Version string
	DryRun  bool
}

func Default() Config {
	return Config{
		AppName: version.Name,
		Version: version.Version,
		DryRun:  true,
	}
}
