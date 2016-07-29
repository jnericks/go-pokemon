package main

type errorString struct {
	msg string
}

func (e *errorString) Error() string {
	return e.msg
}

// RaiseError ...
func RaiseError(msg string) error {
	return &errorString{msg}
}
