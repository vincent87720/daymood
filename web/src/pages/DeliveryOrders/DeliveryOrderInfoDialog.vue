<template>
    <v-dialog v-model="deliveryOrderInfoDialog" hide-overlay @click:outside="onClick_cancel" fullscreen
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
            <v-btn icon dark absolute bottom left @click="onClick_finishDeliveryOrderButton"
                v-if="isEditEnable == true">
                <v-icon>mdi-cart-check</v-icon>
            </v-btn>
            <v-card-title class="text-h5 mt-8">
                {{ text_cardTitle }}
            </v-card-title>
            <v-container class="pa-6 mb-8">
                <v-row>
                    <v-col xs="12" sm="12" class="ml-auto mr-auto">
                        <v-card outlined rounded="lg">
                            <c-data-table :prop_headers="deliveryOrderDetailHeader" :prop_items="deliveryOrderDetails"
                                :prop_search="search" @edit="onClick_editDetailButton"
                                @delete="onClick_deleteDetailButton">
                                <template v-slot:item.ProductID="{ item }">
                                    {{
                                        convertDisplayText_Products(allProductsList, item.ProductID)
                                    }}
                                </template>
                                <template v-slot:item.actions="{ item }">
                                    <v-icon small class="mx-1" @click.stop="onClick_editDetailButton(item)"
                                        v-if="isEditEnable == true">
                                        mdi-pencil
                                    </v-icon>
                                    <v-icon small class="mx-1" @click.stop="onClick_deleteDetailButton(item)"
                                        v-if="isEditEnable == true">
                                        mdi-delete
                                    </v-icon>
                                </template>
                            </c-data-table>
                        </v-card>
                    </v-col>
                </v-row>
                <v-row>
                    <v-col xs="12" sm="6" class="ml-auto mr-auto">
                        <c-card-rounded>
                            <c-data-table :prop_headers="discountHeader" :prop_items="discounts">
                                <template v-slot:item.DiscountType="{ item }">
                                    {{
                                        convertDisplayText(
                                            systemConfigs.DiscountType,
                                        item.DiscountType
                                                                        )
                                    }}
                                </template>
                                <template v-slot:item.actions="{ item }">
                                    <v-icon small class="mx-1" @click.stop="onClick_editDiscountButton(item)"
                                        v-if="isEditEnable == true">
                                        mdi-pencil
                                    </v-icon>
                                    <v-icon small class="mx-1" @click.stop="onClick_deleteDiscountButton(item)"
                                        v-if="isEditEnable == true">
                                        mdi-delete
                                    </v-icon>
                                </template>
                                <template v-slot:footer>
                                    <v-container fluid>
                                        <v-row class="d-flex justify-center">
                                            <v-col>
                                                <v-btn outlined block text
                                                    @click="onClick_newDiscountButton()">新增折扣</v-btn>
                                            </v-col>
                                        </v-row>
                                    </v-container>
                                </template>
                            </c-data-table>
                        </c-card-rounded>
                    </v-col>
                    <v-col xs="12" sm="6" class="ml-auto mr-auto">
                        <c-card-rounded class="pa-3 d-flex justify-end mb-6">
                            <div class="d-flex flex-column justify-end mr-3 ml-5" style="color: gray">
                                <h2>毛利</h2>
                            </div>
                            <div class="d-flex flex-column justify-end align-end">
                                <h2>{{ calcGrossProfit(calc_Total, calc_Cost) }}</h2>
                            </div>
                            <div class="d-flex flex-column justify-end mr-3 ml-5" style="color: gray">
                                <h2>毛利率</h2>
                            </div>
                            <div class="d-flex flex-column justify-end align-end">
                                <h2>{{ calcGrossMargin(calc_Total, calc_Cost) }}%</h2>
                            </div>
                        </c-card-rounded>
                        <c-card-rounded class="pa-3 d-flex justify-end">
                            <div class="d-flex flex-column justify-end mr-3 ml-5" style="color: gray">
                                <h2>商品總數</h2>
                                <h2>商品總計</h2>
                                <h2>貨運運費</h2>
                                <h2>折扣金額</h2>
                            </div>
                            <div class="d-flex flex-column justify-end align-end">
                                <h2>{{ calc_TotalQTY }}</h2>
                                <h2>${{ calc_Subtotals }}</h2>
                                <h2>$-{{ calc_DeliveryFee }}</h2>
                                <h2>$-{{ calc_Discount }}</h2>
                            </div>
                            <div class="d-flex flex-column justify-end mr-3 ml-5" style="color: gray">
                                <h2>總計</h2>
                            </div>
                            <div class="d-flex flex-column justify-end align-end">
                                <h2>$ {{ calc_Total }}</h2>
                            </div>
                        </c-card-rounded>
                    </v-col>
                </v-row>
            </v-container>
        </v-card>
        <DeliveryOrderDetailDialog :prop_deliveryOrderDetailDialog.sync="deliveryOrderDetailDialog"
            :prop_text_cardTitle="text_cardTitle_inner" :prop_text_confirmBtn="text_confirmBtn_inner"
            :prop_deliveryOrderItem="deliveryOrderItem" :prop_deliveryOrderDetailItem="deliveryOrderDetail"
            @confirm="onConfirm_deliveryOrderDetailDialog" />
        <DiscountDialog :prop_discountDialog.sync="discountDialog" :prop_text_cardTitle="text_cardTitle_inner"
            :prop_text_confirmBtn="text_confirmBtn_inner" :prop_deliveryOrderItem="deliveryOrderItem"
            :prop_discountItem="discount" @confirm="onConfirm_discountDialog" />
        <!-- <DeliveryOrderDetailImportDialog :prop_deliveryOrderDetailImportDialog.sync="deliveryOrderDetailImportDialog"
            :prop_text_cardTitle="text_cardTitle_inner" :prop_text_confirmBtn="text_confirmBtn_inner"
            @confirm='onConfirm_deliveryOrderDetailImportDialog' /> -->
        <ConfirmDialog :prop_confirmDialog.sync="confirmDialog" :prop_text_cardTitle="text_cardTitle_inner"
            :prop_text_cardHint="text_cardHint_inner" :prop_text_confirmBtn="text_confirmBtn_inner"
            :prop_confirmTarget.sync="confirmTarget" v-on:confirmClick="onConfirm_confirmDialog" />
        <Alert :prop_alert.sync="alert" :prop_alertType="alertType" :prop_alertText="alertText"></Alert>
    </v-dialog>
</template>

<script>
import Alert from "@/components/Alert/index.vue";
import ConfirmDialog from "@/components/ConfirmDialog/index.vue";
import DatePicker from "@/components/Pickers/DatePicker.vue";
import BtnAdd from "@/components/Buttons/BtnAdd.vue";
import BtnUpload from "@/components/Buttons/BtnUpload.vue";
import CardRounded from "@/components/Cards/CardRounded.vue";
import DataTable from "@/components/DataTables/DataTable.vue";
import DeliveryOrderDetailDialog from "@/pages/DeliveryOrders/DeliveryOrderDetailDialog.vue";
import DiscountDialog from "@/pages/DeliveryOrders/DiscountDialog.vue";
// import DeliveryOrderDetailImportDialog from '../../components/DeliveryOrderDetailImportDialog/index.vue';
import {
    getDeliveryOrderDetails,
    postDeliveryOrderDetails,
    postDeliveryOrderDetail,
    putDeliveryOrderDetail,
    deleteDeliveryOrderDetail,
} from "@/apis/DeliveryOrderDetailsAPI";
import {
    getDiscounts,
    postDiscount,
    putDiscount,
    deleteDiscount,
} from "@/apis/DiscountsAPI";

class DeliveryOrderDetail {
    ID = undefined;
    RetailPrice = undefined;
    QTY = undefined;
    Subtotal = undefined;
    Remark = "";
    DataOrder = undefined;
    CreateAt = "";
    UpdateAt = "";
    DeliveryOrderID = undefined;
    ProductID = undefined;
}

class Discount {
    ID = undefined;
    Name = "";
    Price = undefined;
    DiscountType = undefined;
    Remark = "";
    DataOrder = undefined;
    CreateAt = "";
    UpdateAt = "";
    DeliveryOrderID = undefined;
}

export default {
    name: "deliveryOrderInfoDialog",
    components: {
        Alert,
        ConfirmDialog,
        DeliveryOrderDetailDialog,
        DiscountDialog,
        // DeliveryOrderDetailImportDialog,
        "c-btn-add": BtnAdd,
        "c-btn-upload": BtnUpload,
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

            // ConfirmDialog
            confirmDialog: false,
            text_cardTitle_inner: "",
            text_cardHint_inner: "",
            text_confirmBtn_inner: "",
            confirmTarget: null,

            search: "",
            actionType: "",
            enableEdit: false,

            deliveryOrderDetail: new DeliveryOrderDetail(),
            deliveryOrderDetailDialog: false,
            deliveryOrderDetails: [],
            deliveryOrderDetailHeader: [
                { text: "商品", value: "ProductID" },
                { text: "成本", value: "Cost" },
                { text: "售價", value: "RetailPrice" },
                { text: "數量", value: "QTY" },
                { text: "小計", value: "Subtotal" },
                { text: "備註", value: "Remark" },
                { text: "", value: "actions", sortable: false },
            ],
            discount: new Discount(),
            discounts: [],
            discountDialog: false,
            discountHeader: [
                { text: "折扣名稱", value: "Name" },
                { text: "折扣方式", value: "DiscountType" },
                { text: "折扣金額", value: "Price" },
                { text: "備註", value: "Remark" },
                { text: "", value: "actions", sortable: false },
            ],

            // deliveryOrderDetailImportDialog: false,
        };
    },
    props: {
        prop_deliveryOrderInfoDialog: {
            type: Boolean,
            required: true,
        },
        prop_text_cardTitle: {
            type: String,
            required: true,
        },
        prop_text_confirmBtn: {
            type: String,
            required: true,
        },
        prop_deliveryOrderItem: {
            type: Object,
            required: false,
        },
    },
    mounted() {
        this.$store.dispatch("GetProducts");
    },
    computed: {
        deliveryOrderInfoDialog: {
            get() {
                return this.prop_deliveryOrderInfoDialog;
            },
            set(val) {
                this.$emit("update:prop_deliveryOrderInfoDialog", val);
            },
        },
        text_cardTitle: {
            get() {
                return this.prop_text_cardTitle;
            },
            set(val) {
                this.$emit("update:prop_text_cardTitle", val);
            },
        },
        text_confirmBtn: {
            get() {
                return this.prop_text_confirmBtn;
            },
            set(val) {
                this.$emit("update:prop_text_confirmBtn", val);
            },
        },
        deliveryOrderItem: {
            get() {
                return this.prop_deliveryOrderItem;
            },
            set(val) {
                this.$emit("update:prop_deliveryOrderItem", val);
            },
        },
        allProductsList() {
            return this.$store.state.data.allProducts;
        },
        isEditEnable() {
            if (this.deliveryOrderItem.Status == 1 || this.enableEdit == true) {
                return true;
            }
            return false;
        },
        calc_TotalQTY() {
            //數量
            let result = 0;
            this.deliveryOrderDetails.map(function (item) {
                result += parseFloat(item.QTY);
            });
            return result;
        },
        calc_Subtotals() {
            //商品總計
            let result = 0;
            this.deliveryOrderDetails.map(function (item) {
                result += parseFloat(item.Subtotal);
            });
            return result.toFixed(2);
        },
        calc_DeliveryFee() {
            let result = 0;
            let shippingFee = this.systemConfigs.DeliveryTypeShippingFee.find(
                (x) => x.key == this.deliveryOrderItem.DeliveryType
            );
            if (
                this.deliveryOrderItem.DeliveryFeeStatus == 2 &&
                shippingFee != undefined
            ) {
                //運費由賣家支付
                result += parseFloat(shippingFee.value);
            }
            return result.toFixed(2);
        },
        calc_Discount() {
            let result = 0;
            this.discounts.map((x) => (result += parseFloat(x.Price)));
            return result.toFixed(2);
        },
        calc_Cost() {
            let result = 0;
            this.deliveryOrderDetails.map((x) => (result += parseFloat(x.Cost)));
            return result.toFixed(2);
        },
        calc_Total() {
            //總計 = 商品總計 - 折扣 - 賣家負擔運費
            let result = 0;
            let total =
                parseFloat(this.calc_Subtotals) -
                parseFloat(this.calc_Discount) -
                parseFloat(this.calc_DeliveryFee);
            if (isNaN(total) == false) {
                result += total;
            }
            return result.toFixed(2);
        },
        systemConfigs() {
            return this.$store.state.conf.systemConfigs;
        },
    },
    methods: {
        convertDisplayText(list, key) {
            let result = list.find((x) => x.key == key);
            if (result) {
                return result.value;
            }
            return "";
        },
        convertDisplayText_Products(list, key) {
            let result = list.find((x) => x.key == key);
            if (result) {
                return `${result.SKU} ${result.value}`;
            }
            return "";
        },
        convertDisplayText_Date(datetime) {
            let result = "";
            if (datetime) {
                result = datetime.substring(0, 10);
            }
            return result;
        },
        onClick_cancel() {
            this.deliveryOrderInfoDialog = false;
            this.enableEdit = false;
        },
        onClick_newDetailButton() {
            this.text_cardTitle_inner = "新增出貨明細";
            this.text_confirmBtn_inner = "新增";
            this.actionType = "post";
            this.deliveryOrderDetail = new DeliveryOrderDetail();
            this.deliveryOrderDetailDialog = true;
        },
        onClick_editDetailButton(item) {
            this.text_cardTitle_inner = "編輯出貨明細";
            this.text_confirmBtn_inner = "修改";
            this.actionType = "put";
            this.deliveryOrderDetail = item;
            this.deliveryOrderDetailDialog = true;
        },
        onClick_deleteDetailButton(item) {
            this.text_cardTitle_inner = "確認刪除";
            this.text_confirmBtn_inner = "刪除";
            this.actionType = "deleteDetail";
            this.confirmDialog = true;
            this.confirmTarget = item;
        },
        onClick_finishDeliveryOrderButton() {
            this.text_cardTitle_inner = "是否確定結案？";
            this.text_confirmBtn_inner = "結案";
            this.actionType = "finish";
            this.confirmDialog = true;
        },
        onClick_upload() {
            this.text_cardTitle_inner = "匯入";
            this.text_confirmBtn_inner = "確定";
            this.deliveryOrderDetailImportDialog = true;
        },
        onClick_newDiscountButton() {
            this.text_cardTitle_inner = "新增折扣";
            this.text_confirmBtn_inner = "新增";
            this.actionType = "post";
            this.discount = new Discount();
            this.discountDialog = true;
        },
        onClick_editDiscountButton(item) {
            this.text_cardTitle_inner = "編輯折扣";
            this.text_confirmBtn_inner = "修改";
            this.actionType = "put";
            this.discount = item;
            this.discountDialog = true;
        },
        onClick_deleteDiscountButton(item) {
            this.text_cardTitle_inner = "確認刪除";
            this.text_confirmBtn_inner = "刪除";
            this.actionType = "deleteDiscount";
            this.confirmTarget = item;
            this.confirmDialog = true;
        },
        async onConfirm_deliveryOrderDetailDialog(item) {
            this.deliveryOrderDetailDialog = false;
            if (this.actionType == "post") {
                await this.postDeliveryOrderDetail(item);
            } else if (this.actionType == "put") {
                await this.putDeliveryOrderDetail(item);
            }
        },
        async onConfirm_deliveryOrderDetailImportDialog(item) {
            this.deliveryOrderDetailImportDialog = false;
            await this.postDeliveryOrderDetails(item);
        },
        async onConfirm_discountDialog(item) {
            this.discountDialog = false;
            if (this.actionType == "post") {
                await this.postDiscount(item);
            } else if (this.actionType == "put") {
                await this.putDiscount(item);
            }
        },
        async onConfirm_confirmDialog(item) {
            this.confirmDialog = false;
            if (this.actionType == "deleteDetail") {
                await this.deleteDeliveryOrderDetail(item);
            } else if (this.actionType == "deleteDiscount") {
                await this.deleteDiscount(item);
            } else if (this.actionType == "finish") {
                this.beforeDeliveryOrderFinish();
                this.$emit("finish", this.deliveryOrderItem); //觸發一個在子元件中宣告的事件 childEvnet
            }
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
            let grossMargin =
                ((parseNumber(retailPrice) - parseNumber(cost)) /
                    parseNumber(retailPrice)) *
                100;
            if (isNaN(grossMargin) == false) {
                result += grossMargin;
            }
            return result.toFixed(2);
        },
        beforeDeliveryOrderFinish() {
            this.deliveryOrderItem.QTY = this.calc_TotalQTY;
            this.deliveryOrderItem.TotalOriginal = this.calc_Subtotals;
            this.deliveryOrderItem.Discount = this.calc_Discount;
            this.deliveryOrderItem.TotalDiscounted = this.calc_Total;
            this.deliveryOrderItem.Status = 2;
        },
        preSend(item) {
            item.ID = parseFloat(item.ID);
            item.RetailPrice = parseFloat(item.RetailPrice);
            item.QTY = parseFloat(item.QTY);
            item.Subtotal = parseFloat(item.Subtotal);
            item.DataOrder = parseFloat(item.DataOrder);
            item.DeliveryOrderID = parseFloat(item.DeliveryOrderID);
            item.ProductID = parseFloat(item.ProductID);
            return item;
        },
        async getDeliveryOrderDetails() {
            let filter = {
                DeliveryOrderID: this.deliveryOrderItem.ID,
            };
            await getDeliveryOrderDetails(filter)
                .then((response) => {
                    if (response.data.records != null) {
                        this.deliveryOrderDetails = response.data.records;
                    } else {
                        this.deliveryOrderDetails = [];
                    }
                })
                .catch((error) => { });
        },
        async postDeliveryOrderDetail(item) {
            item = this.preSend(item);
            item.DeliveryOrderID = this.deliveryOrderItem.ID;
            await postDeliveryOrderDetail(item)
                .then(async (response) => {
                    await this.getDeliveryOrderDetails();
                    this.alert = true;
                    this.alertType = "Success";
                    this.alertText = "新增出貨明細成功";
                })
                .catch((error) => {
                    this.alert = true;
                    this.alertType = "Fail";
                    this.alertText = "新增出貨明細失敗";
                });
        },
        async putDeliveryOrderDetail(item) {
            item = this.preSend(item);
            await putDeliveryOrderDetail(item)
                .then(async (response) => {
                    await this.getDeliveryOrderDetails();
                    this.alert = true;
                    this.alertType = "Success";
                    this.alertText = "編輯出貨明細成功";
                })
                .catch((error) => {
                    this.alert = true;
                    this.alertType = "Fail";
                    this.alertText = "編輯出貨明細失敗";
                });
        },
        async deleteDeliveryOrderDetail(item) {
            await deleteDeliveryOrderDetail(item)
                .then(async (response) => {
                    await this.getDeliveryOrderDetails();
                    this.alert = true;
                    this.alertType = "Success";
                    this.alertText = "刪除出貨明細成功";
                })
                .catch((error) => {
                    this.alert = true;
                    this.alertType = "Fail";
                    this.alertText = "刪除出貨明細失敗";
                });
        },
        async postDeliveryOrderDetails(item) {
            item.map((x) => {
                x.DeliveryOrderID = this.deliveryOrderItem.ID;
                x = this.preSend(x);
            });
            await postDeliveryOrderDetails(item)
                .then(async (response) => {
                    await this.getDeliveryOrderDetails();
                    this.alert = true;
                    this.alertType = "Success";
                    this.alertText = "新增出貨明細成功";
                })
                .catch((error) => {
                    this.alert = true;
                    this.alertType = "Fail";
                    this.alertText = "新增出貨明細失敗";
                });
        },
        preSendDiscounts(item) {
            item.ID = parseFloat(item.ID);
            item.Price = parseFloat(item.Price);
            item.DiscountType = parseFloat(item.DiscountType);
            item.DataOrder = parseFloat(item.DataOrder);
            item.DeliveryOrderID = parseFloat(item.DeliveryOrderID);
            return item;
        },
        async getDiscounts() {
            let filter = {
                DeliveryOrderID: this.deliveryOrderItem.ID,
            };
            await getDiscounts(filter)
                .then((response) => {
                    if (response.data.records != null) {
                        this.discounts = response.data.records;
                    } else {
                        this.discounts = [];
                    }
                })
                .catch((error) => { });
        },
        async postDiscount(item) {
            item = this.preSendDiscounts(item);
            item.DeliveryOrderID = this.deliveryOrderItem.ID;
            await postDiscount(item)
                .then(async (response) => {
                    await this.getDiscounts();
                    this.alert = true;
                    this.alertType = "Success";
                    this.alertText = "新增折扣成功";
                })
                .catch((error) => {
                    this.alert = true;
                    this.alertType = "Fail";
                    this.alertText = "新增折扣失敗";
                });
        },
        async putDiscount(item) {
            item = this.preSendDiscounts(item);
            await putDiscount(item)
                .then(async (response) => {
                    await this.getDiscounts();
                    this.alert = true;
                    this.alertType = "Success";
                    this.alertText = "編輯折扣成功";
                })
                .catch((error) => {
                    this.alert = true;
                    this.alertType = "Fail";
                    this.alertText = "編輯折扣失敗";
                });
        },
        async deleteDiscount(item) {
            await deleteDiscount(item)
                .then(async (response) => {
                    await this.getDiscounts();
                    this.alert = true;
                    this.alertType = "Success";
                    this.alertText = "刪除折扣成功";
                })
                .catch((error) => {
                    this.alert = true;
                    this.alertType = "Fail";
                    this.alertText = "刪除折扣失敗";
                });
        },
    },
    watch: {
        deliveryOrderInfoDialog: async function (newVal, oldVal) {
            if (newVal == true) {
                this.getDeliveryOrderDetails();
                this.getDiscounts();
            }
        },
    },
};
const parseNumber = function (x) {
    let parsed = parseFloat(x);
    if (isNaN(parsed) == true) {
        return 0;
    }
    return parsed;
};
</script>
