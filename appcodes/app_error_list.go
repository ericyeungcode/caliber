package appcodes

var OK = RegisterAppError(0, "OK")
var GeneralError = RegisterAppError(1001, "General error")
var InvalidArgumentError = RegisterAppError(1002, "Invalid arguments")
var NotFoundError = RegisterAppError(1003, "Not Found")
var InternalError = RegisterAppError(1100, "General error")
