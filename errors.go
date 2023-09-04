package datetime

type Error struct {
	msg string
}

func (e *Error) Error() string {
	return e.msg
}

func NewDateTimeError(msg string) error {
	return &Error{msg: msg}
}
