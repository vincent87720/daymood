package usecases

import (
	"database/sql"

	"github.com/vincent87720/daymood/app/internal/model"
)

type Discount struct {
	Model *model.Discount
}

func NewDiscount(discountModel *model.Discount) *Discount {
	return &Discount{
		Model: discountModel,
	}
}

func (discount *Discount) Read(db *sql.DB) ([]interface{}, *model.ModelError) {

	discountXi, modelErr := model.Read(discount.Model, db)
	if modelErr != nil {
		return nil, modelErr
	}
	return discountXi, nil
}

func (discount *Discount) ReadAll(db *sql.DB) ([]interface{}, *model.ModelError) {

	discountXi, modelErr := model.ReadAll(discount.Model, db)
	if modelErr != nil {
		return nil, modelErr
	}
	return discountXi, nil
}

func (discount *Discount) Create(db *sql.DB) *model.ModelError {

	modelErr := model.Create(discount.Model, db)
	if modelErr != nil {
		return modelErr
	}
	return nil
}

func (discount *Discount) CreateMultiple(db *sql.DB, discountXi interface{}) *model.ModelError {
	return nil
}

func (discount *Discount) Update(db *sql.DB) *model.ModelError {

	modelErr := model.Update(discount.Model, db)
	if modelErr != nil {
		return modelErr
	}
	return nil
}

func (discount *Discount) Delete(db *sql.DB) *model.ModelError {

	modelErr := model.Delete(discount.Model, db)
	if modelErr != nil {
		return modelErr
	}
	return nil
}
