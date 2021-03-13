package apierror

type ApiError interface {
	ErrorStatusCode() int
	Error() string
}

func New(httpStatusCode int, message string) *apiError {
	return &apiError{
		HttpCode: httpStatusCode,
		Message:  message,
	}
}

type apiError struct {
	Message  string `json:"message"`
	HttpCode int    `json:"http_code"`
}

func (self apiError) ErrorStatusCode() int {
	return self.HttpCode
}

func (self apiError) Error() string {
	return self.Message
}
