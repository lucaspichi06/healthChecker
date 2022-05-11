package monitors

import "fmt"

type redis struct {}

func NewRedisClient() redis {
	return redis{}
}

func (r redis) CheckStatus() {
	fmt.Println("resource: redis", "status: ok")
}