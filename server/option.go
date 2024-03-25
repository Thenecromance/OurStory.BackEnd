package server

//func(httpMethod, absolutePath, handlerName string, nuHandlers int)

type ServerOption struct {
}
type Option func(*ServerOption)
