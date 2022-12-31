package model

import (
	"database/sql"
)

type ProductPurchaseHistory struct {
	PurchaseID         int64    //採購編號
	PurchaseName       string   //採購名稱
	PurchaseDetailID   int64    //採購明細編號
	PurchaseDetailName string   //採購明細名稱
	QTY                int64    //數量
	WholesalePriceKrw  *float32 //韓圓批價
	SubtotalTwd        *float32 //台幣小計
	ShippingArriveAt   *string  //貨運送達日期
	SupplierID         int64    //廠商編號
}

func GetProductPurchaseHistories(db *sql.DB, productID int64) (historyXi []ProductPurchaseHistory, modelErr *ModelError) {
	err := db.Ping()
	if err != nil {
		return nil, normalError("productPurchaseHistories", err)
	}

	row, err := db.Query("SELECT * FROM selectProductPurchaseHistories($1);", productID)
	if err != nil {
		return nil, normalError("productPurchaseHistories", err)
	}
	defer row.Close()

	var history ProductPurchaseHistory
	for row.Next() {
		err := row.Scan(
			&history.PurchaseID, &history.PurchaseName, &history.PurchaseDetailID,
			&history.PurchaseDetailName, &history.QTY, &history.WholesalePriceKrw,
			&history.SubtotalTwd, &history.ShippingArriveAt, &history.SupplierID,
		)
		if err != nil {
			return nil, normalError("purchaseDetails", err)
		}
		historyXi = append(historyXi, history)
	}

	return historyXi, nil
}

// purchase_id
// purchase_Name
// purchase_detail_id
// purchase_detail_name
// qty
// wholesale_price_krw
// subtotal_twd
// shipping_arrive_at
// supplier_id

// PurchaseID
// purchaseName
// PurchaseDetailID
// PurchaseDetailName
// QTY
// WholesalePriceKrw
// SubtotalTwd
// ShippingArriveAt
// SupplierID
