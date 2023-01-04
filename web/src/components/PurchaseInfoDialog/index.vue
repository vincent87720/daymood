<template>
    <v-dialog v-model="purchaseInfoDialog" hide-overlay @click:outside="onClick_cancel" fullscreen
        transition="slide-x-reverse-transition">
        <v-card class="d-flex align-center flex-column justify-center">
            <v-btn icon dark absolute top left @click="onClick_cancel">
                <v-icon>mdi-arrow-left</v-icon>
            </v-btn>
            <v-btn icon dark absolute top right @click="onClick_newDetailButton" v-if="isEditEnable == true">
                <v-icon>mdi-plus</v-icon>
            </v-btn>
            <v-btn icon dark absolute bottom right @click="onClick_upload" v-if="isEditEnable == true">
                <v-icon>mdi-upload</v-icon>
            </v-btn>
            <v-btn icon dark absolute top right @click="enableEdit = true" v-if="isEditEnable == false">
                <v-icon>mdi-pencil</v-icon>
            </v-btn>
            <v-btn icon dark absolute bottom left @click="onClick_finishPurchaseButton" v-if="isEditEnable == true">
                <v-icon>mdi-cart-check</v-icon>
            </v-btn>
            <v-card-title class="text-h5 mt-8">
                {{ text_cardTitle }}
            </v-card-title>
            <v-container class="pa-6 mb-8">
                <v-row>
                    <v-col xs="12" sm="12" class="ml-auto mr-auto">
                        <v-card outlined>
                            <c-data-table :prop_headers="purchaseDetailHeader" :prop_items="purchaseDetails"
                                :prop_search="search" @edit="onClick_editButton" @delete="onClick_deleteButton">
                                <template v-slot:item.Status="{ item }">
                                    <v-chip :color="getStatusChipColor(item.Status)" dark small>
                                        {{ convertDisplayText(systemConfigs.PurchaseDetailStatus, item.Status) }}
                                    </v-chip>
                                </template>
                                <template v-slot:item.SupplierID="{ item }">
                                    {{ convertDisplayText(allSuppliersList, item.SupplierID) }}
                                </template>
                                <template v-slot:item.ProductID="{ item }">
                                    {{ convertDisplayText_Products(allProductsList, item.ProductID) }}
                                </template>
                                <template v-slot:item.actions="{ item }">
                                    <v-icon small class="mx-1" @click.stop="onClick_editButton(item)"
                                        v-if="isEditEnable == true">
                                        mdi-pencil
                                    </v-icon>
                                    <v-icon small class="mx-1" @click.stop="onClick_deleteButton(item)"
                                        v-if="isEditEnable == true">
                                        mdi-delete
                                    </v-icon>
                                </template>
                            </c-data-table>
                        </v-card>
                    </v-col>
                </v-row>
                <v-row>
                    <v-col xs="12" sm="12" class="ml-auto mr-auto">
                        <v-card outlined class="pa-3 d-flex justify-end">
                            <div class="d-flex flex-column justify-end mr-3 ml-5" style="color: gray" v-if="purchaseItem.PurchaseType == 1">
                                <h2>商品總數</h2>
                                <h2>商品總計</h2>
                                <v-tooltip left>
                                    <template v-slot:activator="{ on, attrs }">
                                        <h2 v-bind="attrs" v-on="on">貨運抽成</h2>
                                    </template>
                                    <span>{{ tooltip.calc_Ajeossi }}</span>
                                </v-tooltip>
                                <h2>韓國運費</h2>
                                <v-tooltip left>
                                    <template v-slot:activator="{ on, attrs }">
                                        <h2 v-bind="attrs" v-on="on">國際運費</h2>
                                    </template>
                                    <span>{{ tooltip.calc_ShippingFeeKokusaiKrw }}</span>
                                </v-tooltip>
                                <v-tooltip bottom>
                                    <template v-slot:activator="{ on, attrs }">
                                        <h2 v-bind="attrs" v-on="on">韓圓總計</h2>
                                    </template>
                                    <span>{{ tooltip.calc_TotalKrw }}</span>
                                </v-tooltip>
                            </div>
                            <div class="d-flex flex-column justify-end align-end" v-if="purchaseItem.PurchaseType == 1">
                                <h2>{{ calc_TotalQTY }}</h2>
                                <h2>₩ {{ calc_Subtotals }}</h2>
                                <h2>₩ {{ calc_Ajeossi }}</h2>
                                <h2>₩ {{ calc_ShippingFeeKr }}</h2>
                                <h2>₩ {{ calc_ShippingFeeKokusaiKrw }}</h2>
                                <h2>₩ {{ calc_TotalKrw }}</h2>
                            </div>
                            <div class="d-flex flex-column justify-end mr-3 ml-5" style="color: gray">
                                <h2 v-if="purchaseItem.PurchaseType == 2">商品總數</h2>
                                <h2 v-if="purchaseItem.PurchaseType == 2">商品總計</h2>
                                <h2>台灣運費</h2>
                                <h2 v-if="purchaseItem.PurchaseType == 1">台灣關稅</h2>
                                <h2>台幣總計</h2>
                            </div>
                            <div class="d-flex flex-column justify-end align-end">
                                <h2 v-if="purchaseItem.PurchaseType == 2">{{ calc_TotalQTY }}</h2>
                                <h2 v-if="purchaseItem.PurchaseType == 2">{{ calc_Subtotals }}</h2>
                                <h2>$ {{ calc_ShippingFeeTw }}</h2>
                                <h2 v-if="purchaseItem.PurchaseType == 1">$ {{ calc_Tariff }}</h2>
                                <h2>$ {{ calc_TotalTwd }}</h2>
                            </div>
                            <div class="d-flex flex-column justify-end mr-3 ml-5" style="color: gray">
                                <h2>總計</h2>
                            </div>
                            <div class="d-flex flex-column justify-end align-end">
                                <h2>$ {{ calc_Total }}</h2>
                            </div>
                        </v-card>
                    </v-col>
                </v-row>
            </v-container>
        </v-card>
        <PurchaseDetailDialog :prop_purchaseDetailDialog.sync="purchaseDetailDialog"
            :prop_text_cardTitle="text_cardTitle_inner" :prop_text_confirmBtn="text_confirmBtn_inner"
            :prop_purchaseItem="purchaseItem" :prop_purchaseDetailItem="purchaseDetail"
            @confirm='onConfirm_purchaseDetailDialog' />
        <PurchaseDetailImportDialog :prop_purchaseDetailImportDialog.sync="purchaseDetailImportDialog"
            :prop_text_cardTitle="text_cardTitle_inner" :prop_text_confirmBtn="text_confirmBtn_inner"
            @confirm='onConfirm_purchaseDetailImportDialog' />
        <ConfirmDialog :prop_confirmDialog.sync="confirmDialog" :prop_text_cardTitle="text_cardTitle_inner"
            :prop_text_cardHint="text_cardHint_inner" :prop_text_confirmBtn="text_confirmBtn_inner"
            :prop_confirmTarget.sync="confirmTarget" v-on:confirmClick='onConfirm_confirmDialog' />
        <Alert :prop_alert.sync="alert" :prop_alertType="alertType" :prop_alertText="alertText"></Alert>
    </v-dialog>
</template>

<script>
import Alert from '../../components/Alert/index.vue'
import ConfirmDialog from '../../components/ConfirmDialog/index.vue'
import DatePicker from "../../components/Pickers/DatePicker.vue";
import BtnAdd from "../../components/Buttons/BtnAdd.vue";
import BtnUpload from "../../components/Buttons/BtnUpload.vue";
import DataTable from "../../components/DataTables/DataTable.vue";
import PurchaseDetailDialog from '../../components/PurchaseDetailDialog/index.vue';
import PurchaseDetailImportDialog from '../../components/PurchaseDetailImportDialog/index.vue';
import { getPurchaseDetails, postPurchaseDetails, postPurchaseDetail, putPurchaseDetail, deletePurchaseDetail } from "../../apis/PurchaseDetailsAPI";

class PurchaseDetail {
    ID = undefined;
    NamedID = "";
    Name = "";
    Status = undefined;
    WholesalePrice = undefined;
    QTY = undefined;
    Cost = undefined;
    Currency = undefined;
    Subtotal = undefined;
    Remark = "";
    DataOrder = undefined;
    CreateAt = "";
    UpdateAt = "";
    PurchaseID = undefined;
    SupplierID = undefined;
    ProductID = undefined;
}

export default {
    name: 'purchaseInfoDialog',
    components: {
        Alert,
        ConfirmDialog,
        PurchaseDetailDialog,
        PurchaseDetailImportDialog,
        "c-btn-add": BtnAdd,
        "c-btn-upload": BtnUpload,
        "c-date-picker": DatePicker,
        "c-data-table": DataTable,
    },
    data() {
        return {
            //Alert
            alert: false,
            alertType: "",
            alertText: "",

            // ConfirmDialog
            confirmDialog: false,
            text_cardTitle_inner: '',
            text_cardHint_inner: '',
            text_confirmBtn_inner: '',
            confirmTarget: null,

            search: '',
            actionType: "",
            enableEdit: false,

            purchaseDetail: new PurchaseDetail(),
            purchaseDetailDialog: false,
            purchaseDetails: [],
            purchaseDetailHeader: [
                { text: '是否採用', value: 'Status', align: 'center'},
                { text: '採購商品名稱', value: 'Name' },
                { text: '商品', value: 'ProductID' },
                { text: '廠商', value: 'SupplierID' },
                { text: '批價', value: 'WholesalePrice' },
                { text: '數量', value: 'QTY' },
                { text: '小計', value: 'Subtotal' },
                { text: "備註", value: "Remark"},
                { text: '', value: 'actions', sortable: false},
            ],
            products: [],

            purchaseDetailImportDialog: false,

            tooltip:{
                calc_Ajeossi: "",
                calc_ShippingFeeKokusaiKrw: "",
            }
        };
    },
    props: {
        prop_purchaseInfoDialog: {
            type: Boolean,
            required: true
        },
        prop_text_cardTitle: {
            type: String,
            required: true
        },
        prop_text_confirmBtn: {
            type: String,
            required: true
        },
        prop_purchaseItem: {
            type: Object,
            required: false
        },
    },
    mounted() {
        this.$store.dispatch("GetProducts");
        this.$store.dispatch("GetSuppliers");
    },
    computed: {
        purchaseInfoDialog: {
            get() {
                return this.prop_purchaseInfoDialog
            },
            set(val) {
                this.$emit('update:prop_purchaseInfoDialog', val)

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
        text_confirmBtn: {
            get() {
                return this.prop_text_confirmBtn
            },
            set(val) {
                this.$emit('update:prop_text_confirmBtn', val)
            }
        },
        purchaseItem: {
            get() {
                return this.prop_purchaseItem
            },
            set(val) {
                this.$emit('update:prop_purchaseItem', val)
            }
        },
        systemConfigs() {
            return this.$store.state.systemConfigs;
        },
        tradingSettings() {
            return this.$store.state.tradingSettings;
        },
        allProductsList() {
            return this.$store.state.allProducts;
        },
        allSuppliersList() {
            return this.$store.state.allSuppliers;
        },
        isEditEnable() {
            if (this.purchaseItem.Status == 1 || this.enableEdit == true) {
                return true;
            }
            return false;
        },
        calc_TotalQTY() {
            let result = 0;
            this.purchaseDetails.map(function (item) {
                if (item.Status == 1) {
                    result += item.QTY;
                }
            })
            return result;
        },
        calc_Subtotals() {
            //商品總計
            let result = 0;
            this.purchaseDetails.map(function (item) {
                if (item.Status == 1) {
                    result += parseFloat(item.Subtotal);
                }
            })
            return result.toFixed(2);
        },
        calc_Ajeossi() {
            //貨運行抽成
            let result = 0;
            let subtotals = parseFloat(this.calc_Subtotals);
            let ajeossiPercent = parseFloat(this.purchaseItem.ShippingAgentCutPercent) / 100;
            let ajeossi = subtotals * ajeossiPercent;
            if (isNaN(ajeossi) == false) {
                result += ajeossi
            }
            this.tooltip.calc_Ajeossi = `₩${subtotals}(商品總額) ＊ ${this.purchaseItem.ShippingAgentCutPercent}%(貨運行抽成百分比) ＝ ₩${ajeossi}`
            return result.toFixed(2);

        },
        calc_ShippingFeeKr() {
            //韓國國內運費
            let result = 0;
            let shippingFeeKr = parseFloat(this.purchaseItem.ShippingFeeKr);
            if (isNaN(shippingFeeKr) == false) {
                result += shippingFeeKr
            }
            return result.toFixed(2);
        },
        calc_ShippingFeeKokusaiKrw() {
            //國際運費
            let result = 0;
            let shippingFeeKokusaiKrw = parseFloat(this.purchaseItem.Weight * this.purchaseItem.ShippingFeeKokusaiPerKilo);
            if (isNaN(shippingFeeKokusaiKrw) == false) {
                result += shippingFeeKokusaiKrw
            }
            this.tooltip.calc_ShippingFeeKokusaiKrw = `${this.purchaseItem.Weight}kg(貨運總重) ＊ ₩${this.purchaseItem.ShippingFeeKokusaiPerKilo}(每公斤國際運費) ＝ ₩${shippingFeeKokusaiKrw}`
            return result.toFixed(2);
        },
        calc_TotalKrw() {
            //韓圓總計 = 商品總計（飾品） + 韓國國內運費 + 國際運費（韓圓）
            let result = 0;
            let totalKrw = 0;
            let subtotals = 0;
            let ajeossi = 0;
            let shippingFeeKr = 0;
            let shippingFeeKokusaiKrw = 0;
            this.tooltip.calc_TotalKrw = "";
            if (this.purchaseItem.PurchaseType == 1) {
                subtotals = parseFloat(this.calc_Subtotals);
                ajeossi = parseFloat(this.calc_Ajeossi);
                this.tooltip.calc_TotalKrw += `₩${subtotals}(商品總額) ＋ ₩${ajeossi}(貨運抽成) ＋ `
            }
            shippingFeeKr = parseFloat(this.calc_ShippingFeeKr);
            shippingFeeKokusaiKrw = parseFloat(this.calc_ShippingFeeKokusaiKrw);
            totalKrw = subtotals + ajeossi + shippingFeeKr + shippingFeeKokusaiKrw;
            if (isNaN(totalKrw) == false) {
                result += totalKrw
            }
            this.tooltip.calc_TotalKrw += `₩${shippingFeeKr}(韓國國內運費) ＋ ₩${shippingFeeKokusaiKrw}(國際運費) ＝ ₩${totalKrw}`
            return result.toFixed(2);
        },
        calc_ShippingFeeTw() {
            //台灣國內運費
            let result = 0;
            let shippingFeeTw = parseFloat(this.purchaseItem.ShippingFeeTw);
            if (isNaN(shippingFeeTw) == false) {
                result += shippingFeeTw
            }
            return result.toFixed(2);
        },
        calc_Tariff() {
            //關稅
            let result = 0;
            let tariff = parseFloat(this.purchaseItem.Weight * this.purchaseItem.TariffPerKilo);
            if (isNaN(tariff) == false) {
                result += tariff
            }
            return result.toFixed(2);
        },
        calc_TotalTwd() {
            //台幣總計 = 商品總計（耗材）+ 台灣國內運費 + 關稅
            let result = 0;
            let totalTwd = 0;
            if (this.purchaseItem.PurchaseType == 2) {
                totalTwd += parseFloat(this.calc_Subtotals);
            }
            totalTwd += parseFloat(this.calc_ShippingFeeTw);
            totalTwd += parseFloat(this.calc_Tariff);
            if (isNaN(totalTwd) == false) {
                result += totalTwd
            }
            return result.toFixed(2);
        },
        calc_Total() {
            //總計 = (韓圓總計 / 匯率) + 台幣總計
            let result = 0;
            let total = 0;
            if(this.purchaseItem.ExchangeRateKrw != null){
                let totalKrw = parseFloat(this.calc_TotalKrw);
                let exchangeRateKrw = parseFloat(this.purchaseItem.ExchangeRateKrw);
                let totalExchangedKrw = totalKrw / exchangeRateKrw;
                if (isNaN(totalExchangedKrw) == false) {
                    total += totalExchangedKrw
                }
            }
            let totalTwd = parseFloat(this.calc_TotalTwd);
            total += totalTwd;
            if (isNaN(total) == false) {
                result += total
            }
            return result.toFixed(2);
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
        convertDisplayText_Products(list, key) {
            let result = list.find(x => x.key == key);
            if (result) {
                return `${result.SKU} ${result.value}`
            }
            return "";
        },
        convertDataList(data) {
            data.map(x => x = { key: x.ID, value: x.SKU + " " + x.Name });
            return data
        },
        getStatusChipColor(status) {
            if (status == 1) return '#849A8F'
            // #97A48D
            else if (status == 2) return '#965455'
            else return ''
        },
        onClick_cancel() {
            this.purchaseInfoDialog = false;
            this.enableEdit = false;
        },
        onClick_newDetailButton() {
            this.text_cardTitle_inner = "新增採購明細";
            this.text_confirmBtn_inner = "新增";
            this.actionType = "post";
            this.purchaseDetail = new PurchaseDetail();
            this.purchaseDetailDialog = true;
        },
        onClick_editButton(item) {
            this.text_cardTitle_inner = "編輯採購明細";
            this.text_confirmBtn_inner = "修改";
            this.actionType = "put";
            this.purchaseDetail = item;
            this.purchaseDetailDialog = true;
        },
        onClick_deleteButton(item) {
            this.text_cardTitle_inner = "確認刪除";
            this.text_confirmBtn_inner = "刪除";
            this.actionType = "delete";
            this.confirmDialog = true;
            this.confirmTarget = item;
        },
        onClick_finishPurchaseButton() {
            this.text_cardTitle_inner = "是否確定結案？";
            this.text_confirmBtn_inner = "結案";
            this.actionType = "finish";
            this.confirmDialog = true;
        },
        onClick_upload() {
            this.text_cardTitle_inner = "匯入";
            this.text_confirmBtn_inner = "確定";
            this.purchaseDetailImportDialog = true;
        },
        async onConfirm_purchaseDetailDialog(item) {
            this.purchaseDetailDialog = false;
            if (this.actionType == "post") {
                await this.postPurchaseDetail(item);
            } else if (this.actionType == "put") {
                await this.putPurchaseDetail(item);
            }
        },
        async onConfirm_purchaseDetailImportDialog(item) {
            this.purchaseDetailImportDialog = false;
            await this.postPurchaseDetails(item);
        },
        async onConfirm_confirmDialog(item) {
            this.confirmDialog = false;
            if (this.actionType == "delete") {
                await this.deletePurchaseDetail(item);
            } else if (this.actionType == "finish") {
                this.beforePurchaseFinish();
                this.$emit('finish', this.purchaseItem);//觸發一個在子元件中宣告的事件 childEvnet
            }
        },
        beforePurchaseFinish() {
            this.purchaseItem.QTY = this.calc_TotalQTY;
            this.purchaseItem.ShippingAgentCutKrw = this.calc_Ajeossi;
            this.purchaseItem.ShippingFeeKokusaiKrw = this.calc_ShippingFeeKokusaiKrw;
            this.purchaseItem.TariffTwd = this.calc_Tariff;
            this.purchaseItem.TotalKrw = this.calc_TotalKrw;
            this.purchaseItem.TotalTwd = this.calc_TotalTwd;
            this.purchaseItem.Total = this.calc_Total;
            this.purchaseItem.Status = 2;
        },
        preSend(item) {
            item.ID = parseFloat(item.ID);
            item.Status = parseFloat(item.Status);
            item.QTY = parseFloat(item.QTY);
            item.WholesalePrice = parseFloat(item.WholesalePrice);
            item.Cost = parseFloat(item.Cost);
            item.Currency = parseFloat(item.Currency);
            item.Subtotal = parseFloat(item.Subtotal);
            item.DataOrder = parseFloat(item.DataOrder);
            item.PurchaseID = parseFloat(item.PurchaseID);
            item.SupplierID = parseFloat(item.SupplierID);
            item.ProductID = parseFloat(item.ProductID);
            return item;
        },
        async getPurchaseDetails() {
            let filter = {
                PurchaseID: this.purchaseItem.ID,
            }
            await getPurchaseDetails(filter)
                .then((response) => {
                    if (response.data.records != null) {
                        this.purchaseDetails = response.data.records;
                    }
                    else {
                        this.purchaseDetails = [];
                    }
                })
                .catch((error) => {
                });
        },
        async postPurchaseDetail(item) {
            item = this.preSend(item);
            item.PurchaseID = this.purchaseItem.ID;
            await postPurchaseDetail(item)
                .then(async (response) => {
                    await this.getPurchaseDetails();
                    this.alert = true;
                    this.alertType = "Success";
                    this.alertText = "新增採購明細成功";
                })
                .catch((error) => {
                    this.alert = true;
                    this.alertType = "Fail";
                    this.alertText = "新增採購明細失敗";
                });
        },
        async putPurchaseDetail(item) {
            item = this.preSend(item);
            await putPurchaseDetail(item)
                .then(async (response) => {
                    await this.getPurchaseDetails();
                    this.alert = true;
                    this.alertType = "Success";
                    this.alertText = "編輯採購明細成功";
                })
                .catch((error) => {
                    this.alert = true;
                    this.alertType = "Fail";
                    this.alertText = "編輯採購明細失敗";
                });
        },
        async deletePurchaseDetail(item) {
            await deletePurchaseDetail(item)
                .then(async (response) => {
                    await this.getPurchaseDetails();
                    this.alert = true;
                    this.alertType = "Success";
                    this.alertText = "刪除採購明細成功";
                })
                .catch((error) => {
                    this.alert = true;
                    this.alertType = "Fail";
                    this.alertText = "刪除採購明細失敗";
                });
        },
        async postPurchaseDetails(item) {
            item.map(x => {
                x.PurchaseID = this.purchaseItem.ID;
                x = this.preSend(x);
            });
            await postPurchaseDetails(item)
                .then(async (response) => {
                    await this.getPurchaseDetails();
                    this.alert = true;
                    this.alertType = "Success";
                    this.alertText = "新增採購明細成功";
                })
                .catch((error) => {
                    this.alert = true;
                    this.alertType = "Fail";
                    this.alertText = "新增採購明細失敗";
                });
        },
    },
    watch: {
        purchaseInfoDialog: async function (newVal, oldVal) {
            if (newVal == true) {
                await this.getPurchaseDetails();
            }
        },
    }
}
</script>