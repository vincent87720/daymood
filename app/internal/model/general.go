package model

import (
	"database/sql"
	"fmt"
)

type Model interface {
	Create(db *sql.DB) *ModelError
	Read(db *sql.DB) ([]interface{}, *ModelError)
	ReadAll(db *sql.DB) ([]interface{}, *ModelError)
	Update(db *sql.DB) *ModelError
	Delete(db *sql.DB) *ModelError
}

func Create(model Model, db *sql.DB) *ModelError {
	return model.Create(db)
}

func Read(model Model, db *sql.DB) ([]interface{}, *ModelError) {
	return model.Read(db)
}

func ReadAll(model Model, db *sql.DB) ([]interface{}, *ModelError) {
	return model.ReadAll(db)
}

func Update(model Model, db *sql.DB) *ModelError {
	return model.Update(db)
}

func Delete(model Model, db *sql.DB) *ModelError {
	return model.Delete(db)
}

type ModelError struct {
	Model   string
	Code    int
	Message string
}

func (modelError *ModelError) Error() string {
	return fmt.Sprintf("[Model] %v | Code %v | %v", modelError.Model, modelError.Code, modelError.Message)
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
