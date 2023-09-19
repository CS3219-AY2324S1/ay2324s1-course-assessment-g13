package errors

type ServiceError interface {
	Error() string
	GetMessage() string
	GetStatus() int
}

type serviceError struct {
	Message string
	Status  int
}

func (serviceErr *serviceError) GetMessage() string {
	return serviceErr.Message
}

func (serviceErr *serviceError) GetStatus() int {
	return serviceErr.Status
}
func (serviceError *serviceError) Error() string {
	return serviceError.Message
}

func NewServiceError(message string, status int) ServiceError {
	return &serviceError{message, status}
}
