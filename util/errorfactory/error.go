package errorfactory

type ResponseError struct {
	msg string
}

var _ error = &ResponseError{}

func NewResponseError(msg string) ResponseError {
	return ResponseError{msg: msg}
}

func (re *ResponseError) Error() string {
	return re.msg
}
