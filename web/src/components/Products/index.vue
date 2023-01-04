<template>
    <div>
        <v-container fluid class="d-none d-lg-block">
            <v-row class="ma-0">
                <v-col class="pa-0 d-flex justify-end">
                    <c-btn-add @click="onClick_newButton"></c-btn-add>
                    <c-btn-download @click="onClick_download"></c-btn-download>
                    <c-btn-upload @click="onClick_upload"></c-btn-upload>
                    <c-btn-setting></c-btn-setting>
                </v-col>
            </v-row>
        </v-container>
        <v-container>
            <v-row>
                <v-col xs="12" sm="6" class="mr-auto ml-auto mt-2 d-flex align-center">
                    <v-text-field dark solo ref="searchField" hide-details v-model="search" class="mx-4" label="Search"
                        clearable></v-text-field>
                    <div class="hidden-lg-and-up">
                        <c-btn-add @click="onClick_newButton"></c-btn-add>
                        <c-btn-download @click="onClick_download"></c-btn-download>
                        <c-btn-setting></c-btn-setting>
                    </div>
                </v-col>
            </v-row>
            <v-row>
                <v-col cols="12">
                    <c-data-table :prop_headers="productHeader" :prop_items="products" :prop_search="search"
                        @edit="onClick_editButton" @delete="onClick_deleteButton">
                        <template v-slot:item.ProductType="{ item }">
                            <span>{{ convertDisplayText(systemConfigs.ProductType, item.ProductType) }}</span>
                        </template>
                    </c-data-table>
                </v-col>
            </v-row>
        </v-container>
        <ProductImportDialog :prop_productImportDialog.sync="productImportDialog" :prop_text_cardTitle="text_cardTitle"
            :prop_text_confirmBtn="text_confirmBtn" @confirm='onConfirm_productImportDialog' />
        <ProductInfoDialog :prop_productInfoDialog.sync="productInfoDialog" :prop_text_cardTitle="text_cardTitle"
            :prop_text_confirmBtn="text_confirmBtn" :prop_actionType.sync="actionType" :prop_productItem="product" />
        <ProductDialog :prop_productDialog.sync="productDialog" :prop_text_cardTitle="text_cardTitle"
            :prop_text_confirmBtn="text_confirmBtn" :prop_actionType.sync="actionType" :prop_productItem="product"
            :prop_tradingSettings="tradingSettings" @confirm='onConfirm_productDialog' />
        <ConfirmDialog :prop_confirmDialog.sync="confirmDialog" :prop_text_cardTitle="text_cardTitle"
            :prop_text_cardHint="text_cardHint" :prop_text_confirmBtn="text_confirmBtn"
            :prop_confirmTarget.sync="confirmTarget" v-on:confirmClick='onConfirm_ConfirmDialog'>
            <template v-slot:actions="{ item }">
                <v-btn outlined rounded text @click.stop="onClick_changeDataStatus(item)">
                    刪除並保留歷史紀錄
                </v-btn>
            </template>
        </ConfirmDialog>
        <Alert :prop_alert.sync="alert" :prop_alertType="alertType" :prop_alertText="alertText"></Alert>
    </div>

</template>

<script>
import key from 'keymaster'
import ProductDialog from '../../components/ProductDialog/index.vue'
import ProductInfoDialog from '../../components/ProductInfoDialog/index.vue'
import ProductImportDialog from '../../components/ProductImportDialog/index.vue'
import Alert from '../../components/Alert/index.vue'
import ConfirmDialog from '../../components/ConfirmDialog/index.vue'
import BtnAdd from "../../components/Buttons/BtnAdd.vue";
import BtnDownload from "../../components/Buttons/BtnDownload.vue";
import BtnSetting from "../../components/Buttons/BtnSetting.vue";
import DataTable from "../../components/DataTables/DataTable.vue";
import {
    getProducts,
    postProduct,
    putProduct,
    deleteProduct,
} from "../../apis/ProductsAPI";
const FileDownload = require('js-file-download');

class Product {
    ID = undefined;
    SKU = "";
    Name = "";
    ProductType = undefined;
    ImgName = "";
    ImgID = "";
    Stocks = undefined;
    Weight = undefined;
    RetailPrice = undefined;
    Remark = "";
    DataOrder = "";
    CreateAt = "";
    UpdateAt = "";
}

export default {
    name: 'Products',
    components: {
        ProductDialog,
        Alert,
        ConfirmDialog,
        ProductDialog,
        ProductInfoDialog,
        ProductImportDialog,
        "c-btn-add": BtnAdd,
        "c-btn-download": BtnDownload,
        "c-btn-upload": BtnUpload,
        "c-btn-setting": BtnSetting,
        "c-data-table": DataTable,
    },
    data() {
        return {
            search: '',
            text_cardTitle: "新增",
            text_confirmBtn: "新增",

            //Alert
            alert: false,
            alertType: "",
            alertText: "",

            confirmDialog: false,
            confirmTarget: null,

            product: new Product(),
            productDialog: false,
            products: [],
            productHeader: [
                { text: '商品編號', value: 'SKU', width: "12%" },
                { text: '商品名稱', value: 'Name' },
                { text: '商品種類', value: 'ProductType' },
                // { text: '採購商品編號', value: 'PurchaseProductID' },
                { text: '庫存(個)', value: 'Stocks' },
                { text: '重量(g)', value: 'Weight' },
                { text: '售價', value: 'RetailPrice' },
                { text: '備註', value: 'Remark', width: "10%" },
                // { text: '批價(KRW)', value: 'KrwWholesalePrice', width: "11%" },
                // { text: '批價(TWD)', value: 'NtdWholesalePrice', width: "11%" },
                // { text: '定價(TWD)', value: 'NtdListPrice', width: "11%" },
                // { text: '售價(TWD)', value: 'NtdSellingPrice', width: "11%" },
                // { text: '含稅運成本(TWD)', value: 'NtdCnf' },
                // { text: '總成本(TWD)', value: 'NtdCost', width: "11%" },
                // { text: '建立時間', value: 'CreateAt' },
                // { text: '更新時間', value: 'UpdateAt' },
                { text: '', value: 'actions', sortable: false, width: "10%" },
            ],
            productInfoDialog: false,

            tradingSettings: {},
            actionType: "",

            productImportDialog: false,
        };
    },
    async mounted() {
        key('command+/', this.onFocus_searchFields);
        key('ctrl+/', this.onFocus_searchFields);
        await this.getProducts();
        // this.dumpProducts();
        // this.dumpFirms();
    },
    props: {
    },
    computed: {
        systemConfigs() {
            return this.$store.state.systemConfigs;
        },
    },
    methods: {
        convertDisplayText(systemConfig, key) {
            return systemConfig[key].value
        },
        onFocus_searchFields() {
            this.$refs.searchField.focus();
        },
        onClick_search() {
            this.getDeliveryOrder();
        },
        onClick_row(val) {
            if (this.activeType == 0) {
                this.productID = val.ProductID;
                this.productSku = val.ProductSku;
                this.productName = val.ProductName;
                this.productStocks = val.Stocks;
                this.text_cardTitle = "新增庫存";
                this.text_confirmBtn = "增加";
                this.purchaseDialog = true;
            }
            else if (this.activeType == 1) {
                this.productID = val.ProductID;
                this.productSku = val.ProductSku;
                this.productName = val.ProductName;
                this.productType = val.ProductType;
                this.productImgID = val.ProductImgID;
                this.productStocks = val.Stocks;
                this.productWeight = val.Weight.toString();
                this.krwWholesalePrice = val.KrwWholesalePrice.toString();
                this.ntdWholesalePrice = val.NtdWholesalePrice.toString();
                this.ntdListPrice = val.NtdListPrice.toString();
                this.ntdSellingPrice = val.NtdSellingPrice.toString();
                this.ntdCnf = val.NtdCnf.toString();
                this.ntdCost = val.NtdCost.toString();
                this.purchaseProductID = val.PurchaseProductID;
                this.firmID = val.FirmInfo.ID;
                this.productInfoDialog = true;
            }
        },
        onClick_newButton() {
            this.text_cardTitle = "新增商品";
            this.text_confirmBtn = "新增";
            this.actionType = "post";
            this.product = new Product();
            this.productDialog = true;
        },
        onClick_editButton(item) {
            this.text_cardTitle = "編輯商品";
            this.text_confirmBtn = "修改";
            this.actionType = "put";
            this.product = item;
            this.productDialog = true;
        },
        onClick_deleteButton(item) {
            this.text_cardTitle = "確認刪除";
            this.text_confirmBtn = "刪除";
            this.confirmDialog = true;
            this.confirmTarget = item;
        },
        onClick_download() {
            if (this.activeType == 1) {
                this.dumpProducts();
            }
            else if (this.activeType == 2) {
                this.dumpFirms();
            }
        },
        onClick_removeDeliveryInfo() {
            this.deliveryObj = {};
        },
        onClick_delivery() {
            this.postDeliveryOrder();
        },
        onClick_setting() {
            this.ajeossi = (this.tradingSettings.Ajeossi).toString();
            this.shippingFee = (this.tradingSettings.ShippingFee).toString();
            this.exchangeRate = (this.tradingSettings.ExchangeRate).toString();
            this.tariff = (this.tradingSettings.Tariff).toString();
            this.markup = (this.tradingSettings.Markup).toString();
            this.text_cardTitle = "設定";
            this.text_confirmBtn = "確定";
            this.settingDialog = true;
        },
        async onClick_changeDataStatus(item) {
            this.confirmDialog = false;
            item.DataStatus = 0;
            await this.putProduct(item);
        },
        async onConfirm_productDialog(item) {
            this.productDialog = false;
            if (this.actionType == "post") {
                await this.postProduct(item);
            }
            else if (this.actionType == "put") {
                await this.putProduct(item);
            }
        },
        async onConfirm_productImportDialog(item) {
            this.productImportDialog = false;
            await this.postProducts(item);
        },
        async onConfirm_ConfirmDialog(item) {
            this.confirmDialog = false;
            await this.deleteProduct(item);
        },
        onSettingDialogConfirmClick(item) {
            this.settingDialog = false;
            this.putTradingSettings(item);

        },
        getDeliveryOrder() {
            let vthis = this;

            this.axios.get(process.env.VUE_APP_BASE_URL + 'deliveryorders/' + vthis.search)
                .then(function (response) {

                    if (response.data != null) {
                        vthis.deliveryObj = response.data.order;
                        // vthis.toolbarTitle = vthis.guideInfoObj.name+"步驟列表";
                        // console.log(response.data.order);
                        vthis.search = "";
                    }
                    else {
                        vthis.deliveryObj = {};
                    }
                    // console.log(response.data);
                })
                .catch((error) => {
                    // vthis.productList = [];
                    if (error.response.data.code == "DERR0") {
                        vthis.alert = true;
                        vthis.alertType = "Fail";
                        vthis.alertText = "無法連線到EasyStore，請開啟EasyPick";
                    }
                    else if (error.response.data.code == "DERR1") {
                        vthis.alert = true;
                        vthis.alertType = "Fail";
                        vthis.alertText = "商品資料缺少，請進入商品管理新增商品資訊";
                    }
                    // console.log(error.response);
                });
        },
        postDeliveryOrder() {
            let vthis = this;

            this.axios.post(process.env.VUE_APP_BASE_URL + "deliveryorders", JSON.stringify({ order: vthis.deliveryObj }))
                .then(function (response) {
                    // console.log(response.data);
                    vthis.getProductList();
                    vthis.alert = true;
                    vthis.alertType = "Success";
                    vthis.alertText = "出貨成功";
                })
                .catch((error) => {
                    if (error.response.data.code == "DERR2") {
                        vthis.alert = true;
                        vthis.alertType = "Fail";
                        vthis.alertText = "扣除商品失敗";
                    }
                    else {
                        vthis.alert = true;
                        vthis.alertType = "Fail";
                        vthis.alertText = "出貨失敗";
                    }
                    // console.log(error.response);
                });
        },
        getProductList() {
            let vthis = this;

            this.axios.get(process.env.VUE_APP_BASE_URL + 'products')
                .then(function (response) {

                    if (response.data.products != null) {
                        vthis.productList = response.data.products;
                        // vthis.toolbarTitle = vthis.guideInfoObj.name+"步驟列表";
                        // console.log(vthis.productList);
                    }
                    else {
                        vthis.productList = [];
                    }
                    // console.log(response.data);
                })
                .catch((error) => {
                    // vthis.productList = [];
                    // console.log(error.response);
                });
        },
        getFirmList() {
            let vthis = this;

            this.axios.get(process.env.VUE_APP_BASE_URL + 'firm')
                .then(function (response) {

                    if (response.data.firms != null) {
                        vthis.firmList = response.data.firms;
                        // vthis.toolbarTitle = vthis.guideInfoObj.name+"步驟列表";
                        // console.log(vthis.firmList);
                    }
                    else {
                        vthis.firmList = [];
                    }
                    // console.log(response.data);
                })
                .catch((error) => {
                    // console.log(error.response);
                });
        },
        getTradingSettings() {
            let vthis = this;

            this.axios.get(process.env.VUE_APP_BASE_URL + 'tradings')
                .then(function (response) {
                    // console.log(response.data);

                    if (response.data.trading != null) {
                        vthis.tradingSettings = response.data.trading;
                        // vthis.ajeossi = response.data.trading.Ajeossi;
                        // vthis.shippingFee = response.data.trading.ShippingFee;
                        // vthis.exchangeRate = response.data.trading.ExchangeRate;
                        // vthis.tariff = response.data.trading.Tariff;
                        // vthis.markup = response.data.trading.Markup;
                        // vthis.toolbarTitle = vthis.guideInfoObj.name+"步驟列表";
                        // console.log(vthis.tradingSettings);
                    }
                    else {
                        vthis.tradingSettings = {};
                    }
                    // console.log(response.data);
                })
                .catch((error) => {
                    // console.log(error.response);
                });
        },
        putTradingSettings(item) {
            let vthis = this;

            let formData = new FormData();
            formData.append('ajeossi', item.ajeossi);
            formData.append('shippingFee', item.shippingFee);
            formData.append('exchangeRate', item.exchangeRate);
            formData.append('tariff', item.tariff);
            formData.append('markup', item.markup);

            this.axios.put(process.env.VUE_APP_BASE_URL + "tradings", formData)
                .then(function (response) {
                    // console.log(response.data);
                    vthis.getTradingSettings();
                    vthis.alert = true;
                    vthis.alertType = "Success";
                    vthis.alertText = "修改設定成功";
                })
                .catch((error) => {
                    // console.log(error.response);
                    vthis.alert = true;
                    vthis.alertType = "Fail";
                    vthis.alertText = "修改設定失敗";
                });
        },
        postFirm(name, address, remark) {
            let vthis = this;

            let formData = new FormData();
            formData.append('name', name);
            formData.append('address', address);
            formData.append('remark', remark);

            this.axios.post(process.env.VUE_APP_BASE_URL + "firm", formData)
                .then(function (response) {
                    // console.log(response.data);
                    vthis.getFirmList();
                    vthis.alert = true;
                    vthis.alertType = "Success";
                    vthis.alertText = "新增廠商成功";
                })
                .catch((error) => {
                    // console.log(error.response);
                    vthis.alert = true;
                    vthis.alertType = "Fail";
                    vthis.alertText = "新增廠商失敗";
                });
        },
        putFirm(item) {
            let vthis = this;

            let formData = new FormData();
            formData.append('name', item.firmName);
            formData.append('address', item.firmAddress);
            formData.append('remark', item.remark);

            this.axios.put(process.env.VUE_APP_BASE_URL + "firm/" + item.firmID, formData)
                .then(function (response) {
                    // console.log(response.data);
                    vthis.getFirmList();
                    vthis.alert = true;
                    vthis.alertType = "Success";
                    vthis.alertText = "編輯廠商成功";
                })
                .catch((error) => {
                    // console.log(error.response);
                    vthis.alert = true;
                    vthis.alertType = "Fail";
                    vthis.alertText = "編輯廠商失敗";
                });
        },
        deleteFirm(item) {
            let vthis = this;

            this.axios.delete(process.env.VUE_APP_BASE_URL + "firm/" + item.ID)
                .then(function (response) {
                    // console.log(response.data);
                    vthis.getFirmList();
                    vthis.alert = true;
                    vthis.alertType = "Success";
                    vthis.alertText = "刪除廠商成功";
                })
                .catch((error) => {
                    if (error.response.data.code == "FERR1") {
                        vthis.alert = true;
                        vthis.alertType = "Fail";
                        vthis.alertText = "尚有商品屬於此廠商，請先移除相關商品後再移除此廠商";
                    }
                });
        },
        dumpProducts() {
            let vthis = this;

            this.axios({
                url: process.env.VUE_APP_BASE_URL + 'products/dumping',
                method: 'GET',
                responseType: 'blob', // Important
            })
                .then(function (response) {
                    if (response.data) {
                        FileDownload(response.data, 'products.csv');
                    }
                    // console.log(response.data);
                })
                .catch((error) => {
                    // console.log(error.response);
                });
        },
        dumpFirms() {
            let vthis = this;

            this.axios({
                url: process.env.VUE_APP_BASE_URL + 'firm/dumping',
                method: 'GET',
                responseType: 'blob', // Important
            })
                .then(function (response) {
                    if (response.data != null) {
                        FileDownload(response.data, 'firms.csv');
                    }
                    // console.log(response.data);
                })
                .catch((error) => {
                    // console.log(error.response);
                });
        },
        postStocks(productID, purchaseQty) {
            let vthis = this;

            let formData = new FormData();
            formData.append('qty', purchaseQty);

            this.axios.post(process.env.VUE_APP_BASE_URL + "stocks/" + productID, formData)
                .then(function (response) {
                    // console.log(response.data);
                    vthis.getProductList();
                    vthis.alert = true;
                    vthis.alertType = "Success";
                    vthis.alertText = "新增庫存成功";
                })
                .catch((error) => {
                    // console.log(error.response);
                    vthis.alert = true;
                    vthis.alertType = "Fail";
                    vthis.alertText = "新增庫存失敗";
                });
        },
        preSend(item) {
            item.ID = parseFloat(item.ID);
            item.ProductType = parseFloat(item.ProductType);
            item.Stocks = parseFloat(item.Stocks);
            item.Weight = parseFloat(item.Weight);
            item.RetailPrice = parseFloat(item.RetailPrice);
            item.DataOrder = parseFloat(item.DataOrder);
            return item;
        },
        postGet(item) {
            item.ID = String(item.ID);
            item.Status = String(item.Status);
            item.PurchaseType = String(item.PurchaseType);
            item.Weight = String(item.Weight);
            item.ShippingFeeKr = String(item.ShippingFeeKr);
            item.ShippingFeeTw = String(item.ShippingFeeTw);
            item.ShippingFeeKokusai = String(item.ShippingFeeKokusai);
            item.ExchangeRateKrw = String(item.ExchangeRateKrw);
            item.TotalKrw = String(item.TotalKrw);
            item.TotalTwd = String(item.TotalTwd);
            item.DataOrder = String(item.DataOrder);
            return item;
        },
        filterEnableData(item) {
            return item.filter(x => x.DataStatus == 1);
        },
        async getProducts() {
            await getProducts()
                .then((response) => {
                    if (response.data.records != null) {
                        this.products = this.filterEnableData(response.data.records);
                    }
                    else {
                        this.products = [];
                    }
                })
                .catch((error) => {
                });
        },
        async postProduct(item) {
            item = this.preSend(item);
            await postProduct(item)
                .then(async (response) => {
                    await this.getProducts();
                    this.alert = true;
                    this.alertType = "Success";
                    this.alertText = "新增商品成功";
                })
                .catch((error) => {
                    this.alert = true;
                    this.alertType = "Fail";
                    this.alertText = "新增商品失敗";
                });
        },
        async putProduct(item) {
            item = this.preSend(item);
            await putProduct(item)
                .then(async (response) => {
                    await this.getProducts();
                    this.alert = true;
                    this.alertType = "Success";
                    this.alertText = "編輯商品成功";
                })
                .catch((error) => {
                    this.alert = true;
                    this.alertType = "Fail";
                    this.alertText = "編輯商品失敗";
                });
        },
        async deleteProduct(item) {
            await deleteProduct(item)
                .then(async (response) => {
                    await this.getProducts();
                    this.alert = true;
                    this.alertType = "Success";
                    this.alertText = "刪除商品成功";
                })
                .catch((error) => {
                    this.alert = true;
                    this.alertType = "Fail";
                    this.alertText = "刪除商品失敗";
                });
        },
    },
}
</script>