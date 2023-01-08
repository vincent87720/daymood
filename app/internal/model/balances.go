package model

import "database/sql"

type Balance struct {
	PurchaseTotal float32
	DeliveryTotal float32
}

func GetBalances(db *sql.DB) (balance []Balance, modelErr *ModelError) {
	err := db.Ping()
	if err != nil {
		return nil, normalError("balance", err)
	}

	row, err := db.Query("SELECT * FROM selectBalance();")
	if err != nil {
		return nil, normalError("balance", err)
	}
	defer row.Close()

	var balanceXi []Balance
	var tempRow Balance
	for row.Next() {
		err := row.Scan(
			&tempRow.PurchaseTotal,
			&tempRow.DeliveryTotal,
		)
		if err != nil {
			return nil, normalError("balance", err)
		}
		balanceXi = append(balanceXi, tempRow)
	}

	return balanceXi, nil
}
