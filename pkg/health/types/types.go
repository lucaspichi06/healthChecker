package types

import "github.com/lucaspichi06/healthChecker/pkg/health/monitors"

type Monitor struct {
	Name     string
	Handler  monitors.Monitor
	Critical bool
}
