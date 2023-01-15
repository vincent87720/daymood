<template>
    <v-dialog v-model="productInfoDialog" hide-overlay @click:outside="onClick_cancel" fullscreen
        transition="slide-x-reverse-transition">
        <v-card class="d-flex align-center flex-column justify-center">
            <v-btn icon dark absolute top left @click="productInfoDialog = false">
                <v-icon>mdi-arrow-left</v-icon>
            </v-btn>
            <v-card-title class="text-h3">
                {{ text_cardTitle }}
            </v-card-title>
            <v-container class="pa-6 mb-8">
                <v-row>
                    <v-col cols="12" sm="12" lg="12" class="d-flex justify-center align-center">
                        <v-card class="elevation-0" width="400">
                            <!-- <v-img v-if="imgSrc" class="ma-0" :lazy-src="imgSrc" :src="imgSrc" max-height="150"
                                max-width="344"></v-img> -->
                            <!-- <v-card-title class="d-flex justify-center">
                                <span class="text-h4 mx-1" style="color:gray;">{{ productItem.SKU }}</span>
                                <span class="text-h4 mx-1">{{ productItem.Name }}</span>
                            </v-card-title> -->
                            <div class="d-flex justify-space-between align-center my-2">
                                <span class="text-subtitle-1 font-weight-medium">商品名稱</span>
                                <span class="text-subtitle-1 font-weight-bold"> {{ productItem.Name }}</span>
                            </div>
                            <div class="d-flex justify-space-between align-center my-2">
                                <span class="text-subtitle-1 font-weight-medium">SKU</span>
                                <span class="text-subtitle-1 font-weight-bold"> {{ productItem.SKU }}</span>
                            </div>
                            <div class="d-flex justify-space-between align-center my-2">
                                <span class="text-subtitle-1 font-weight-medium">商品種類</span>
                                <span class="text-subtitle-1 font-weight-bold"> {{
                                    convertDisplayText(systemConfigs.ProductType, productItem.ProductType)
                                }}</span>
                            </div>
                            <div class="d-flex justify-space-between align-center my-2">
                                <span class="text-subtitle-1 font-weight-medium">供應商</span>
                                <span class="text-subtitle-1 font-weight-bold"> {{
                                    convertDisplayText(allSuppliersList, productItem.SupplierID)
                                }}</span>
                            </div>
                            <div class="d-flex justify-space-between align-center my-2">
                                <span class="text-subtitle-1 font-weight-medium">商品重量(g)</span>
                                <span class="text-subtitle-1 font-weight-bold"> {{ productItem.Weight }} g</span>
                            </div>
                            <div class="d-flex justify-space-between align-center my-2">
                                <span class="text-subtitle-1 font-weight-medium">售價</span>
                                <span class="text-subtitle-1 font-weight-bold"> {{ productItem.RetailPrice }}</span>
                            </div>
                            <div class="d-flex justify-space-between align-center my-2">
                                <span class="text-subtitle-1 font-weight-medium">庫存</span>
                                <span class="text-subtitle-1 font-weight-bold"> {{ calcStocks() }}</span>
                            </div>
                            <div class="d-flex justify-space-between my-2">
                                <span class="text-subtitle-1 font-weight-medium">備註</span>
                                <span class="text-subtitle-1 font-weight-bold d-flex justify-end"
                                    style="white-space: pre-line;"> {{ productItem.Remark }}</span>
                            </div>
                        </v-card>
                    </v-col>
                </v-row>
                <v-row>
                    <v-col cols="12" sm="12" lg="12">
                        <c-card-rounded>
                            <c-data-table :prop_headers="purchaseHistoryHeader" :prop_items="purchaseHistories" :prop_search="search">
                                <template v-slot:top="{ item }">
                                    <div class="pa-4 text-h5 d-flex justify-center">採購歷史紀錄</div>
                                    <v-divider></v-divider>
                                </template>
                                <template v-slot:item.PurchaseSupplierID="{ item }">
                                    <span>{{ convertDisplayText(allSuppliersList, item.SupplierID) }}</span>
                                </template>
                                <template v-slot:item.ShippingArriveAt="{ item }">
                                    <span v-if="item.ShippingArriveAt">{{ convertDateFormat(item.ShippingArriveAt) }}</span>
                                </template>
                                <template v-slot:item.actions="{ item }">
                                    <v-icon small class="mx-1" @click.stop="onClick_checkoutPurchaseHistory(item)">
                                        mdi-arrow-right-circle
                                    </v-icon>
                                </template>
                            </c-data-table>
                        </c-card-rounded>
                    </v-col>
                </v-row>
                <v-row>
                    <v-col cols="12" sm="12" lg="12">
                        <c-card-rounded>
                            <c-data-table :prop_headers="historyHeader_delivery" :prop_items="deliveryHistories" :prop_search="search">
                                <template v-slot:top="{ item }">
                                    <div class="pa-4 text-h5 d-flex justify-center">出貨歷史紀錄</div>
                                    <v-divider></v-divider>
                                </template>
                                <template v-slot:item.PurchaseSupplierID="{ item }">
                                    <span>{{ convertDisplayText(allSuppliersList, item.SupplierID) }}</span>
                                </template>
                                <template v-slot:item.OrderAt="{ item }">
                                    <span v-if="item.OrderAt">{{ convertDateFormat(item.OrderAt) }}</span>
                                </template>
                                <template v-slot:item.actions="{ item }">
                                    <v-icon small class="mx-1" @click.stop="onClick_checkoutDeliveryHistory(item)">
                                        mdi-arrow-right-circle
                                    </v-icon>
                                </template>
                            </c-data-table>
                        </c-card-rounded>
                    </v-col>
                </v-row>
            </v-container>
        </v-card>
        <PurchaseInfoDialog :prop_purchaseInfoDialog.sync="purchaseInfoDialog"
            :prop_text_cardTitle="text_cardTitle_inner" :prop_text_confirmBtn="text_confirmBtn_inner"
            :prop_purchaseItem="purchase" @finish='onFinish_purchaseInfoDialog' />
        <DeliveryOrderInfoDialog :prop_deliveryOrderInfoDialog.sync="deliveryOrderInfoDialog"
            :prop_text_cardTitle="text_cardTitle_inner" :prop_text_confirmBtn="text_confirmBtn_inner"
            :prop_deliveryOrderItem="deliveryOrder" @finish='onFinish_deliveryOrderInfoDialog' />
        <Alert :prop_alert.sync="alert" :prop_alertType="alertType" :prop_alertText="alertText"></Alert>
    </v-dialog>
</template>

<script>
import Alert from '../../components/Alert/index.vue'
import CardRounded from "../../components/Cards/CardRounded.vue";
import DatePicker from "../../components/Pickers/DatePicker.vue";
import DataTable from "../../components/DataTables/DataTable.vue";
import PurchaseInfoDialog from '../../components/PurchaseInfoDialog/index.vue'
import DeliveryOrderInfoDialog from '../../components/DeliveryOrderInfoDialog/index.vue'
import { getProductPurchaseHistories, getProductDeliveryHistories } from "../../apis/ProductsAPI";
import { getPurchase, putPurchase } from "../../apis/PurchasesAPI";
import { getDeliveryOrder, putDeliveryOrder } from "../../apis/DeliveryOrdersAPI";

export default {
    name: 'productInfoDialog',
    components: {
        Alert,
        PurchaseInfoDialog,
        DeliveryOrderInfoDialog,
        "c-card-rounded": CardRounded,
        "c-date-picker": DatePicker,
        "c-data-table": DataTable,
    },
    data() {
        return {
            //Alert
            alert: false,
            alertType: "",
            alertText: "",

            //PurchaseInfoDialog
            text_cardTitle_inner: '',
            text_cardHint_inner: '',
            text_confirmBtn_inner: '',
            purchase: {},
            deliveryOrder: {},
            purchaseInfoDialog: false,
            deliveryOrderInfoDialog: false,

            //c-data-table
            search: '',
            purchaseHistories: [],
            deliveryHistories: [],
            historyHeader_product: [
                { text: '採購名稱', value: 'PurchaseName' },
                { text: '進貨廠商', value: 'PurchaseSupplierID' },
                { text: '批價', value: 'WholesalePrice' },
                { text: '批價', value: 'WholesalePriceTwd' },
                { text: '數量', value: 'PurchaseDetailQTY' },
                { text: '進口成本', value: 'ImportCost' },
                { text: '成本', value: 'Costs' },
                { text: '毛利', value: 'GrossProfit' },
                { text: '毛利率(%)', value: 'GrossMargin' },
                { text: '', value: 'actions', sortable: false, width: "1%" },
            ],
            historyHeader_material: [
                { text: '採購名稱', value: 'PurchaseName' },
                { text: '進貨廠商', value: 'PurchaseSupplierID' },
                { text: '批價', value: 'WholesalePrice' },
                { text: '數量', value: 'PurchaseDetailQTY' },
                { text: '', value: 'actions', sortable: false, width: "1%" },
            ],
            historyHeader_delivery: [
                { text: '出貨單編號', value: 'DeliveryOrderID' },
                { text: '下單日期', value: 'OrderAt' },
                { text: '出貨時售價', value: 'RetailPrice' },
                { text: '數量', value: 'QTY' },
                { text: '總計', value: 'Subtotal' },
                { text: '出貨時成本', value: 'Cost' },
                { text: '毛利', value: 'GrossProfit' },
                { text: '毛利率(%)', value: 'GrossMargin' },
                { text: '', value: 'actions', sortable: false, width: "1%" },
            ],
        };
    },
    props: {
        prop_productInfoDialog: {
            type: Boolean,
            required: true
        },
        prop_text_cardTitle: {
            type: String,
            required: true
        },
        prop_productItem: {
            type: Object,
            required: false
        },
    },
    mounted() {
        this.$store.dispatch("GetSuppliers");
    },
    computed: {
        productInfoDialog: {
            get() {
                return this.prop_productInfoDialog
            },
            set(val) {
                this.$emit('update:prop_productInfoDialog', val)

            }
        },
        text_cardTitle: {
            get() {
                return this.prop_text_cardTitle
            },
            set(val) {
                this.$emit('update:prop_text_cardTitle', val)
            }
        },
        productItem: {
            get() {
                return this.prop_productItem
            },
            set(val) {
                this.$emit('update:prop_productItem', val)
            }
        },
        purchaseHistoryHeader() {
            return this.isMaterials ? this.historyHeader_material : this.historyHeader_product;
        },
        isMaterials() {
            let type_products = [1, 2, 3, 4, 5];//商品
            let type_materials = [6, 7, 8, 9, 10, 11, 12, 13];//耗材
            if (type_products.includes(this.productItem.ProductType)) {
                return false;
            }
            else if (type_materials.includes(this.productItem.ProductType)) {
                return true
            }
        },
        systemConfigs() {
            return this.$store.state.systemConfigs;
        },
        tradingSettings() {
            return this.$store.state.tradingSettings;
        },
        allSuppliersList() {
            return this.$store.state.allSuppliers;
        },
    },
    methods: {
        convertDisplayText(list, key) {
            let result = list.find(x => x.key == key);
            if (result) {
                return result.value
            }
            return "";
        },
        convertDataList(data) {
            data.map(x => x = { key: x.ID, value: x.SKU + " " + x.Name });
            return data
        },
        convertDateFormat(datetime) {
            if (datetime) {
                return datetime.substring(0, 10);
            }
            return ""
        },
        onClick_cancel() {
            this.productInfoDialog = false;
        },
        async onClick_checkoutPurchaseHistory(item) {
            this.text_cardTitle_inner = "採購案" + item.PurchaseName;
            this.text_confirmBtn_inner = "";
            await this.getPurchase(item.PurchaseID);
            this.purchaseInfoDialog = true;
        },
        async onClick_checkoutDeliveryHistory(item) {
            this.text_cardTitle_inner = "出貨單" + item.DeliveryOrderID;
            this.text_confirmBtn_inner = "";
            await this.getDeliveryOrder(item.DeliveryOrderID);
            this.deliveryOrderInfoDialog = true;
        },
        async onFinish_purchaseInfoDialog(item) {
            this.purchaseInfoDialog = false;
            await this.putPurchase(item);
        },
        async onFinish_deliveryOrderInfoDialog(item) {
            this.deliveryOrderInfoDialog = false;
            await this.putDeliveryOrder(item);
        },
        calcPurchaseInfo(item) {
            let vthis = this;
            item.map(function (x) {
                x.ImportCost = vthis.calcImportCost(x);

                x.WholesalePriceTwd = vthis.calcWholesalePriceTwd(x.WholesalePrice, x.ExchangeRateKrw);

                x.Costs = vthis.calcCost(x);

                //計算毛利
                x.GrossProfit = vthis.calcGrossProfit(vthis.productItem.RetailPrice, x.Costs);

                //計算毛利率
                x.GrossMargin = vthis.calcGrossMargin(vthis.productItem.RetailPrice, x.Costs);

            });
            return item;
        },
        calcDeliveryInfo(item) {
            let vthis = this;
            item.map(function (x) {

                //計算毛利
                x.GrossProfit = vthis.calcGrossProfit(x.RetailPrice, x.Cost);

                //計算毛利率
                x.GrossMargin = vthis.calcGrossMargin(x.RetailPrice, x.Cost);

            });
            return item;
        },
        calcWholesalePriceTwd(wholesalePrice, exchangeRateKrw) {
            if (isNaN(exchangeRateKrw) == true || exchangeRateKrw == null) {
                return undefined;
            }
            let result = 0;
            let wholesalePriceTwd = parseNumber(wholesalePrice) / parseNumber(exchangeRateKrw);
            if (isNaN(wholesalePriceTwd) == false) {
                result += wholesalePriceTwd;
            }
            return result.toFixed(2);
        },
        calcImportCost(item) {
            let result = 0;
            if (isNaN(item.ExchangeRateKrw) == true || item.ExchangeRateKrw == null) {
                return undefined;
            }

            //檢查分母不得為0
            let purchaseQTY = checkZero(item.PurchaseQTY);

            //計算貨運行抽成（商品金額*貨運行抽成百分比）
            let shippingAgentCutKrw = parseNumber(item.WholesalePrice) * (parseNumber(item.ShippingAgentCutPercent) / 100);

            //計算每個商品的國際運費（國際運費總額/進貨商品數量）
            let shippingFeeKokusaiDivideByPurchaseQTY = parseNumber(item.ShippingFeeKokusaiKrw) / purchaseQTY;

            //計算韓幣開銷總金額
            //商品批價（item.WholesalePrice）
            //貨運行抽成（shippingAgentCutKrw）
            //每個商品的國際運費（shippingFeeKokusaiDivideByPurchaseQTY）
            //韓國國內運費（item.ShippingFeeKr）
            let subtotalKrw = item.WholesalePrice + shippingAgentCutKrw + shippingFeeKokusaiDivideByPurchaseQTY + item.ShippingFeeKr;

            //計算台幣開銷總金額
            //關稅（韓國國內運費（item.TariffTwd）
            //台灣國內運費（韓國國內運費（item.ShippingFeeTw）
            let costTwd = parseNumber(item.TariffTwd) + parseNumber(item.ShippingFeeTw);

            //計算貨運關稅成本
            //換為台幣後的韓幣開銷（subtotalKrw / 韓國國內運費（item.ExchangeRateKrw)
            //每個商品的台幣開銷（台幣開銷總金額/商品個數）（costTwd / purchaseQTY）
            let importCost = (subtotalKrw / item.ExchangeRateKrw) + (costTwd / purchaseQTY);
            if (isNaN(importCost) == false) {
                result += importCost;
            }
            return result.toFixed(2);
        },
        calcCost(item) {
            //計算總成本
            //進口成本+其他成本（包裝廣告）
            let costs = 0;
            costs += parseFloat(item.ImportCost);
            this.tradingSettings.Costs.map(function (y) {
                costs += parseFloat(y.Value);
            });
            return costs.toFixed(2);
        },
        calcGrossProfit(retailPrice, cost) {
            let result = 0;
            let grossProfit = parseNumber(retailPrice) - parseNumber(cost);
            if (isNaN(grossProfit) == false) {
                result += grossProfit;
            }
            return result.toFixed(2);
        },
        calcGrossMargin(retailPrice, cost) {
            if (isNaN(retailPrice) == true || retailPrice == null) {
                x.GrossMargin = undefined;
                return;
            }
            let result = 0;
            let grossMargin = (parseNumber(retailPrice) - parseNumber(cost)) / parseNumber(retailPrice) * 100;
            if (isNaN(grossMargin) == false) {
                result += grossMargin;
            }
            return result.toFixed(2);
        },
        calcStocks(){
            let purchaseDetailQTYCount = 0;
            let deliveryDetailQTYCount = 0;
            this.purchaseHistories.map((x)=>{purchaseDetailQTYCount+=x.PurchaseDetailQTY});
            this.deliveryHistories.map((x)=>{deliveryDetailQTYCount+=x.QTY});
            return purchaseDetailQTYCount - deliveryDetailQTYCount;
        },
        preSendPurchase(item) {
            item.ID = parseFloat(item.ID);
            item.Status = parseFloat(item.Status);
            item.PurchaseType = parseFloat(item.PurchaseType);
            item.Weight = parseFloat(item.Weight);
            item.ShippingFeeKr = parseFloat(item.ShippingFeeKr);
            item.ShippingFeeTw = parseFloat(item.ShippingFeeTw);
            item.ShippingFeeKokusai = parseFloat(item.ShippingFeeKokusai);
            item.ExchangeRateKrw = parseFloat(item.ExchangeRateKrw);
            item.TotalKrw = parseFloat(item.TotalKrw);
            item.TotalTwd = parseFloat(item.TotalTwd);
            item.DataOrder = parseFloat(item.DataOrder);
            return item;
        },
        async getPurchase(purchaseID) {
            let filter = {
                ID: purchaseID,
            }
            await getPurchase(filter)
                .then(async (response) => {
                    if (response.data.records != null) {
                        this.purchase = response.data.records[0];
                    }
                })
                .catch((error) => {
                });
        },
        async putPurchase(item) {
            item = this.preSendPurchase(item);
            await putPurchase(item)
                .then(async (response) => {
                    await this.getPurchases();
                    this.alert = true;
                    this.alertType = "Success";
                    this.alertText = "編輯採購案成功";
                })
                .catch((error) => {
                    this.alert = true;
                    this.alertType = "Fail";
                    this.alertText = "編輯採購案失敗";
                });
        },
        preSendDelivery(item) {
            item.ID = parseFloat(item.ID);
            item.DeliveryType = parseFloat(item.DeliveryType);
            item.DeliveryStatus = parseFloat(item.DeliveryStatus);
            item.DeliveryFeeStatus = parseFloat(item.DeliveryFeeStatus);
            item.PaymentType = parseFloat(item.PaymentType);
            item.PaymentStatus = parseFloat(item.PaymentStatus);
            item.TotalOriginal = parseFloat(item.TotalOriginal);
            item.Discount = parseFloat(item.Discount);
            item.TotalDiscounted = parseFloat(item.TotalDiscounted);
            item.DataOrder = parseFloat(item.DataOrder);
            return item;
        },
        async getDeliveryOrder(deliveryOrderID) {
            let filter = {
                ID: deliveryOrderID,
            }
            await getDeliveryOrder(filter)
                .then((response) => {
                    if (response.data.records != null) {
                        this.deliveryOrder = response.data.records[0];
                    }
                    else {
                        this.deliveryOrder = {};
                    }
                })
                .catch((error) => {
                });
        },
        async putDeliveryOrder(item) {
            item = this.preSendDelivery(item);
            await putDeliveryOrder(item)
                .then(async (response) => {
                    await this.getDeliveryOrders();
                    this.alert = true;
                    this.alertType = "Success";
                    this.alertText = "編輯出貨單成功";
                })
                .catch((error) => {
                    this.alert = true;
                    this.alertType = "Fail";
                    this.alertText = "編輯出貨單失敗";
                });
        },
        async getProductPurchaseHistories() {
            await getProductPurchaseHistories(this.productItem)
                .then((response) => {
                    if (response.data.records != null) {
                        let item = this.calcPurchaseInfo(response.data.records);
                        this.purchaseHistories = item;
                    }
                    else {
                        this.purchaseHistories = [];
                    }
                });
        },
        async getProductDeliveryHistories() {
            await getProductDeliveryHistories(this.productItem)
                .then((response) => {
                    if (response.data.records != null) {
                        let item = this.calcDeliveryInfo(response.data.records);
                        this.deliveryHistories = item;
                    }
                    else {
                        this.deliveryHistories = [];
                    }
                });
        },
    },
    watch: {
        productInfoDialog: async function (newVal, oldVal) {
            if (newVal == true) {
                this.getProductPurchaseHistories();
                this.getProductDeliveryHistories();
            }
        },
    }
}


const parseNumber = function (x) {
    let parsed = parseFloat(x);
    if (isNaN(parsed) == true) {
        return 0;
    }
    return parsed;
};
const checkZero = function (x) {
    let parsed = parseFloat(x);
    if (isNaN(parsed) == true || parsed == 0) {
        return 1;
    }
    return parsed;
};
</script>