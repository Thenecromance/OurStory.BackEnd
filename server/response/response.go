package response

func Reset(r *Response) {
	r.Meta.Count = 0
	r.Data = nil
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
