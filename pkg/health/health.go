package health

import (
	"errors"
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

type Health interface {
	Check()
	Monitor(types.Monitor) error
}

type service struct {
	monitors map[string]types.Monitor
}

func (s *service) Check() {
	for _, v := range s.monitors {
		go v.Handler.CheckStatus()
	}
}

func (s *service) Monitor(monitor types.Monitor) error {
	if _, ok := s.monitors[monitor.Name]; ok {
		return errors.New("monitor already exists")
	}

	s.monitors[monitor.Name] = monitor
	return nil
}
