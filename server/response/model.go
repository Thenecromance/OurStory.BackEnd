package response

/*
	type ginCtx struct {
		statuCode int          `json:"-"`
		ctx       *gin.Context `json:"-"`
	}
*/
type meta struct {
	TraceId string `json:"trace_id"`
	Count   int    `json:"count"`
}

type Response struct {
	//ginCtx `json:"-"`
	Code int         `json:"code"`
	Meta meta        `json:"meta"`
	Data interface{} `json:"data"`
}
