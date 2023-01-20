<template>
    <div>
        <v-container fluid class="d-none d-lg-block">
            <v-row class="ma-0">
                <v-col class="pa-0 d-flex justify-end">
                    <c-btn-add @click="onClick_newButton"></c-btn-add>
                    <c-btn-download @click="onClick_download"></c-btn-download>
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
                    <c-data-table :prop_headers="isSmalldevice ? deliveryOrderHeaderLess : deliveryOrderHeader"
                        :prop_items="deliveryOrders" :prop_search="search" @edit="onClick_editButton"
                        @delete="onClick_deleteButton">
                        <template v-slot:item.Status="{ item }">
                            <v-chip :color="getStatusChipColor(item.Status)" dark small>
                                {{ convertDisplayText(systemConfigs.DeliveryOrderStatus, item.Status) }}
                            </v-chip>
                        </template>
                        <template v-slot:item.DeliveryStatus="{ item }">
                            <v-chip :color="getDeliveryStatusChipColor(item.DeliveryStatus)" dark small>
                                {{ convertDisplayText(systemConfigs.DeliveryStatus, item.DeliveryStatus) }}
                            </v-chip>
                        </template>
                        <template v-slot:item.DeliveryType="{ item }">
                            <span>{{
                                convertDisplayText(systemConfigs.DeliveryType, item.DeliveryType)
                            }}</span>
                        </template>
                        <template v-slot:item.PaymentStatus="{ item }">
                            <span>{{
                                convertDisplayText(systemConfigs.PaymentStatus, item.PaymentStatus)
                            }}</span>
                        </template>
                        <template v-slot:item.OrderAt="{ item }">
                            <span>{{ item.OrderAt.substring(0, 10) }}</span>
                        </template>
                        <template v-slot:item.SendAt="{ item }">
                            <span>{{ item.SendAt.substring(0, 10) }}</span>
                        </template>
                        <template v-slot:item.ArriveAt="{ item }">
                            <span>{{ item.ArriveAt.substring(0, 10) }}</span>
                        </template>
                        <template v-slot:item.actions.plus="{ item }">
                            <v-icon small class="mx-1" @click.stop="onClick_checkoutDeliveryOrderInfo(item)">
                                mdi-arrow-right-circle
                            </v-icon>
                        </template>
                    </c-data-table>
                </v-col>
            </v-row>
        </v-container>
        <DeliveryOrderDialog :prop_deliveryOrderDialog.sync="deliveryOrderDialog" :prop_text_cardTitle="text_cardTitle"
            :prop_text_confirmBtn="text_confirmBtn" :prop_actionType="actionType"
            :prop_deliveryOrderItem="deliveryOrder" @confirm='onConfirm_deliveryOrderDialog' />
        <DeliveryOrderInfoDialog :prop_deliveryOrderInfoDialog.sync="deliveryOrderInfoDialog"
            :prop_text_cardTitle="text_cardTitle" :prop_text_confirmBtn="text_confirmBtn"
            :prop_deliveryOrderItem="deliveryOrder" @finish='onFinish_deliveryOrderInfoDialog' />
        <ConfirmDialog :prop_confirmDialog.sync="confirmDialog" :prop_text_cardTitle="text_cardTitle"
            :prop_text_cardHint="text_cardHint" :prop_text_confirmBtn="text_confirmBtn"
            :prop_confirmTarget.sync="confirmTarget" v-on:confirmClick='onConfirm_confirmDialog' />
        <Alert :prop_alert.sync="alert" :prop_alertType="alertType" :prop_alertText="alertText"></Alert>
    </div>

</template>

<script>
import key from 'keymaster'
import Alert from '@/components/Alert/index.vue'
import ConfirmDialog from '@/components/ConfirmDialog/index.vue'
import BtnAdd from "@/components/Buttons/BtnAdd.vue";
import BtnDownload from "@/components/Buttons/BtnDownload.vue";
import BtnSetting from "@/components/Buttons/BtnSetting.vue";
import DataTable from "@/components/DataTables/DataTable.vue";
import DeliveryOrderDialog from '@/pages/DeliveryOrders/DeliveryOrderDialog.vue'
import DeliveryOrderInfoDialog from '@/pages/DeliveryOrders/DeliveryOrderInfoDialog.vue'
import {
    getDeliveryOrders,
    postDeliveryOrder,
    putDeliveryOrder,
    deleteDeliveryOrder,
} from "@/apis/DeliveryOrdersAPI";


class DeliveryOrder {
    ID = undefined;//流水號
    DeliveryType = undefined;//出貨方式
    DeliveryStatus = undefined;//出貨狀態
    DeliveryFeeStatus = undefined;//運費狀態
    PaymentType = undefined;//付款方式
    PaymentStatus = undefined;//付款狀態
    TotalOriginal = undefined;//原價
    Discount = undefined;//折扣
    TotalDiscounted = undefined;//總價
    Remark = "";//備註
    DataOrder = undefined;//順序
    OrderAt = "";//下訂日期
    SendAt = "";//出貨日期
    ArriveAt = "";//送達日期
    CreateAt = "";//建立時間
    UpdateAt = "";//最後編輯時間
}

export default {
    name: 'DeliveryOrders',
    components: {
        Alert,
        ConfirmDialog,
        DeliveryOrderDialog,
        DeliveryOrderInfoDialog,
        "c-btn-add": BtnAdd,
        "c-btn-download": BtnDownload,
        "c-btn-setting": BtnSetting,
        "c-data-table": DataTable,
    },
    data() {
        return {
            search: '',
            text_cardTitle: "新增",
            text_cardHint: "",
            text_confirmBtn: "新增",

            //Alert
            alert: false,
            alertType: "",
            alertText: "",
            alertTimeoutID: null,

            confirmDialog: false,
            confirmTarget: null,
            actionType: "",

            deliveryOrder: new DeliveryOrder(),
            deliveryOrderDialog: false,
            deliveryOrders: [],
            deliveryOrderHeader: [
                { text: '流水號', value: 'ID' },
                { text: '出貨單狀態', value: 'Status', align: 'center' },
                { text: '出貨狀態', value: 'DeliveryStatus', align: 'center' },
                { text: '出貨方式', value: 'DeliveryType' },
                { text: '付款狀態', value: 'PaymentStatus' },
                { text: '原價', value: 'TotalOriginal' },
                { text: '折扣', value: 'Discount' },
                { text: '總價', value: 'TotalDiscounted' },
                { text: '下訂日期', value: 'OrderAt' },
                { text: '出貨日期', value: 'SendAt' },
                { text: '送達日期', value: 'ArriveAt' },
                { text: '備註', value: 'Remark' },
                { text: '', value: 'actions', sortable: false, width: "10%" },
            ],
            deliveryOrderHeaderLess: [
                { text: '流水號', value: 'ID' },
                { text: '出貨單狀態', value: 'Status', align: 'center' },
                { text: '原價', value: 'TotalOriginal' },
                { text: '折扣', value: 'Discount' },
                { text: '總價', value: 'TotalDiscounted' },
                { text: '', value: 'actions', sortable: false, width: "10%" },
            ],
            deliveryOrderInfoDialog: false,
        };
    },
    async mounted() {
        key('command+/', this.onFocus_searchFields);
        key('ctrl+/', this.onFocus_searchFields);
        await this.getDeliveryOrders();
    },
    props: {
    },
    computed: {
        systemConfigs() {
            return this.$store.state.systemConfigs;
        },
        isSmalldevice() {
            if (this.$vuetify.breakpoint.name == "xs") {
                return true;
            }
            return false;
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
        getStatusChipColor(status) {
            if (status == 1) return '#849A8F'
            // #97A48D
            else if (status == 2) return '#7E7E7E'
            else return ''
        },
        getDeliveryStatusChipColor(status) {
            if (status == 1) return '#965455'
            else if (status == 2) return '#849A8F'
            else return ''
        },
        onFocus_searchFields() {
            this.$refs.searchField.focus();
        },
        onClick_download() {

        },
        onClick_newButton() {
            this.text_cardTitle = "新增出貨單";
            this.text_confirmBtn = "新增";
            this.actionType = "post";
            this.deliveryOrder = new DeliveryOrder();
            this.deliveryOrderDialog = true;
        },
        onClick_editButton(item) {
            this.text_cardTitle = "編輯出貨單";
            this.text_confirmBtn = "修改";
            this.actionType = "put";
            this.deliveryOrder = item;
            this.deliveryOrderDialog = true;
        },
        onClick_deleteButton(item) {
            this.text_cardTitle = "確認刪除";
            this.text_confirmBtn = "刪除";
            this.confirmDialog = true;
            this.confirmTarget = item;
        },
        onClick_checkoutDeliveryOrderInfo(item) {
            this.text_cardTitle = "出貨單" + item.ID;
            this.text_confirmBtn = "";
            this.deliveryOrder = item;
            this.deliveryOrderInfoDialog = true;
        },
        async onConfirm_deliveryOrderDialog(item) {
            this.deliveryOrderDialog = false;
            if (this.actionType == "post") {
                await this.postDeliveryOrder(item);
            } else if (this.actionType == "put") {
                await this.putDeliveryOrder(item);
            }
        },
        async onConfirm_confirmDialog(item) {
            this.confirmDialog = false;
            await this.deleteDeliveryOrder(item);
        },
        async onFinish_deliveryOrderInfoDialog(item) {
            this.deliveryOrderInfoDialog = false;
            await this.putDeliveryOrder(item);
        },
        preSend(item) {
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
        async getDeliveryOrders() {
            await getDeliveryOrders()
                .then((response) => {
                    if (response.data.records != null) {
                        this.deliveryOrders = response.data.records;
                    }
                    else {
                        this.deliveryOrders = [];
                    }
                })
                .catch((error) => {
                });
        },
        async postDeliveryOrder(item) {
            item = this.preSend(item);
            item.OrderAt = null;
            item.SendAt = null;
            item.ArriveAt = null;
            item.Status = 1;
            await postDeliveryOrder(item)
                .then(async (response) => {
                    await this.getDeliveryOrders();
                    this.alert = true;
                    this.alertType = "Success";
                    this.alertText = "新增出貨單成功";
                })
                .catch((error) => {
                    this.alert = true;
                    this.alertType = "Fail";
                    this.alertText = "新增出貨單失敗";
                });
        },
        async putDeliveryOrder(item) {
            item = this.preSend(item);
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
        async deleteDeliveryOrder(item) {
            await deleteDeliveryOrder(item)
                .then(async (response) => {
                    await this.getDeliveryOrders();
                    this.alert = true;
                    this.alertType = "Success";
                    this.alertText = "刪除出貨單成功";
                })
                .catch((error) => {
                    if (error.response.data.role == "model" && error.response.data.code == 1) {
                        this.alert = true;
                        this.alertType = "Fail";
                        this.alertText = "此出貨單尚包含商品明細，請先移除相關商品明細後再移除此出貨單";
                    }
                    else {
                        this.alert = true;
                        this.alertType = "Fail";
                        this.alertText = "刪除出貨單失敗";
                    }
                });
        },
    },
}
</script>