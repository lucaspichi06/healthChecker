package monitors

type Monitor interface {
	CheckStatus() error
}
