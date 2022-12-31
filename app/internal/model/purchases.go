package model

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/lib/pq"
)

type Purchase struct {
	ID                        int64    //流水號
	Name                      string   //採購名稱
	Status                    int64    //採購狀態
	PurchaseType              int64    //採購種類
	ShippingAgent             *string  //貨運行
	ShippingAgentCutKrw       *float32 //貨運行抽成
	ShippingAgentCutPercent   *float32 //貨運行抽成趴數
	ShippingInitiator         *string  //貨運團主
	ShippingCreateAt          *string  //貨運開團日期
	ShippingEndAt             *string  //貨運結單日期
	ShippingArriveAt          *string  //貨運送達日期
	Weight                    *float32 //貨運總重
	ShippingFeeKr             *float32 //貨運國內運費_韓國
	ShippingFeeTw             *float32 //貨運國內運費_台灣
	ShippingFeeKokusaiKrw     *float32 //貨運國際運費
	ShippingFeeKokusaiPerKilo *float32 //每公斤貨運國際運費
	ExchangeRateKrw           *float32 //韓圓匯率
	TariffTwd                 *float32 //關稅
	TariffPerKilo             *float32 //每公斤關稅
	TotalKrw                  *float32 //韓圓總價
	TotalTwd                  *float32 //台幣總價
	Total                     *float32 //總金額
	Remark                    *string  //備註
	DataOrder                 *int64   //順序
	CreateAt                  string   //建立時間
	UpdateAt                  string   //最後編輯時間
}

func NewPurchase(name string, status int64, purchaseType int64) (Purchase, error) {
	var purchase Purchase

	if name == "" {
		return purchase, errors.New("name field should not be empty")
	}

	purchase = Purchase{
		Name:         name,
		Status:       status,
		PurchaseType: purchaseType,
	}

	return purchase, nil
}

func GetAllPurchases(db *sql.DB) (purchaseXi []Purchase, modelErr *ModelError) {
	err := db.Ping()
	if err != nil {
		return nil, &ModelError{Model: "purchases", Code: 0, Message: err.Error()}
	}

	row, err := db.Query("SELECT * FROM purchases ORDER BY id DESC;")
	if err != nil {
		return nil, &ModelError{Model: "purchases", Code: 0, Message: err.Error()}
	}
	defer row.Close()

	var purchase Purchase
	for row.Next() {
		err := row.Scan(
			&purchase.ID, &purchase.Name, &purchase.Status, &purchase.PurchaseType,
			&purchase.ShippingAgent, &purchase.ShippingAgentCutKrw, &purchase.ShippingAgentCutPercent, &purchase.ShippingInitiator,
			&purchase.ShippingCreateAt, &purchase.ShippingEndAt, &purchase.ShippingArriveAt, &purchase.Weight,
			&purchase.ShippingFeeKr, &purchase.ShippingFeeTw, &purchase.ShippingFeeKokusaiKrw, &purchase.ShippingFeeKokusaiPerKilo,
			&purchase.ExchangeRateKrw, &purchase.TariffTwd, &purchase.TariffPerKilo, &purchase.TotalKrw,
			&purchase.TotalTwd, &purchase.Total, &purchase.Remark, &purchase.DataOrder,
			&purchase.CreateAt, &purchase.UpdateAt,
		)
		if err != nil {
			return nil, &ModelError{Model: "suppliers", Code: 0, Message: err.Error()}
		}

		purchaseXi = append(purchaseXi, purchase)
	}

	return purchaseXi, nil
}

func (purchase *Purchase) Create(db *sql.DB) (modelErr *ModelError) {
	err := db.Ping()
	if err != nil {
		return &ModelError{Model: "purchases", Code: 0, Message: err.Error()}
	}

	qryString := `INSERT INTO purchases(
		name, status, purchase_type, shipping_agent,
		shipping_agent_cut_krw, shipping_agent_cut_percent, shipping_initiator, shipping_create_at,
		shipping_end_at, shipping_arrive_at, weight, shipping_fee_kr,
		shipping_fee_tw, shipping_fee_kokusai_krw, shipping_fee_kokusai_per_kilo, exchange_rate_krw,
		tariff_twd, tariff_per_kilo, total_krw, total_twd,
		total, remark, data_order
	) VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18,$19,$20,$21,$22,$23);`

	stmt, err := db.Prepare(qryString)
	if err != nil {
		return &ModelError{Model: "purchases", Code: 0, Message: err.Error()}
	}
	defer stmt.Close()

	res, err := stmt.Exec(purchase.Name, purchase.Status, purchase.PurchaseType, purchase.ShippingAgent,
		purchase.ShippingAgentCutKrw, purchase.ShippingAgentCutPercent, purchase.ShippingInitiator, purchase.ShippingCreateAt,
		purchase.ShippingEndAt, purchase.ShippingArriveAt, purchase.Weight, purchase.ShippingFeeKr,
		purchase.ShippingFeeTw, purchase.ShippingFeeKokusaiKrw, purchase.ShippingFeeKokusaiPerKilo, purchase.ExchangeRateKrw,
		purchase.TariffTwd, purchase.TariffPerKilo, purchase.TotalKrw, purchase.TotalTwd,
		purchase.Total, purchase.Remark, purchase.DataOrder)
	if err != nil {
		return &ModelError{Model: "purchases", Code: 0, Message: err.Error()}
	}

	rowsAff, err := res.RowsAffected()
	if rowsAff != 1 {
		return &ModelError{Model: "purchases", Code: 2, Message: "RowsAffected incorrect."}
	}
	return nil
}

func (purchase *Purchase) Update(db *sql.DB) (modelErr *ModelError) {
	err := db.Ping()
	if err != nil {
		return &ModelError{Model: "purchases", Code: 0, Message: err.Error()}
	}

	_, err = db.Exec(
		"CALL updatePurchases($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18,$19,$20,$21,$22,$23,$24)",
		purchase.ID, purchase.Name, purchase.Status, purchase.PurchaseType, purchase.ShippingAgent,
		purchase.ShippingAgentCutKrw, purchase.ShippingAgentCutPercent, purchase.ShippingInitiator, purchase.ShippingCreateAt,
		purchase.ShippingEndAt, purchase.ShippingArriveAt, purchase.Weight, purchase.ShippingFeeKr,
		purchase.ShippingFeeTw, purchase.ShippingFeeKokusaiKrw, purchase.ShippingFeeKokusaiPerKilo, purchase.ExchangeRateKrw,
		purchase.TariffTwd, purchase.TariffPerKilo, purchase.TotalKrw, purchase.TotalTwd,
		purchase.Total, purchase.Remark, purchase.DataOrder,
	)
	if err != nil {
		return &ModelError{Model: "purchases", Code: 0, Message: err.Error()}
	}

	return nil
}

func (purchase *Purchase) Delete(db *sql.DB) (modelErr *ModelError) {
	err := db.Ping()
	if err != nil {
		return &ModelError{Model: "purchases", Code: 0, Message: err.Error()}
	}

	stmt, err := db.Prepare("DELETE FROM purchases WHERE id = $1;")
	if err != nil {
		return &ModelError{Model: "purchases", Code: 0, Message: err.Error()}
	}
	defer stmt.Close()

	res, err := stmt.Exec(purchase.ID)
	if err, ok := err.(*pq.Error); ok {
		fmt.Println(err.Code.Name())
		return &ModelError{Model: "purchases", Code: 1, Message: "supplier still have children."}
	}

	rowsAff, err := res.RowsAffected()
	if rowsAff != 1 {
		return &ModelError{Model: "purchases", Code: 2, Message: "RowsAffected incorrect."}
	}
	return nil
}

// ID
// Name
// Status
// PurchaseType
// ShippingAgent
// ShippingAgentCutKrw
// ShippingAgentCutPercent
// ShippingInitiator
// ShippingCreateAt
// ShippingEndAt
// ShippingArriveAt
// Weight
// ShippingFeeKr
// ShippingFeeTw
// ShippingFeeKokusaiKrw
// ShippingFeeKokusaiPerKilo
// ExchangeRateKrw
// TariffTwd
// TariffPerKilo
// TotalKrw
// TotalTwd
// Total
// Remark
// DataOrder
// CreateAt
// UpdateAt

// id
// name
// status
// purchase_type
// shipping_agent
// shipping_agent_cut_krw
// shipping_agent_cut_percent
// shipping_initiator
// shipping_create_at
// shipping_end_at
// shipping_arrive_at
// weight
// shipping_fee_kr
// shipping_fee_tw
// shipping_fee_kokusai_krw
// shipping_fee_kokusai_per_kilo
// exchange_rate_krw
// tariff_twd
// tariff_per_kilo
// total_krw
// total_twd
// total
// remark
// data_order
// create_at
// update_at
