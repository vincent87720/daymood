package model

import (
	"database/sql"
	"fmt"

	"github.com/lib/pq"
)

type DeliveryOrder struct {
	ID                int64   //流水號
	DeliveryType      *int64  //出貨方式
	DeliveryStatus    *int64  //出貨狀態
	DeliveryFeeStatus *int64  //運費狀態
	PaymentType       *int64  //付款方式
	PaymentStatus     int64   //付款狀態
	TotalOriginal     float32 //原價
	Discount          float32 //折扣
	TotalDiscounted   float32 //總價
	Remark            *string //備註
	DataOrder         *int64  //順序
	OrderAt           *string //下訂日期
	SendAt            *string //出貨日期
	ArriveAt          *string //送達日期
	CreateAt          string  //建立時間
	UpdateAt          string  //最後編輯時間
}

func GetAllDeliveryOrders(db *sql.DB) (deliveryOrderXi []DeliveryOrder, modelErr *ModelError) {
	err := db.Ping()
	if err != nil {
		return nil, &ModelError{Model: "deliveryOrders", Code: 0, Message: err.Error()}
	}

	row, err := db.Query("SELECT * FROM deliveryOrders ORDER BY id DESC;")
	if err != nil {
		return nil, &ModelError{Model: "deliveryOrders", Code: 0, Message: err.Error()}
	}
	defer row.Close()

	var deliveryOrder DeliveryOrder
	for row.Next() {
		err := row.Scan(
			&deliveryOrder.ID, &deliveryOrder.DeliveryType, &deliveryOrder.DeliveryStatus,
			&deliveryOrder.DeliveryFeeStatus, &deliveryOrder.PaymentType, &deliveryOrder.PaymentStatus,
			&deliveryOrder.TotalOriginal, &deliveryOrder.Discount, &deliveryOrder.TotalDiscounted,
			&deliveryOrder.Remark, &deliveryOrder.DataOrder, &deliveryOrder.OrderAt,
			&deliveryOrder.SendAt, &deliveryOrder.ArriveAt, &deliveryOrder.CreateAt,
			&deliveryOrder.UpdateAt,
		)
		if err != nil {
			return nil, &ModelError{Model: "deliveryOrders", Code: 0, Message: err.Error()}
		}

		deliveryOrderXi = append(deliveryOrderXi, deliveryOrder)
	}

	return deliveryOrderXi, nil
}
func (deliveryOrder *DeliveryOrder) GetDeliveryOrder(db *sql.DB) (deliveryOrderXi []DeliveryOrder, modelErr *ModelError) {
	err := db.Ping()
	if err != nil {
		return nil, &ModelError{Model: "deliveryOrders", Code: 0, Message: err.Error()}
	}

	row, err := db.Query("SELECT * FROM deliveryOrders WHERE id = $1;", deliveryOrder.ID)
	if err != nil {
		return nil, &ModelError{Model: "deliveryOrders", Code: 0, Message: err.Error()}
	}
	defer row.Close()

	var purchaseRow DeliveryOrder
	for row.Next() {
		err := row.Scan(
			&deliveryOrder.ID, &deliveryOrder.DeliveryType, &deliveryOrder.DeliveryStatus,
			&deliveryOrder.DeliveryFeeStatus, &deliveryOrder.PaymentType, &deliveryOrder.PaymentStatus,
			&deliveryOrder.TotalOriginal, &deliveryOrder.Discount, &deliveryOrder.TotalDiscounted,
			&deliveryOrder.Remark, &deliveryOrder.DataOrder, &deliveryOrder.OrderAt,
			&deliveryOrder.SendAt, &deliveryOrder.ArriveAt, &deliveryOrder.CreateAt,
			&deliveryOrder.UpdateAt,
		)
		if err != nil {
			return nil, &ModelError{Model: "suppliers", Code: 0, Message: err.Error()}
		}

		deliveryOrderXi = append(deliveryOrderXi, purchaseRow)
	}

	return deliveryOrderXi, nil
}

func (deliveryOrder *DeliveryOrder) Create(db *sql.DB) (modelErr *ModelError) {
	err := db.Ping()
	if err != nil {
		return &ModelError{Model: "deliveryOrders", Code: 0, Message: err.Error()}
	}

	qryString := `INSERT INTO deliveryOrders(
		delivery_type, delivery_status, delivery_fee_status,
		payment_type, payment_status, total_original,
		discount, total_discounted, remark,
		data_order, order_at, send_at,
		arrive_at
	) VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13);`

	stmt, err := db.Prepare(qryString)
	if err != nil {
		return &ModelError{Model: "deliveryOrders", Code: 0, Message: err.Error()}
	}
	defer stmt.Close()

	res, err := stmt.Exec(
		deliveryOrder.DeliveryType, deliveryOrder.DeliveryStatus,
		deliveryOrder.DeliveryFeeStatus, deliveryOrder.PaymentType, deliveryOrder.PaymentStatus,
		deliveryOrder.TotalOriginal, deliveryOrder.Discount, deliveryOrder.TotalDiscounted,
		deliveryOrder.Remark, deliveryOrder.DataOrder, deliveryOrder.OrderAt,
		deliveryOrder.SendAt, deliveryOrder.ArriveAt,
	)
	if err != nil {
		return &ModelError{Model: "deliveryOrders", Code: 0, Message: err.Error()}
	}

	rowsAff, err := res.RowsAffected()
	if rowsAff != 1 {
		return &ModelError{Model: "deliveryOrders", Code: 2, Message: "RowsAffected incorrect."}
	}
	return nil
}

func (deliveryOrder *DeliveryOrder) Update(db *sql.DB) (modelErr *ModelError) {
	err := db.Ping()
	if err != nil {
		return &ModelError{Model: "deliveryOrders", Code: 0, Message: err.Error()}
	}

	_, err = db.Exec(
		"CALL updateDeliveryOrders($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14)",
		deliveryOrder.ID, deliveryOrder.DeliveryType, deliveryOrder.DeliveryStatus,
		deliveryOrder.DeliveryFeeStatus, deliveryOrder.PaymentType, deliveryOrder.PaymentStatus,
		deliveryOrder.TotalOriginal, deliveryOrder.Discount, deliveryOrder.TotalDiscounted,
		deliveryOrder.Remark, deliveryOrder.DataOrder, deliveryOrder.OrderAt,
		deliveryOrder.SendAt, deliveryOrder.ArriveAt,
	)
	if err != nil {
		return &ModelError{Model: "deliveryOrders", Code: 0, Message: err.Error()}
	}

	return nil
}

func (deliveryOrder *DeliveryOrder) Delete(db *sql.DB) (modelErr *ModelError) {
	err := db.Ping()
	if err != nil {
		return &ModelError{Model: "deliveryOrders", Code: 0, Message: err.Error()}
	}

	stmt, err := db.Prepare("DELETE FROM deliveryOrders WHERE id = $1;")
	if err != nil {
		return &ModelError{Model: "deliveryOrders", Code: 0, Message: err.Error()}
	}
	defer stmt.Close()

	res, err := stmt.Exec(deliveryOrder.ID)
	if err, ok := err.(*pq.Error); ok {
		fmt.Println(err.Code.Name())
		return &ModelError{Model: "deliveryOrders", Code: 1, Message: "deliveryOrders still have children."}
	}

	rowsAff, err := res.RowsAffected()
	if rowsAff != 1 {
		return &ModelError{Model: "deliveryOrders", Code: 2, Message: "RowsAffected incorrect."}
	}
	return nil
}

// ID
// DeliveryType
// DeliveryStatus
// DeliveryFeeStatus
// PaymentType
// PaymentStatus
// TotalOriginal
// Discount
// TotalDiscounted
// Remark
// DataOrder
// OrderAt
// SendAt
// ArriveAt
// CreateAt
// UpdateAt

// id
// delivery_type
// delivery_status
// delivery_fee_status
// payment_type
// payment_status
// total_original
// discount
// total_discounted
// remark
// data_order
// order_at
// send_at
// arrive_at
// create_at
// update_at

// &deliveryOrder.ID,
// &deliveryOrder.DeliveryType,
// &deliveryOrder.DeliveryStatus,
// &deliveryOrder.DeliveryFeeStatus,
// &deliveryOrder.PaymentType,
// &deliveryOrder.PaymentStatus,
// &deliveryOrder.TotalOriginal,
// &deliveryOrder.Discount,
// &deliveryOrder.TotalDiscounted,
// &deliveryOrder.Remark,
// &deliveryOrder.DataOrder,
// &deliveryOrder.OrderAt,
// &deliveryOrder.SendAt,
// &deliveryOrder.ArriveAt,
// &deliveryOrder.CreateAt,
// &deliveryOrder.UpdateAt,
