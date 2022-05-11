package monitors

type redis struct {}

// NewRedisClient returns a new instance of the redis client
func NewRedisClient() redis {
	return redis{}
}

// CheckStatus returns the status of the client
func (r redis) CheckStatus() error {
	return nil
}