package response

func New() *Response {
	return &Response{
		Code: NotFound,
		Meta: meta{
			Count: 0,
		},
		Data: nil,
	}
}
