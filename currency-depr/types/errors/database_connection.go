package errors

type DatabaseConnectionError struct {
	s string
}

func NewDatabaseConnectionError(text string) error {
	return &DatabaseConnectionError{text}
}

func (e *DatabaseConnectionError) Error() string {
	return e.s
}
