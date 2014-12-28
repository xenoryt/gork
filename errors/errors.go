package errors

//GenericError is a generic error
type GenericError string

func (err GenericError) Error() string {
	return string(err)
}

//New creates a new error with the given message
func New(msg string) GenericError {
	return GenericError(msg)
}
