package monitors

// Monitor interface for checking status on the different clients
type Monitor interface {
	CheckStatus() error
}
