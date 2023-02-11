# Daymood Database

- [版本資訊](#版本資訊)
- [指令](#指令)
- [建置](#建置dockerimages)
- [初始化資料庫](#初始化資料庫)
- [建立通用函式](#建立通用函式)
- [廠商（suppliers）](#廠商suppliers)
- [採購（purchases）](#採購（purchases）)
- [採購明細（purchaseDetails）](#採購明細（purchaseDetails）)
- [上架商品（products）](#上架商品（products）)
- [商品進貨歷史紀錄（productPurchaseHistories）](#商品進貨歷史紀錄（productPurchaseHistories）)
- [商品出貨歷史紀錄（productDeliveryHistories）](#商品出貨歷史紀錄（productDeliveryHistories）)
- [出貨（deliveryOrders）](#出貨（deliveryOrders）)
- [出貨明細（deliveryOrderDetails）](#出貨明細（deliveryOrderDetails）)
- [折扣（discounts）](#折扣（discounts）)
- [使用者（users）](#使用者（users）)
- [報表（reports）](#報表（reports）)

## 版本資訊
- psql (PostgreSQL) 15.1 (Debian 15.1-1.pgdg110+1)

## 指令

### 使用 Docker 執行 PostgreSQL
```sh
make dockerrun
```

### 使用 Docker 建立 PostgreSQL 資料庫

```bash

#查看所安裝的 PostgreSQL 版本
docker exec daymood-database psql -V

#查看當前存在的 Database Name
#-U：指定使用者
#-l：列出資料庫名稱
docker exec daymood-database psql -U postgres -l

#進入 PostgreSQL 的 CLI 命令列介面
#輸入\q離開
docker exec -it daymood-database psql -U postgres
```

### 查看狀態指令

```sql
--show database list
\l

--show tables
\dt

--show functions
\df

--show table info(include trigger info)
\dS [tableName]
```

## 初始化資料庫

```sql
--create user
CREATE USER daymood WITH PASSWORD 'daymooddev';

--create database
CREATE DATABASE daymood WITH OWNER daymood;
```

## 建立通用函式

```sql
--create timestamp update trigger
CREATE OR REPLACE FUNCTION onUpdate()
RETURNS TRIGGER AS $$
BEGIN
    NEW.update_at = now();
    RETURN NEW;
END;
$$ language 'plpgsql';
```

## 廠商（suppliers）

### CreateTableSQL

```sql
CREATE TABLE suppliers (
    id SERIAL PRIMARY KEY,--流水號
    name VARCHAR(256) NOT NULL,--廠商名稱
    address VARCHAR(256),--廠商地址
    remark VARCHAR(256),--備註
    data_status INTEGER DEFAULT 1,--是否啟用
    create_at TIMESTAMP(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,--建立時間
    update_at TIMESTAMP(6) NOT NULL DEFAULT CURRENT_TIMESTAMP--最後編輯時間
);
CREATE TRIGGER suppliers_update_at BEFORE UPDATE ON suppliers FOR EACH ROW EXECUTE PROCEDURE onUpdate();
```

### 更新 suppliers 資料表預存程序

```sql
CREATE OR REPLACE PROCEDURE updateSuppliers(
    id INTEGER,
    name TEXT,
    address TEXT,
    remark TEXT,
    data_status INTEGER
)
LANGUAGE plpgsql
AS $$
BEGIN
    UPDATE suppliers
       SET name = COALESCE(updateSuppliers.name, suppliers.name),
		   address = COALESCE(updateSuppliers.address, suppliers.address),
		   remark = COALESCE(updateSuppliers.remark, suppliers.remark),
		   data_status = COALESCE(updateSuppliers.data_status, suppliers.data_status)
     WHERE suppliers.id = updateSuppliers.id;
END;
$$;
```

## 採購（purchases）

### CreateTableSQL

```sql
CREATE TABLE purchases (
    id SERIAL PRIMARY KEY,--流水號
    name VARCHAR(256) NOT NULL,--採購名稱
    status INTEGER NOT NULL,--採購狀態(0:採購中,1:結案)
    purchase_type INTEGER NOT NULL,--採購類型(0:商品,1:耗材)
    qty INTEGER,--商品總數
    shipping_agent VARCHAR(256),--貨運行
    shipping_agent_cut_krw REAL,--貨運行抽成
    shipping_agent_cut_percent REAL,--貨運行抽成趴數
    shipping_initiator VARCHAR(256),--貨運團主
    shipping_create_at TIMESTAMP(6),--貨運開團日期
    shipping_end_at TIMESTAMP(6),--貨運結單日期
    shipping_arrive_at TIMESTAMP(6),--貨運送達日期
    weight REAL,--貨運總重
    shipping_fee_kr REAL,--貨運國內運費_韓國
    shipping_fee_tw REAL,--貨運國內運費_台灣
    shipping_fee_kokusai_krw REAL,--貨運國際運費
    shipping_fee_kokusai_per_kilo REAL,--每公斤貨運國際運費
    exchange_rate_krw REAL,--韓圓匯率
    tariff_twd REAL,--關稅
    tariff_per_kilo REAL,--每公斤關稅
    total_krw REAL,--韓圓總價
    total_twd REAL,--台幣總價
    total REAL,--總金額
    remark VARCHAR(256),--備註
    data_order INTEGER,--順序
    create_at TIMESTAMP(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,--建立時間
    update_at TIMESTAMP(6) NOT NULL DEFAULT CURRENT_TIMESTAMP--最後編輯時間
);
CREATE TRIGGER purchases_update_at BEFORE UPDATE ON purchases FOR EACH ROW EXECUTE PROCEDURE onUpdate();
```

### 更新 purchases 資料表預存程序

```sql
CREATE OR REPLACE PROCEDURE updatePurchases(
    id INTEGER,
    name TEXT,
    status INTEGER,
    purchase_type INTEGER,
    qty INTEGER,
    shipping_agent TEXT,
    shipping_agent_cut_krw REAL,
    shipping_agent_cut_percent REAL,
    shipping_initiator TEXT,
    shipping_create_at TIMESTAMP,
    shipping_end_at TIMESTAMP,
    shipping_arrive_at TIMESTAMP,
    weight REAL,
    shipping_fee_kr REAL,
    shipping_fee_tw REAL,
    shipping_fee_kokusai_krw REAL,
    shipping_fee_kokusai_per_kilo REAL,
    exchange_rate_krw REAL,
    tariff_twd REAL,
    tariff_per_kilo REAL,
    total_krw REAL,
    total_twd REAL,
    total REAL,
    remark TEXT,
    data_order INTEGER
)
LANGUAGE plpgsql
AS $$
BEGIN
    UPDATE purchases
       SET name = COALESCE(updatePurchases.name, purchases.name),
           status = COALESCE(updatePurchases.status, purchases.status),
           purchase_type = COALESCE(updatePurchases.purchase_type, purchases.purchase_type),
           qty = COALESCE(updatePurchases.qty, purchases.qty),
           shipping_agent = COALESCE(updatePurchases.shipping_agent, purchases.shipping_agent),
           shipping_agent_cut_krw = COALESCE(updatePurchases.shipping_agent_cut_krw, purchases.shipping_agent_cut_krw),
           shipping_agent_cut_percent = COALESCE(updatePurchases.shipping_agent_cut_percent, purchases.shipping_agent_cut_percent),
           shipping_initiator = COALESCE(updatePurchases.shipping_initiator, purchases.shipping_initiator),
           shipping_create_at = COALESCE(updatePurchases.shipping_create_at, purchases.shipping_create_at),
           shipping_end_at = COALESCE(updatePurchases.shipping_end_at, purchases.shipping_end_at),
           shipping_arrive_at = COALESCE(updatePurchases.shipping_arrive_at, purchases.shipping_arrive_at),
           weight = COALESCE(updatePurchases.weight, purchases.weight),
           shipping_fee_kr = COALESCE(updatePurchases.shipping_fee_kr, purchases.shipping_fee_kr),
           shipping_fee_tw = COALESCE(updatePurchases.shipping_fee_tw, purchases.shipping_fee_tw),
           shipping_fee_kokusai_krw = COALESCE(updatePurchases.shipping_fee_kokusai_krw, purchases.shipping_fee_kokusai_krw),
           shipping_fee_kokusai_per_kilo = COALESCE(updatePurchases.shipping_fee_kokusai_per_kilo, purchases.shipping_fee_kokusai_per_kilo),
           exchange_rate_krw = COALESCE(updatePurchases.exchange_rate_krw, purchases.exchange_rate_krw),
           tariff_twd = COALESCE(updatePurchases.tariff_twd, purchases.tariff_twd),
           tariff_per_kilo = COALESCE(updatePurchases.tariff_per_kilo, purchases.tariff_per_kilo),
           total_krw = COALESCE(updatePurchases.total_krw, purchases.total_krw),
           total_twd = COALESCE(updatePurchases.total_twd, purchases.total_twd),
           total = COALESCE(updatePurchases.total, purchases.total),
           remark = COALESCE(updatePurchases.remark, purchases.remark),
           data_order = COALESCE(updatePurchases.data_order, purchases.data_order)
     WHERE purchases.id = updatePurchases.id;
END;
$$;
```

## 上架商品（products）

### CreateTableSQL

```sql
CREATE TABLE products (
    id SERIAL PRIMARY KEY,--商品編號
    sku VARCHAR(36) UNIQUE, --商品顯示編號
    name VARCHAR(256) NOT NULL,--商品名稱
    type INTEGER NOT NULL,--商品種類（0:耳環, 1:耳骨夾, 2:戒指, 3:項鍊, 4:手鍊, 5:紙盒, 6:氣泡紙, 7:拭銀布, 8:破壞袋, 9:夾鏈袋, 10:飾品卡, 11:留言卡, 12:香芬）
    img_id VARCHAR(256),--商品照檔案實際名稱
    img_name VARCHAR(256),--商品照檔案顯示名稱
    stocks INTEGER NOT NULL,--庫存
    weight REAL,--商品重量
    retail_price REAL NOT NULL,--售價（若商品種類為耗材，售價為零）
    remark VARCHAR(256),--備註
    data_status INTEGER DEFAULT 1,--是否啟用
    create_at TIMESTAMP(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,--建立時間
    update_at TIMESTAMP(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,--最後編輯時間
    supplier_id INTEGER,--廠商編號
    CONSTRAINT fk_supplier
      FOREIGN KEY(supplier_id)
	  REFERENCES suppliers(id)
	  ON DELETE SET NULL
      ON UPDATE NO ACTION
);
CREATE TRIGGER products_update_at BEFORE UPDATE ON products FOR EACH ROW EXECUTE PROCEDURE onUpdate();
```

### 更新 products 資料表預存程序

```sql
CREATE OR REPLACE PROCEDURE updateProducts(
    id INTEGER,
    sku TEXT,
    name TEXT,
    type INTEGER,
    img_id TEXT,
    img_name TEXT,
    stocks INTEGER,
    weight REAL,
    retail_price REAL,
    remark TEXT,
    data_status INTEGER,
	supplier_id INTEGER
)
LANGUAGE plpgsql
AS $$
BEGIN
    UPDATE products
       SET sku = COALESCE(updateProducts.sku, products.sku),
           name = COALESCE(updateProducts.name, products.name),
           type = COALESCE(updateProducts.type, products.type),
           img_id = COALESCE(updateProducts.img_id, products.img_id),
           img_name = COALESCE(updateProducts.img_name, products.img_name),
           stocks = COALESCE(updateProducts.stocks, products.stocks),
           weight = COALESCE(updateProducts.weight, products.weight),
           retail_price = COALESCE(updateProducts.retail_price, products.retail_price),
           remark = COALESCE(updateProducts.remark, products.remark),
           data_status = COALESCE(updateProducts.data_status, products.data_status),
	       supplier_id = updateProducts.supplier_id
     WHERE products.id = updateProducts.id;
END;
$$;
```

## 商品進貨歷史紀錄（productPurchaseHistories）

### 查詢 productPurchaseHistories 資料表預存程序

```sql
DROP FUNCTION selectProductPurchaseHistories;
CREATE OR REPLACE FUNCTION selectProductPurchaseHistories(product_id INTEGER)
RETURNS TABLE (
	purchase_id INTEGER,
	purchase_name VARCHAR(256),
    purchase_qty INTEGER,
	purchase_detail_id INTEGER,
	purchase_detail_name VARCHAR(256),
	purchase_detail_qty INTEGER,
	wholesale_price REAL,
	subtotal REAL,
	shipping_arrive_at TIMESTAMP(6),
	supplier_id INTEGER,
	exchange_rate_krw REAL,
	shipping_agent_cut_percent REAL,
	shipping_fee_kr REAL,
	shipping_fee_tw REAL,
	shipping_fee_kokusai_krw REAL,
	tariff_twd REAL
)
LANGUAGE plpgsql AS
$func$
BEGIN
   RETURN QUERY
   SELECT
        a.purchase_id,
        b.name AS purchase_name,
		b.qty AS purchase_qty,
        a.id AS purchase_detail_id,
        a.name AS purchase_detail_name,
        a.qty AS purchase_detail_qty,
        a.wholesale_price,
        a.subtotal,
        b.shipping_arrive_at,
        a.supplier_id,
		b.exchange_rate_krw,
		b.shipping_agent_cut_percent,
		b.shipping_fee_kr,
		b.shipping_fee_tw,
		b.shipping_fee_kokusai_krw,
		b.tariff_twd
    FROM
        purchaseDetails AS a
        INNER JOIN
        purchases AS b
            ON a.purchase_id = b.id
    WHERE a.product_id = selectProductPurchaseHistories.product_id AND a.Status = 1;
END
$func$;
```

## 商品出貨歷史紀錄（productDeliveryHistories）

### 查詢 ProductDeliveryHistories 資料表預存程序

```sql
DROP FUNCTION selectProductDeliveryHistories;
CREATE OR REPLACE FUNCTION selectProductDeliveryHistories(product_id INTEGER)
RETURNS TABLE (
	delivery_order_id INTEGER,
	delivery_order_detail_id INTEGER,
	retail_price REAL,
	cost REAL,
	qty INTEGER,
	subtotal REAL,
    order_at TIMESTAMP
)
LANGUAGE plpgsql AS
$func$
BEGIN
   RETURN QUERY
   SELECT
        a.delivery_order_id,
        a.id AS delivery_order_detail_id,
		a.retail_price,
		a.cost,
        a.qty,
        a.subtotal,
        b.order_at
    FROM
        deliveryOrderDetails AS a
        INNER JOIN
        deliveryOrders AS b
            ON a.delivery_order_id = b.id
    WHERE a.product_id = selectProductDeliveryHistories.product_id AND b.Status = 2;
END
$func$;
```

## 採購明細（purchaseDetails）

### CreateTableSQL

```sql
CREATE TABLE purchaseDetails (
    id SERIAL PRIMARY KEY,--流水號
    named_id VARCHAR(256),--採購明細編號（識別編號）
    name VARCHAR(256) NOT NULL,--商品名稱
    status INTEGER NOT NULL,--是否採用（0:不採用,1:採用）
    wholesale_price REAL NOT NULL,--批價
    qty INTEGER NOT NULL,--數量
    cost REAL,--成本
    currency INTEGER,--幣別
    subtotal REAL NOT NULL,--小計
    remark VARCHAR(256),--備註
    data_order INTEGER,--順序
    create_at TIMESTAMP(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,--建立時間
    update_at TIMESTAMP(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,--最後編輯時間
    purchase_id INTEGER NOT NULL,--採購編號（為採購明細對應的主檔，必須存在）
    supplier_id INTEGER,--廠商編號（因商品廠商可能會變動，此欄位用於記載採購時的進貨廠商）
    product_id INTEGER,--商品編號（必須知道此採購商品對應到的上架商品，以便新增庫存量）
    CONSTRAINT fk_purchase
      FOREIGN KEY(purchase_id)
	  REFERENCES purchases(id)
	  ON DELETE NO ACTION
      ON UPDATE NO ACTION,
    CONSTRAINT fk_product
      FOREIGN KEY(product_id)
	  REFERENCES products(id)
	  ON DELETE SET NULL
      ON UPDATE NO ACTION,
    CONSTRAINT fk_supplier
      FOREIGN KEY(supplier_id)
	  REFERENCES suppliers(id)
	  ON DELETE SET NULL
      ON UPDATE NO ACTION
);
CREATE TRIGGER purchase_details_update_at BEFORE UPDATE ON purchaseDetails FOR EACH ROW EXECUTE PROCEDURE onUpdate();
```

### 更新 purchaseDetails 資料表預存程序

```sql
CREATE OR REPLACE PROCEDURE updatePurchaseDetails(
	id INTEGER,
	named_id TEXT,
	name TEXT,
	status INTEGER,
	wholesale_price REAL,
	qty INTEGER,
	cost REAL,
	currency INTEGER,
	subtotal REAL,
	remark TEXT,
	data_order INTEGER,
	purchase_id INTEGER,
	supplier_id INTEGER,
	product_id INTEGER
)
LANGUAGE plpgsql
AS $$
BEGIN
    UPDATE purchaseDetails
       SET named_id = COALESCE(updatePurchaseDetails.named_id, purchaseDetails.named_id),
	       name = COALESCE(updatePurchaseDetails.name, purchaseDetails.name),
	       status = COALESCE(updatePurchaseDetails.status, purchaseDetails.status),
	       wholesale_price = COALESCE(updatePurchaseDetails.wholesale_price, purchaseDetails.wholesale_price),
	       qty = COALESCE(updatePurchaseDetails.qty, purchaseDetails.qty),
	       cost = COALESCE(updatePurchaseDetails.cost, purchaseDetails.cost),
	       currency = COALESCE(updatePurchaseDetails.currency, purchaseDetails.currency),
	       subtotal = COALESCE(updatePurchaseDetails.subtotal, purchaseDetails.subtotal),
	       remark = COALESCE(updatePurchaseDetails.remark, purchaseDetails.remark),
	       data_order = COALESCE(updatePurchaseDetails.data_order, purchaseDetails.data_order),
	       purchase_id = updatePurchaseDetails.purchase_id,
	       supplier_id = updatePurchaseDetails.supplier_id,
	       product_id = updatePurchaseDetails.product_id
     WHERE purchaseDetails.id = updatePurchaseDetails.id;
END;
$$;
```

## 出貨（deliveryOrders）

### CreateTableSQL

```sql
CREATE TABLE deliveryOrders (
    id SERIAL PRIMARY KEY,--出貨明細編號
    status INTEGER NOT NULL,--處理狀態
    delivery_type INTEGER ,--出貨方式（0:宅配, 1:7-11, 2:全家, 3:萊爾富）
    delivery_status INTEGER,--出貨狀態（0:未出貨, 1:已出貨）
    delivery_fee_status INTEGER,--運費狀態（0:運費由平台支付, 1:運費由賣家支付, 2:運費由買家支付）
    payment_type INTEGER,--付款方式（0:ATM轉帳, 1:信用卡支付, 2:超商取貨付款）
    payment_status INTEGER NOT NULL,--付款狀態（0:未付款, 1:已付款）
    total_original REAL NOT NULL,--原價
    discount REAL NOT NULL,--折扣
    total_discounted REAL NOT NULL,--總價
    remark VARCHAR(256),--備註
    data_order INTEGER,--順序
    order_at TIMESTAMP(6),--下訂日期
    send_at TIMESTAMP(6),--出貨日期
    arrive_at TIMESTAMP(6),--送達日期
    create_at TIMESTAMP(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,--建立時間
    update_at TIMESTAMP(6) NOT NULL DEFAULT CURRENT_TIMESTAMP--最後編輯時間
);
CREATE TRIGGER delivery_orders_update_at BEFORE UPDATE ON deliveryOrders FOR EACH ROW EXECUTE PROCEDURE onUpdate();
```

### 更新 deliveryOrders 資料表預存程序

```sql
CREATE OR REPLACE PROCEDURE updateDeliveryOrders(
    id INTEGER,
    status INTEGER,
    delivery_type INTEGER,
    delivery_status INTEGER,
    delivery_fee_status INTEGER,
    payment_type INTEGER,
    payment_status INTEGER,
    total_original REAL,
    discount REAL,
    total_discounted REAL,
    remark TEXT,
    data_order INTEGER,
    order_at TIMESTAMP,
    send_at TIMESTAMP,
    arrive_at TIMESTAMP
)
LANGUAGE plpgsql
AS $$
BEGIN
    UPDATE deliveryOrders
       SET status = COALESCE(updateDeliveryOrders.status, deliveryOrders.status),
           delivery_type = COALESCE(updateDeliveryOrders.delivery_type, deliveryOrders.delivery_type),
           delivery_status = COALESCE(updateDeliveryOrders.delivery_status, deliveryOrders.delivery_status),
           delivery_fee_status = COALESCE(updateDeliveryOrders.delivery_fee_status, deliveryOrders.delivery_fee_status),
           payment_type = COALESCE(updateDeliveryOrders.payment_type, deliveryOrders.payment_type),
           payment_status = COALESCE(updateDeliveryOrders.payment_status, deliveryOrders.payment_status),
           total_original = COALESCE(updateDeliveryOrders.total_original, deliveryOrders.total_original),
           discount = COALESCE(updateDeliveryOrders.discount, deliveryOrders.discount),
           total_discounted = COALESCE(updateDeliveryOrders.total_discounted, deliveryOrders.total_discounted),
           remark = COALESCE(updateDeliveryOrders.remark, deliveryOrders.remark),
           data_order = COALESCE(updateDeliveryOrders.data_order, deliveryOrders.data_order),
           order_at = COALESCE(updateDeliveryOrders.order_at, deliveryOrders.order_at),
           send_at = COALESCE(updateDeliveryOrders.send_at, deliveryOrders.send_at),
           arrive_at = COALESCE(updateDeliveryOrders.arrive_at, deliveryOrders.arrive_at)
     WHERE deliveryOrders.id = updateDeliveryOrders.id;
END;
$$;
```

## 出貨明細（deliveryOrderDetails）

### CreateTableSQL

```sql
CREATE TABLE deliveryOrderDetails (
    id SERIAL PRIMARY KEY,--流水號
    retail_price REAL NOT NULL,--出貨時售價（因商品售價可能會變動，此欄位紀錄出貨時的商品售價）
    cost REAL NOT NULL,--出貨時成本
    qty INTEGER NOT NULL,--數量
    subtotal REAL NOT NULL,--小計
    remark VARCHAR(256),--備註
    data_order INTEGER,--順序
    create_at TIMESTAMP(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,--建立時間
    update_at TIMESTAMP(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,--最後編輯時間
    delivery_order_id INTEGER NOT NULL,--出貨編號（為出貨明細對應的主檔，必須存在）
    product_id INTEGER,--商品編號（不一定需要存在，便於查找目前對應到的商品）
    CONSTRAINT fk_delivery_order
      FOREIGN KEY(delivery_order_id)
	  REFERENCES deliveryOrders(id)
	  ON DELETE NO ACTION
      ON UPDATE NO ACTION,
    CONSTRAINT fk_product
      FOREIGN KEY(product_id)
	  REFERENCES products(id)
	  ON DELETE SET NULL
      ON UPDATE NO ACTION
);
CREATE TRIGGER delivery_order_details_update_at BEFORE UPDATE ON deliveryOrderDetails FOR EACH ROW EXECUTE PROCEDURE onUpdate();
```

### 更新 deliveryOrderDetails 資料表預存程序

```sql
CREATE OR REPLACE PROCEDURE updateDeliveryOrderDetails(
    id INTEGER,
	retail_price REAL,
    cost REAL,
	qty INTEGER,
	subtotal REAL,
	remark TEXT,
	data_order INTEGER,
	delivery_order_id INTEGER,
	product_id INTEGER
)
LANGUAGE plpgsql
AS $$
BEGIN
    UPDATE deliveryOrderDetails
       SET retail_price = COALESCE(updateDeliveryOrderDetails.retail_price, deliveryOrderDetails.retail_price),
           cost = COALESCE(updateDeliveryOrderDetails.cost, deliveryOrderDetails.cost),
		   qty = COALESCE(updateDeliveryOrderDetails.qty, deliveryOrderDetails.qty),
		   subtotal = COALESCE(updateDeliveryOrderDetails.subtotal, deliveryOrderDetails.subtotal),
		   remark = COALESCE(updateDeliveryOrderDetails.remark, deliveryOrderDetails.remark),
		   data_order = COALESCE(updateDeliveryOrderDetails.data_order, deliveryOrderDetails.data_order),
		   delivery_order_id = COALESCE(updateDeliveryOrderDetails.delivery_order_id, deliveryOrderDetails.delivery_order_id),
		   product_id = COALESCE(updateDeliveryOrderDetails.product_id, deliveryOrderDetails.product_id)
     WHERE deliveryOrderDetails.id = updateDeliveryOrderDetails.id;
END;
$$;
```

## 折扣（discounts）

### CreateTableSQL

```sql
CREATE TABLE discounts (
    id SERIAL PRIMARY KEY,--流水號
    name VARCHAR(256) NOT NULL,--折扣名稱
    price REAL NOT NULL,--折扣金額
    discount_type INTEGER NOT NULL,--折扣方式（0:整筆訂單折扣, 1:單件商品折扣）
    remark VARCHAR(256),--備註
    data_order INTEGER,--順序
    create_at TIMESTAMP(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,--建立時間
    update_at TIMESTAMP(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,--最後編輯時間
    delivery_order_id INTEGER NOT NULL,--出貨編號
    CONSTRAINT fk_delivery_order
      FOREIGN KEY(delivery_order_id)
	  REFERENCES deliveryOrders(id)
	  ON DELETE CASCADE
      ON UPDATE NO ACTION
);
CREATE TRIGGER discounts_update_at BEFORE UPDATE ON discounts FOR EACH ROW EXECUTE PROCEDURE onUpdate();
```

### 更新 discounts 資料表預存程序

```sql
CREATE OR REPLACE PROCEDURE updateDiscounts(
	id INTEGER,
	name TEXT,
	price REAL,
	discount_type INTEGER,
	remark TEXT,
	data_order INTEGER,
	delivery_order_id INTEGER
)
LANGUAGE plpgsql
AS $$
BEGIN
    UPDATE discounts
       SET name = COALESCE(updateDiscounts.name, discounts.name),
		   price = COALESCE(updateDiscounts.price, discounts.price),
		   discount_type = COALESCE(updateDiscounts.discount_type, discounts.discount_type),
		   remark = COALESCE(updateDiscounts.remark, discounts.remark),
		   data_order = COALESCE(updateDiscounts.data_order, discounts.data_order),
		   delivery_order_id = COALESCE(updateDiscounts.delivery_order_id, discounts.delivery_order_id)
     WHERE discounts.id = updateDiscounts.id;
END;
$$;
```

## 使用者（users）

### CreateTableSQL

```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,--流水號
    username VARCHAR(256) UNIQUE NOT NULL,--帳號
    password VARCHAR(256) NOT NULL,--密碼
    name VARCHAR(256) NOT NULL,--名稱
    email VARCHAR(256) UNIQUE NOT NULL,--E-mail
    create_at TIMESTAMP(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,--建立時間
    update_at TIMESTAMP(6) NOT NULL DEFAULT CURRENT_TIMESTAMP--最後編輯時間
);
CREATE TRIGGER users_update_at BEFORE UPDATE ON users FOR EACH ROW EXECUTE PROCEDURE onUpdate();
```

### 更新 users 資料表預存程序

```sql
CREATE OR REPLACE PROCEDURE updateUsers(
	id INTEGER,
	username TEXT,
	name TEXT,
	email TEXT
)
LANGUAGE plpgsql
AS $$
BEGIN
    UPDATE users
       SET username = COALESCE(updateUsers.username, users.username),
		   name = COALESCE(updateUsers.name, users.name),
		   email = COALESCE(updateUsers.email, users.email)
    WHERE users.id = updateUsers.id;
END;
$$;
```

### 更新 users 資料表 password 欄位預存程序

```sql
CREATE OR REPLACE PROCEDURE updateUserPasswords(
	id INTEGER,
	password TEXT
)
LANGUAGE plpgsql
AS $$
BEGIN
    UPDATE users
       SET password = COALESCE(updateUserPasswords.password, users.password)
    WHERE users.id = updateUserPasswords.id;
END;
$$;
```

## 報表（reports）

### 查詢總收支預存程序

```sql
DROP FUNCTION selectBalance;
CREATE OR REPLACE FUNCTION selectBalance()
RETURNS TABLE (
	purchase_total REAL,
	delivery_total REAL
)
LANGUAGE plpgsql AS
$func$
BEGIN
   RETURN QUERY
    SELECT
        purchases.purchase_total, deliveryOrders.delivery_total
    FROM
        (
            SELECT SUM(total) AS purchase_total
            FROM purchases
        ) AS purchases,
        (
            SELECT SUM(total_original) AS delivery_total
            FROM deliveryOrders
        ) AS deliveryOrders;
END
$func$;
```
## 匯出
```
make dump
```

## 建置
[DaymoodReadme#Build](../README.md#建置)