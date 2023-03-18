package errors

type DatabaseNotFoundError struct {
	s string
}

func NewDatabaseNotFoundError(text string) error {
	return &DatabaseQueryFailError{text}
}

func (e *DatabaseNotFoundError) Error() string {
	return e.s
}
