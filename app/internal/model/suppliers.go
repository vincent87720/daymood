package model

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/lib/pq"
)

type Supplier struct {
	ID        int64   //流水號
	Name      string  //廠商名稱
	Address   *string //廠商地址
	Remark    *string //備註
	DataOrder *int64  //順序
	CreateAt  string  //建立時間
	UpdateAt  string  //最後編輯時間
}

func NewSupplier(name string, address *string, remark *string, dataOrder *int64) (Supplier, error) {
	var supplier Supplier

	if name == "" {
		return supplier, errors.New("name field should not be empty")
	}

	supplier = Supplier{
		Name:      name,
		Address:   address,
		Remark:    remark,
		DataOrder: dataOrder,
	}

	return supplier, nil
}

func GetAllSuppliers(db *sql.DB) (supplierXi []Supplier, modelErr *ModelError) {
	err := db.Ping()
	if err != nil {
		return nil, &ModelError{Model: "suppliers", Code: 0, Message: err.Error()}
	}

	row, err := db.Query("SELECT * FROM suppliers;")
	if err != nil {
		return nil, &ModelError{Model: "suppliers", Code: 0, Message: err.Error()}
	}
	defer row.Close()

	var supplierSchema Supplier
	for row.Next() {
		err := row.Scan(&supplierSchema.ID, &supplierSchema.Name, &supplierSchema.Address, &supplierSchema.Remark, &supplierSchema.DataOrder, &supplierSchema.CreateAt, &supplierSchema.UpdateAt)
		if err != nil {
			return nil, &ModelError{Model: "suppliers", Code: 0, Message: err.Error()}
		}
		tempSupplier := Supplier{
			ID:        supplierSchema.ID,
			Name:      supplierSchema.Name,
			Address:   supplierSchema.Address,
			Remark:    supplierSchema.Remark,
			DataOrder: supplierSchema.DataOrder,
			CreateAt:  supplierSchema.CreateAt,
			UpdateAt:  supplierSchema.UpdateAt,
		}

		// if supplierSchema.Address == nil {
		// 	tempSupplier.Address = ""
		// } else {
		// 	tempSupplier.Address = *supplierSchema.Address
		// }
		// if supplierSchema.Remark == nil {
		// 	tempSupplier.Remark = ""
		// } else {
		// 	tempSupplier.Remark = *supplierSchema.Remark
		// }
		// if supplierSchema.DataOrder == nil {
		// 	tempSupplier.DataOrder = 0
		// } else {
		// 	tempSupplier.DataOrder = *supplierSchema.DataOrder
		// }

		supplierXi = append(supplierXi, tempSupplier)
	}

	return supplierXi, nil
}

func (supplier *Supplier) Create(db *sql.DB) (modelErr *ModelError) {
	err := db.Ping()
	if err != nil {
		return &ModelError{Model: "suppliers", Code: 0, Message: err.Error()}
	}

	stmt, err := db.Prepare("INSERT INTO suppliers(name, address, remark, data_order) VALUES($1,$2,$3,$4);")
	if err != nil {
		return &ModelError{Model: "suppliers", Code: 0, Message: err.Error()}
	}
	defer stmt.Close()

	res, err := stmt.Exec(supplier.Name, supplier.Address, supplier.Remark, supplier.DataOrder)
	if err != nil {
		return &ModelError{Model: "suppliers", Code: 0, Message: err.Error()}
	}

	rowsAff, err := res.RowsAffected()
	if rowsAff != 1 {
		return &ModelError{Model: "suppliers", Code: 2, Message: "RowsAffected incorrect."}
	}
	return nil
}

func (supplier *Supplier) Update(db *sql.DB) (modelErr *ModelError) {
	err := db.Ping()
	if err != nil {
		return &ModelError{Model: "suppliers", Code: 0, Message: err.Error()}
	}

	_, err = db.Exec("CALL updateSuppliers($1,$2,$3,$4,$5)", supplier.ID, supplier.Name, supplier.Address, supplier.Remark, supplier.DataOrder)
	if err != nil {
		return &ModelError{Model: "suppliers", Code: 0, Message: err.Error()}
	}

	return nil
}

func (supplier *Supplier) Delete(db *sql.DB) (modelErr *ModelError) {
	err := db.Ping()
	if err != nil {
		return &ModelError{Model: "suppliers", Code: 0, Message: err.Error()}
	}

	stmt, err := db.Prepare("DELETE FROM suppliers WHERE id = $1;")
	if err != nil {
		return &ModelError{Model: "suppliers", Code: 0, Message: err.Error()}
	}
	defer stmt.Close()

	res, err := stmt.Exec(supplier.ID)
	if err, ok := err.(*pq.Error); ok {
		fmt.Println(err.Code.Name())
		return &ModelError{Model: "suppliers", Code: 1, Message: "supplier still have children."}
	}

	rowsAff, err := res.RowsAffected()
	if rowsAff != 1 {
		return &ModelError{Model: "suppliers", Code: 2, Message: "RowsAffected incorrect."}
	}
	return nil
}
