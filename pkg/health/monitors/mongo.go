package monitors

import "errors"

type mongo struct {}

// NewMongoClient returns a new instance of the mongo client
func NewMongoClient() mongo {
	return mongo{}
}

// CheckStatus returns the status of the client
func (m mongo) CheckStatus() error {
	return errors.New("ups... something is wrong")
}