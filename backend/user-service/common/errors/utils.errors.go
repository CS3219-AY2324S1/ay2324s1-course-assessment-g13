package errors

import "net/http"

const (
	INTERNAL_SERVER_ERROR_MESSAGE = "Internal Server Error"
	INVALID_TOKEN_MESSAGE         = "Invalid Token"
)

func InternalServerError() ServiceError {
	return NewServiceError(
		INTERNAL_SERVER_ERROR_MESSAGE,
		http.StatusInternalServerError,
	)
}

func UnauthorisedError(message string) ServiceError {
	return NewServiceError(
		message,
		http.StatusUnauthorized,
	)
}

func InvalidTokenError() ServiceError {
	return NewServiceError(
		INVALID_TOKEN_MESSAGE,
		http.StatusBadRequest,
	)
}

func MethodNotAllowedError(message string) ServiceError {
	return NewServiceError(
		message,
		http.StatusMethodNotAllowed,
	)
}

func ParseErrorToServiceError(err error) (int, string) {
	serviceError, ok := err.(ServiceError)
	if ok {
		return serviceError.GetStatus(), serviceError.Error()
	}
	return http.StatusInternalServerError, INTERNAL_SERVER_ERROR_MESSAGE
}
