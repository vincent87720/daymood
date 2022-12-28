package model

import (
	"database/sql"
	"errors"

	"github.com/lib/pq"
)

//PERR0: Connection to the Database is lost.
//PERR1: Firm instance dosen't exists.
//PERR2: Invalid input
//PERR3: Update fail
//PERR4: ProductSku must be unique.

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
	DataOrder   *int64   //順序
	CreateAt    string   //建立時間
	UpdateAt    string   //最後編輯時間
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

func GetAllProducts(db *sql.DB) (productXi []Product, modelErr *ModelError) {
	err := db.Ping()
	if err != nil {
		return nil, connectionError("products")
	}

	row, err := db.Query("SELECT * FROM products ORDER BY id DESC;")
	if err != nil {
		return nil, normalError("products", err)
	}
	defer row.Close()

	var product Product
	for row.Next() {
		err := row.Scan(&product.ID, &product.SKU, &product.Name,
			&product.ProductType, &product.ImgName, &product.ImgID,
			&product.Stocks, &product.Weight, &product.RetailPrice,
			&product.Remark, &product.DataOrder, &product.CreateAt,
			&product.UpdateAt)
		if err != nil {
			return nil, normalError("products", err)
		}
		productXi = append(productXi, product)
	}

	return productXi, nil
}

// func GetProducts(db *sql.DB, productSkuXi []string) (map[string]Product, error) {

// 	productMap := make(map[string]Product)
// 	err := db.Ping()
// 	if err != nil {
// 		return nil, err
// 	}

// 	porductSkuStr := strings.Join(productSkuXi, "','")
// 	stmt := `
// 	SELECT
// 	p.product_id, p.product_sku,
// 	p.product_name, p.product_type,
// 	p.product_img_name, p.product_img_id,
// 	p.stocks, p.weight,
// 	p.krw_wholesale_price, p.ntd_wholesale_price,
// 	p.ntd_list_price, p.ntd_selling_price,
// 	p.ntd_cnf, p.ntd_cost,
// 	p.data_order, p.create_at, p.update_at,
// 	p.purchase_product_id, f.firm_id, f.firm_name
// 	FROM product AS p INNER JOIN firm AS f ON p.firm_id = f.firm_id
// 	WHERE p.product_sku
// 	IN ('` + porductSkuStr + `')
// 	ORDER BY p.update_at DESC;`

// 	row, err := db.Query(stmt)
// 	if err, ok := err.(*pq.Error); ok {
// 		fmt.Println(err)
// 	}
// 	defer row.Close()

// 	var productSchema ProductSchema
// 	// var productXi []Product
// 	for row.Next() {
// 		err := row.Scan(&productSchema.ProductID, &productSchema.ProductSku,
// 			&productSchema.ProductName, &productSchema.ProductType,
// 			&productSchema.ProductImgName, &productSchema.ProductImgID,
// 			&productSchema.Stocks, &productSchema.Weight,
// 			&productSchema.KrwWholesalePrice, &productSchema.NtdWholesalePrice,
// 			&productSchema.NtdListPrice, &productSchema.NtdSellingPrice,
// 			&productSchema.NtdCnf, &productSchema.NtdCost,
// 			&productSchema.DataOrder, &productSchema.CreateAt, &productSchema.UpdateAt,
// 			&productSchema.PurchaseProductID, &productSchema.FirmID, &productSchema.FirmName)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		tempProduct := Product{
// 			FirmInfo: Firm{
// 				ID:   productSchema.FirmID,
// 				Name: productSchema.FirmName,
// 			},
// 			ProductID:         productSchema.ProductID,
// 			ProductSku:        productSchema.ProductSku,
// 			ProductName:       productSchema.ProductName,
// 			ProductType:       productSchema.ProductType,
// 			ProductImgName:    *productSchema.ProductImgName,
// 			ProductImgID:      *productSchema.ProductImgID,
// 			Stocks:            productSchema.Stocks,
// 			Weight:            productSchema.Weight,
// 			KrwWholesalePrice: productSchema.KrwWholesalePrice,
// 			NtdWholesalePrice: productSchema.NtdWholesalePrice,
// 			NtdListPrice:      productSchema.NtdListPrice,
// 			NtdSellingPrice:   productSchema.NtdSellingPrice,
// 			NtdCnf:            productSchema.NtdCnf,
// 			NtdCost:           productSchema.NtdCost,
// 			CreateAt:          productSchema.CreateAt,
// 			UpdateAt:          productSchema.UpdateAt,
// 			PurchaseProductID: productSchema.PurchaseProductID,
// 		}

// 		if productSchema.ProductImgName == nil {
// 			tempProduct.ProductImgName = ""
// 		} else {
// 			tempProduct.ProductImgName = *productSchema.ProductImgName
// 		}

// 		if productSchema.ProductImgID == nil {
// 			tempProduct.ProductImgID = ""
// 		} else {
// 			tempProduct.ProductImgID = *productSchema.ProductImgID
// 		}

// 		if productSchema.DataOrder == nil {
// 			tempProduct.DataOrder = ""
// 		} else {
// 			tempProduct.DataOrder = *productSchema.DataOrder
// 		}
// 		productMap[strconv.FormatInt(productSchema.ProductID, 10)] = tempProduct

// 		// productXi = append(productXi, tempProduct)
// 	}

// 	return productMap, nil
// }

// func GetProductsByID(db *sql.DB, productIdXi []string) (map[string]Product, error) {

// 	productMap := make(map[string]Product)
// 	err := db.Ping()
// 	if err != nil {
// 		return nil, err
// 	}

// 	porductIdStr := strings.Join(productIdXi, "','")
// 	stmt := `
// 	SELECT
// 	p.product_id, p.product_sku,
// 	p.product_name, p.product_type,
// 	p.product_img_name, p.product_img_id,
// 	p.stocks, p.weight,
// 	p.krw_wholesale_price, p.ntd_wholesale_price,
// 	p.ntd_list_price, p.ntd_selling_price,
// 	p.ntd_cnf, p.ntd_cost,
// 	p.data_order, p.create_at, p.update_at,
// 	p.purchase_product_id, f.firm_id, f.firm_name
// 	FROM product AS p INNER JOIN firm AS f ON p.firm_id = f.firm_id
// 	WHERE p.product_id
// 	IN ('` + porductIdStr + `')
// 	ORDER BY p.update_at DESC;`

// 	row, err := db.Query(stmt)
// 	if err, ok := err.(*pq.Error); ok {
// 		fmt.Println(err)
// 	}
// 	defer row.Close()

// 	var productSchema ProductSchema
// 	// var productXi []Product
// 	for row.Next() {
// 		err := row.Scan(&productSchema.ProductID, &productSchema.ProductSku,
// 			&productSchema.ProductName, &productSchema.ProductType,
// 			&productSchema.ProductImgName, &productSchema.ProductImgID,
// 			&productSchema.Stocks, &productSchema.Weight,
// 			&productSchema.KrwWholesalePrice, &productSchema.NtdWholesalePrice,
// 			&productSchema.NtdListPrice, &productSchema.NtdSellingPrice,
// 			&productSchema.NtdCnf, &productSchema.NtdCost,
// 			&productSchema.DataOrder, &productSchema.CreateAt, &productSchema.UpdateAt,
// 			&productSchema.PurchaseProductID, &productSchema.FirmID, &productSchema.FirmName)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		tempProduct := Product{
// 			FirmInfo: Firm{
// 				ID:   productSchema.FirmID,
// 				Name: productSchema.FirmName,
// 			},
// 			ProductID:         productSchema.ProductID,
// 			ProductSku:        productSchema.ProductSku,
// 			ProductName:       productSchema.ProductName,
// 			ProductType:       productSchema.ProductType,
// 			ProductImgName:    *productSchema.ProductImgName,
// 			ProductImgID:      *productSchema.ProductImgID,
// 			Stocks:            productSchema.Stocks,
// 			Weight:            productSchema.Weight,
// 			KrwWholesalePrice: productSchema.KrwWholesalePrice,
// 			NtdWholesalePrice: productSchema.NtdWholesalePrice,
// 			NtdListPrice:      productSchema.NtdListPrice,
// 			NtdSellingPrice:   productSchema.NtdSellingPrice,
// 			NtdCnf:            productSchema.NtdCnf,
// 			NtdCost:           productSchema.NtdCost,
// 			CreateAt:          productSchema.CreateAt,
// 			UpdateAt:          productSchema.UpdateAt,
// 			PurchaseProductID: productSchema.PurchaseProductID,
// 		}

// 		if productSchema.ProductImgName == nil {
// 			tempProduct.ProductImgName = ""
// 		} else {
// 			tempProduct.ProductImgName = *productSchema.ProductImgName
// 		}

// 		if productSchema.ProductImgID == nil {
// 			tempProduct.ProductImgID = ""
// 		} else {
// 			tempProduct.ProductImgID = *productSchema.ProductImgID
// 		}

// 		if productSchema.DataOrder == nil {
// 			tempProduct.DataOrder = ""
// 		} else {
// 			tempProduct.DataOrder = *productSchema.DataOrder
// 		}
// 		productMap[strconv.FormatInt(productSchema.ProductID, 10)] = tempProduct

// 		// productXi = append(productXi, tempProduct)
// 	}

// 	return productMap, nil
// }

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
		data_order
	)
	VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10);`

	stmt, err := db.Prepare(qstr)
	if err != nil {
		return normalError("products", err)
	}
	defer stmt.Close()

	res, err := stmt.Exec(
		product.SKU, product.Name, product.ProductType,
		product.ImgName, product.ImgID, product.Stocks,
		product.Weight, product.RetailPrice, product.Remark,
		product.DataOrder)
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

func (product *Product) Update(db *sql.DB) (modelErr *ModelError) {
	err := db.Ping()
	if err != nil {
		return connectionError("products")
	}

	stmt, err := db.Prepare("CALL updateProducts($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)")
	if err != nil {
		return normalError("products", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		product.ID, product.SKU, product.Name, product.ProductType,
		product.ImgName, product.ImgID, product.Stocks,
		product.Weight, product.RetailPrice, product.Remark,
		product.DataOrder)

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

// ID
// SKU
// Name
// ProductType
// ImgName
// ImgID
// Stocks
// Weight
// RetailPrice
// Remark
// DataOrder
// CreateAt
// UpdateAt

// id
// sku
// name
// productType
// imgName
// imgID
// stocks
// weight
// retailPrice
// remark
// dataOrder
// createAt
// updateAt

//流水號
//商品顯示編號
//商品名稱
//商品種類
//商品圖片
//商品圖片編號
//庫存
//重量
//售價
//備註
//順序
//建立時間
//最後編輯時間

// id
// sku
// name
// type
// img_id
// img_name
// stocks
// weight
// retail_price
// remark
// data_order
// create_at
// update_at
