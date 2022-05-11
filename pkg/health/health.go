package health

import (
	"errors"
	"fmt"
	"github.com/lucaspichi06/healthChecker/pkg/health/types"
	"sync"
)

var (
	h    Health
	once sync.Once
)

// GetInstance returns a unique instance of the health service
func GetInstance() Health {
	once.Do(func() {
		h = &service{
			monitors: make(map[string]types.Monitor),
		}
	})
	return h
}

// Health interface for checking the status of the different clients
type Health interface {
	Check()
	Monitor(types.Monitor) error
}

type service struct {
	monitors map[string]types.Monitor
}

// Check return the status of the registered clients
func (s *service) Check() {
	for _, v := range s.monitors {
		v := v
		go func() {
			status := "ok"
			if err := v.Handler.CheckStatus(); err != nil {
				if v.Critical {
					// do something
				}
				status = err.Error()
			}
			fmt.Println(fmt.Sprintf("resourceName: %s - status: %s", v.Name, status))
		}()
	}
}

// Monitor registers a new monitor into the health service
func (s *service) Monitor(monitor types.Monitor) error {
	if _, ok := s.monitors[monitor.Name]; ok {
		return errors.New("monitor already exists")
	}

	s.monitors[monitor.Name] = monitor
	return nil
}
