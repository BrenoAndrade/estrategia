package model

// Error struct for error management
type Error struct {
	ID      string `json:"id"`
	Message string `json:"message"`
	Status  int    `json:"status"`
}

// NewError intance handle error
func NewError(id, message string, status int) *Error {
	return &Error{
		ID:      id,
		Message: message,
		Status:  status,
	}
}
