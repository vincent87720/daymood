package usecases

import (
	"database/sql"

	"github.com/vincent87720/daymood/app/internal/model"
)

type PurchaseDetail struct {
	Model *model.PurchaseDetail
}

func NewPurchaseDetail(purchaseDetailModel *model.PurchaseDetail) *PurchaseDetail {
	return &PurchaseDetail{
		Model: purchaseDetailModel,
	}
}

func (purchaseDetail *PurchaseDetail) Read(db *sql.DB) ([]interface{}, *model.ModelError) {

	purchaseDetailXi, modelErr := model.Read(purchaseDetail.Model, db)
	if modelErr != nil {
		return nil, modelErr
	}
	return purchaseDetailXi, nil
}

func (purchaseDetail *PurchaseDetail) ReadAll(db *sql.DB) ([]interface{}, *model.ModelError) {

	purchaseDetailXi, modelErr := model.ReadAll(purchaseDetail.Model, db)
	if modelErr != nil {
		return nil, modelErr
	}
	return purchaseDetailXi, nil
}

func (purchaseDetail *PurchaseDetail) Create(db *sql.DB) *model.ModelError {

	modelErr := model.Create(purchaseDetail.Model, db)
	if modelErr != nil {
		return modelErr
	}
	return nil
}

func (purchaseDetail *PurchaseDetail) CreateMultiple(db *sql.DB, purchaseDetailXi interface{}) *model.ModelError {
	detailXi, ok := purchaseDetailXi.([]model.PurchaseDetail)
	if !ok {
		return nil
	}

	var purchaseDetailModel model.PurchaseDetail

	modelErr := purchaseDetailModel.CreateMultiple(db, detailXi)
	if modelErr != nil {
		return modelErr
	}
	return nil
}

func (purchaseDetail *PurchaseDetail) Update(db *sql.DB) *model.ModelError {

	modelErr := model.Update(purchaseDetail.Model, db)
	if modelErr != nil {
		return modelErr
	}
	return nil
}

func (purchaseDetail *PurchaseDetail) Delete(db *sql.DB) *model.ModelError {

	modelErr := model.Delete(purchaseDetail.Model, db)
	if modelErr != nil {
		return modelErr
	}
	return nil
}
