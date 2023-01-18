package model

import (
	"database/sql"
	"fmt"

	"github.com/lib/pq"
)

type Discount struct {
	ID              int64   //流水號
	Name            string  //折扣名稱
	Price           float32 //折扣金額
	DiscountType    int64   //折扣方式
	Remark          *string //備註
	DataOrder       *int64  //順序
	CreateAt        string  //建立時間
	UpdateAt        string  //最後編輯時間
	DeliveryOrderID int64   //出貨編號
}

func GetDiscounts(db *sql.DB, deliveryOrderID int64) (discountXi []Discount, modelErr *ModelError) {
	err := db.Ping()
	if err != nil {
		return nil, normalError("discounts", err)
	}

	row, err := db.Query("SELECT * FROM discounts WHERE delivery_order_id = $1 ORDER BY id DESC;", deliveryOrderID)
	if err != nil {
		return nil, normalError("discounts", err)
	}
	defer row.Close()

	var discount Discount
	for row.Next() {
		err := row.Scan(
			&discount.ID, &discount.Name, &discount.Price,
			&discount.DiscountType, &discount.Remark, &discount.DataOrder,
			&discount.CreateAt, &discount.UpdateAt, &discount.DeliveryOrderID,
		)
		if err != nil {
			return nil, normalError("discounts", err)
		}

		discountXi = append(discountXi, discount)
	}

	return discountXi, nil
}

func (discount *Discount) Create(db *sql.DB) (modelErr *ModelError) {
	err := db.Ping()
	if err != nil {
		return normalError("discounts", err)
	}

	qryString := `INSERT INTO discounts(
		name, price, discount_type,
		remark, data_order, delivery_order_id,
	) VALUES($1,$2,$3,$4,$5,$6,$7,$8);`

	stmt, err := db.Prepare(qryString)
	if err != nil {
		return normalError("discounts", err)
	}
	defer stmt.Close()

	res, err := stmt.Exec(
		discount.Name, discount.Price, discount.DiscountType,
		discount.Remark, discount.DataOrder, discount.DeliveryOrderID,
	)
	if err != nil {
		return normalError("discounts", err)
	}

	rowsAff, err := res.RowsAffected()
	if rowsAff != 1 {
		return rowsAffectError("discounts")
	}
	return nil
}

func (discount *Discount) Update(db *sql.DB) (modelErr *ModelError) {
	err := db.Ping()
	if err != nil {
		return normalError("discounts", err)
	}

	_, err = db.Exec(
		"CALL updateDiscounts($1,$2,$3,$4,$5,$6,$7,$8,$9)",
		discount.ID, discount.Name, discount.Price,
		discount.DiscountType, discount.Remark, discount.DataOrder,
		discount.DeliveryOrderID,
	)
	if err != nil {
		return normalError("discounts", err)
	}

	return nil
}

func (discount *Discount) Delete(db *sql.DB) (modelErr *ModelError) {
	err := db.Ping()
	if err != nil {
		return normalError("discounts", err)
	}

	stmt, err := db.Prepare("DELETE FROM discounts WHERE id = $1;")
	if err != nil {
		return normalError("discounts", err)
	}
	defer stmt.Close()

	res, err := stmt.Exec(discount.ID)
	if err, ok := err.(*pq.Error); ok {
		fmt.Println(err.Code.Name())
	}

	rowsAff, err := res.RowsAffected()
	if rowsAff != 1 {
		return rowsAffectError("discounts")
	}
	return nil
}

// ID
// Name
// Price
// DiscountType
// Remark
// DataOrder
// CreateAt
// UpdateAt
// DeliveryOrderID

// id
// name
// price
// discount_type
// remark
// data_order
// create_at
// update_at
// delivery_order_id

// name, price, discount_type,
// remark, data_order, delivery_order_id,
