package monitors

type mysql struct {}

// NewMySQLClient returns a new instance of the mysql client
func NewMySQLClient() mysql {
	return mysql{}
}

// CheckStatus returns the status of the client
func (m mysql) CheckStatus() error {
	return nil
}
