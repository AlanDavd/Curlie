package ports

// CurlRepository for data persistence.
type CurlRepository interface {
	SaveCommand(command string) error
	GetCommands() ([]string, error)
}
