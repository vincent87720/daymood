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
                        <template v-slot:item.actions.plus="{ item }">
                            <v-icon small class="mx-1" @click.stop="onClick_checkoutProductInfo(item)">
                                mdi-arrow-right-circle
                            </v-icon>
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
import BtnUpload from "../../components/Buttons/BtnUpload.vue";
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
                { text: 'SKU', value: 'SKU', width: "12%" },
                { text: '商品名稱', value: 'Name' },
                { text: '商品種類', value: 'ProductType' },
                { text: '庫存(個)', value: 'Stocks' },
                { text: '重量(g)', value: 'Weight' },
                { text: '售價', value: 'RetailPrice' },
                { text: '備註', value: 'Remark', width: "10%" },
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
        },
        onClick_upload() {
            this.text_cardTitle = "匯入";
            this.text_confirmBtn = "確定";
            this.productImportDialog = true;
        },
        onClick_checkoutProductInfo(item) {
            this.text_cardTitle = item.Name;
            this.text_confirmBtn = "";
            this.product = item;
            this.productInfoDialog = true;
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
        async postProducts(item) {
            await postProducts(item)
                .then(async (response) => {
                    await this.getProducts();
                    this.alert = true;
                    this.alertType = "Success";
                    this.alertText = "新增商品成功";
                })
                .catch((error) => {
                    if (error.response.data.model == "products" && error.response.data.code == 3) {
                        this.alert = true;
                        this.alertType = "Fail";
                        this.alertText = "新增項目與資料庫中的SKU重複，請檢查後重新上傳";
                    }
                    else if (error.response.data.role == "router" && error.response.data.code == 1) {
                        this.alert = true;
                        this.alertType = "Fail";
                        this.alertText = "商品名稱缺漏，請檢查商品名稱是否填寫";
                    }
                    else {
                        
                        this.alert = true;
                        this.alertType = "Fail";
                        this.alertText = "新增商品失敗";
                    }
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