package message

import "api-gateway/models"

func CreateErrorMessage(message string) models.ErrorMessage {
	return models.ErrorMessage{
		Message: message,
	}
}

func CreateSuccessMessage(message string, data ...interface{}) models.SuccessMessage {
	var incomingData interface{}

	if len(data) > 0 {
		incomingData = data[0]
	}

	return models.SuccessMessage{
		Message: message,
		Data:    incomingData,
	}
}
