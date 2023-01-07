package serializer

const (
	CodeParamError    = 40001
	CodeRecordRepeat  = 40002
	CodeDBError       = 50001
	CodePasswordError = 50002
)

type AppError struct {
	Code     int
	Msg      string
	RawError error
}

// DBErr 数据库操作失败
func DBErr(msg string, err error) Response {
	if msg == "" {
		msg = "Database operation failed."
	}
	return Err(CodeDBError, msg, err)
}

// ParamErr 各种参数错误
func ParamErr(msg string, err error) Response {
	if msg == "" {
		msg = "Invalid parameters."
	}
	return Err(CodeParamError, msg, err)
}

// Err 通用错误处理
func Err(errCode int, msg string, err error) Response {
	return Response{
		Code:  errCode,
		Msg:   msg,
		Error: err,
	}
}
