package crash

type ErrorMessage struct {
	ErrorCode int    `json:"errorCode"`
	Message   string `json:"message"`
	Error     string `json:"error"`
}

func GenerateError(errorCode int, errorSystem error) ErrorMessage {
	return ErrorMessage{
		Message:   ErrorCodes[errorCode],
		ErrorCode: errorCode,
		Error:     errorSystem.Error(),
	}
}
