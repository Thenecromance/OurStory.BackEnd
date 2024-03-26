package api

const (
	Success = iota
	Failed
)

type Response struct {
	Code int `json:"code"`

	// the real data that will be sent to client
	Result interface{} `json:"result"`
}
