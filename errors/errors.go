package errors

import (
	"errors"
	"fmt"
)

// Code — машинно‑читаемый код ошибки для контрактов между сервисами.
// По смыслу должен быть стабильным и не зависеть от текста сообщения.
type Code string

const (
	// Общие коды (можно расширять по мере необходимости в сервисах).
	CodeUnknown            Code = "UNKNOWN"
	CodeInvalidArgument    Code = "INVALID_ARGUMENT"
	CodeNotFound           Code = "NOT_FOUND"
	CodeUnauthenticated    Code = "UNAUTHENTICATED"
	CodePermissionDenied   Code = "PERMISSION_DENIED"
	CodeAlreadyExists      Code = "ALREADY_EXISTS"
	CodeFailedPrecondition Code = "FAILED_PRECONDITION"
	CodeInternal           Code = "INTERNAL"
)

// Error — общая обёртка для ошибок, которые ходят между сервисами.
// Внутри сервиса поверх неё можно навешивать свои sentinel‑ошибки,
// но снаружи (через gRPC/HTTP) обычно достаточно кода и сообщения.
type Error struct {
	Code    Code
	Message string
}

func (e *Error) Error() string {
	if e == nil {
		return ""
	}
	if e.Code == "" {
		return e.Message
	}
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

// New создаёт новую ошибку с кодом и сообщением.
func New(code Code, message string) *Error {
	return &Error{Code: code, Message: message}
}

// Wrap оборачивает существующую ошибку в Error, если это ещё не Error.
// Если err == nil, возвращает nil.
func Wrap(code Code, message string, err error) error {
	if err == nil {
		return nil
	}
	var e *Error
	if errors.As(err, &e) {
		// Уже Error — не теряем исходный код, но добавляем контекст в Message.
		if message == "" {
			return e
		}
		return &Error{
			Code:    e.Code,
			Message: message + ": " + e.Message,
		}
	}
	if message == "" {
		return &Error{Code: code, Message: err.Error()}
	}
	return &Error{Code: code, Message: message + ": " + err.Error()}
}

// CodeOf возвращает код ошибки, если err совместим с *Error.
func CodeOf(err error) (Code, bool) {
	var e *Error
	if errors.As(err, &e) {
		return e.Code, true
	}
	return "", false
}

// IsCode проверяет, что у ошибки задан конкретный код.
func IsCode(err error, code Code) bool {
	if err == nil {
		return false
	}
	c, ok := CodeOf(err)
	return ok && c == code
}
