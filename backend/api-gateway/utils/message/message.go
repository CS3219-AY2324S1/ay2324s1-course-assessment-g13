package message

import "api-gateway/models"

func CreateErrorMessage(message string) models.ErrorMessage {
	return models.ErrorMessage{
		Message: message,
	}
}

func CreateSuccessMessage(message string) models.SuccessMessage {
	return models.SuccessMessage{
		Message: message,
	}
}

func CreateSuccessUserMessage(message string, user ...models.User) models.SuccessUserMessage {
	var userData models.User

	if len(user) > 0 {
		userData = user[0]
	}

	return models.SuccessUserMessage{
		Message: message,
		User:    userData,
	}
}

func CreateSuccessUsersMessage(message string, users []models.User) models.SuccessUsersMessage {
	return models.SuccessUsersMessage{
		Message: message,
		Users:   users,
	}
}

func CreateSuccessUupdateUserMessage(message string, user models.UserServiceUser) models.SuccessUpdateUserMessage {
	return models.SuccessUpdateUserMessage{
		Message: message,
		User:    user,
	}
}
