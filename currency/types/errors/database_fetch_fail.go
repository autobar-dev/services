package errors

type DatabaseQueryFailError struct {
	s string
}

func NewDatabaseQueryFailError(text string) error {
	return &DatabaseQueryFailError{text}
}

func (e *DatabaseQueryFailError) Error() string {
	return e.s
}
