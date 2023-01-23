package usecases

import (
	"database/sql"

	"github.com/vincent87720/daymood/app/internal/model"
)

type Supplier struct {
	Model *model.Supplier
}

func NewSupplier(supplierModel *model.Supplier) *Supplier {
	return &Supplier{
		Model: supplierModel,
	}
}

func (supplier *Supplier) Read(db *sql.DB) ([]interface{}, *model.ModelError) {

	supplierXi, modelErr := model.Read(supplier.Model, db)
	if modelErr != nil {
		return nil, modelErr
	}
	return supplierXi, nil
}

func (supplier *Supplier) ReadAll(db *sql.DB) ([]interface{}, *model.ModelError) {

	supplierXi, modelErr := model.ReadAll(supplier.Model, db)
	if modelErr != nil {
		return nil, modelErr
	}
	return supplierXi, nil
}

func (supplier *Supplier) Create(db *sql.DB) *model.ModelError {

	//DataStatus預設為使用中
	supplier.Model.DataStatus = 1

	modelErr := model.Create(supplier.Model, db)
	if modelErr != nil {
		return modelErr
	}
	return nil
}

func (supplier *Supplier) CreateMultiple(db *sql.DB, supplierXi interface{}) *model.ModelError {
	return nil
}

func (supplier *Supplier) Update(db *sql.DB) *model.ModelError {

	modelErr := model.Update(supplier.Model, db)
	if modelErr != nil {
		return modelErr
	}
	return nil
}

func (supplier *Supplier) Delete(db *sql.DB) *model.ModelError {

	modelErr := model.Delete(supplier.Model, db)
	if modelErr != nil {
		return modelErr
	}
	return nil
}
