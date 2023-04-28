package errors

import "fmt"

type (
	Error struct {
		code    string
		message string
	}
)

func New(code, message string) *Error {
	e := &Error{
		code:    code,
		message: message,
	}
	return e
}

func (e Error) Error() string {
	return fmt.Sprintf("code:%s message:%s", e.code, e.message)
}

func (e Error) Code() string {
	return e.code
}

func (e Error) Message() string {
	return e.message
}
