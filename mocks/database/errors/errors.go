package errors

import "fmt"

var (
	ErrUserAlreadyExists = fmt.Errorf("user already exists")
	ErrUserNotFound      = fmt.Errorf("user not found")
)
