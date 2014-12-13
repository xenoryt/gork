package ui

type GenericError string
type InitError string

func (err GenericError) Error() string {
	return string(err)
}

func (err InitError) Error() string {
	return "Failed to initialize: " + string(err)
}
