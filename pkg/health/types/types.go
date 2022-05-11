package types

import "github.com/lucaspichi06/healthChecker/pkg/health/monitors"

// Monitor struct for the health checker
type Monitor struct {
	Name     string
	Handler  monitors.Monitor
	Critical bool
}
