package defs

type Err struct {
	Error string `json:"error"`
	ErrorCode string `json:"error_code"`
}

type ErrResponse struct {
	HttpSC int
	Erros Err
}

var (
	ErrorRequestBodyParseFailed = ErrResponse{HttpSC: 400, Error:Err{Error:"Request faild"}}
	ErrorNotAuthUser = ErrResponse{HttpSC: 401, Error:Err{Error:"User authentical faild"}}
)