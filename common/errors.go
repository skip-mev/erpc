package common

import "fmt"

//
// Base Types
//

type BaseError struct {
	Code    string                 `json:"code"`
	Message string                 `json:"message"`
	Cause   error                  `json:"cause"`
	Details map[string]interface{} `json:"details"`
}

func (e *BaseError) Unwrap() error {
	return e.Cause
}

func (e *BaseError) Error() string {
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

func (e *BaseError) CodeChain() string {
	if e.Cause != nil {
		if be, ok := e.Cause.(*BaseError); ok {
			return fmt.Sprintf("%s <- %s", e.Code, be.CodeChain())
		}
	}

	return e.Code
}

type ErrorWithStatusCode interface {
	ErrorStatusCode() int
}

type ErrorWithBody interface {
	ErrorResponseBody() interface{}
}

type RetryableError interface {
	RetryAfter() int
}

//
// Common Errors
//

type ErrProjectNotFound struct {
	BaseError
	ProjectId string
}

func (e *ErrProjectNotFound) Error() string {
	return fmt.Sprintf("project not found: %s", e.ProjectId)
}
