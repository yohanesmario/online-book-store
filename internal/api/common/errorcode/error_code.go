package errorcode

type ErrorCode string

const (
	UnknownError ErrorCode = "UNKNOWN_ERROR"
	InvalidInput ErrorCode = "INVALID_INPUT"
	NotFound     ErrorCode = "NOT_FOUND"
	Unauthorized ErrorCode = "UNAUTHORIZED"
)
