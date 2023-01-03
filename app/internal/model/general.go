package model

type ModelError struct {
	Model   string
	Code    int
	Message string
}

var normalError = func(model string, err error) (modelErr *ModelError) {
	return &ModelError{Model: model, Code: 0, Message: err.Error()}
}
var connectionError = func(model string) (modelErr *ModelError) {
	return &ModelError{Model: model, Code: 1, Message: "Connection to the Database is lost."}
}
var rowsAffectError = func(model string) (modelErr *ModelError) {
	return &ModelError{Model: model, Code: 2, Message: "RowsAffected incorrect."}
}
var uniqueError = func(model string, varName string) (modelErr *ModelError) {
	return &ModelError{Model: model, Code: 3, Message: varName + " must be unique."}
}
var transactionError = func(model string) (modelErr *ModelError) {
	return &ModelError{Model: model, Code: 4, Message: "Transaction fail."}
}
