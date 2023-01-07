package model

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/lib/pq"
)

type DeliveryOrderDetail struct {
	ID               int64   //流水號
	RetailPrice      float32 //出貨時售價
	QTY              int64   //數量
	Subtotal         float32 //小計
	Remark           *string //備註
	DataOrder        *int64  //順序
	CreateAt         string  //建立時間
	UpdateAt         string  //最後編輯時間
	DeliveryOrderID  int64   //出貨編號
	ProductID        *int64  //商品編號
	PurchaseDetailID *int64  //採購明細編號
}

func GetAllDeliveryOrderDetails(db *sql.DB) (deliveryOrderDetailXi []DeliveryOrderDetail, modelErr *ModelError) {
	err := db.Ping()
	if err != nil {
		return nil, normalError("deliveryOrderDetails", err)
	}

	row, err := db.Query("SELECT * FROM deliveryOrderDetails ORDER BY id DESC;")
	if err != nil {
		return nil, normalError("deliveryOrderDetails", err)
	}
	defer row.Close()

	var deliveryOrderDetail DeliveryOrderDetail
	for row.Next() {
		err := row.Scan(
			&deliveryOrderDetail.ID, &deliveryOrderDetail.RetailPrice, &deliveryOrderDetail.QTY,
			&deliveryOrderDetail.Subtotal, &deliveryOrderDetail.Remark, &deliveryOrderDetail.DataOrder,
			&deliveryOrderDetail.CreateAt, &deliveryOrderDetail.UpdateAt, &deliveryOrderDetail.DeliveryOrderID,
			&deliveryOrderDetail.ProductID, &deliveryOrderDetail.PurchaseDetailID,
		)
		if err != nil {
			return nil, normalError("deliveryOrderDetails", err)
		}

		deliveryOrderDetailXi = append(deliveryOrderDetailXi, deliveryOrderDetail)
	}

	return deliveryOrderDetailXi, nil
}

func GetDeliveryOrderDetails(db *sql.DB, deliveryOrderID int64) (deliveryOrderDetailXi []DeliveryOrderDetail, modelErr *ModelError) {
	err := db.Ping()
	if err != nil {
		return nil, normalError("deliveryOrderDetails", err)
	}

	row, err := db.Query("SELECT * FROM deliveryOrderDetails WHERE delivery_order_id = $1 ORDER BY id DESC;", deliveryOrderID)
	if err != nil {
		return nil, normalError("deliveryOrderDetails", err)
	}
	defer row.Close()

	var deliveryOrderDetail DeliveryOrderDetail
	for row.Next() {
		err := row.Scan(
			&deliveryOrderDetail.ID, &deliveryOrderDetail.RetailPrice, &deliveryOrderDetail.QTY,
			&deliveryOrderDetail.Subtotal, &deliveryOrderDetail.Remark, &deliveryOrderDetail.DataOrder,
			&deliveryOrderDetail.CreateAt, &deliveryOrderDetail.UpdateAt, &deliveryOrderDetail.DeliveryOrderID,
			&deliveryOrderDetail.ProductID, &deliveryOrderDetail.PurchaseDetailID,
		)
		if err != nil {
			return nil, normalError("deliveryOrderDetails", err)
		}

		deliveryOrderDetailXi = append(deliveryOrderDetailXi, deliveryOrderDetail)
	}

	return deliveryOrderDetailXi, nil
}

func (deliveryOrderDetail *DeliveryOrderDetail) Create(db *sql.DB) (modelErr *ModelError) {
	err := db.Ping()
	if err != nil {
		return normalError("deliveryOrderDetails", err)
	}

	qryString := `INSERT INTO deliveryOrderDetails(
		retail_price, qty, subtotal,
		remark, data_order, delivery_order_id,
		product_id, purchase_detail_id
	) VALUES($1,$2,$3,$4,$5,$6,$7,$8);`

	stmt, err := db.Prepare(qryString)
	if err != nil {
		return normalError("deliveryOrderDetails", err)
	}
	defer stmt.Close()

	res, err := stmt.Exec(
		deliveryOrderDetail.RetailPrice, deliveryOrderDetail.QTY, deliveryOrderDetail.Subtotal,
		deliveryOrderDetail.Remark, deliveryOrderDetail.DataOrder, deliveryOrderDetail.DeliveryOrderID,
		deliveryOrderDetail.ProductID, deliveryOrderDetail.PurchaseDetailID,
	)
	if err != nil {
		return normalError("deliveryOrderDetails", err)
	}

	rowsAff, err := res.RowsAffected()
	if rowsAff != 1 {
		return rowsAffectError("deliveryOrderDetails")
	}
	return nil
}

func (deliveryOrderDetail *DeliveryOrderDetail) CreateMultiple(db *sql.DB, deliveryOrderDetailXi []DeliveryOrderDetail) (modelErr *ModelError) {
	err := db.Ping()
	if err != nil {
		return connectionError("deliveryOrderDetails")
	}

	qryString := `INSERT INTO deliveryOrderDetails(
		retail_price, qty,
		subtotal, remark, data_order,
		delivery_order_id, product_id, purchase_detail_id,
	) VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13);`

	stmt, err := db.Prepare(qryString)
	if err != nil {
		return normalError("deliveryOrderDetails", err)
	}
	defer stmt.Close()

	ctx := context.Background()

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return normalError("deliveryOrderDetails", err)
	}

	for _, deliveryOrderDetail := range deliveryOrderDetailXi {

		res, err := stmt.Exec(
			deliveryOrderDetail.RetailPrice, deliveryOrderDetail.QTY, deliveryOrderDetail.Subtotal,
			deliveryOrderDetail.Remark, deliveryOrderDetail.DataOrder, deliveryOrderDetail.DeliveryOrderID,
			deliveryOrderDetail.ProductID, deliveryOrderDetail.PurchaseDetailID,
		)
		if err, ok := err.(*pq.Error); ok {
			_ = tx.Rollback()
			return normalError("deliveryOrderDetails", err)
		}
		rowsAff, execErr := res.RowsAffected()
		if execErr != nil || err != nil || rowsAff != 1 {
			_ = tx.Rollback()
			return transactionError("deliveryOrderDetails")
		}
	}
	if err := tx.Commit(); err != nil {
		return normalError("deliveryOrderDetails", err)
	}

	return nil
}

func (deliveryOrderDetail *DeliveryOrderDetail) Update(db *sql.DB) (modelErr *ModelError) {
	err := db.Ping()
	if err != nil {
		return normalError("deliveryOrderDetails", err)
	}

	_, err = db.Exec(
		"CALL updateDeliveryOrderDetails($1,$2,$3,$4,$5,$6,$7,$8,$9)",
		deliveryOrderDetail.ID, deliveryOrderDetail.RetailPrice, deliveryOrderDetail.QTY, deliveryOrderDetail.Subtotal,
		deliveryOrderDetail.Remark, deliveryOrderDetail.DataOrder, deliveryOrderDetail.DeliveryOrderID,
		deliveryOrderDetail.ProductID, deliveryOrderDetail.PurchaseDetailID,
	)
	if err != nil {
		return normalError("deliveryOrderDetails", err)
	}

	return nil
}

func (deliveryOrderDetail *DeliveryOrderDetail) Delete(db *sql.DB) (modelErr *ModelError) {
	err := db.Ping()
	if err != nil {
		return normalError("deliveryOrderDetails", err)
	}

	stmt, err := db.Prepare("DELETE FROM deliveryOrderDetails WHERE id = $1;")
	if err != nil {
		return normalError("deliveryOrderDetails", err)
	}
	defer stmt.Close()

	res, err := stmt.Exec(deliveryOrderDetail.ID)
	if err, ok := err.(*pq.Error); ok {
		fmt.Println(err.Code.Name())
		// return &ModelError{Model: "deliveryOrders", Code: 1, Message: "supplier still have children."}
	}

	rowsAff, err := res.RowsAffected()
	if rowsAff != 1 {
		return rowsAffectError("deliveryOrderDetails")
	}
	return nil
}

// ID
// RetailPrice
// QTY
// Subtotal
// Remark
// DataOrder
// CreateAt
// UpdateAt
// DeliveryOrderID
// ProductID
// PurchaseDetailID

// id
// retail_price
// qty
// subtotal
// remark
// data_order
// create_at
// update_at
// delivery_order_id
// product_id
// purchase_detail_id
