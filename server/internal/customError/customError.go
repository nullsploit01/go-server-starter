package customError

import "fmt"

type ValidationError struct {
	Field   string
	Message string
}

type WrongCredentialsError struct {
	Message string
}

type UnauthorizedError struct {
	Message string
}

type ForbiddenError struct {
	Message string
}

type NotFoundError struct {
	Message string
}

func (e NotFoundError) Error() string {
	return e.Message
}

func (e ForbiddenError) Error() string {
	return e.Message
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("validation failed for field %s: %s", e.Field, e.Message)
}

func (e WrongCredentialsError) Error() string {
	return e.Message
}

func (e UnauthorizedError) Error() string {
	return e.Message
}
