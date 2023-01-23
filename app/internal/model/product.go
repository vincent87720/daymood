package model

import (
	"context"
	"database/sql"
	"errors"

	"github.com/lib/pq"
)

type Product struct {
	ID          int64    //流水號
	SKU         *string  //商品顯示編號
	Name        string   //商品名稱
	ProductType int64    //商品種類
	ImgName     *string  //商品圖片
	ImgID       *string  //商品圖片編號
	Stocks      int64    //庫存
	Weight      *float32 //重量
	RetailPrice float32  //售價
	Remark      *string  //備註
	DataStatus  int64    //是否啟用 0:已刪除,1:使用中
	CreateAt    string   //建立時間
	UpdateAt    string   //最後編輯時間
	SupplierID  *int64   //廠商編號
}

func NewProduct(name string, productType int64, stocks int64, retailPrice float32) (Product, error) {
	var product Product

	if checkEmpty(name) == true ||
		productType < 0 ||
		stocks < 0 ||
		retailPrice < 0 {
		return product, errors.New("Column is null or empty.")
	}

	product = Product{
		Name:        name,
		ProductType: productType,
		Stocks:      stocks,
		RetailPrice: retailPrice,
	}

	return product, nil
}

func (product *Product) ReadAll(db *sql.DB) (productXi []interface{}, modelErr *ModelError) {
	err := db.Ping()
	if err != nil {
		return nil, connectionError("products")
	}

	row, err := db.Query("SELECT * FROM products ORDER BY id DESC;")
	if err != nil {
		return nil, normalError("products", err)
	}
	defer row.Close()

	var productRow Product
	for row.Next() {
		err := row.Scan(&productRow.ID, &productRow.SKU, &productRow.Name,
			&productRow.ProductType, &productRow.ImgName, &productRow.ImgID,
			&productRow.Stocks, &productRow.Weight, &productRow.RetailPrice,
			&productRow.Remark, &productRow.DataStatus, &productRow.CreateAt,
			&productRow.UpdateAt, &productRow.SupplierID)
		if err != nil {
			return nil, normalError("products", err)
		}
		productXi = append(productXi, productRow)
	}

	return productXi, nil
}

func (product *Product) Read(db *sql.DB) (productXi []interface{}, modelErr *ModelError) {
	err := db.Ping()
	if err != nil {
		return nil, connectionError("products")
	}

	row, err := db.Query("SELECT * FROM products WHERE id = $1 ORDER BY id DESC;", product.ID)
	if err != nil {
		return nil, normalError("products", err)
	}
	defer row.Close()

	var productRow Product
	for row.Next() {
		err := row.Scan(&productRow.ID, &productRow.SKU, &productRow.Name,
			&productRow.ProductType, &productRow.ImgName, &productRow.ImgID,
			&productRow.Stocks, &productRow.Weight, &productRow.RetailPrice,
			&productRow.Remark, &productRow.DataStatus, &productRow.CreateAt,
			&productRow.UpdateAt, &productRow.SupplierID)
		if err != nil {
			return nil, normalError("products", err)
		}
		productXi = append(productXi, productRow)
	}

	return productXi, nil
}

func (product *Product) Create(db *sql.DB) (modelErr *ModelError) {
	err := db.Ping()
	if err != nil {
		return connectionError("products")
	}

	qstr := `
	INSERT INTO
	products(
		sku, name, type, img_id, img_name,
		stocks, weight, retail_price, remark, 
		data_status, supplier_id
	)
	VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11);`

	stmt, err := db.Prepare(qstr)
	if err != nil {
		return normalError("products", err)
	}
	defer stmt.Close()

	res, err := stmt.Exec(
		product.SKU, product.Name, product.ProductType,
		product.ImgName, product.ImgID, product.Stocks,
		product.Weight, product.RetailPrice, product.Remark,
		product.DataStatus, product.SupplierID)
	if err, ok := err.(*pq.Error); ok {
		if err.Code == "23505" {
			return uniqueError("products", "SKU")

		} else {
			return normalError("products", err)
		}
	}

	rowsAff, err := res.RowsAffected()
	if rowsAff != 1 {
		return rowsAffectError("products")
	}
	return nil
}

func (product *Product) CreateMultiple(db *sql.DB, productXi []Product) (modelErr *ModelError) {
	err := db.Ping()
	if err != nil {
		return connectionError("products")
	}

	qstr := `
	INSERT INTO
	products(
		sku, name, type, img_id, img_name,
		stocks, weight, retail_price, remark, 
		data_status, supplier_id
	)
	VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11);`

	stmt, err := db.Prepare(qstr)
	if err != nil {
		return normalError("products", err)
	}
	defer stmt.Close()

	ctx := context.Background()

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return normalError("products", err)
	}

	for _, product := range productXi {

		res, err := stmt.Exec(
			product.SKU, product.Name, product.ProductType,
			product.ImgName, product.ImgID, product.Stocks,
			product.Weight, product.RetailPrice, product.Remark,
			product.DataStatus, product.SupplierID)
		if err, ok := err.(*pq.Error); ok {
			if err.Code == "23505" {
				_ = tx.Rollback()
				return uniqueError("products", "SKU")

			} else {
				_ = tx.Rollback()
				return normalError("products", err)
			}
		}
		rowsAff, execErr := res.RowsAffected()
		if execErr != nil || err != nil || rowsAff != 1 {
			_ = tx.Rollback()
			return transactionError("products")
		}
	}
	if err := tx.Commit(); err != nil {
		return normalError("products", err)
	}

	return nil
}

func (product *Product) Update(db *sql.DB) (modelErr *ModelError) {
	err := db.Ping()
	if err != nil {
		return connectionError("products")
	}

	stmt, err := db.Prepare("CALL updateProducts($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12)")
	if err != nil {
		return normalError("products", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		product.ID, product.SKU, product.Name, product.ProductType,
		product.ImgName, product.ImgID, product.Stocks,
		product.Weight, product.RetailPrice, product.Remark,
		product.DataStatus, product.SupplierID)

	if err != nil {
		return normalError("products", err)
	}

	return nil
}

func (product *Product) Delete(db *sql.DB) (modelErr *ModelError) {
	err := db.Ping()
	if err != nil {
		return connectionError("products")
	}

	stmt, err := db.Prepare("DELETE FROM products WHERE id = $1;")
	if err != nil {
		return normalError("products", err)
	}
	defer stmt.Close()

	res, err := stmt.Exec(product.ID)
	if err, ok := err.(*pq.Error); ok {
		return normalError("products", err)
	}

	rowsAff, err := res.RowsAffected()
	if rowsAff != 1 {
		return rowsAffectError("products")
	}
	return nil
}

func IncreaseStocks(db *sql.DB, productID string, qty int) error {
	err := db.Ping()
	if err != nil {
		return err
	}

	stmt, err := db.Prepare("UPDATE product SET stocks = stocks + $1 WHERE product_id = $2;")
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(qty, productID)
	if err != nil {
		return err
	}
	rowsAff, err := res.RowsAffected()
	if err != nil || rowsAff != 1 {
		return err
	}
	return nil
}

func checkEmpty(s string) bool {
	if s == "" || len(s) <= 0 {
		return true
	}
	return false
}
