package model

type (
	OKResponse struct {
		Status  string `json:"status,omitempty"`
		Message string `json:"message,omitempty"`
	}
	ErrorResponse struct {
		Error   string `json:"error,omitempty"`
		Message string `json:"message,omitempty"`
	}
)
