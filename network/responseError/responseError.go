package responseError

import ()

type ResponseError struct {
	Code 		int			`json:"error_code"`
	Message		string		`json:"error_msg"`
}