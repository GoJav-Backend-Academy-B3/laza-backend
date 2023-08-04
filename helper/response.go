package helper

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
	Code        int    `json:"-"`
	Status      string `json:"status"`
	IsError     bool   `json:"isError"`
	Data        any    `json:"data,omitempty"`
	Description any    `json:"description,omitempty"`
}

func (res *Response) Send(ctx *gin.Context) {

	ctx.Header("Content-type", "application/json")
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.Header("Access-Control-Allow-Methods", "OPTIONS, GET, POST, PUT, DELETE")
	ctx.Header("Access-Control-Allow-Headers", "Content-Type, X-CSRF-Token")

	if res.IsError {
		ctx.Status(res.Code)
	}

	ctx.JSON(res.Code, res)

}

func GetResponse(data any, code int, isError bool) *Response {

	if isError {
		return &Response{
			Code:        code,
			Status:      getStatus(code),
			IsError:     isError,
			Description: data,
		}

	}
	return &Response{
		Code:    code,
		Status:  getStatus(code),
		IsError: isError,
		Data:    data,
	}

}

func getStatus(code int) (desc string) {

	switch code {
	case 200:
		desc = "OK"
	case 201:
		desc = "Created"
	case 202:
		desc = "Accepted"
	case 304:
		desc = "Not Modified"
	case 400:
		desc = "Bad Request"
	case 401:
		desc = "Unauthorized"
	case 403:
		desc = "Forbidden"
	case 404:
		desc = "Not Found"
	case 415:
		desc = "Unsupported Media Type"
	case 500:
		desc = "Internal Server Error"
	case 502:
		desc = "Bad Gateway"
	default:
		desc = "Status Code Undefined"
	}

	return

}
