package model

type Error struct {
	Number       uint
	ResponseCode uint
	ErrString    string
}

func (err *Error) Error() string {
	return err.ErrString
}
