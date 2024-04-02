package server

//func(httpMethod, absolutePath, handlerName string, nuHandlers int)

type ServerOption struct {
	CertFile string
	KeyFile  string

	TLS bool
}
type Option func(*ServerOption)

func RunningWithCA(certFile string, keyFile string) Option {
	return func(o *ServerOption) {
		o.CertFile = certFile
		o.KeyFile = keyFile
		o.TLS = true
	}
}
