package server

type tls struct {
	certFile string
	keyFile  string
}

func newTls(certFile, keyFile string) *tls {
	return &tls{
		certFile: certFile,
		keyFile:  keyFile,
	}
}
