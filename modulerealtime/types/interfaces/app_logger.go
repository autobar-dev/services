package interfaces

// AppLogger is an interface for the logger used in the application.
type AppLogger interface {
	// Debug logs a debug message.
	Debug(interface{}, ...interface{})
	// Info logs an info message.
	Info(interface{}, ...interface{})
	// Error logs an error message.
	Error(interface{}, ...interface{})
}
