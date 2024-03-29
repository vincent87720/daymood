package model

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/lib/pq"
)

type DeliveryOrderDetail struct {
	ID              int64   //流水號
	RetailPrice     float32 //出貨時售價
	Cost            float32 //出貨時成本
	QTY             int64   //數量
	Subtotal        float32 //小計
	Remark          *string //備註
	DataOrder       *int64  //順序
	CreateAt        string  //建立時間
	UpdateAt        string  //最後編輯時間
	DeliveryOrderID int64   //出貨編號
	ProductID       *int64  //商品編號
}

func (deliveryOrderDetail *DeliveryOrderDetail) ReadAll(db *sql.DB) (deliveryOrderDetailXi []interface{}, modelErr *ModelError) {
	err := db.Ping()
	if err != nil {
		return nil, normalError("deliveryOrderDetails", err)
	}

	row, err := db.Query("SELECT * FROM deliveryOrderDetails ORDER BY id DESC;")
	if err != nil {
		return nil, normalError("deliveryOrderDetails", err)
	}
	defer row.Close()

	var deliveryOrderDetailRow DeliveryOrderDetail
	for row.Next() {
		err := row.Scan(
			&deliveryOrderDetailRow.ID, &deliveryOrderDetailRow.RetailPrice, &deliveryOrderDetailRow.Cost, &deliveryOrderDetailRow.QTY,
			&deliveryOrderDetailRow.Subtotal, &deliveryOrderDetailRow.Remark, &deliveryOrderDetailRow.DataOrder,
			&deliveryOrderDetailRow.CreateAt, &deliveryOrderDetailRow.UpdateAt, &deliveryOrderDetailRow.DeliveryOrderID,
			&deliveryOrderDetailRow.ProductID,
		)
		if err != nil {
			return nil, normalError("deliveryOrderDetails", err)
		}

		deliveryOrderDetailXi = append(deliveryOrderDetailXi, deliveryOrderDetailRow)
	}

	return deliveryOrderDetailXi, nil
}

func (deliveryOrderDetail *DeliveryOrderDetail) Read(db *sql.DB) (deliveryOrderDetailXi []interface{}, modelErr *ModelError) {
	err := db.Ping()
	if err != nil {
		return nil, normalError("deliveryOrderDetails", err)
	}

	row, err := db.Query("SELECT * FROM deliveryOrderDetails WHERE delivery_order_id = $1 ORDER BY id DESC;", deliveryOrderDetail.DeliveryOrderID)
	if err != nil {
		return nil, normalError("deliveryOrderDetails", err)
	}
	defer row.Close()

	var deliveryOrderDetailRow DeliveryOrderDetail
	for row.Next() {
		err := row.Scan(
			&deliveryOrderDetailRow.ID, &deliveryOrderDetailRow.RetailPrice, &deliveryOrderDetailRow.Cost, &deliveryOrderDetailRow.QTY,
			&deliveryOrderDetailRow.Subtotal, &deliveryOrderDetailRow.Remark, &deliveryOrderDetailRow.DataOrder,
			&deliveryOrderDetailRow.CreateAt, &deliveryOrderDetailRow.UpdateAt, &deliveryOrderDetailRow.DeliveryOrderID,
			&deliveryOrderDetailRow.ProductID,
		)
		if err != nil {
			return nil, normalError("deliveryOrderDetails", err)
		}

		deliveryOrderDetailXi = append(deliveryOrderDetailXi, deliveryOrderDetailRow)
	}

	return deliveryOrderDetailXi, nil
}

func (deliveryOrderDetail *DeliveryOrderDetail) Create(db *sql.DB) (modelErr *ModelError) {
	err := db.Ping()
	if err != nil {
		return normalError("deliveryOrderDetails", err)
	}

	qryString := `INSERT INTO deliveryOrderDetails(
		retail_price, cost, qty, subtotal, remark, 
		data_order, delivery_order_id, product_id
	) VALUES($1,$2,$3,$4,$5,$6,$7,$8);`

	stmt, err := db.Prepare(qryString)
	if err != nil {
		return normalError("deliveryOrderDetails", err)
	}
	defer stmt.Close()

	res, err := stmt.Exec(
		deliveryOrderDetail.RetailPrice, deliveryOrderDetail.Cost, deliveryOrderDetail.QTY, deliveryOrderDetail.Subtotal,
		deliveryOrderDetail.Remark, deliveryOrderDetail.DataOrder, deliveryOrderDetail.DeliveryOrderID,
		deliveryOrderDetail.ProductID,
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
		retail_price, cost, qty, subtotal, remark, data_order,
		delivery_order_id, product_id
	) VALUES($1,$2,$3,$4,$5,$6,$7,$8);`

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
			deliveryOrderDetail.RetailPrice, deliveryOrderDetail.Cost, deliveryOrderDetail.QTY, deliveryOrderDetail.Subtotal,
			deliveryOrderDetail.Remark, deliveryOrderDetail.DataOrder, deliveryOrderDetail.DeliveryOrderID,
			deliveryOrderDetail.ProductID,
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
		deliveryOrderDetail.ID, deliveryOrderDetail.RetailPrice, deliveryOrderDetail.Cost, deliveryOrderDetail.QTY, deliveryOrderDetail.Subtotal,
		deliveryOrderDetail.Remark, deliveryOrderDetail.DataOrder, deliveryOrderDetail.DeliveryOrderID,
		deliveryOrderDetail.ProductID,
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
