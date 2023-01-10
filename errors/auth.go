package errors

var (
	EMPTY_TOKEN        = NewError(10001, "empty token")
	INVALID_TOKEN      = NewError(10002, "invalid token")
	WRONG_TOKEN_FORMAT = NewError(10003, "incorrect token format")
)
