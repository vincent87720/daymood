package usecases

import (
	"database/sql"

	"github.com/vincent87720/daymood/app/internal/model"
)

type Purchase struct {
	Model *model.Purchase
}

func NewPurchase(purchaseModel *model.Purchase) *Purchase {
	return &Purchase{
		Model: purchaseModel,
	}
}

func (purchase *Purchase) Read(db *sql.DB) ([]interface{}, *model.ModelError) {

	purchaseXi, modelErr := model.Read(purchase.Model, db)
	if modelErr != nil {
		return nil, modelErr
	}
	return purchaseXi, nil
}

func (purchase *Purchase) ReadAll(db *sql.DB) ([]interface{}, *model.ModelError) {

	purchaseXi, modelErr := model.ReadAll(purchase.Model, db)
	if modelErr != nil {
		return nil, modelErr
	}
	return purchaseXi, nil
}

func (purchase *Purchase) Create(db *sql.DB) *model.ModelError {

	modelErr := model.Create(purchase.Model, db)
	if modelErr != nil {
		return modelErr
	}
	return nil
}

func (purchase *Purchase) CreateMultiple(db *sql.DB, purchaseXi interface{}) *model.ModelError {
	return nil
}

func (purchase *Purchase) Update(db *sql.DB) *model.ModelError {

	modelErr := model.Update(purchase.Model, db)
	if modelErr != nil {
		return modelErr
	}
	return nil
}

func (purchase *Purchase) Delete(db *sql.DB) *model.ModelError {

	modelErr := model.Delete(purchase.Model, db)
	if modelErr != nil {
		return modelErr
	}
	return nil
}
