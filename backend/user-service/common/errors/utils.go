package errors

import "net/http"

func InternalServerError() ServiceError {
	return NewServiceError(
		"Internal Server Erorr",
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
		"Invalid Token",
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
	return http.StatusInternalServerError, "Internal Server Error"
}
