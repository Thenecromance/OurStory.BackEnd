package response

type meta struct {
	Count int `json:"count"`
}

type Response struct {
	//ginCtx `json:"-"`
	Code int         `json:"code"`
	Meta meta        `json:"meta"`
	Data interface{} `json:"data"`
}
