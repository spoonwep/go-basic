package errors

type BaseError struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

var (
	NORMAL        = NewError(200, "success")
	NO_PERMISSION = NewError(403, "No Permission")
	SERVER_ERROR  = NewError(500, "Server Error")
)

func (e *BaseError) Error() string {
	return e.Msg
}

func NewError(code int, msg string) *BaseError {
	return &BaseError{
		Msg:  msg,
		Code: code,
	}
}

func GetError(e *BaseError, data interface{}) *BaseError {
	return &BaseError{
		Msg:  e.Msg,
		Code: e.Code,
		Data: data,
	}
}
