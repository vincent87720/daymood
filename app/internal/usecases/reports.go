package usecases

import (
	"database/sql"

	"github.com/vincent87720/daymood/app/internal/model"
)

type Balance struct {
}

func NewBalance() *Balance {
	return &Balance{}
}

func (balance *Balance) Read(db *sql.DB) ([]model.Balance, *model.ModelError) {

	reportXi, modelErr := model.GetBalances(db)
	if modelErr != nil {
		return nil, modelErr
	}
	return reportXi, nil
}
