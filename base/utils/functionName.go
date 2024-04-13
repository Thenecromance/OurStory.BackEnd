package utils

import (
	"reflect"
	"runtime"
	"strings"
)

// CurrentFunctionName returns the name of the function that called it
func CurrentFunctionName() string {
	pc := make([]uintptr, 1)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	return f.Name()
}

// GetFunctionName returns the name of the function
func GetFunctionName(fn interface{} /*, seps ...rune*/) string {
	fullName := runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name()
	parts := strings.Split(fullName, ".")
	name := parts[len(parts)-1]

	//end with -fm
	if strings.HasSuffix(name, "-fm") {
		name = name[:len(name)-3]
	}

	/*	if len(seps) > 0 {
		return strings.Map(func(r rune) rune {
			for _, sep := range seps {
				if r == sep {
					return sep
				}
			}
			return -1
		}, name)
	}*/
	return name
}

// GetPackageName returns the name of the package
func GetPackageName(fn interface{}) string {

	fullName := runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name()
	parts := strings.Split(fullName, ".")
	fullPath := parts[len(parts)-2]
	parts = strings.Split(fullPath, "/")
	return parts[len(parts)-1]
}
