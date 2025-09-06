package appcodes

import "fmt"

// ///////////////////////////////////////////////
// register app error
// ///////////////////////////////////////////////
var AppErrorPool = map[int]*AppError{}

// RegisterAppError is an exported function that returns an instance of appError
func RegisterAppError(code int, message string) *AppError {
	if _, found := AppErrorPool[code]; found {
		panic(fmt.Sprintf("Duplicated app error code %v", code))
	}

	appErr := NewAppError(code, message)
	AppErrorPool[code] = appErr
	return appErr
}
