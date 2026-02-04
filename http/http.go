package http

// HTTP-методы (для документации, роутинга и проверок в связанных сервисах).
const (
	MethodGet    = "GET"
	MethodPost   = "POST"
	MethodPut    = "PUT"
	MethodPatch  = "PATCH"
	MethodDelete = "DELETE"
)

// Заголовки.
const (
	HeaderAuthorization = "Authorization"
	HeaderContentType   = "Content-Type"
)

// Content-Type.
const (
	ContentTypeJSON = "application/json"
)
