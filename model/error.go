package model

// Error estrutura para o manejo dos erros
type Error struct {
	ID      string `json:"id"`
	Message string `json:"message"`
	Status  int    `json:"status"`
}

// NewError instancia um novo erro
func NewError(id, message string, status int) *Error {
	return &Error{
		ID:      id,
		Message: message,
		Status:  status,
	}
}
