package errorlib

import (
	"encoding/json"
	"fmt"
)

type Err struct {
	errorType Type
	code      string
	errors    []error
	message   string
	detail    map[string]interface{}
}

type Error interface {
	json.Marshaler
	error
	Type() Type
	WithError(error) Error
	WithMessage(string) Error
	WithDetail(string, interface{}) Error
}

type Type string

const (
	TypeInfo     Type = "INFO"
	TypeWarn     Type = "WARNING"
	TypeError    Type = "ERROR"
	TypeCritical Type = "CRITICAL"
)

func NewError(errType Type, code string) Error {
	return Err{
		errorType: errType,
		code:      code,
	}
}

func (e Err) Type() Type {
	return e.errorType
}

func (e Err) WithError(err error) Error {
	errCast, ok := err.(Err)
	if ok {
		e.errors = append(e.errors, errCast)
	} else {
		e.errors = append(e.errors, Err{
			code:    "NATIVE_ERROR",
			message: err.Error(),
		})
	}
	return e
}

func (e Err) WithMessage(message string) Error {
	e.message = message
	return e
}

func (e Err) WithDetail(detailName string, detailData interface{}) Error {
	e.detail[detailName] = detailData
	return e
}

func (e Err) Error() string {
	if e.message != "" {
		return fmt.Sprintf("[%s] %s : %s", e.errorType, e.code, e.message)
	}
	return fmt.Sprintf("[%s] %s", e.errorType, e.code)
}

func (e Err) MarshalJSON() ([]byte, error) {
	err := struct {
		Code    string                 `json:"code"`
		Type    Type                   `json:"type"`
		Errors  []error                `json:"errors"`
		Message string                 `json:"message"`
		Detail  map[string]interface{} `json:"detail"`
	}{
		Code:    e.code,
		Type:    e.errorType,
		Errors:  e.errors,
		Message: e.message,
		Detail:  e.detail,
	}
	return json.Marshal(err)
}
