package monitors

type redis struct {}

func NewRedisClient() redis {
	return redis{}
}

func (r redis) CheckStatus() error {
	return nil
}