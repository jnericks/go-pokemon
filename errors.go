package gopokemon

type errorString struct {
	msg string
}

func (e *errorString) Error() string {
	return e.msg
}

func raiseError(msg string) error {
	return &errorString{msg}
}
