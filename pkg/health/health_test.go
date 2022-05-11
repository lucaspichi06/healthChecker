package health

import (
	"errors"
	"github.com/lucaspichi06/healthChecker/pkg/health/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

type dummyRedis struct {
	checkStatus func() error
}

func (d dummyRedis) CheckStatus() error {
	return d.checkStatus()
}

func Test_service_Check(t *testing.T) {
	t.Run("checking status - status ok", func(t *testing.T) {
		service := &service{
			monitors: make(map[string]types.Monitor),
		}
		if err := service.Monitor(types.Monitor{
			Name:     "redis",
			Critical: true,
			Handler: dummyRedis{
				checkStatus: func() error {
					return nil
				},
			},
		}); err != nil {
			t.Errorf("Monitor() error = %v", err)
		}

		output := service.Check()

		assert.Len(t, output, 1)
		assert.Equal(t, "resourceName: redis - status: ok", output[0])
	})
	t.Run("checking status - status error", func(t *testing.T) {
		service := &service{
			monitors: make(map[string]types.Monitor),
		}
		if err := service.Monitor(types.Monitor{
			Name:     "redis",
			Critical: true,
			Handler: dummyRedis{
				checkStatus: func() error {
					return errors.New("test error")
				},
			},
		}); err != nil {
			t.Errorf("Monitor() error = %v", err)
		}

		output := service.Check()

		assert.Len(t, output, 1)
		assert.Equal(t, "resourceName: redis - status: test error", output[0])
	})
}

func Test_service_Monitor(t *testing.T) {
	service := GetInstance()
	t.Run("adding a new monitor - success", func(t *testing.T) {
		if err := service.Monitor(types.Monitor{
			Name:     "redis",
			Critical: false,
			Handler:  dummyRedis{},
		}); err != nil {
			t.Errorf("Monitor() error = %v", err)
		}
	})
	t.Run("adding a new monitor - error", func(t *testing.T) {
		err := service.Monitor(types.Monitor{
			Name:     "redis",
			Critical: false,
			Handler:  dummyRedis{},
		})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "monitor already exists")
	})
}
