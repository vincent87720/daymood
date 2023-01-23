package usecases

import (
	"database/sql"

	"github.com/vincent87720/daymood/app/internal/model"
)

type Product struct {
	Model *model.Product
}

func NewProduct(productModel *model.Product) *Product {
	return &Product{
		Model: productModel,
	}
}

func (product *Product) Read(db *sql.DB) ([]interface{}, *model.ModelError) {

	productXi, modelErr := model.Read(product.Model, db)
	if modelErr != nil {
		return nil, modelErr
	}
	return productXi, nil
}

func (product *Product) ReadAll(db *sql.DB) ([]interface{}, *model.ModelError) {

	productXi, modelErr := model.ReadAll(product.Model, db)
	if modelErr != nil {
		return nil, modelErr
	}
	return productXi, nil
}

func (product *Product) ReadDeliveryHistories(db *sql.DB, productID int64) ([]model.ProductDeliveryHistory, *model.ModelError) {

	historyModel := &model.ProductDeliveryHistory{}
	historyXi, modelErr := historyModel.ReadAll(db, productID)
	if modelErr != nil {
		return nil, modelErr
	}
	return historyXi, nil
}

func (product *Product) ReadPurchaseHistories(db *sql.DB, productID int64) ([]model.ProductPurchaseHistory, *model.ModelError) {

	historyModel := &model.ProductPurchaseHistory{}
	historyXi, modelErr := historyModel.ReadAll(db, productID)
	if modelErr != nil {
		return nil, modelErr
	}
	return historyXi, nil
}

func (product *Product) Create(db *sql.DB) *model.ModelError {
	product.Model.DataStatus = 1
	modelErr := model.Create(product.Model, db)
	if modelErr != nil {
		return modelErr
	}
	return nil
}

func (product *Product) CreateMultiple(db *sql.DB, productXi interface{}) *model.ModelError {
	productSlice, ok := productXi.([]model.Product)
	if !ok {
		return nil
	}

	for _, val := range productSlice {
		val.DataStatus = 1
	}

	var productModel model.Product

	modelErr := productModel.CreateMultiple(db, productSlice)
	if modelErr != nil {
		return modelErr
	}
	return nil
}

func (product *Product) Update(db *sql.DB) *model.ModelError {

	modelErr := model.Update(product.Model, db)
	if modelErr != nil {
		return modelErr
	}
	return nil
}

func (product *Product) Delete(db *sql.DB) *model.ModelError {

	modelErr := model.Delete(product.Model, db)
	if modelErr != nil {
		return modelErr
	}
	return nil
}
