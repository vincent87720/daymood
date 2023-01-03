package model

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/lib/pq"
)

type PurchaseDetail struct {
	ID             int64    //流水號
	NamedID        *string  //採購商品編號
	Name           string   //商品名稱
	Status         int64    //是否採用
	WholesalePrice float32  //批價
	QTY            int64    //數量
	Cost           *float32 //成本
	Currency       *int64   //幣別
	Subtotal       float32  //小計
	Remark         *string  //備註
	DataOrder      *int64   //順序
	CreateAt       string   //建立時間
	UpdateAt       string   //最後編輯時間
	PurchaseID     int64    //採購編號
	SupplierID     *int64   //廠商編號
	ProductID      *int64   //商品編號
}

func NewPurchaseDetail(name string, status int64, qty int64, subtotalTwd float32) (PurchaseDetail, error) {
	var purchaseDetail PurchaseDetail

	if name == "" {
		return purchaseDetail, errors.New("name field should not be empty")
	}

	purchaseDetail = PurchaseDetail{
		Name:   name,
		Status: status,
		QTY:    qty,
	}

	return purchaseDetail, nil
}

func GetAllPurchaseDetails(db *sql.DB) (purchaseDetailXi []PurchaseDetail, modelErr *ModelError) {
	err := db.Ping()
	if err != nil {
		return nil, normalError("purchaseDetails", err)
	}

	row, err := db.Query("SELECT * FROM purchaseDetails ORDER BY id DESC;")
	if err != nil {
		return nil, normalError("purchaseDetails", err)
	}
	defer row.Close()

	var purchaseDetail PurchaseDetail
	for row.Next() {
		err := row.Scan(
			&purchaseDetail.ID, &purchaseDetail.NamedID, &purchaseDetail.Name,
			&purchaseDetail.Status, &purchaseDetail.WholesalePrice, &purchaseDetail.QTY,
			&purchaseDetail.Cost, &purchaseDetail.Currency, &purchaseDetail.Subtotal,
			&purchaseDetail.Remark, &purchaseDetail.DataOrder, &purchaseDetail.CreateAt,
			&purchaseDetail.UpdateAt, &purchaseDetail.PurchaseID, &purchaseDetail.SupplierID,
			&purchaseDetail.ProductID,
		)
		if err != nil {
			return nil, normalError("purchaseDetails", err)
		}

		purchaseDetailXi = append(purchaseDetailXi, purchaseDetail)
	}

	return purchaseDetailXi, nil
}

func GetPurchaseDetails(db *sql.DB, purchaseID int64) (purchaseDetailXi []PurchaseDetail, modelErr *ModelError) {
	err := db.Ping()
	if err != nil {
		return nil, normalError("purchaseDetails", err)
	}

	row, err := db.Query("SELECT * FROM purchaseDetails WHERE purchase_id = $1 ORDER BY id DESC;", purchaseID)
	if err != nil {
		return nil, normalError("purchaseDetails", err)
	}
	defer row.Close()

	var purchaseDetail PurchaseDetail
	for row.Next() {
		err := row.Scan(
			&purchaseDetail.ID, &purchaseDetail.NamedID, &purchaseDetail.Name,
			&purchaseDetail.Status, &purchaseDetail.WholesalePrice, &purchaseDetail.QTY,
			&purchaseDetail.Cost, &purchaseDetail.Currency, &purchaseDetail.Subtotal,
			&purchaseDetail.Remark, &purchaseDetail.DataOrder, &purchaseDetail.CreateAt,
			&purchaseDetail.UpdateAt, &purchaseDetail.PurchaseID, &purchaseDetail.SupplierID,
			&purchaseDetail.ProductID,
		)
		if err != nil {
			return nil, normalError("purchaseDetails", err)
		}

		purchaseDetailXi = append(purchaseDetailXi, purchaseDetail)
	}

	return purchaseDetailXi, nil
}

func (purchaseDetail *PurchaseDetail) Create(db *sql.DB) (modelErr *ModelError) {
	err := db.Ping()
	if err != nil {
		return normalError("purchaseDetails", err)
	}

	qryString := `INSERT INTO purchaseDetails(
		named_id, name, status,
		wholesale_price, qty, cost,
		currency, subtotal, remark,
		data_order, purchase_id, supplier_id,
		product_id
	) VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13);`

	stmt, err := db.Prepare(qryString)
	if err != nil {
		return normalError("purchaseDetails", err)
	}
	defer stmt.Close()

	res, err := stmt.Exec(purchaseDetail.NamedID, purchaseDetail.Name,
		purchaseDetail.Status, purchaseDetail.WholesalePrice, purchaseDetail.QTY,
		purchaseDetail.Cost, purchaseDetail.Currency, purchaseDetail.Subtotal,
		purchaseDetail.Remark, purchaseDetail.DataOrder, purchaseDetail.PurchaseID,
		purchaseDetail.SupplierID, purchaseDetail.ProductID)
	if err != nil {
		return normalError("purchaseDetails", err)
	}

	rowsAff, err := res.RowsAffected()
	if rowsAff != 1 {
		return rowsAffectError("purchaseDetails")
	}
	return nil
}

func (purchaseDetail *PurchaseDetail) CreateMultiple(db *sql.DB, purchaseDetailXi []PurchaseDetail) (modelErr *ModelError) {
	err := db.Ping()
	if err != nil {
		return connectionError("purchaseDetails")
	}

	qryString := `INSERT INTO purchaseDetails(
		named_id, name, status,
		wholesale_price, qty, cost,
		currency, subtotal, remark,
		data_order, purchase_id, supplier_id,
		product_id
	) VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13);`

	stmt, err := db.Prepare(qryString)
	if err != nil {
		return normalError("purchaseDetails", err)
	}
	defer stmt.Close()

	ctx := context.Background()

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return normalError("purchaseDetails", err)
	}

	for _, purchaseDetail := range purchaseDetailXi {

		res, err := stmt.Exec(purchaseDetail.NamedID, purchaseDetail.Name,
			purchaseDetail.Status, purchaseDetail.WholesalePrice, purchaseDetail.QTY,
			purchaseDetail.Cost, purchaseDetail.Currency, purchaseDetail.Subtotal,
			purchaseDetail.Remark, purchaseDetail.DataOrder, purchaseDetail.PurchaseID,
			purchaseDetail.SupplierID, purchaseDetail.ProductID)
		if err, ok := err.(*pq.Error); ok {
			_ = tx.Rollback()
			return normalError("purchaseDetails", err)
		}
		rowsAff, execErr := res.RowsAffected()
		if execErr != nil || err != nil || rowsAff != 1 {
			_ = tx.Rollback()
			return transactionError("purchaseDetails")
		}
	}
	if err := tx.Commit(); err != nil {
		return normalError("purchaseDetails", err)
	}

	return nil
}

func (purchaseDetail *PurchaseDetail) Update(db *sql.DB) (modelErr *ModelError) {
	err := db.Ping()
	if err != nil {
		return normalError("purchaseDetails", err)
	}

	_, err = db.Exec(
		"CALL updatePurchaseDetails($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14)",
		purchaseDetail.ID, purchaseDetail.NamedID, purchaseDetail.Name,
		purchaseDetail.Status, purchaseDetail.WholesalePrice, purchaseDetail.QTY,
		purchaseDetail.Cost, purchaseDetail.Currency, purchaseDetail.Subtotal,
		purchaseDetail.Remark, purchaseDetail.DataOrder, purchaseDetail.PurchaseID,
		purchaseDetail.SupplierID, purchaseDetail.ProductID,
	)
	if err != nil {
		return normalError("purchaseDetails", err)
	}

	return nil
}

func (purchaseDetail *PurchaseDetail) Delete(db *sql.DB) (modelErr *ModelError) {
	err := db.Ping()
	if err != nil {
		return normalError("purchaseDetails", err)
	}

	stmt, err := db.Prepare("DELETE FROM purchaseDetails WHERE id = $1;")
	if err != nil {
		return normalError("purchaseDetails", err)
	}
	defer stmt.Close()

	res, err := stmt.Exec(purchaseDetail.ID)
	if err, ok := err.(*pq.Error); ok {
		fmt.Println(err.Code.Name())
		// return &ModelError{Model: "purchases", Code: 1, Message: "supplier still have children."}
	}

	rowsAff, err := res.RowsAffected()
	if rowsAff != 1 {
		return rowsAffectError("purchaseDetails")
	}
	return nil
}

// ID
// NamedID
// Name
// Status
// WholesalePrice
// QTY
// Cost
// Currency
// Subtotal
// Remark
// DataOrder
// CreateAt
// UpdateAt
// PurchaseID
// SupplierID
// ProductID

// id
// named_id
// name
// status
// wholesale_price
// qty
// cost
// currency
// subtotal
// remark
// data_order
// create_at
// update_at
// purchase_id
// supplier_id
// product_id
