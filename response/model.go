package response

import "github.com/gin-gonic/gin"

type meta struct {
	Count int `json:"count"`
}

type Response struct {
	//ginCtx `json:"-"`
	Code int         `json:"code"`
	Meta meta        `json:"meta"`
	Data interface{} `json:"data"`
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
