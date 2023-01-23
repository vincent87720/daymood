package usecases

import (
	"database/sql"

	"github.com/vincent87720/daymood/app/internal/model"
)

type Usecase interface {
	Read(db *sql.DB) ([]interface{}, *model.ModelError)
	ReadAll(db *sql.DB) ([]interface{}, *model.ModelError)
	Create(db *sql.DB) *model.ModelError
	CreateMultiple(db *sql.DB, model interface{}) *model.ModelError
	Update(db *sql.DB) *model.ModelError
	Delete(db *sql.DB) *model.ModelError
}

func ReadAll(usecase Usecase, db *sql.DB) ([]interface{}, *model.ModelError) {
	return usecase.ReadAll(db)
}

func Read(usecase Usecase, db *sql.DB) ([]interface{}, *model.ModelError) {
	return usecase.Read(db)
}

func CreateMultiple(usecase Usecase, db *sql.DB, model interface{}) *model.ModelError {
	return usecase.CreateMultiple(db, model)
}

func Create(usecase Usecase, db *sql.DB) *model.ModelError {
	return usecase.Create(db)
}

func Update(usecase Usecase, db *sql.DB) *model.ModelError {
	return usecase.Update(db)
}

func Delete(usecase Usecase, db *sql.DB) *model.ModelError {
	return usecase.Delete(db)
}
