package model

import (
	"database/sql"
)

type ProductPurchaseHistory struct {
	PurchaseID              int64    //採購編號
	PurchaseName            string   //採購名稱
	PurchaseQTY             *int64   //採購商品總數
	PurchaseDetailID        int64    //採購明細編號
	PurchaseDetailName      string   //採購明細名稱
	PurchaseDetailQTY       int64    //採購明細商品總數
	WholesalePrice          *float32 //批價
	Subtotal                *float32 //小計
	ShippingArriveAt        *string  //貨運送達日期
	SupplierID              *int64   //廠商編號
	ExchangeRateKrw         *float32 //韓圓匯率
	ShippingAgentCutPercent *float32 //貨運行百分比
	ShippingFeeKr           *float32 //貨運國內運費_韓國
	ShippingFeeTw           *float32 //貨運國內運費_台灣
	ShippingFeeKokusaiKrw   *float32 //貨運國際運費
	TariffTwd               *float32 //關稅
}

func (productPurchaseHistory *ProductPurchaseHistory) ReadAll(db *sql.DB, productID int64) (historyXi []ProductPurchaseHistory, modelErr *ModelError) {
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
			&history.PurchaseID, &history.PurchaseName, &history.PurchaseQTY, &history.PurchaseDetailID,
			&history.PurchaseDetailName, &history.PurchaseDetailQTY, &history.WholesalePrice, &history.Subtotal,
			&history.ShippingArriveAt, &history.SupplierID, &history.ExchangeRateKrw, &history.ShippingAgentCutPercent,
			&history.ShippingFeeKr, &history.ShippingFeeTw, &history.ShippingFeeKokusaiKrw, &history.TariffTwd,
		)
		if err != nil {
			return nil, normalError("productPurchaseHistories", err)
		}
		historyXi = append(historyXi, history)
	}

	return historyXi, nil
}
