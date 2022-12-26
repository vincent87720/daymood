package model

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/lib/pq"
)

type Firm struct {
	ID        int64
	Name      string
	Address   string
	Remark    string
	DataOrder string
	CreateAt  string
	UpdateAt  string
}

type FirmSchema struct {
	ID        int64
	Name      string
	Address   *string
	Remark    *string
	DataOrder *string
	CreateAt  string
	UpdateAt  string
}

func NewFirm(name string, address string, remark string) (Firm, error) {
	var firm Firm

	if name == "" {
		return firm, errors.New("Column is null or empty.")
	}

	// uuid := uuid.New()
	// id := uuid.String()

	firm = Firm{
		// ID:      id[0:29],
		Name:    name,
		Address: address,
		Remark:  remark,
	}

	return firm, nil
}

func GetAllFirm(db *sql.DB) ([]Firm, error) {
	err := db.Ping()
	if err != nil {
		return nil, err
	}

	row, err := db.Query("SELECT * FROM firm;")
	if err != nil {
		fmt.Println(err)
	}
	defer row.Close()

	var firmSchema FirmSchema
	var firmXi []Firm
	for row.Next() {
		err := row.Scan(&firmSchema.ID, &firmSchema.Name, &firmSchema.Address, &firmSchema.Remark, &firmSchema.DataOrder, &firmSchema.CreateAt, &firmSchema.UpdateAt)
		if err != nil {
			log.Fatal(err)
		}
		tempFirm := Firm{
			ID:       firmSchema.ID,
			Name:     firmSchema.Name,
			CreateAt: firmSchema.CreateAt,
			UpdateAt: firmSchema.UpdateAt,
		}

		if firmSchema.Address == nil {
			tempFirm.Address = ""
		} else {
			tempFirm.Address = *firmSchema.Address
		}
		if firmSchema.Remark == nil {
			tempFirm.Remark = ""
		} else {
			tempFirm.Remark = *firmSchema.Remark
		}
		if firmSchema.DataOrder == nil {
			tempFirm.DataOrder = ""
		} else {
			tempFirm.DataOrder = *firmSchema.DataOrder
		}

		firmXi = append(firmXi, tempFirm)
	}

	return firmXi, nil
}

func (firm *Firm) Create(db *sql.DB) error {
	err := db.Ping()
	if err != nil {
		return err
	}

	// stmt, err := db.Prepare("INSERT INTO firm(firm_id, firm_name, firm_address, remark) VALUES($1,$2,$3,$4);")
	stmt, err := db.Prepare("INSERT INTO firm(firm_name, firm_address, remark) VALUES($1,$2,$3);")
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(firm.Name, firm.Address, firm.Remark)
	if err != nil {
		return err
	}

	rowsAff, err := res.RowsAffected()
	if rowsAff != 1 {
		return errors.New("RowsAffected incorrect.")
	}
	return nil
}

func (firm *Firm) Update(db *sql.DB) (modelErr *ModelError, err error) {
	err = db.Ping()
	if err != nil {
		return &ModelError{Code: 0, Message: "Connection to the Database is lost."}, err
	}

	res, err := db.Exec("CALL update_firm($1,$2,$3,$4)", firm.ID, firm.Name, firm.Address, firm.Remark)
	if err != nil {
		return &ModelError{Code: 3, Message: "Update fail"}, err
	}
	fmt.Println(res)

	return nil, nil
}

func (firm *Firm) Delete(db *sql.DB) (modelErr *ModelError, err error) {
	err = db.Ping()
	if err != nil {
		return &ModelError{Code: 0, Message: "?"}, err
	}

	stmt, err := db.Prepare("DELETE FROM firm WHERE firm_id = $1;")
	if err != nil {
		return &ModelError{Code: 0, Message: "?"}, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(firm.ID)
	if err, ok := err.(*pq.Error); ok {
		fmt.Println(err.Code.Name())
		return &ModelError{Code: 1, Message: "Firm still have children."}, err
	}
	// if err != nil {
	// 	return err
	// }

	rowsAff, err := res.RowsAffected()
	if rowsAff != 1 {
		return &ModelError{Code: 2, Message: "RowsAffected incorrect."}, err
	}
	return nil, nil
}
