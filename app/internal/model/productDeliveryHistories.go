package model

import (
	"database/sql"
)

type ProductDeliveryHistory struct {
	DeliveryOrderID       int64    //出貨編號
	DeliveryOrderDetailID int64    //出貨明細編號
	RetailPrice           float32  //出貨時售價
	QTY                   int64    //商品總數
	Subtotal              *float32 //小計
	OrderAt               *string  //下訂日期
}

func GetProductDeliveryHistories(db *sql.DB, productID int64) (historyXi []ProductDeliveryHistory, modelErr *ModelError) {
	err := db.Ping()
	if err != nil {
		return nil, normalError("productDeliveryHistories", err)
	}

	row, err := db.Query("SELECT * FROM selectProductDeliveryHistories($1);", productID)
	if err != nil {
		return nil, normalError("productDeliveryHistories", err)
	}
	defer row.Close()

	var history ProductDeliveryHistory
	for row.Next() {
		err := row.Scan(
			&history.DeliveryOrderID,
			&history.DeliveryOrderDetailID,
			&history.RetailPrice,
			&history.QTY,
			&history.Subtotal,
			&history.OrderAt,
		)
		if err != nil {
			return nil, normalError("productDeliveryHistories", err)
		}
		historyXi = append(historyXi, history)
	}

	return historyXi, nil
}
