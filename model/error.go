package model

type Error struct {
	Number       uint
	ResponseCode uint
	ErrString    string
}

func (err *Error) Error() string {
	return err.ErrString
}

const (
	BadId = iota
	BadRequestBody
	ServerErr
)

func NewError(errorNo uint, responseCode uint, err error) *Error {
	return &Error{
		Number:       errorNo,
		ResponseCode: responseCode,
		ErrString:    err.Error(),
	}
}
