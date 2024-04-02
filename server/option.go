package server

//func(httpMethod, absolutePath, handlerName string, nuHandlers int)

type ServerOption struct {
	CertFile string
	KeyFile  string

	TLS bool
}
type Option func(*ServerOption)
