package model

import (
	"context"
	"database/sql"
	"errors"
)

type DeliveryProduct struct {
	// DeliveryProductID       int `json:"DeliveryProductID,string,omitempty"`
	DeliveryProductName     string
	NtdSellingPrice         float32
	DeliveryProductQty      int
	DeliveryProductSubtotal float32
	ProductSku              string
}

type DeliveryOrder struct {
	DeliveryOrderID     string
	DeliveryType        string
	Discount            float32
	DeliveryOrderTotal  float32
	EasyStoreOrderTotal float32
	Products            []DeliveryProduct
}

func (deliveryOrder *DeliveryOrder) DecreaseStocks(db *sql.DB) error {
	err := db.Ping()
	if err != nil {
		return err
	}

	ctx := context.Background()

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	for _, val := range deliveryOrder.Products {

		stmt, err := tx.Prepare("UPDATE product SET stocks = stocks - $1 WHERE product_sku = $2 AND (stocks - $1) >= 0;")
		if err != nil {
			return err
		}
		defer stmt.Close()

		res, execErr := stmt.Exec(val.DeliveryProductQty, val.ProductSku)
		rowsAff, err := res.RowsAffected()
		if execErr != nil || err != nil || rowsAff != 1 {
			_ = tx.Rollback()
			return errors.New("Transaction fail.")
		}
	}
	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}
