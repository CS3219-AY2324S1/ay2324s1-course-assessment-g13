package model

type ErrorMessage struct {
	Message string `json:"error"`
}

type SuccessUserMessage struct {
	Message string `json:"message"`
	User    User   `json:"user,omitempty"`
}

type SuccessUsersMessage struct {
	Message string `json:"message"`
	Users   []User `json:"users,omitempty"`
}

type SuccessMessage struct {
	Message string `json:"message"`
}
