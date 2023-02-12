package model

import (
	"context"
	"database/sql"

	"github.com/lib/pq"
)

type DeliveryOrder struct {
	ID                int64   //流水號
	Name              string  //出貨名稱
	Status            int64   //採購狀態
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

func (deliveryOrder *DeliveryOrder) ReadAll(db *sql.DB) (deliveryOrderXi []interface{}, modelErr *ModelError) {
	err := db.Ping()
	if err != nil {
		return nil, &ModelError{Model: "deliveryOrders", Code: 0, Message: err.Error()}
	}

	row, err := db.Query("SELECT * FROM deliveryOrders ORDER BY id DESC;")
	if err != nil {
		return nil, &ModelError{Model: "deliveryOrders", Code: 0, Message: err.Error()}
	}
	defer row.Close()

	var deliveryOrderRow DeliveryOrder
	for row.Next() {
		err := row.Scan(
			&deliveryOrderRow.ID, &deliveryOrderRow.Name, &deliveryOrderRow.Status, &deliveryOrderRow.DeliveryType, &deliveryOrderRow.DeliveryStatus,
			&deliveryOrderRow.DeliveryFeeStatus, &deliveryOrderRow.PaymentType, &deliveryOrderRow.PaymentStatus,
			&deliveryOrderRow.TotalOriginal, &deliveryOrderRow.Discount, &deliveryOrderRow.TotalDiscounted,
			&deliveryOrderRow.Remark, &deliveryOrderRow.DataOrder, &deliveryOrderRow.OrderAt,
			&deliveryOrderRow.SendAt, &deliveryOrderRow.ArriveAt, &deliveryOrderRow.CreateAt,
			&deliveryOrderRow.UpdateAt,
		)
		if err != nil {
			return nil, &ModelError{Model: "deliveryOrders", Code: 0, Message: err.Error()}
		}

		deliveryOrderXi = append(deliveryOrderXi, deliveryOrderRow)
	}

	return deliveryOrderXi, nil
}
func (deliveryOrder *DeliveryOrder) Read(db *sql.DB) (deliveryOrderXi []interface{}, modelErr *ModelError) {
	err := db.Ping()
	if err != nil {
		return nil, &ModelError{Model: "deliveryOrders", Code: 0, Message: err.Error()}
	}

	row, err := db.Query("SELECT * FROM deliveryOrders WHERE id = $1;", deliveryOrder.ID)
	if err != nil {
		return nil, &ModelError{Model: "deliveryOrders", Code: 0, Message: err.Error()}
	}
	defer row.Close()

	var deliveryOrderRow DeliveryOrder
	for row.Next() {
		err := row.Scan(
			&deliveryOrderRow.ID, &deliveryOrderRow.Name, &deliveryOrderRow.Status, &deliveryOrderRow.DeliveryType, &deliveryOrderRow.DeliveryStatus,
			&deliveryOrderRow.DeliveryFeeStatus, &deliveryOrderRow.PaymentType, &deliveryOrderRow.PaymentStatus,
			&deliveryOrderRow.TotalOriginal, &deliveryOrderRow.Discount, &deliveryOrderRow.TotalDiscounted,
			&deliveryOrderRow.Remark, &deliveryOrderRow.DataOrder, &deliveryOrderRow.OrderAt,
			&deliveryOrderRow.SendAt, &deliveryOrderRow.ArriveAt, &deliveryOrderRow.CreateAt,
			&deliveryOrderRow.UpdateAt,
		)
		if err != nil {
			return nil, &ModelError{Model: "deliveryOrders", Code: 0, Message: err.Error()}
		}

		deliveryOrderXi = append(deliveryOrderXi, deliveryOrderRow)
	}

	return deliveryOrderXi, nil
}

func (deliveryOrder *DeliveryOrder) Create(db *sql.DB) (modelErr *ModelError) {
	err := db.Ping()
	if err != nil {
		return &ModelError{Model: "deliveryOrders", Code: 0, Message: err.Error()}
	}

	qryString := `INSERT INTO deliveryOrders(
		name, status, delivery_type, delivery_status, delivery_fee_status,
		payment_type, payment_status, total_original,
		discount, total_discounted, remark,
		data_order, order_at, send_at,
		arrive_at
	) VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15);`

	stmt, err := db.Prepare(qryString)
	if err != nil {
		return &ModelError{Model: "deliveryOrders", Code: 0, Message: err.Error()}
	}
	defer stmt.Close()

	res, err := stmt.Exec(
		deliveryOrder.Name, deliveryOrder.Status, deliveryOrder.DeliveryType, deliveryOrder.DeliveryStatus,
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
		"CALL updateDeliveryOrders($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16)",
		deliveryOrder.ID, deliveryOrder.Name, deliveryOrder.Status, deliveryOrder.DeliveryType, deliveryOrder.DeliveryStatus,
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

	ctx := context.Background()

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return transactionError("deliveryOrders")
	}
	defer tx.Rollback()

	deleteDeliveryOrderDetailsStmt := "DELETE FROM deliveryOrderDetails WHERE delivery_order_id = $1;"
	res, err := tx.Exec(deleteDeliveryOrderDetailsStmt, deliveryOrder.ID)
	if err, ok := err.(*pq.Error); ok {
		return &ModelError{Model: "deliveryOrders", Code: 1, Message: err.Message}
	}
	rowsAff, execErr := res.RowsAffected()
	if execErr != nil || err != nil {
		return transactionError("deliveryOrders")
	}

	deleteDiscountsStmt := "DELETE FROM discounts WHERE delivery_order_id = $1;"
	res, err = tx.Exec(deleteDiscountsStmt, deliveryOrder.ID)
	if err, ok := err.(*pq.Error); ok {
		return &ModelError{Model: "deliveryOrders", Code: 1, Message: err.Message}
	}
	rowsAff, execErr = res.RowsAffected()
	if execErr != nil || err != nil {
		return transactionError("deliveryOrders")
	}

	deleteDeliveryOrdersStmt := "DELETE FROM deliveryOrders WHERE id = $1;"
	res, err = tx.Exec(deleteDeliveryOrdersStmt, deliveryOrder.ID)
	if err, ok := err.(*pq.Error); ok {
		_ = tx.Rollback()
		return &ModelError{Model: "deliveryOrders", Code: 1, Message: err.Message}
	}
	rowsAff, execErr = res.RowsAffected()
	if execErr != nil || err != nil || rowsAff != 1 {
		_ = tx.Rollback()
		return transactionError("deliveryOrders")
	}

	if err := tx.Commit(); err != nil {
		return normalError("deliveryOrders", err)
	}
	return nil
}
