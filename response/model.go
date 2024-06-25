package response

import (
	"github.com/gin-gonic/gin"
)

type meta struct {
	Count int `json:"count"`
}

type Response struct {
	//ginCtx `json:"-"`
	Code int         `json:"code"`
	Meta meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type ErrorResponse struct {
	Message string `json:"err_message"`
}

func (r *Response) Reset() {
	r.Meta.Count = 0
	r.Data = nil
}

func (r *Response) Send(ctx *gin.Context) {
	ctx.JSON(r.Code, r)
}

func (r *Response) AddData(data interface{}) {
	r.Meta.Count++
	r.Data = data
}

func (r *Response) SetCode(code int) *Response {
	r.Code = code
	return r
}

func (r *Response) Error(errMsg string) {
	r.SetCode(BadRequest)

	r.AddData(ErrorResponse{Message: errMsg})
}

func (r *Response) Success(data interface{}) {
	r.SetCode(OK).AddData(data)
}

func (r *Response) NotFound() {
	r.SetCode(NotFound)

}

func (r *Response) Unauthorized(msg string) {
	r.SetCode(Unauthorized).AddData(ErrorResponse{Message: msg})
}
