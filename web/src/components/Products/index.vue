<script>
import key from 'keymaster'
import PurchaseDialog from '../../components/PurchaseDialog/index.vue'
import ProductDialog from '../../components/ProductDialog/index.vue'
import ProductInfoDialog from '../../components/ProductInfoDialog/index.vue'
import SettingDialog from '../../components/SettingDialog/index.vue'
import Alert from '../../components/Alert/index.vue'
import ConfirmDialog from '../../components/ConfirmDialog/index.vue'
const FileDownload = require('js-file-download');

    export default {
        name: 'Products',
        components:{
            PurchaseDialog,
            ProductDialog,
            ProductInfoDialog,
            SettingDialog,
            Alert,
            ConfirmDialog,
        },
        data(){
            return{
                oldActiveType: null,
                search: '',
                id:'',
                productList: [],
                firmList: [],
                tradingSettings: {},
                deliveryObj: {},
                productHeader:[
                    { text: '商品編號', value: 'ProductSku', width: "12%" },
                    { text: '商品名稱', value: 'ProductName'},
                    // { text: '採購商品編號', value: 'PurchaseProductID' },
                    // { text: '庫存(個)', value: 'Stocks' },
                    // { text: '重量(g)', value: 'Weight' },
                    { text: '批價(KRW)', value: 'KrwWholesalePrice', width: "11%" },
                    { text: '批價(TWD)', value: 'NtdWholesalePrice', width: "11%" },
                    { text: '定價(TWD)', value: 'NtdListPrice', width: "11%" },
                    { text: '售價(TWD)', value: 'NtdSellingPrice', width: "11%" },
                    // { text: '含稅運成本(TWD)', value: 'NtdCnf' },
                    { text: '總成本(TWD)', value: 'NtdCost', width: "11%" },
                    // { text: '建立時間', value: 'CreateAt' },
                    // { text: '更新時間', value: 'UpdateAt' },
                    { text: '', value: 'actions', sortable: false, width: "11%" },
                ],
                productHeaderLess:[
                    { text: '商品編號', value: 'ProductSku'},
                    { text: '商品名稱', value: 'ProductName'},
                    { text: '庫存(個)', value: 'Stocks'},
                ],
                firmHeader:[
                    { text: '廠商名稱', value: 'Name' },
                    { text: '廠商地址', value: 'Address' },
                    { text: '備註', value: 'Remark' },
                    { text: '', value: 'actions', sortable: false, width: "10%" },
                ],
                deliveryHeader:[
                    { text: '商品編號', value: 'ProductSku' },
                    { text: '商品名稱', value: 'DeliveryProductName'},
                    { text: '商品售價', value: 'NtdRetailPrice' },
                    { text: '購買數量', value: 'DeliveryProductQty' },
                    { text: '商品小計', value: 'DeliveryProductSubtotal' },
                ],
                text_cardTitle: "新增",
                text_confirmBtn: "新增",

                purchaseDialog: false,
                productDialog: false,
                productInfoDialog: false,
                actionType: null,
                productID: null,
                productSku: null,
                purchaseProductID: null,
                productName: null,
                productType: null,
                productImgName: null,
                productImgID: null,
                productStocks: null,
                productWeight: null,
                krwWholesalePrice: null,
                ntdWholesalePrice: null,
                ntdListPrice: null,
                ntdSellingPrice: null,
                ntdCnf: null,
                ntdCost: null,
                firmID: null,

                firmDialog: false,
                firmName:null,
                firmAddress: null,
                remark: null,

                settingDialog: false,
                ajeossi: null,
                shippingFee: null,
                exchangeRate: null,
                tariff: null,
                markup: null,

                //Alert
                alert: false,
                alertType: "",
                alertText: "",
                alertTimeoutID: null,

                confirmDialog: false,
                deleteTarget: null,
            };
        },
        mounted(){
            key('command+/', this.onFocus_searchFields);
            key('ctrl+/', this.onFocus_searchFields);
            this.getProductList();
            this.getFirmList();
            this.getTradingSettings();
            // this.dumpProducts();
            // this.dumpFirms();
        },
        props:{
            prop_activeType:{
                type:Number,
                required:true
            },
        },
        computed:{
            activeType:{
                get(){
                    return this.prop_activeType
                },
                set(val){
                    this.$emit('update:prop_activeType',val)
                    
                }
            },
        },
        methods:{
            // filterText (value, search, item) {
            //     return value != null &&
            //     search != null &&
            //     value.toString().indexOf(search) !== -1
            // },
            deleteItem(item){
                if(this.activeType == 1){
                    this.deleteProduct(item);
                }
                else if(this.activeType == 2){
                    this.deleteFirm(item);
                }
            },
            headerSelector(){
                if(this.activeType == 0){
                    return this.productHeaderLess;
                }
                else if(this.activeType == 1){
                    return this.productHeader;
                }
                else if(this.activeType == 2){
                    return this.firmHeader;
                }
                else if(this.activeType == 3){
                    return this.deliveryHeader;
                };
            },
            contentSelector(){
                if(this.activeType == 0){
                    return this.productList;
                }
                else if(this.activeType == 1){
                    return this.productList;
                }
                else if(this.activeType == 2){
                    return this.firmList;
                }
                else if(this.activeType == 3){
                    return this.deliveryObj.Products;
                };
            },
            onFocus_searchFields(){
                this.$refs.searchField.focus();
            },
            onClick_search(){
                this.getDeliveryOrder();
            },
            onClick_row(val){
                if(this.activeType == 0){
                    this.productID = val.ProductID;
                    this.productSku = val.ProductSku;
                    this.productName = val.ProductName;
                    this.productStocks = val.Stocks;
                    this.text_cardTitle = "新增庫存";
                    this.text_confirmBtn = "增加";
                    this.purchaseDialog = true;
                }
                else if(this.activeType == 1){
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
            onClick_newButton(){
                if(this.activeType == 1){
                    this.text_cardTitle = "新增商品";
                    this.text_confirmBtn = "新增";
                    this.actionType = "post";
                    this.productDialog = true;
                }
                else if(this.activeType == 2){
                    this.text_cardTitle = "新增廠商";
                    this.text_confirmBtn = "新增";
                    this.actionType = "post";
                    this.firmDialog = true;
                }
            },
            onClick_editButton(item){
                if(this.activeType == 1){
                    this.text_cardTitle = "編輯商品";
                    this.text_confirmBtn = "修改";
                    this.actionType = "put";
                    this.productID = item.ProductID;
                    this.productSku = item.ProductSku;
                    this.productName = item.ProductName;
                    this.productType = item.ProductType;
                    this.productImgName = item.ProductImgName;
                    this.productImgID = item.ProductImgID;
                    this.productStocks = item.Stocks;
                    this.productWeight = item.Weight.toString();
                    this.krwWholesalePrice = item.KrwWholesalePrice.toString();
                    this.ntdWholesalePrice = item.NtdWholesalePrice.toString();
                    this.ntdListPrice = item.NtdListPrice.toString();
                    this.ntdSellingPrice = item.NtdSellingPrice.toString();
                    this.ntdCnf = item.NtdCnf.toString();
                    this.ntdCost = item.NtdCost.toString();
                    this.purchaseProductID = item.PurchaseProductID;
                    this.firmID = item.FirmInfo.ID;
                    this.productDialog = true;
                }
                else if(this.activeType == 2){
                    this.text_cardTitle = "編輯廠商";
                    this.text_confirmBtn = "修改";
                    this.actionType = "put";
                    this.firmID = item.ID;
                    this.firmName = item.Name;
                    this.firmAddress = item.Address;
                    this.remark = item.Remark;
                    this.firmDialog = true;
                }
            },
            onClick_deleteButton(item){
                this.text_cardTitle = "確認刪除";
                this.text_confirmBtn = "刪除";
                this.confirmDialog = true;
                this.deleteTarget = item;
            },
            onClick_download(){
              if(this.activeType == 1){
                  this.dumpProducts();
              }
              else if(this.activeType == 2){
                  this.dumpFirms();
              }
            },
            onClick_removeDeliveryInfo(){
                this.deliveryObj = {};
            },
            onClick_delivery(){
                this.postDeliveryOrder();
            },
            onClick_setting(){
                this.ajeossi = (this.tradingSettings.Ajeossi).toString();
                this.shippingFee = (this.tradingSettings.ShippingFee).toString();
                this.exchangeRate = (this.tradingSettings.ExchangeRate).toString();
                this.tariff = (this.tradingSettings.Tariff).toString();
                this.markup = (this.tradingSettings.Markup).toString();
                this.text_cardTitle = "設定";
                this.text_confirmBtn = "確定";
                this.settingDialog = true;
            },
            onPurchaseDialogConfirmClick(val){
                this.purchaseDialog = false;
                this.postStocks(val.productID,val.quantity);
            },
            onProductDialogConfirmClick(val){
                this.productDialog = false;
                if(this.actionType == "post"){
                    this.postProduct(val);
                }
                else if(this.actionType == "put"){
                    this.putProduct(val);
                }
            },
            onFirmDialogConfirmClick(val){
                this.firmDialog = false;
                if(this.actionType == "post"){
                    this.postFirm(val.firmName,val.firmAddress,val.remark);
                }
                else if(this.actionType == "put"){
                    this.putFirm(val);
                }
            },
            onConfirmDialogConfirmClick(item){
                this.confirmDialog = false;
                this.deleteItem(item);
            },
            onSettingDialogConfirmClick(item){
                this.settingDialog = false;
                this.putTradingSettings(item);
                
            },
            getDeliveryOrder(){
                let vthis = this;

                this.axios.get(process.env.VUE_APP_BASE_URL+'deliveryorders/'+vthis.search)
                .then(function (response) {

                    if(response.data != null){
                    vthis.deliveryObj = response.data.order;
                    // vthis.toolbarTitle = vthis.guideInfoObj.name+"步驟列表";
                    // console.log(response.data.order);
                    vthis.search = "";
                    }
                    else{
                        vthis.deliveryObj = {};
                    }
                    // console.log(response.data);
                })
                .catch((error) => {
                    // vthis.productList = [];
                    if(error.response.data.code == "DERR0"){
                        vthis.alert = true;
                        vthis.alertType = "Fail";
                        vthis.alertText = "無法連線到EasyStore，請開啟EasyPick";
                    }
                    else if(error.response.data.code == "DERR1"){
                        vthis.alert = true;
                        vthis.alertType = "Fail";
                        vthis.alertText = "商品資料缺少，請進入商品管理新增商品資訊";
                    }
                    // console.log(error.response);
                });
            },
            postDeliveryOrder(){
                let vthis = this;

                this.axios.post(process.env.VUE_APP_BASE_URL+"deliveryorders",JSON.stringify({order:vthis.deliveryObj}))
                .then(function (response) {
                    // console.log(response.data);
                    vthis.getProductList();
                    vthis.alert = true;
                    vthis.alertType = "Success";
                    vthis.alertText = "出貨成功";
                })
                .catch((error) => {
                    if(error.response.data.code == "DERR2"){
                        vthis.alert = true;
                        vthis.alertType = "Fail";
                        vthis.alertText = "扣除商品失敗";
                    }
                    else{
                        vthis.alert = true;
                        vthis.alertType = "Fail";
                        vthis.alertText = "出貨失敗";
                    }
                    // console.log(error.response);
                });
            },
            getProductList(){
                let vthis = this;

                this.axios.get(process.env.VUE_APP_BASE_URL+'products')
                .then(function (response) {

                    if(response.data.products != null){
                    vthis.productList = response.data.products;
                    // vthis.toolbarTitle = vthis.guideInfoObj.name+"步驟列表";
                    // console.log(vthis.productList);
                    }
                    else{
                        vthis.productList = [];
                    }
                    // console.log(response.data);
                })
                .catch((error) => {
                    // vthis.productList = [];
                    // console.log(error.response);
                });
            },
            getFirmList(){
                let vthis = this;

                this.axios.get(process.env.VUE_APP_BASE_URL+'firm')
                .then(function (response) {

                    if(response.data.firms != null){
                    vthis.firmList = response.data.firms;
                    // vthis.toolbarTitle = vthis.guideInfoObj.name+"步驟列表";
                    // console.log(vthis.firmList);
                    }
                    else{
                        vthis.firmList = [];
                    }
                    // console.log(response.data);
                })
                .catch((error) => {
                    // console.log(error.response);
                });
            },
            getTradingSettings(){
                let vthis = this;

                this.axios.get(process.env.VUE_APP_BASE_URL+'tradings')
                .then(function (response) {
                    // console.log(response.data);

                    if(response.data.trading != null){
                    vthis.tradingSettings = response.data.trading;
                    // vthis.ajeossi = response.data.trading.Ajeossi;
                    // vthis.shippingFee = response.data.trading.ShippingFee;
                    // vthis.exchangeRate = response.data.trading.ExchangeRate;
                    // vthis.tariff = response.data.trading.Tariff;
                    // vthis.markup = response.data.trading.Markup;
                    // vthis.toolbarTitle = vthis.guideInfoObj.name+"步驟列表";
                    // console.log(vthis.tradingSettings);
                    }
                    else{
                        vthis.tradingSettings = {};
                    }
                    // console.log(response.data);
                })
                .catch((error) => {
                    // console.log(error.response);
                });
            },
            putTradingSettings(item){
                let vthis = this;

                let formData = new FormData();
                formData.append('ajeossi',item.ajeossi);
                formData.append('shippingFee',item.shippingFee);
                formData.append('exchangeRate',item.exchangeRate);
                formData.append('tariff',item.tariff);
                formData.append('markup',item.markup);

                this.axios.put(process.env.VUE_APP_BASE_URL+"tradings",formData)
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
            postFirm(name,address,remark){
                let vthis = this;

                let formData = new FormData();
                formData.append('name', name);
                formData.append('address', address);
                formData.append('remark', remark);

                this.axios.post(process.env.VUE_APP_BASE_URL+"firm",formData)
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
            putFirm(item){
                let vthis = this;

                let formData = new FormData();
                formData.append('name', item.firmName);
                formData.append('address', item.firmAddress);
                formData.append('remark', item.remark);

                this.axios.put(process.env.VUE_APP_BASE_URL+"firm/"+item.firmID,formData)
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
            deleteFirm(item){
                let vthis = this;

                this.axios.delete(process.env.VUE_APP_BASE_URL+"firm/"+item.ID)
                .then(function (response) {
                    // console.log(response.data);
                    vthis.getFirmList();
                    vthis.alert = true;
                    vthis.alertType = "Success";
                    vthis.alertText = "刪除廠商成功";
                })
                .catch((error) => {
                    if(error.response.data.code == "FERR1"){
                        vthis.alert = true;
                        vthis.alertType = "Fail";
                        vthis.alertText = "尚有商品屬於此廠商，請先移除相關商品後再移除此廠商";
                    }
                });
            },
            postProduct(val){

                let vthis = this;

                let formData = new FormData();
                formData.append('productSku', val.productSku);
                formData.append('purchaseProductID', val.purchaseProductID);
                formData.append('productName', val.productName);
                formData.append('productType', val.productType);
                formData.append('weight', val.productWeight);
                formData.append('krwWholesalePrice', val.krwWholesalePrice);
                formData.append('ntdWholesalePrice', val.ntdWholesalePrice);
                formData.append('ntdListPrice', val.ntdListPrice);
                formData.append('ntdSellingPrice', val.ntdSellingPrice);
                formData.append('ntdCnf', val.ntdCnf);
                formData.append('ntdCost', val.ntdCost);
                formData.append('firmID', val.firmID);
                if(val.imgFile != val.fooFile){
                    formData.append('imgFile', val.imgFile);
                }

                this.axios.post(process.env.VUE_APP_BASE_URL+"products",formData)
                .then(function (response) {
                    // console.log(response);
                    vthis.getProductList();
                    vthis.alert = true;
                    vthis.alertType = "Success";
                    vthis.alertText = "新增商品成功";
                })
                .catch((error) => {
                    // console.log(error.response);

                    if(error.response.data.code == "PERR1"){
                        vthis.alert = true;
                        vthis.alertType = "Fail";
                        vthis.alertText = "商品編號已存在";
                    }
                    else{
                        vthis.alert = true;
                        vthis.alertType = "Fail";
                        vthis.alertText = "新增商品失敗";
                    }
                });
            },
            putProduct(val){

                let vthis = this;

                let formData = new FormData();
                formData.append('purchaseProductID', val.purchaseProductID);
                formData.append('productSku', val.productSku);
                formData.append('productName', val.productName);
                formData.append('productType', val.productType);
                formData.append('weight', val.productWeight);
                formData.append('krwWholesalePrice', val.krwWholesalePrice);
                formData.append('ntdWholesalePrice', val.ntdWholesalePrice);
                formData.append('ntdListPrice', val.ntdListPrice);
                formData.append('ntdSellingPrice', val.ntdSellingPrice);
                formData.append('ntdCnf', val.ntdCnf);
                formData.append('ntdCost', val.ntdCost);
                formData.append('firmID', val.firmID);

                if(val.imgFile != val.fooFile){
                    if(val.fooFile != null && val.imgFile == null){
                        //刪除舊檔且不新增新檔
                        // console.log("刪除舊檔");
                        formData.append('delImgID', val.productImgID);
    
                        formData.append('productImgName', "");
                        formData.append('productImgID', "");
                    }
                    else if(val.fooFile == null && val.imgFile != null){
                        //僅新增新檔
                        // console.log("僅新增新檔");
                        formData.append('imgFile', val.imgFile);
                    }
                    else if(val.fooFile != null && val.imgFile != null){
                        //刪除舊檔且新增新檔
                        // console.log("刪除舊檔且新增新檔");
                        formData.append('delImgID', val.productImgID);
                        formData.append('imgFile', val.imgFile);
                    }
                    // formData.append('delImgID', val.delImgID);
                    // formData.append('imgFile', val.imgFile);
                }
                else{
                    //不更動檔案
                    // console.log("不更動檔案");
                    formData.append('productImgName', val.productImgName);
                    formData.append('productImgID', val.productImgID);
                }

                this.axios.put(process.env.VUE_APP_BASE_URL+"products/"+val.productID,formData)
                .then(function (response) {
                    // console.log(response);
                    vthis.getProductList();
                    vthis.alert = true;
                    vthis.alertType = "Success";
                    vthis.alertText = "編輯商品成功";
                })
                .catch((error) => {
                    // console.log(error.response);
                    vthis.alert = true;
                    vthis.alertType = "Fail";
                    vthis.alertText = "編輯商品失敗";
                });
            },
            deleteProduct(item){
                let vthis = this;

                this.axios.delete(process.env.VUE_APP_BASE_URL+"products/"+item.ProductID)
                .then(function (response) {
                    vthis.getProductList();
                    vthis.alert = true;
                    vthis.alertType = "Success";
                    vthis.alertText = "刪除商品成功";
                })
                .catch((error) => {
                    // console.log(error.response);
                    vthis.alert = true;
                    vthis.alertType = "Fail";
                    vthis.alertText = "刪除商品失敗";
                });
            },
            dumpProducts(){
                let vthis = this;

                this.axios({
                    url: process.env.VUE_APP_BASE_URL+'products/dumping',
                    method: 'GET',
                    responseType: 'blob', // Important
                })
                .then(function (response) {
                    if(response.data){
                        FileDownload(response.data, 'products.csv');
                    }
                    // console.log(response.data);
                })
                .catch((error) => {
                    // console.log(error.response);
                });
            },
            dumpFirms(){
                let vthis = this;

                this.axios({
                    url: process.env.VUE_APP_BASE_URL+'firm/dumping',
                    method: 'GET',
                    responseType: 'blob', // Important
                })
                .then(function (response) {
                    if(response.data != null){
                        FileDownload(response.data, 'firms.csv');
                    }
                    // console.log(response.data);
                })
                .catch((error) => {
                    // console.log(error.response);
                });
            },
            postStocks(productID,purchaseQty){
                let vthis = this;

                let formData = new FormData();
                formData.append('qty', purchaseQty);

                this.axios.post(process.env.VUE_APP_BASE_URL+"stocks/"+productID,formData)
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
        },
        watch:  {
            activeType: function(){
                if(this.activeType != 3 && this.oldActiveType == 3){
                    this.search = "";
                }
                this.oldActiveType = this.activeType;
            },
        }
    }
</script>
<template src="./template.html"></template>