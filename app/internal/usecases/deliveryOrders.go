package usecases

import (
	"database/sql"

	"github.com/vincent87720/daymood/app/internal/model"
)

type DeliveryOrder struct {
	Model *model.DeliveryOrder
}

func NewDeliveryOrder(deliveryOrderModel *model.DeliveryOrder) *DeliveryOrder {
	return &DeliveryOrder{
		Model: deliveryOrderModel,
	}
}

func (deliveryOrder *DeliveryOrder) Read(db *sql.DB) ([]interface{}, *model.ModelError) {

	deliveryOrderXi, modelErr := model.Read(deliveryOrder.Model, db)
	if modelErr != nil {
		return nil, modelErr
	}
	return deliveryOrderXi, nil
}

func (deliveryOrder *DeliveryOrder) ReadAll(db *sql.DB) ([]interface{}, *model.ModelError) {

	deliveryOrderXi, modelErr := model.ReadAll(deliveryOrder.Model, db)
	if modelErr != nil {
		return nil, modelErr
	}
	return deliveryOrderXi, nil
}

func (deliveryOrder *DeliveryOrder) Create(db *sql.DB) *model.ModelError {

	modelErr := model.Create(deliveryOrder.Model, db)
	if modelErr != nil {
		return modelErr
	}
	return nil
}

func (deliveryOrder *DeliveryOrder) CreateMultiple(db *sql.DB, deliveryOrderXi interface{}) *model.ModelError {
	return nil
}

func (deliveryOrder *DeliveryOrder) Update(db *sql.DB) *model.ModelError {

	modelErr := model.Update(deliveryOrder.Model, db)
	if modelErr != nil {
		return modelErr
	}
	return nil
}

func (deliveryOrder *DeliveryOrder) Delete(db *sql.DB) *model.ModelError {

	modelErr := model.Delete(deliveryOrder.Model, db)
	if modelErr != nil {
		return modelErr
	}
	return nil
}
