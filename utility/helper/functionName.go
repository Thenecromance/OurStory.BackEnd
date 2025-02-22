package helper

import (
	"reflect"
	"runtime"
	"strings"
)

// CurrentFunctionName returns the name of the function that called it
// Example: helper.CurrentFunctionName() -> CurrentFunctionName
func CurrentFunctionName() string {
	pc := make([]uintptr, 1)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	return f.Name()
}

// GetFunctionName returns the name of the function
// Example: helper.GetFunctionName(helper.GetFunctionName) -> GetFunctionName
func GetFunctionName(fn interface{} /*, seps ...rune*/) string {
	fullName := runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name()
	parts := strings.Split(fullName, ".")
	name := parts[len(parts)-1]

	//end with -fm
	if strings.HasSuffix(name, "-fm") {
		name = name[:len(name)-3]
	}

	return name
}

// GetPackageName returns the name of the package
// Example: helper.GetPackageName(helper.GetPackageName) -> helper
func GetPackageName(fn interface{}) string {

	fullName := runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name()
	parts := strings.Split(fullName, ".")
	fullPath := parts[len(parts)-2]
	parts = strings.Split(fullPath, "/")
	return parts[len(parts)-1]
}
