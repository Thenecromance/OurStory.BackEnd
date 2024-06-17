package response

import "github.com/gin-gonic/gin"

func Reset(r *Response) {
	r.Meta.Count = 0
	r.Data = nil
}

func Send(g *gin.Context, r *Response) {
	g.JSON(r.Code, r)
}

func New() *Response {
	return &Response{
		Code: NotFound,
		Meta: meta{
			Count: 0,
		},
		Data: nil,
	}
}
