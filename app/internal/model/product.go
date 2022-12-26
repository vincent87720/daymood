package model

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/lib/pq"
)

type ModelError struct {
	Model   string
	Code    int
	Message string
}

//PERR0: Connection to the Database is lost.
//PERR1: Firm instance dosen't exists.
//PERR2: Invalid input
//PERR3: Update fail
//PERR4: ProductSku must be unique.

type Product struct {
	FirmInfo          Firm
	ProductID         int64
	ProductSku        string
	ProductName       string
	ProductType       string
	ProductImgName    string
	ProductImgID      string
	Stocks            int
	Weight            float32
	KrwWholesalePrice float32
	NtdWholesalePrice float32
	NtdListPrice      float32
	NtdSellingPrice   float32
	NtdCnf            float32
	NtdCost           float32
	PurchaseProductID string
	DataOrder         string
	CreateAt          string
	UpdateAt          string
}

type ProductSchema struct {
	ProductID         int64
	ProductSku        string
	ProductName       string
	ProductType       string
	ProductImgName    *string
	ProductImgID      *string
	Stocks            int
	Weight            float32
	KrwWholesalePrice float32
	NtdWholesalePrice float32
	NtdListPrice      float32
	NtdSellingPrice   float32
	NtdCnf            float32
	NtdCost           float32
	PurchaseProductID string
	DataOrder         *string
	CreateAt          string
	UpdateAt          string
	FirmID            int64
	FirmName          string
}

func NewProduct(productSku string, productName string, productType string,
	productImgName string, productImgID string,
	weight float32, krwWholesalePrice float32, ntdWholesalePrice float32,
	ntdListPrice float32, ntdSellingPrice float32,
	ntdCnf float32, ntdCost float32, purchaseProductID string, firmID int64) (Product, error) {
	var product Product

	if checkEmpty(productSku) == true ||
		checkEmpty(productName) == true ||
		checkEmpty(productType) == true ||
		weight < 0 ||
		krwWholesalePrice < 0 ||
		ntdWholesalePrice < 0 ||
		ntdListPrice < 0 ||
		ntdSellingPrice < 0 ||
		ntdCnf < 0 ||
		ntdCost < 0 ||
		checkEmpty(purchaseProductID) == true ||
		firmID < 0 {

		return product, errors.New("Column is null or empty.")
	}

	// uuid := uuid.New()
	// id := uuid.String()

	// if checkEmpty(productID) == true {
	// 	productID = id
	// }

	product = Product{
		ProductSku:        productSku,
		ProductName:       productName,
		ProductType:       productType,
		ProductImgName:    productImgName,
		ProductImgID:      productImgID,
		Stocks:            0,
		Weight:            weight,
		KrwWholesalePrice: krwWholesalePrice,
		NtdWholesalePrice: ntdWholesalePrice,
		NtdListPrice:      ntdListPrice,
		NtdSellingPrice:   ntdSellingPrice,
		NtdCnf:            ntdCnf,
		NtdCost:           ntdCost,
		PurchaseProductID: purchaseProductID,
		FirmInfo: Firm{
			ID: firmID,
		},
	}

	return product, nil
}

func GetAllProducts(db *sql.DB) ([]Product, error) {
	err := db.Ping()
	if err != nil {
		return nil, err
	}

	stmt := `
	SELECT 
	p.product_id, p.product_sku, 
	p.product_name, p.product_type, 
	p.product_img_name, p.product_img_id, 
	p.stocks, p.weight, 
	p.krw_wholesale_price, p.ntd_wholesale_price, 
	p.ntd_list_price, p.ntd_selling_price, 
	p.ntd_cnf, p.ntd_cost,
	p.data_order, p.create_at, p.update_at, 
	p.purchase_product_id, f.firm_id, f.firm_name 
	FROM product AS p INNER JOIN firm AS f ON p.firm_id = f.firm_id 
	ORDER BY p.update_at DESC;`

	row, err := db.Query(stmt)
	if err != nil {
		fmt.Println(err)
	}
	defer row.Close()

	var productSchema ProductSchema
	var productXi []Product
	for row.Next() {
		err := row.Scan(&productSchema.ProductID, &productSchema.ProductSku,
			&productSchema.ProductName, &productSchema.ProductType,
			&productSchema.ProductImgName, &productSchema.ProductImgID,
			&productSchema.Stocks, &productSchema.Weight,
			&productSchema.KrwWholesalePrice, &productSchema.NtdWholesalePrice,
			&productSchema.NtdListPrice, &productSchema.NtdSellingPrice,
			&productSchema.NtdCnf, &productSchema.NtdCost,
			&productSchema.DataOrder, &productSchema.CreateAt, &productSchema.UpdateAt,
			&productSchema.PurchaseProductID, &productSchema.FirmID, &productSchema.FirmName)
		if err != nil {
			log.Fatal(err)
		}
		tempProduct := Product{
			FirmInfo: Firm{
				ID:   productSchema.FirmID,
				Name: productSchema.FirmName,
			},
			ProductID:         productSchema.ProductID,
			ProductSku:        productSchema.ProductSku,
			ProductName:       productSchema.ProductName,
			ProductType:       productSchema.ProductType,
			Stocks:            productSchema.Stocks,
			Weight:            productSchema.Weight,
			KrwWholesalePrice: productSchema.KrwWholesalePrice,
			NtdWholesalePrice: productSchema.NtdWholesalePrice,
			NtdListPrice:      productSchema.NtdListPrice,
			NtdSellingPrice:   productSchema.NtdSellingPrice,
			NtdCnf:            productSchema.NtdCnf,
			NtdCost:           productSchema.NtdCost,
			PurchaseProductID: productSchema.PurchaseProductID,
			CreateAt:          productSchema.CreateAt,
			UpdateAt:          productSchema.UpdateAt,
		}

		if productSchema.ProductImgName == nil {
			tempProduct.ProductImgName = ""
		} else {
			tempProduct.ProductImgName = *productSchema.ProductImgName
		}

		if productSchema.ProductImgID == nil {
			tempProduct.ProductImgID = ""
		} else {
			tempProduct.ProductImgID = *productSchema.ProductImgID
		}

		if productSchema.DataOrder == nil {
			tempProduct.DataOrder = ""
		} else {
			tempProduct.DataOrder = *productSchema.DataOrder
		}

		productXi = append(productXi, tempProduct)
	}

	return productXi, nil
}

func GetProducts(db *sql.DB, productSkuXi []string) (map[string]Product, error) {

	productMap := make(map[string]Product)
	err := db.Ping()
	if err != nil {
		return nil, err
	}

	porductSkuStr := strings.Join(productSkuXi, "','")
	stmt := `
	SELECT 
	p.product_id, p.product_sku, 
	p.product_name, p.product_type, 
	p.product_img_name, p.product_img_id, 
	p.stocks, p.weight, 
	p.krw_wholesale_price, p.ntd_wholesale_price, 
	p.ntd_list_price, p.ntd_selling_price, 
	p.ntd_cnf, p.ntd_cost, 
	p.data_order, p.create_at, p.update_at, 
	p.purchase_product_id, f.firm_id, f.firm_name 
	FROM product AS p INNER JOIN firm AS f ON p.firm_id = f.firm_id 
	WHERE p.product_sku
	IN ('` + porductSkuStr + `')
	ORDER BY p.update_at DESC;`

	row, err := db.Query(stmt)
	if err, ok := err.(*pq.Error); ok {
		fmt.Println(err)
	}
	defer row.Close()

	var productSchema ProductSchema
	// var productXi []Product
	for row.Next() {
		err := row.Scan(&productSchema.ProductID, &productSchema.ProductSku,
			&productSchema.ProductName, &productSchema.ProductType,
			&productSchema.ProductImgName, &productSchema.ProductImgID,
			&productSchema.Stocks, &productSchema.Weight,
			&productSchema.KrwWholesalePrice, &productSchema.NtdWholesalePrice,
			&productSchema.NtdListPrice, &productSchema.NtdSellingPrice,
			&productSchema.NtdCnf, &productSchema.NtdCost,
			&productSchema.DataOrder, &productSchema.CreateAt, &productSchema.UpdateAt,
			&productSchema.PurchaseProductID, &productSchema.FirmID, &productSchema.FirmName)
		if err != nil {
			log.Fatal(err)
		}
		tempProduct := Product{
			FirmInfo: Firm{
				ID:   productSchema.FirmID,
				Name: productSchema.FirmName,
			},
			ProductID:         productSchema.ProductID,
			ProductSku:        productSchema.ProductSku,
			ProductName:       productSchema.ProductName,
			ProductType:       productSchema.ProductType,
			ProductImgName:    *productSchema.ProductImgName,
			ProductImgID:      *productSchema.ProductImgID,
			Stocks:            productSchema.Stocks,
			Weight:            productSchema.Weight,
			KrwWholesalePrice: productSchema.KrwWholesalePrice,
			NtdWholesalePrice: productSchema.NtdWholesalePrice,
			NtdListPrice:      productSchema.NtdListPrice,
			NtdSellingPrice:   productSchema.NtdSellingPrice,
			NtdCnf:            productSchema.NtdCnf,
			NtdCost:           productSchema.NtdCost,
			CreateAt:          productSchema.CreateAt,
			UpdateAt:          productSchema.UpdateAt,
			PurchaseProductID: productSchema.PurchaseProductID,
		}

		if productSchema.ProductImgName == nil {
			tempProduct.ProductImgName = ""
		} else {
			tempProduct.ProductImgName = *productSchema.ProductImgName
		}

		if productSchema.ProductImgID == nil {
			tempProduct.ProductImgID = ""
		} else {
			tempProduct.ProductImgID = *productSchema.ProductImgID
		}

		if productSchema.DataOrder == nil {
			tempProduct.DataOrder = ""
		} else {
			tempProduct.DataOrder = *productSchema.DataOrder
		}
		productMap[strconv.FormatInt(productSchema.ProductID, 10)] = tempProduct

		// productXi = append(productXi, tempProduct)
	}

	return productMap, nil
}

func GetProductsByID(db *sql.DB, productIdXi []string) (map[string]Product, error) {

	productMap := make(map[string]Product)
	err := db.Ping()
	if err != nil {
		return nil, err
	}

	porductIdStr := strings.Join(productIdXi, "','")
	stmt := `
	SELECT 
	p.product_id, p.product_sku, 
	p.product_name, p.product_type, 
	p.product_img_name, p.product_img_id, 
	p.stocks, p.weight, 
	p.krw_wholesale_price, p.ntd_wholesale_price, 
	p.ntd_list_price, p.ntd_selling_price, 
	p.ntd_cnf, p.ntd_cost, 
	p.data_order, p.create_at, p.update_at, 
	p.purchase_product_id, f.firm_id, f.firm_name 
	FROM product AS p INNER JOIN firm AS f ON p.firm_id = f.firm_id 
	WHERE p.product_id
	IN ('` + porductIdStr + `')
	ORDER BY p.update_at DESC;`

	row, err := db.Query(stmt)
	if err, ok := err.(*pq.Error); ok {
		fmt.Println(err)
	}
	defer row.Close()

	var productSchema ProductSchema
	// var productXi []Product
	for row.Next() {
		err := row.Scan(&productSchema.ProductID, &productSchema.ProductSku,
			&productSchema.ProductName, &productSchema.ProductType,
			&productSchema.ProductImgName, &productSchema.ProductImgID,
			&productSchema.Stocks, &productSchema.Weight,
			&productSchema.KrwWholesalePrice, &productSchema.NtdWholesalePrice,
			&productSchema.NtdListPrice, &productSchema.NtdSellingPrice,
			&productSchema.NtdCnf, &productSchema.NtdCost,
			&productSchema.DataOrder, &productSchema.CreateAt, &productSchema.UpdateAt,
			&productSchema.PurchaseProductID, &productSchema.FirmID, &productSchema.FirmName)
		if err != nil {
			log.Fatal(err)
		}
		tempProduct := Product{
			FirmInfo: Firm{
				ID:   productSchema.FirmID,
				Name: productSchema.FirmName,
			},
			ProductID:         productSchema.ProductID,
			ProductSku:        productSchema.ProductSku,
			ProductName:       productSchema.ProductName,
			ProductType:       productSchema.ProductType,
			ProductImgName:    *productSchema.ProductImgName,
			ProductImgID:      *productSchema.ProductImgID,
			Stocks:            productSchema.Stocks,
			Weight:            productSchema.Weight,
			KrwWholesalePrice: productSchema.KrwWholesalePrice,
			NtdWholesalePrice: productSchema.NtdWholesalePrice,
			NtdListPrice:      productSchema.NtdListPrice,
			NtdSellingPrice:   productSchema.NtdSellingPrice,
			NtdCnf:            productSchema.NtdCnf,
			NtdCost:           productSchema.NtdCost,
			CreateAt:          productSchema.CreateAt,
			UpdateAt:          productSchema.UpdateAt,
			PurchaseProductID: productSchema.PurchaseProductID,
		}

		if productSchema.ProductImgName == nil {
			tempProduct.ProductImgName = ""
		} else {
			tempProduct.ProductImgName = *productSchema.ProductImgName
		}

		if productSchema.ProductImgID == nil {
			tempProduct.ProductImgID = ""
		} else {
			tempProduct.ProductImgID = *productSchema.ProductImgID
		}

		if productSchema.DataOrder == nil {
			tempProduct.DataOrder = ""
		} else {
			tempProduct.DataOrder = *productSchema.DataOrder
		}
		productMap[strconv.FormatInt(productSchema.ProductID, 10)] = tempProduct

		// productXi = append(productXi, tempProduct)
	}

	return productMap, nil
}

func (product *Product) Create(db *sql.DB) (modelErr *ModelError, err error) {
	err = db.Ping()
	if err != nil {
		return &ModelError{Code: 0, Message: "Connection to the Database is lost."}, err
	}

	qstr := `
	INSERT INTO 
	product(
		product_sku, product_name, product_type,
		product_img_name, product_img_id,
		stocks, weight, 
		krw_wholesale_price, ntd_wholesale_price, 
		ntd_list_price, ntd_selling_price, 
		ntd_cnf, ntd_cost, 
		purchase_product_id, firm_id
	) 
	VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15);`

	stmt, err := db.Prepare(qstr)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(
		product.ProductSku, product.ProductName, product.ProductType,
		product.ProductImgName, product.ProductImgID,
		product.Stocks, product.Weight,
		product.KrwWholesalePrice, product.NtdWholesalePrice,
		product.NtdListPrice, product.NtdSellingPrice,
		product.NtdCnf, product.NtdCost,
		product.PurchaseProductID, product.FirmInfo.ID)
	if err, ok := err.(*pq.Error); ok {
		if err.Code == "23505" {
			return &ModelError{Code: 4, Message: "ProductSku must be unique."}, err

		} else if err.Code == "23503" {
			return &ModelError{Code: 1, Message: "Firm instance dosen't exists."}, err
		} else {
			return &ModelError{Code: -1, Message: "?"}, err
		}
	}

	rowsAff, err := res.RowsAffected()
	if rowsAff != 1 {
		return nil, err
	}
	return nil, nil
}

func (product *Product) Update(db *sql.DB) (modelErr *ModelError, err error) {
	err = db.Ping()
	if err != nil {
		return &ModelError{Code: 0, Message: "Connection to the Database is lost."}, err
	}

	stmt, err := db.Prepare("CALL update_product($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16)")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(
		product.ProductID, product.ProductSku,
		product.ProductName, product.ProductType,
		product.ProductImgName, product.ProductImgID,
		product.Stocks, product.Weight,
		product.KrwWholesalePrice, product.NtdWholesalePrice,
		product.NtdListPrice, product.NtdSellingPrice,
		product.NtdCnf, product.NtdCost,
		product.PurchaseProductID, product.FirmInfo.ID)

	// res, err := db.Exec("CALL update_product($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14)", product.ProductID, product.ProductSku, product.ProductName, product.ProductType, product.Stocks, product.Weight, product.KrwWholesalePrice, product.NtdWholesalePrice, product.NtdListPrice, product.NtdSellingPrice, product.NtdCnf, product.NtdCost, product.PurchaseProductID, product.FirmInfo.ID)
	if err != nil {
		return &ModelError{Code: 3, Message: "Update fail"}, err
	}
	fmt.Println(res.RowsAffected())

	return nil, nil
}

func (product *Product) Delete(db *sql.DB) error {
	err := db.Ping()
	if err != nil {
		return err
	}

	stmt, err := db.Prepare("DELETE FROM product WHERE product_id = $1;")
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(product.ProductID)
	if err != nil {
		return err
	}

	rowsAff, err := res.RowsAffected()
	if rowsAff != 1 {
		return err
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
