package main

import (
	"fmt"
	"github.com/lucaspichi06/healthChecker/pkg/health"
	"github.com/lucaspichi06/healthChecker/pkg/health/monitors"
	"github.com/lucaspichi06/healthChecker/pkg/health/types"
	"time"
)

func main() {
	healthService := health.GetInstance()
	redisClient := monitors.NewRedisClient()
	mongoClient := monitors.NewMongoClient()
	mysqlClient := monitors.NewMySQLClient()

	healthService.Monitor(types.Monitor{
		Name: "Redis",
		Critical: true,
		Handler: redisClient,
	})

	healthService.Monitor(types.Monitor{
		Name: "MySQL",
		Critical: false,
		Handler: mysqlClient,
	})

	healthService.Monitor(types.Monitor{
		Name: "MongoDB",
		Critical: false,
		Handler: mongoClient,
	})

	ticker := time.NewTicker(5 * time.Second)

	for {
		select {
		case <-ticker.C:
			fmt.Println("--------------------- check status start ---------------------")
			for _, v := range healthService.Check() {
				fmt.Println(v)
			}
			fmt.Println("--------------------- check status end ---------------------")
		}
	}
}
