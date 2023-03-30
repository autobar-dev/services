package types

// Generic error
//
// swagger:model RestError
type RestError struct {
	// Error message
	// example: could not fetch available currencies
	Error string `json:"error"`
}
