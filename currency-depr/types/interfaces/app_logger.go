package interfaces

type AppLogger interface {
	Info(interface{}, ...interface{})
	Error(interface{}, ...interface{})
}
