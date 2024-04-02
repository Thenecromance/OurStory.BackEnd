package server

//func(httpMethod, absolutePath, handlerName string, nuHandlers int)

type CoreOption struct {
	CertFile string
	KeyFile  string
}

type Option func(*CoreOption)

func RunningWithCA(certFile string, keyFile string) Option {
	return func(o *CoreOption) {
		o.CertFile = certFile
		o.KeyFile = keyFile
	}
}
