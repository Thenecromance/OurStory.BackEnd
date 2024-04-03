package response

import "github.com/gin-gonic/gin"

const (
	SUCCESS = iota
	FAIL
	Unauthorized
	NotFound
)

type GinCtx struct {
	statuCode int          `json:"-"`
	ctx       *gin.Context `json:"-"`
}
type meta struct {
	Count int `json:"count"`
}
type Response struct {
	GinCtx `json:"-"`
	Code   int         `json:"code"`
	Meta   meta        `json:"meta"`
	Data   interface{} `json:"data"`
}

func (r *Response) Clear() *Response {
	r.Meta.Count = 0
	r.Data = nil
	return r
}
func (r *Response) AddData(d interface{}) *Response {
	if r.Data == nil || r.Meta.Count == 0 {
		r.Data = d
		r.Meta.Count = 1
	} else {
		r.Data = append(r.Data.([]interface{}), d)
		r.Meta.Count++
	}

	return r
}

func (r *Response) SetCode(code int) *Response {
	r.Code = code
	return r
}

func (r *Response) Send() {
	r.ctx.JSON(r.statuCode, r)
}

func New(c *gin.Context) *Response {
	return &Response{
		GinCtx: GinCtx{
			statuCode: 200,
			ctx:       c,
		},
		Code: FAIL,
		Meta: meta{
			Count: 0,
		},
		Data: nil,
	}
}
