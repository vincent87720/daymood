--
-- PostgreSQL database dump
--

-- Dumped from database version 15.1 (Debian 15.1-1.pgdg110+1)
-- Dumped by pg_dump version 15.1 (Debian 15.1-1.pgdg110+1)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: onupdate(); Type: FUNCTION; Schema: public; Owner: daymood
--

CREATE FUNCTION public.onupdate() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    NEW.update_at = now();
    RETURN NEW;
END;
$$;


ALTER FUNCTION public.onupdate() OWNER TO daymood;

--
-- Name: selectbalance(); Type: FUNCTION; Schema: public; Owner: daymood
--

CREATE FUNCTION public.selectbalance() RETURNS TABLE(purchase_total real, delivery_total real)
    LANGUAGE plpgsql
    AS $$
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
$$;


ALTER FUNCTION public.selectbalance() OWNER TO daymood;

--
-- Name: selectproductdeliveryhistories(integer); Type: FUNCTION; Schema: public; Owner: daymood
--

CREATE FUNCTION public.selectproductdeliveryhistories(product_id integer) RETURNS TABLE(delivery_order_id integer, delivery_order_detail_id integer, retail_price real, cost real, qty integer, subtotal real, order_at timestamp without time zone)
    LANGUAGE plpgsql
    AS $$
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
$$;


ALTER FUNCTION public.selectproductdeliveryhistories(product_id integer) OWNER TO daymood;

--
-- Name: selectproductpurchasehistories(integer); Type: FUNCTION; Schema: public; Owner: daymood
--

CREATE FUNCTION public.selectproductpurchasehistories(product_id integer) RETURNS TABLE(purchase_id integer, purchase_name character varying, purchase_qty integer, purchase_detail_id integer, purchase_detail_name character varying, purchase_detail_qty integer, wholesale_price real, subtotal real, shipping_arrive_at timestamp without time zone, supplier_id integer, exchange_rate_krw real, shipping_agent_cut_percent real, shipping_fee_kr real, shipping_fee_tw real, shipping_fee_kokusai_krw real, tariff_twd real)
    LANGUAGE plpgsql
    AS $$
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
$$;


ALTER FUNCTION public.selectproductpurchasehistories(product_id integer) OWNER TO daymood;

--
-- Name: truncatetables(character varying); Type: FUNCTION; Schema: public; Owner: daymood
--

CREATE FUNCTION public.truncatetables(username character varying) RETURNS void
    LANGUAGE plpgsql
    AS $$
DECLARE
    statements CURSOR FOR
        SELECT tablename FROM pg_tables
        WHERE tableowner = username AND schemaname = 'public';
BEGIN
    FOR stmt IN statements LOOP
        EXECUTE 'TRUNCATE TABLE ' || quote_ident(stmt.tablename) || ' CASCADE;';
    END LOOP;
END;
$$;


ALTER FUNCTION public.truncatetables(username character varying) OWNER TO daymood;

--
-- Name: updatedeliveryorderdetails(integer, real, real, integer, real, text, integer, integer, integer); Type: PROCEDURE; Schema: public; Owner: daymood
--

CREATE PROCEDURE public.updatedeliveryorderdetails(IN id integer, IN retail_price real, IN cost real, IN qty integer, IN subtotal real, IN remark text, IN data_order integer, IN delivery_order_id integer, IN product_id integer)
    LANGUAGE plpgsql
    AS $$
BEGIN
    UPDATE deliveryOrderDetails
       SET retail_price         = COALESCE(updateDeliveryOrderDetails.retail_price, deliveryOrderDetails.retail_price),
           cost                 = COALESCE(updateDeliveryOrderDetails.cost, deliveryOrderDetails.cost),
		   qty                  = COALESCE(updateDeliveryOrderDetails.qty, deliveryOrderDetails.qty),
		   subtotal             = COALESCE(updateDeliveryOrderDetails.subtotal, deliveryOrderDetails.subtotal),
		   remark               = COALESCE(updateDeliveryOrderDetails.remark, deliveryOrderDetails.remark),
		   data_order           = COALESCE(updateDeliveryOrderDetails.data_order, deliveryOrderDetails.data_order),
		   delivery_order_id    = COALESCE(updateDeliveryOrderDetails.delivery_order_id, deliveryOrderDetails.delivery_order_id),
		   product_id           = COALESCE(updateDeliveryOrderDetails.product_id, deliveryOrderDetails.product_id)
     WHERE deliveryOrderDetails.id = updateDeliveryOrderDetails.id;
END;
$$;


ALTER PROCEDURE public.updatedeliveryorderdetails(IN id integer, IN retail_price real, IN cost real, IN qty integer, IN subtotal real, IN remark text, IN data_order integer, IN delivery_order_id integer, IN product_id integer) OWNER TO daymood;

--
-- Name: updatedeliveryorders(integer, integer, integer, integer, integer, integer, integer, real, real, real, text, integer, timestamp without time zone, timestamp without time zone, timestamp without time zone); Type: PROCEDURE; Schema: public; Owner: daymood
--

CREATE PROCEDURE public.updatedeliveryorders(IN id integer, IN status integer, IN delivery_type integer, IN delivery_status integer, IN delivery_fee_status integer, IN payment_type integer, IN payment_status integer, IN total_original real, IN discount real, IN total_discounted real, IN remark text, IN data_order integer, IN order_at timestamp without time zone, IN send_at timestamp without time zone, IN arrive_at timestamp without time zone)
    LANGUAGE plpgsql
    AS $$
BEGIN
    UPDATE deliveryOrders
       SET status                     = COALESCE(updateDeliveryOrders.status, deliveryOrders.status),
           delivery_type              = COALESCE(updateDeliveryOrders.delivery_type, deliveryOrders.delivery_type),
           delivery_status            = COALESCE(updateDeliveryOrders.delivery_status, deliveryOrders.delivery_status),
           delivery_fee_status        = COALESCE(updateDeliveryOrders.delivery_fee_status, deliveryOrders.delivery_fee_status),
           payment_type               = COALESCE(updateDeliveryOrders.payment_type, deliveryOrders.payment_type),
           payment_status             = COALESCE(updateDeliveryOrders.payment_status, deliveryOrders.payment_status),
           total_original             = COALESCE(updateDeliveryOrders.total_original, deliveryOrders.total_original),
           discount                   = COALESCE(updateDeliveryOrders.discount, deliveryOrders.discount),
           total_discounted           = COALESCE(updateDeliveryOrders.total_discounted, deliveryOrders.total_discounted),
           remark                     = COALESCE(updateDeliveryOrders.remark, deliveryOrders.remark),
           data_order                 = COALESCE(updateDeliveryOrders.data_order, deliveryOrders.data_order),
           order_at                   = COALESCE(updateDeliveryOrders.order_at, deliveryOrders.order_at),
           send_at                    = COALESCE(updateDeliveryOrders.send_at, deliveryOrders.send_at),
           arrive_at                  = COALESCE(updateDeliveryOrders.arrive_at, deliveryOrders.arrive_at)
     WHERE deliveryOrders.id = updateDeliveryOrders.id;
END;
$$;


ALTER PROCEDURE public.updatedeliveryorders(IN id integer, IN status integer, IN delivery_type integer, IN delivery_status integer, IN delivery_fee_status integer, IN payment_type integer, IN payment_status integer, IN total_original real, IN discount real, IN total_discounted real, IN remark text, IN data_order integer, IN order_at timestamp without time zone, IN send_at timestamp without time zone, IN arrive_at timestamp without time zone) OWNER TO daymood;

--
-- Name: updatediscounts(integer, text, real, integer, text, integer, integer); Type: PROCEDURE; Schema: public; Owner: daymood
--

CREATE PROCEDURE public.updatediscounts(IN id integer, IN name text, IN price real, IN discount_type integer, IN remark text, IN data_order integer, IN delivery_order_id integer)
    LANGUAGE plpgsql
    AS $$
BEGIN
    UPDATE discounts
       SET name                    = COALESCE(updateDiscounts.name, discounts.name),
		   price                   = COALESCE(updateDiscounts.price, discounts.price),
		   discount_type           = COALESCE(updateDiscounts.discount_type, discounts.discount_type),
		   remark                  = COALESCE(updateDiscounts.remark, discounts.remark),
		   data_order              = COALESCE(updateDiscounts.data_order, discounts.data_order),
		   delivery_order_id       = COALESCE(updateDiscounts.delivery_order_id, discounts.delivery_order_id)
     WHERE discounts.id = updateDiscounts.id;
END;
$$;


ALTER PROCEDURE public.updatediscounts(IN id integer, IN name text, IN price real, IN discount_type integer, IN remark text, IN data_order integer, IN delivery_order_id integer) OWNER TO daymood;

--
-- Name: updateproducts(integer, text, text, integer, text, text, integer, real, real, text, integer); Type: PROCEDURE; Schema: public; Owner: daymood
--

CREATE PROCEDURE public.updateproducts(IN id integer, IN sku text, IN name text, IN type integer, IN img_id text, IN img_name text, IN stocks integer, IN weight real, IN retail_price real, IN remark text, IN data_status integer)
    LANGUAGE plpgsql
    AS $$
BEGIN
    UPDATE products
       SET sku            = COALESCE(updateProducts.sku, products.sku),
           name           = COALESCE(updateProducts.name, products.name),
           type           = COALESCE(updateProducts.type, products.type),
           img_id         = COALESCE(updateProducts.img_id, products.img_id),
           img_name       = COALESCE(updateProducts.img_name, products.img_name),
           stocks         = COALESCE(updateProducts.stocks, products.stocks),
           weight         = COALESCE(updateProducts.weight, products.weight),
           retail_price   = COALESCE(updateProducts.retail_price, products.retail_price),
           remark         = COALESCE(updateProducts.remark, products.remark),
           data_status     = COALESCE(updateProducts.data_status, products.data_status)
     WHERE products.id = updateProducts.id;
END;
$$;


ALTER PROCEDURE public.updateproducts(IN id integer, IN sku text, IN name text, IN type integer, IN img_id text, IN img_name text, IN stocks integer, IN weight real, IN retail_price real, IN remark text, IN data_status integer) OWNER TO daymood;

--
-- Name: updateproducts(integer, text, text, integer, text, text, integer, real, real, text, integer, integer); Type: PROCEDURE; Schema: public; Owner: daymood
--

CREATE PROCEDURE public.updateproducts(IN id integer, IN sku text, IN name text, IN type integer, IN img_id text, IN img_name text, IN stocks integer, IN weight real, IN retail_price real, IN remark text, IN data_status integer, IN supplier_id integer)
    LANGUAGE plpgsql
    AS $$
BEGIN
    UPDATE products
       SET sku            = COALESCE(updateProducts.sku, products.sku),
           name           = COALESCE(updateProducts.name, products.name),
           type           = COALESCE(updateProducts.type, products.type),
           img_id         = COALESCE(updateProducts.img_id, products.img_id),
           img_name       = COALESCE(updateProducts.img_name, products.img_name),
           stocks         = COALESCE(updateProducts.stocks, products.stocks),
           weight         = COALESCE(updateProducts.weight, products.weight),
           retail_price   = COALESCE(updateProducts.retail_price, products.retail_price),
           remark         = COALESCE(updateProducts.remark, products.remark),
           data_status     = COALESCE(updateProducts.data_status, products.data_status),
	       supplier_id          = updateProducts.supplier_id
     WHERE products.id = updateProducts.id;
END;
$$;


ALTER PROCEDURE public.updateproducts(IN id integer, IN sku text, IN name text, IN type integer, IN img_id text, IN img_name text, IN stocks integer, IN weight real, IN retail_price real, IN remark text, IN data_status integer, IN supplier_id integer) OWNER TO daymood;

--
-- Name: updatepurchasedetails(integer, text, text, integer, real, integer, real, integer, real, text, integer, integer, integer, integer); Type: PROCEDURE; Schema: public; Owner: daymood
--

CREATE PROCEDURE public.updatepurchasedetails(IN id integer, IN named_id text, IN name text, IN status integer, IN wholesale_price real, IN qty integer, IN cost real, IN currency integer, IN subtotal real, IN remark text, IN data_order integer, IN purchase_id integer, IN supplier_id integer, IN product_id integer)
    LANGUAGE plpgsql
    AS $$
BEGIN
    UPDATE purchaseDetails
       SET named_id             =COALESCE(updatePurchaseDetails.named_id, purchaseDetails.named_id),
	       name                 =COALESCE(updatePurchaseDetails.name, purchaseDetails.name),
	       status               =COALESCE(updatePurchaseDetails.status, purchaseDetails.status),
	       wholesale_price  =COALESCE(updatePurchaseDetails.wholesale_price, purchaseDetails.wholesale_price),
	       qty                  =COALESCE(updatePurchaseDetails.qty, purchaseDetails.qty),
	       cost             =COALESCE(updatePurchaseDetails.cost, purchaseDetails.cost),
	       currency         =COALESCE(updatePurchaseDetails.currency, purchaseDetails.currency),
	       subtotal         =COALESCE(updatePurchaseDetails.subtotal, purchaseDetails.subtotal),
	       remark               =COALESCE(updatePurchaseDetails.remark, purchaseDetails.remark),
	       data_order           =COALESCE(updatePurchaseDetails.data_order, purchaseDetails.data_order),
	       purchase_id          =updatePurchaseDetails.purchase_id,
	       supplier_id          =updatePurchaseDetails.supplier_id,
	       product_id           =updatePurchaseDetails.product_id 
     WHERE purchaseDetails.id = updatePurchaseDetails.id;
END;
$$;


ALTER PROCEDURE public.updatepurchasedetails(IN id integer, IN named_id text, IN name text, IN status integer, IN wholesale_price real, IN qty integer, IN cost real, IN currency integer, IN subtotal real, IN remark text, IN data_order integer, IN purchase_id integer, IN supplier_id integer, IN product_id integer) OWNER TO daymood;

--
-- Name: updatepurchasedetails(integer, text, text, integer, integer, real, real, real, real, real, text, integer, integer, integer, integer); Type: PROCEDURE; Schema: public; Owner: daymood
--

CREATE PROCEDURE public.updatepurchasedetails(IN id integer, IN named_id text, IN name text, IN status integer, IN qty integer, IN wholesale_price_krw real, IN cost_twd real, IN cnf_twd real, IN subtotal_krw real, IN subtotal_twd real, IN remark text, IN data_order integer, IN purchase_id integer, IN supplier_id integer, IN product_id integer)
    LANGUAGE plpgsql
    AS $$
BEGIN
    UPDATE purchaseDetails
       SET named_id             =COALESCE(updatePurchaseDetails.named_id, purchaseDetails.named_id),
	       name                 =COALESCE(updatePurchaseDetails.name, purchaseDetails.name),
	       status               =COALESCE(updatePurchaseDetails.status, purchaseDetails.status),
	       qty                  =COALESCE(updatePurchaseDetails.qty, purchaseDetails.qty),
	       wholesale_price_krw  =COALESCE(updatePurchaseDetails.wholesale_price_krw, purchaseDetails.wholesale_price_krw),
	       cost_twd             =COALESCE(updatePurchaseDetails.cost_twd, purchaseDetails.cost_twd),
	       cnf_twd              =COALESCE(updatePurchaseDetails.cnf_twd, purchaseDetails.cnf_twd),
	       subtotal_krw         =COALESCE(updatePurchaseDetails.subtotal_krw, purchaseDetails.subtotal_krw),
	       subtotal_twd         =COALESCE(updatePurchaseDetails.subtotal_twd, purchaseDetails.subtotal_twd),
	       remark               =COALESCE(updatePurchaseDetails.remark, purchaseDetails.remark),
	       data_order           =COALESCE(updatePurchaseDetails.data_order, purchaseDetails.data_order),
	       purchase_id          =COALESCE(updatePurchaseDetails.purchase_id, purchaseDetails.purchase_id),
	       supplier_id          =COALESCE(updatePurchaseDetails.supplier_id, purchaseDetails.supplier_id),
	       product_id           =COALESCE(updatePurchaseDetails.product_id, purchaseDetails.product_id)
     WHERE purchaseDetails.id = updatePurchaseDetails.id;
END;
$$;


ALTER PROCEDURE public.updatepurchasedetails(IN id integer, IN named_id text, IN name text, IN status integer, IN qty integer, IN wholesale_price_krw real, IN cost_twd real, IN cnf_twd real, IN subtotal_krw real, IN subtotal_twd real, IN remark text, IN data_order integer, IN purchase_id integer, IN supplier_id integer, IN product_id integer) OWNER TO daymood;

--
-- Name: updatepurchases(integer, text, integer, integer, text, text, timestamp without time zone, timestamp without time zone, timestamp without time zone, real, real, real, real, real, real, real, text, integer); Type: PROCEDURE; Schema: public; Owner: daymood
--

CREATE PROCEDURE public.updatepurchases(IN id integer, IN name text, IN status integer, IN type integer, IN shipping_agent text, IN shipping_initiator text, IN shipping_create_at timestamp without time zone, IN shipping_end_at timestamp without time zone, IN shipping_arrive_at timestamp without time zone, IN weight real, IN shipping_fee_kr real, IN shipping_fee_tw real, IN shipping_fee_kokusai real, IN exchange_rate_krw real, IN total_krw real, IN total_twd real, IN remark text, IN data_order integer)
    LANGUAGE plpgsql
    AS $$
BEGIN
    UPDATE purchases
       SET name               = COALESCE(updatePurchases.name, purchases.name),
           status             = COALESCE(updatePurchases.status, purchases.status),
           type               = COALESCE(updatePurchases.type, purchases.type),
           shipping_agent      = COALESCE(updatePurchases.shipping_agent, purchases.shipping_agent),
           shipping_initiator  = COALESCE(updatePurchases.shipping_initiator, purchases.shipping_initiator),
           shipping_create_at   = COALESCE(updatePurchases.shipping_create_at, purchases.shipping_create_at),
           shipping_end_at      = COALESCE(updatePurchases.shipping_end_at, purchases.shipping_end_at),
           shipping_arrive_at   = COALESCE(updatePurchases.shipping_arrive_at, purchases.shipping_arrive_at),
           weight             = COALESCE(updatePurchases.weight, purchases.weight),
           shipping_fee_kr      = COALESCE(updatePurchases.shipping_fee_kr, purchases.shipping_fee_kr),
           shipping_fee_tw      = COALESCE(updatePurchases.shipping_fee_tw, purchases.shipping_fee_tw),
           shipping_fee_kokusai = COALESCE(updatePurchases.shipping_fee_kokusai, purchases.shipping_fee_kokusai),
           exchange_rate_krw    = COALESCE(updatePurchases.exchange_rate_krw, purchases.exchange_rate_krw),
           total_krw           = COALESCE(updatePurchases.total_krw, purchases.total_krw),
           total_twd           = COALESCE(updatePurchases.total_twd, purchases.total_twd),
           remark             = COALESCE(updatePurchases.remark, purchases.remark),
           data_order          = COALESCE(updatePurchases.data_order, purchases.data_order)
     WHERE purchases.id = updatePurchases.id;
END;
$$;


ALTER PROCEDURE public.updatepurchases(IN id integer, IN name text, IN status integer, IN type integer, IN shipping_agent text, IN shipping_initiator text, IN shipping_create_at timestamp without time zone, IN shipping_end_at timestamp without time zone, IN shipping_arrive_at timestamp without time zone, IN weight real, IN shipping_fee_kr real, IN shipping_fee_tw real, IN shipping_fee_kokusai real, IN exchange_rate_krw real, IN total_krw real, IN total_twd real, IN remark text, IN data_order integer) OWNER TO daymood;

--
-- Name: updatepurchases(integer, text, integer, integer, text, real, real, text, timestamp without time zone, timestamp without time zone, timestamp without time zone, real, real, real, real, real, real, real, real, real, real, real, text, integer); Type: PROCEDURE; Schema: public; Owner: daymood
--

CREATE PROCEDURE public.updatepurchases(IN id integer, IN name text, IN status integer, IN purchase_type integer, IN shipping_agent text, IN shipping_agent_cut_krw real, IN shipping_agent_cut_percent real, IN shipping_initiator text, IN shipping_create_at timestamp without time zone, IN shipping_end_at timestamp without time zone, IN shipping_arrive_at timestamp without time zone, IN weight real, IN shipping_fee_kr real, IN shipping_fee_tw real, IN shipping_fee_kokusai_krw real, IN shipping_fee_kokusai_per_kilo real, IN exchange_rate_krw real, IN tariff_twd real, IN tariff_per_kilo real, IN total_krw real, IN total_twd real, IN total real, IN remark text, IN data_order integer)
    LANGUAGE plpgsql
    AS $$
BEGIN
    UPDATE purchases
       SET name                         = COALESCE(updatePurchases.name, purchases.name),
           status                       = COALESCE(updatePurchases.status, purchases.status),
           purchase_type                = COALESCE(updatePurchases.purchase_type, purchases.purchase_type),
           shipping_agent               = COALESCE(updatePurchases.shipping_agent, purchases.shipping_agent),
           shipping_agent_cut_krw       = COALESCE(updatePurchases.shipping_agent_cut_krw, purchases.shipping_agent_cut_krw),
           shipping_agent_cut_percent   = COALESCE(updatePurchases.shipping_agent_cut_percent, purchases.shipping_agent_cut_percent),
           shipping_initiator           = COALESCE(updatePurchases.shipping_initiator, purchases.shipping_initiator),
           shipping_create_at           = COALESCE(updatePurchases.shipping_create_at, purchases.shipping_create_at),
           shipping_end_at              = COALESCE(updatePurchases.shipping_end_at, purchases.shipping_end_at),
           shipping_arrive_at           = COALESCE(updatePurchases.shipping_arrive_at, purchases.shipping_arrive_at),
           weight                       = COALESCE(updatePurchases.weight, purchases.weight),
           shipping_fee_kr              = COALESCE(updatePurchases.shipping_fee_kr, purchases.shipping_fee_kr),
           shipping_fee_tw              = COALESCE(updatePurchases.shipping_fee_tw, purchases.shipping_fee_tw),
           shipping_fee_kokusai_krw         = COALESCE(updatePurchases.shipping_fee_kokusai_krw, purchases.shipping_fee_kokusai_krw),
           shipping_fee_kokusai_per_Kilo         = COALESCE(updatePurchases.shipping_fee_kokusai_per_Kilo, purchases.shipping_fee_kokusai_per_Kilo),
           exchange_rate_krw            = COALESCE(updatePurchases.exchange_rate_krw, purchases.exchange_rate_krw),
           tariff_twd            = COALESCE(updatePurchases.tariff_twd, purchases.tariff_twd),
           tariff_per_Kilo            = COALESCE(updatePurchases.tariff_per_Kilo, purchases.tariff_per_Kilo),
           total_krw                    = COALESCE(updatePurchases.total_krw, purchases.total_krw),
           total_twd                    = COALESCE(updatePurchases.total_twd, purchases.total_twd),
           total                    = COALESCE(updatePurchases.total, purchases.total),
           remark                       = COALESCE(updatePurchases.remark, purchases.remark),
           data_order                   = COALESCE(updatePurchases.data_order, purchases.data_order)
     WHERE purchases.id = updatePurchases.id;
END;
$$;


ALTER PROCEDURE public.updatepurchases(IN id integer, IN name text, IN status integer, IN purchase_type integer, IN shipping_agent text, IN shipping_agent_cut_krw real, IN shipping_agent_cut_percent real, IN shipping_initiator text, IN shipping_create_at timestamp without time zone, IN shipping_end_at timestamp without time zone, IN shipping_arrive_at timestamp without time zone, IN weight real, IN shipping_fee_kr real, IN shipping_fee_tw real, IN shipping_fee_kokusai_krw real, IN shipping_fee_kokusai_per_kilo real, IN exchange_rate_krw real, IN tariff_twd real, IN tariff_per_kilo real, IN total_krw real, IN total_twd real, IN total real, IN remark text, IN data_order integer) OWNER TO daymood;

--
-- Name: updatepurchases(integer, text, integer, integer, integer, text, real, real, text, timestamp without time zone, timestamp without time zone, timestamp without time zone, real, real, real, real, real, real, real, real, real, real, real, text, integer); Type: PROCEDURE; Schema: public; Owner: daymood
--

CREATE PROCEDURE public.updatepurchases(IN id integer, IN name text, IN status integer, IN purchase_type integer, IN qty integer, IN shipping_agent text, IN shipping_agent_cut_krw real, IN shipping_agent_cut_percent real, IN shipping_initiator text, IN shipping_create_at timestamp without time zone, IN shipping_end_at timestamp without time zone, IN shipping_arrive_at timestamp without time zone, IN weight real, IN shipping_fee_kr real, IN shipping_fee_tw real, IN shipping_fee_kokusai_krw real, IN shipping_fee_kokusai_per_kilo real, IN exchange_rate_krw real, IN tariff_twd real, IN tariff_per_kilo real, IN total_krw real, IN total_twd real, IN total real, IN remark text, IN data_order integer)
    LANGUAGE plpgsql
    AS $$
BEGIN
    UPDATE purchases
       SET name                         = COALESCE(updatePurchases.name, purchases.name),
           status                       = COALESCE(updatePurchases.status, purchases.status),
           purchase_type                = COALESCE(updatePurchases.purchase_type, purchases.purchase_type),
           qty                = COALESCE(updatePurchases.qty, purchases.qty),
           shipping_agent               = COALESCE(updatePurchases.shipping_agent, purchases.shipping_agent),
           shipping_agent_cut_krw       = COALESCE(updatePurchases.shipping_agent_cut_krw, purchases.shipping_agent_cut_krw),
           shipping_agent_cut_percent   = COALESCE(updatePurchases.shipping_agent_cut_percent, purchases.shipping_agent_cut_percent),
           shipping_initiator           = COALESCE(updatePurchases.shipping_initiator, purchases.shipping_initiator),
           shipping_create_at           = COALESCE(updatePurchases.shipping_create_at, purchases.shipping_create_at),
           shipping_end_at              = COALESCE(updatePurchases.shipping_end_at, purchases.shipping_end_at),
           shipping_arrive_at           = COALESCE(updatePurchases.shipping_arrive_at, purchases.shipping_arrive_at),
           weight                       = COALESCE(updatePurchases.weight, purchases.weight),
           shipping_fee_kr              = COALESCE(updatePurchases.shipping_fee_kr, purchases.shipping_fee_kr),
           shipping_fee_tw              = COALESCE(updatePurchases.shipping_fee_tw, purchases.shipping_fee_tw),
           shipping_fee_kokusai_krw         = COALESCE(updatePurchases.shipping_fee_kokusai_krw, purchases.shipping_fee_kokusai_krw),
           shipping_fee_kokusai_per_kilo         = COALESCE(updatePurchases.shipping_fee_kokusai_per_kilo, purchases.shipping_fee_kokusai_per_kilo),
           exchange_rate_krw            = COALESCE(updatePurchases.exchange_rate_krw, purchases.exchange_rate_krw),
           tariff_twd            = COALESCE(updatePurchases.tariff_twd, purchases.tariff_twd),
           tariff_per_kilo            = COALESCE(updatePurchases.tariff_per_kilo, purchases.tariff_per_kilo),
           total_krw                    = COALESCE(updatePurchases.total_krw, purchases.total_krw),
           total_twd                    = COALESCE(updatePurchases.total_twd, purchases.total_twd),
           total                    = COALESCE(updatePurchases.total, purchases.total),
           remark                       = COALESCE(updatePurchases.remark, purchases.remark),
           data_order                   = COALESCE(updatePurchases.data_order, purchases.data_order)
     WHERE purchases.id = updatePurchases.id;
END;
$$;


ALTER PROCEDURE public.updatepurchases(IN id integer, IN name text, IN status integer, IN purchase_type integer, IN qty integer, IN shipping_agent text, IN shipping_agent_cut_krw real, IN shipping_agent_cut_percent real, IN shipping_initiator text, IN shipping_create_at timestamp without time zone, IN shipping_end_at timestamp without time zone, IN shipping_arrive_at timestamp without time zone, IN weight real, IN shipping_fee_kr real, IN shipping_fee_tw real, IN shipping_fee_kokusai_krw real, IN shipping_fee_kokusai_per_kilo real, IN exchange_rate_krw real, IN tariff_twd real, IN tariff_per_kilo real, IN total_krw real, IN total_twd real, IN total real, IN remark text, IN data_order integer) OWNER TO daymood;

--
-- Name: updatesuppliers(integer, text, text, text); Type: PROCEDURE; Schema: public; Owner: daymood
--

CREATE PROCEDURE public.updatesuppliers(IN id integer, IN name text, IN address text, IN remark text)
    LANGUAGE plpgsql
    AS $$
BEGIN
    UPDATE suppliers
       SET name       = COALESCE(updateSuppliers.name, suppliers.name),
		   address    = COALESCE(updateSuppliers.address, suppliers.address),
		   remark     = COALESCE(updateSuppliers.remark, suppliers.remark)
     WHERE suppliers.id = updateSuppliers.id;
END;
$$;


ALTER PROCEDURE public.updatesuppliers(IN id integer, IN name text, IN address text, IN remark text) OWNER TO daymood;

--
-- Name: updatesuppliers(integer, text, text, text, integer); Type: PROCEDURE; Schema: public; Owner: daymood
--

CREATE PROCEDURE public.updatesuppliers(IN id integer, IN name text, IN address text, IN remark text, IN data_status integer)
    LANGUAGE plpgsql
    AS $$
BEGIN
    UPDATE suppliers
       SET name       = COALESCE(updateSuppliers.name, suppliers.name),
		   address    = COALESCE(updateSuppliers.address, suppliers.address),
		   remark     = COALESCE(updateSuppliers.remark, suppliers.remark),
		   data_status = COALESCE(updateSuppliers.data_status, suppliers.data_status)
     WHERE suppliers.id = updateSuppliers.id;
END;
$$;


ALTER PROCEDURE public.updatesuppliers(IN id integer, IN name text, IN address text, IN remark text, IN data_status integer) OWNER TO daymood;

--
-- Name: updateuserpasswords(integer, text); Type: PROCEDURE; Schema: public; Owner: daymood
--

CREATE PROCEDURE public.updateuserpasswords(IN id integer, IN password text)
    LANGUAGE plpgsql
    AS $$
BEGIN
    UPDATE users
       SET password      = COALESCE(updateUserPasswords.password, users.password)
     WHERE users.id = updateUserPasswords.id;
END;
$$;


ALTER PROCEDURE public.updateuserpasswords(IN id integer, IN password text) OWNER TO daymood;

--
-- Name: updateusers(integer, text, text, text); Type: PROCEDURE; Schema: public; Owner: daymood
--

CREATE PROCEDURE public.updateusers(IN id integer, IN username text, IN name text, IN email text)
    LANGUAGE plpgsql
    AS $$
BEGIN
    UPDATE users
       SET username      = COALESCE(updateUsers.username, users.username),
		   name          = COALESCE(updateUsers.name, users.name),
		   email         = COALESCE(updateUsers.email, users.email)
     WHERE users.id = updateUsers.id;
END;
$$;


ALTER PROCEDURE public.updateusers(IN id integer, IN username text, IN name text, IN email text) OWNER TO daymood;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: deliveryorderdetails; Type: TABLE; Schema: public; Owner: daymood
--

CREATE TABLE public.deliveryorderdetails (
    id integer NOT NULL,
    retail_price real NOT NULL,
    cost real NOT NULL,
    qty integer NOT NULL,
    subtotal real NOT NULL,
    remark character varying(256),
    data_order integer,
    create_at timestamp(6) without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    update_at timestamp(6) without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    delivery_order_id integer NOT NULL,
    product_id integer
);


ALTER TABLE public.deliveryorderdetails OWNER TO daymood;

--
-- Name: deliveryorderdetails_id_seq; Type: SEQUENCE; Schema: public; Owner: daymood
--

CREATE SEQUENCE public.deliveryorderdetails_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.deliveryorderdetails_id_seq OWNER TO daymood;

--
-- Name: deliveryorderdetails_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: daymood
--

ALTER SEQUENCE public.deliveryorderdetails_id_seq OWNED BY public.deliveryorderdetails.id;


--
-- Name: deliveryorders; Type: TABLE; Schema: public; Owner: daymood
--

CREATE TABLE public.deliveryorders (
    id integer NOT NULL,
    status integer NOT NULL,
    delivery_type integer,
    delivery_status integer,
    delivery_fee_status integer,
    payment_type integer,
    payment_status integer NOT NULL,
    total_original real NOT NULL,
    discount real NOT NULL,
    total_discounted real NOT NULL,
    remark character varying(256),
    data_order integer,
    order_at timestamp(6) without time zone,
    send_at timestamp(6) without time zone,
    arrive_at timestamp(6) without time zone,
    create_at timestamp(6) without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    update_at timestamp(6) without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);


ALTER TABLE public.deliveryorders OWNER TO daymood;

--
-- Name: deliveryorders_id_seq; Type: SEQUENCE; Schema: public; Owner: daymood
--

CREATE SEQUENCE public.deliveryorders_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.deliveryorders_id_seq OWNER TO daymood;

--
-- Name: deliveryorders_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: daymood
--

ALTER SEQUENCE public.deliveryorders_id_seq OWNED BY public.deliveryorders.id;


--
-- Name: discounts; Type: TABLE; Schema: public; Owner: daymood
--

CREATE TABLE public.discounts (
    id integer NOT NULL,
    name character varying(256) NOT NULL,
    price real NOT NULL,
    discount_type integer NOT NULL,
    remark character varying(256),
    data_order integer,
    create_at timestamp(6) without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    update_at timestamp(6) without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    delivery_order_id integer NOT NULL
);


ALTER TABLE public.discounts OWNER TO daymood;

--
-- Name: discounts_id_seq; Type: SEQUENCE; Schema: public; Owner: daymood
--

CREATE SEQUENCE public.discounts_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.discounts_id_seq OWNER TO daymood;

--
-- Name: discounts_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: daymood
--

ALTER SEQUENCE public.discounts_id_seq OWNED BY public.discounts.id;


--
-- Name: products; Type: TABLE; Schema: public; Owner: daymood
--

CREATE TABLE public.products (
    id integer NOT NULL,
    sku character varying(36),
    name character varying(256) NOT NULL,
    type integer NOT NULL,
    img_id character varying(256),
    img_name character varying(256),
    stocks integer NOT NULL,
    weight real,
    retail_price real NOT NULL,
    remark character varying(256),
    data_status integer DEFAULT 1,
    create_at timestamp(6) without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    update_at timestamp(6) without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    supplier_id integer
);


ALTER TABLE public.products OWNER TO daymood;

--
-- Name: products_id_seq; Type: SEQUENCE; Schema: public; Owner: daymood
--

CREATE SEQUENCE public.products_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.products_id_seq OWNER TO daymood;

--
-- Name: products_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: daymood
--

ALTER SEQUENCE public.products_id_seq OWNED BY public.products.id;


--
-- Name: purchasedetails; Type: TABLE; Schema: public; Owner: daymood
--

CREATE TABLE public.purchasedetails (
    id integer NOT NULL,
    named_id character varying(256),
    name character varying(256) NOT NULL,
    status integer NOT NULL,
    wholesale_price real NOT NULL,
    qty integer NOT NULL,
    cost real,
    currency integer,
    subtotal real NOT NULL,
    remark character varying(256),
    data_order integer,
    create_at timestamp(6) without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    update_at timestamp(6) without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    purchase_id integer NOT NULL,
    supplier_id integer,
    product_id integer
);


ALTER TABLE public.purchasedetails OWNER TO daymood;

--
-- Name: purchasedetails_id_seq; Type: SEQUENCE; Schema: public; Owner: daymood
--

CREATE SEQUENCE public.purchasedetails_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.purchasedetails_id_seq OWNER TO daymood;

--
-- Name: purchasedetails_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: daymood
--

ALTER SEQUENCE public.purchasedetails_id_seq OWNED BY public.purchasedetails.id;


--
-- Name: purchases; Type: TABLE; Schema: public; Owner: daymood
--

CREATE TABLE public.purchases (
    id integer NOT NULL,
    name character varying(256) NOT NULL,
    status integer NOT NULL,
    purchase_type integer NOT NULL,
    qty integer,
    shipping_agent character varying(256),
    shipping_agent_cut_krw real,
    shipping_agent_cut_percent real,
    shipping_initiator character varying(256),
    shipping_create_at timestamp(6) without time zone,
    shipping_end_at timestamp(6) without time zone,
    shipping_arrive_at timestamp(6) without time zone,
    weight real,
    shipping_fee_kr real,
    shipping_fee_tw real,
    shipping_fee_kokusai_krw real,
    shipping_fee_kokusai_per_kilo real,
    exchange_rate_krw real,
    tariff_twd real,
    tariff_per_kilo real,
    total_krw real,
    total_twd real,
    total real,
    remark character varying(256),
    data_order integer,
    create_at timestamp(6) without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    update_at timestamp(6) without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);


ALTER TABLE public.purchases OWNER TO daymood;

--
-- Name: purchases_id_seq; Type: SEQUENCE; Schema: public; Owner: daymood
--

CREATE SEQUENCE public.purchases_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.purchases_id_seq OWNER TO daymood;

--
-- Name: purchases_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: daymood
--

ALTER SEQUENCE public.purchases_id_seq OWNED BY public.purchases.id;


--
-- Name: suppliers; Type: TABLE; Schema: public; Owner: daymood
--

CREATE TABLE public.suppliers (
    id integer NOT NULL,
    name character varying(256) NOT NULL,
    address character varying(256),
    remark character varying(256),
    data_status integer DEFAULT 1 NOT NULL,
    create_at timestamp(6) without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    update_at timestamp(6) without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);


ALTER TABLE public.suppliers OWNER TO daymood;

--
-- Name: suppliers_id_seq; Type: SEQUENCE; Schema: public; Owner: daymood
--

CREATE SEQUENCE public.suppliers_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.suppliers_id_seq OWNER TO daymood;

--
-- Name: suppliers_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: daymood
--

ALTER SEQUENCE public.suppliers_id_seq OWNED BY public.suppliers.id;


--
-- Name: users; Type: TABLE; Schema: public; Owner: daymood
--

CREATE TABLE public.users (
    id integer NOT NULL,
    username character varying(256) NOT NULL,
    password character varying(256) NOT NULL,
    name character varying(256) NOT NULL,
    email character varying(256) NOT NULL,
    create_at timestamp(6) without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    update_at timestamp(6) without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);


ALTER TABLE public.users OWNER TO daymood;

--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: daymood
--

CREATE SEQUENCE public.users_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.users_id_seq OWNER TO daymood;

--
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: daymood
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- Name: deliveryorderdetails id; Type: DEFAULT; Schema: public; Owner: daymood
--

ALTER TABLE ONLY public.deliveryorderdetails ALTER COLUMN id SET DEFAULT nextval('public.deliveryorderdetails_id_seq'::regclass);


--
-- Name: deliveryorders id; Type: DEFAULT; Schema: public; Owner: daymood
--

ALTER TABLE ONLY public.deliveryorders ALTER COLUMN id SET DEFAULT nextval('public.deliveryorders_id_seq'::regclass);


--
-- Name: discounts id; Type: DEFAULT; Schema: public; Owner: daymood
--

ALTER TABLE ONLY public.discounts ALTER COLUMN id SET DEFAULT nextval('public.discounts_id_seq'::regclass);


--
-- Name: products id; Type: DEFAULT; Schema: public; Owner: daymood
--

ALTER TABLE ONLY public.products ALTER COLUMN id SET DEFAULT nextval('public.products_id_seq'::regclass);


--
-- Name: purchasedetails id; Type: DEFAULT; Schema: public; Owner: daymood
--

ALTER TABLE ONLY public.purchasedetails ALTER COLUMN id SET DEFAULT nextval('public.purchasedetails_id_seq'::regclass);


--
-- Name: purchases id; Type: DEFAULT; Schema: public; Owner: daymood
--

ALTER TABLE ONLY public.purchases ALTER COLUMN id SET DEFAULT nextval('public.purchases_id_seq'::regclass);


--
-- Name: suppliers id; Type: DEFAULT; Schema: public; Owner: daymood
--

ALTER TABLE ONLY public.suppliers ALTER COLUMN id SET DEFAULT nextval('public.suppliers_id_seq'::regclass);


--
-- Name: users id; Type: DEFAULT; Schema: public; Owner: daymood
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- Name: deliveryorderdetails deliveryorderdetails_pkey; Type: CONSTRAINT; Schema: public; Owner: daymood
--

ALTER TABLE ONLY public.deliveryorderdetails
    ADD CONSTRAINT deliveryorderdetails_pkey PRIMARY KEY (id);


--
-- Name: deliveryorders deliveryorders_pkey; Type: CONSTRAINT; Schema: public; Owner: daymood
--

ALTER TABLE ONLY public.deliveryorders
    ADD CONSTRAINT deliveryorders_pkey PRIMARY KEY (id);


--
-- Name: discounts discounts_pkey; Type: CONSTRAINT; Schema: public; Owner: daymood
--

ALTER TABLE ONLY public.discounts
    ADD CONSTRAINT discounts_pkey PRIMARY KEY (id);


--
-- Name: products products_pkey; Type: CONSTRAINT; Schema: public; Owner: daymood
--

ALTER TABLE ONLY public.products
    ADD CONSTRAINT products_pkey PRIMARY KEY (id);


--
-- Name: products products_sku_key; Type: CONSTRAINT; Schema: public; Owner: daymood
--

ALTER TABLE ONLY public.products
    ADD CONSTRAINT products_sku_key UNIQUE (sku);


--
-- Name: purchasedetails purchasedetails_pkey; Type: CONSTRAINT; Schema: public; Owner: daymood
--

ALTER TABLE ONLY public.purchasedetails
    ADD CONSTRAINT purchasedetails_pkey PRIMARY KEY (id);


--
-- Name: purchases purchases_pkey; Type: CONSTRAINT; Schema: public; Owner: daymood
--

ALTER TABLE ONLY public.purchases
    ADD CONSTRAINT purchases_pkey PRIMARY KEY (id);


--
-- Name: suppliers suppliers_pkey; Type: CONSTRAINT; Schema: public; Owner: daymood
--

ALTER TABLE ONLY public.suppliers
    ADD CONSTRAINT suppliers_pkey PRIMARY KEY (id);


--
-- Name: users users_email_key; Type: CONSTRAINT; Schema: public; Owner: daymood
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_email_key UNIQUE (email);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: daymood
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: users users_username_key; Type: CONSTRAINT; Schema: public; Owner: daymood
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_username_key UNIQUE (username);


--
-- Name: deliveryorderdetails delivery_order_details_update_at; Type: TRIGGER; Schema: public; Owner: daymood
--

CREATE TRIGGER delivery_order_details_update_at BEFORE UPDATE ON public.deliveryorderdetails FOR EACH ROW EXECUTE FUNCTION public.onupdate();


--
-- Name: deliveryorders delivery_orders_update_at; Type: TRIGGER; Schema: public; Owner: daymood
--

CREATE TRIGGER delivery_orders_update_at BEFORE UPDATE ON public.deliveryorders FOR EACH ROW EXECUTE FUNCTION public.onupdate();


--
-- Name: discounts discounts_update_at; Type: TRIGGER; Schema: public; Owner: daymood
--

CREATE TRIGGER discounts_update_at BEFORE UPDATE ON public.discounts FOR EACH ROW EXECUTE FUNCTION public.onupdate();


--
-- Name: products products_update_at; Type: TRIGGER; Schema: public; Owner: daymood
--

CREATE TRIGGER products_update_at BEFORE UPDATE ON public.products FOR EACH ROW EXECUTE FUNCTION public.onupdate();


--
-- Name: purchasedetails purchase_details_update_at; Type: TRIGGER; Schema: public; Owner: daymood
--

CREATE TRIGGER purchase_details_update_at BEFORE UPDATE ON public.purchasedetails FOR EACH ROW EXECUTE FUNCTION public.onupdate();


--
-- Name: purchases purchases_update_at; Type: TRIGGER; Schema: public; Owner: daymood
--

CREATE TRIGGER purchases_update_at BEFORE UPDATE ON public.purchases FOR EACH ROW EXECUTE FUNCTION public.onupdate();


--
-- Name: suppliers suppliers_update_at; Type: TRIGGER; Schema: public; Owner: daymood
--

CREATE TRIGGER suppliers_update_at BEFORE UPDATE ON public.suppliers FOR EACH ROW EXECUTE FUNCTION public.onupdate();


--
-- Name: users users_update_at; Type: TRIGGER; Schema: public; Owner: daymood
--

CREATE TRIGGER users_update_at BEFORE UPDATE ON public.users FOR EACH ROW EXECUTE FUNCTION public.onupdate();


--
-- Name: deliveryorderdetails fk_delivery_order; Type: FK CONSTRAINT; Schema: public; Owner: daymood
--

ALTER TABLE ONLY public.deliveryorderdetails
    ADD CONSTRAINT fk_delivery_order FOREIGN KEY (delivery_order_id) REFERENCES public.deliveryorders(id);


--
-- Name: discounts fk_delivery_order; Type: FK CONSTRAINT; Schema: public; Owner: daymood
--

ALTER TABLE ONLY public.discounts
    ADD CONSTRAINT fk_delivery_order FOREIGN KEY (delivery_order_id) REFERENCES public.deliveryorders(id) ON DELETE CASCADE;


--
-- Name: purchasedetails fk_product; Type: FK CONSTRAINT; Schema: public; Owner: daymood
--

ALTER TABLE ONLY public.purchasedetails
    ADD CONSTRAINT fk_product FOREIGN KEY (product_id) REFERENCES public.products(id) ON DELETE SET NULL;


--
-- Name: deliveryorderdetails fk_product; Type: FK CONSTRAINT; Schema: public; Owner: daymood
--

ALTER TABLE ONLY public.deliveryorderdetails
    ADD CONSTRAINT fk_product FOREIGN KEY (product_id) REFERENCES public.products(id) ON DELETE SET NULL;


--
-- Name: purchasedetails fk_purchase; Type: FK CONSTRAINT; Schema: public; Owner: daymood
--

ALTER TABLE ONLY public.purchasedetails
    ADD CONSTRAINT fk_purchase FOREIGN KEY (purchase_id) REFERENCES public.purchases(id);


--
-- Name: products fk_supplier; Type: FK CONSTRAINT; Schema: public; Owner: daymood
--

ALTER TABLE ONLY public.products
    ADD CONSTRAINT fk_supplier FOREIGN KEY (supplier_id) REFERENCES public.suppliers(id) ON DELETE SET NULL;


--
-- Name: purchasedetails fk_supplier; Type: FK CONSTRAINT; Schema: public; Owner: daymood
--

ALTER TABLE ONLY public.purchasedetails
    ADD CONSTRAINT fk_supplier FOREIGN KEY (supplier_id) REFERENCES public.suppliers(id) ON DELETE SET NULL;


--
-- PostgreSQL database dump complete
--

