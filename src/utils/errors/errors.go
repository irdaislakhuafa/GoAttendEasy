package errors

import "github.com/irdaislakhuafa/GoAttendEasy/src/utils/errors/codes"

type Error struct {
	code    codes.Code
	message string
}

func NewWithCode(code codes.Code, message string, args ...any) error {
	return &Error{code: code, message: message}
}

func (e *Error) Error() string {
	return e.message
}

func GetCode(err error) codes.Code {
	if e, isOk := err.(*Error); !isOk {
		return codes.NoCode
	} else {
		return e.code
	}
}
