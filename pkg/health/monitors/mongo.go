package monitors

import "errors"

type mongo struct {}

func NewMongoClient() mongo {
	return mongo{}
}

func (m mongo) CheckStatus() error {
	return errors.New("ups... something is wrong")
}