package models

type ErrorMessage struct {
	Message string `json:"error"`
}

type SuccessMessage struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}
