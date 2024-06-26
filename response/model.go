package response

import (
	"github.com/gin-gonic/gin"
	"sync"
)

type meta struct {
	Count int `json:"count"`
}

type Response struct {
	//ginCtx `json:"-"`
	Code int         `json:"code"`
	Meta meta        `json:"meta"`
	Data interface{} `json:"data"`
	mu   sync.Mutex
}

type ErrorResponse struct {
	Message string `json:"err_message"`
}

func (r *Response) Reset() {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.Meta.Count = 0
	r.Data = nil
}
func (r *Response) Send(ctx *gin.Context) {
	ctx.JSON(r.Code, r)
}

func (r *Response) AddData(data interface{}) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.Meta.Count == 0 {
		r.Meta.Count = 1
		r.Data = data
	} else if r.Meta.Count == 1 {
		// when the data has 1 element, we need to convert it to an array
		r.Meta.Count++
		obj := r.Data
		r.Data = []interface{}{obj, data}
	} else {
		r.Meta.Count++
		r.Data = append(r.Data.([]interface{}), data)
	}
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
