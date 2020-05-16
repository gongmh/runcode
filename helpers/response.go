package helpers

type Error struct {
	ErrNo  int
	ErrMsg string
}

func (err Error) Error() string {
	return err.ErrMsg
}

var Success = Error{0, "成功"}
var CommonError = Error{1000, "%s"}
var ParamsError = Error{1001, "参数错误"}
var CodeTypeError = Error{1002, "代码类型不支持"}
var ShareIDError = Error{1003, "分享链接不正确，请确认~"}

func GetResponseJson(err error) map[string]interface{} {
	errNo := -1
	errMsg := err.Error()
	if _, ok := err.(Error); ok {
		errNo = err.(Error).ErrNo
		errMsg = err.(Error).ErrMsg
	}

	resp := map[string]interface{}{
		"errno":  errNo,
		"errmsg": errMsg,
	}

	return resp
}
