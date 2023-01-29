<template>
    <div>
        <v-container fluid class="d-none d-lg-block">
            <v-row class="ma-0">
                <v-col class="pa-0 d-flex justify-end">
                    <c-btn-add @click="onClick_newButton"></c-btn-add>
                    <c-btn-download @click="onClick_download"></c-btn-download>
                    <c-btn-setting></c-btn-setting>
                    <c-btn-logout></c-btn-logout>
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
                    <c-data-table :prop_headers="isSmalldevice ? purchaseHeaderLess : purchaseHeader"
                        :prop_items="purchases" :prop_search="search" @edit="onClick_editButton"
                        @delete="onClick_deleteButton">
                        <template v-slot:item.Status="{ item }">
                            <v-chip :color="getStatusChipColor(item.Status)" dark small>
                                {{ convertDisplayText(systemConfigs.PurchaseStatus, item.Status) }}
                            </v-chip>
                        </template>
                        <template v-slot:item.PurchaseType="{ item }">
                            <span>{{ convertDisplayText(systemConfigs.PurchaseType, item.PurchaseType) }}</span>
                        </template>
                        <template v-slot:item.actions.plus="{ item }">
                            <v-icon small class="mx-1" @click.stop="onClick_checkoutPurchaseInfo(item)">
                                mdi-arrow-right-circle
                            </v-icon>
                        </template>
                    </c-data-table>
                </v-col>
            </v-row>
        </v-container>
        <PurchaseDialog :prop_purchaseDialog.sync="purchaseDialog" :prop_text_cardTitle="text_cardTitle"
            :prop_text_confirmBtn="text_confirmBtn" :prop_actionType="actionType" :prop_purchaseItem="purchase"
            @confirm='onConfirm_purchaseDialog' />
        <PurchaseInfoDialog :prop_purchaseInfoDialog.sync="purchaseInfoDialog" :prop_text_cardTitle="text_cardTitle"
            :prop_text_confirmBtn="text_confirmBtn" :prop_purchaseItem="purchase"
            @finish='onFinish_purchaseInfoDialog' />
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
import BtnLogout from "@/components/Buttons/BtnLogout.vue";
import DataTable from "@/components/DataTables/DataTable.vue";
import PurchaseDialog from '@/pages/Purchases/PurchaseDialog.vue'
import PurchaseInfoDialog from '@/pages/Purchases/PurchaseInfoDialog.vue'
import {
    getPurchases,
    postPurchase,
    putPurchase,
    deletePurchase,
} from "@/apis/PurchasesAPI";


class Purchase {
    ID = undefined;
    Name = "";
    Status = undefined;
    PurchaseType = undefined;
    QTY = undefined;
    ShippingAgent = "";
    ShippingAgentCutKrw = undefined;
    ShippingAgentCutPercent = undefined;
    ShippingInitiator = "";
    ShippingCreateAt = "";
    ShippingEndAt = "";
    ShippingArriveAt = "";
    Weight = undefined;
    ShippingFeeKr = undefined;
    ShippingFeeTw = undefined;
    ShippingFeeKokusaiKrw = undefined;
    ShippingFeeKokusaiPerKilo = undefined;
    ExchangeRateKrw = undefined;
    TariffTwd = undefined;
    TariffPerKilo = undefined;
    TotalKrw = undefined;
    TotalTwd = undefined;
    Total = undefined;
    Remark = "";
    DataOrder = undefined;
    CreateAt = "";
    UpdateAt = "";
}

export default {
    name: 'Purchases',
    components: {
        Alert,
        ConfirmDialog,
        PurchaseDialog,
        PurchaseInfoDialog,
        "c-btn-add": BtnAdd,
        "c-btn-download": BtnDownload,
        "c-btn-setting": BtnSetting,
        "c-btn-logout": BtnLogout,
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

            purchase: new Purchase(),
            purchaseDialog: false,
            purchases: [],
            purchaseHeader: [
                { text: '採購狀態', value: 'Status', align: 'center', width: "10%" },
                { text: '採購名稱', value: 'Name' },
                { text: '採購種類', value: 'PurchaseType' },
                { text: '韓圓總計', value: 'TotalKrw' },
                { text: '總計', value: 'Total' },
                { text: "備註", value: "Remark", width: "10%" },
                { text: '', value: 'actions', sortable: false, width: "10%" },
            ],
            purchaseHeaderLess: [
                { text: '採購狀態', value: 'Status', align: 'center' },
                { text: '採購名稱', value: 'Name' },
                { text: '', value: 'actions', sortable: false, width: "10%" },
            ],
            purchaseInfoDialog: false,
        };
    },
    async mounted() {
        key('command+/', this.onFocus_searchFields);
        key('ctrl+/', this.onFocus_searchFields);
        await this.getPurchases();
    },
    props: {
    },
    computed: {
        systemConfigs() {
            return this.$store.state.conf.systemConfigs;
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
        onFocus_searchFields() {
            this.$refs.searchField.focus();
        },
        onClick_download() {

        },
        onClick_newButton() {
            this.text_cardTitle = "新增採購案";
            this.text_confirmBtn = "新增";
            this.actionType = "post";
            this.purchase = new Purchase();
            this.purchaseDialog = true;
        },
        onClick_editButton(item) {
            this.text_cardTitle = "編輯採購案";
            this.text_confirmBtn = "修改";
            this.actionType = "put";
            this.purchase = item;
            this.purchaseDialog = true;
        },
        onClick_deleteButton(item) {
            this.text_cardTitle = "確認刪除";
            this.text_confirmBtn = "刪除";
            this.confirmDialog = true;
            this.confirmTarget = item;
        },
        onClick_checkoutPurchaseInfo(item) {
            this.text_cardTitle = "採購案" + item.Name;
            this.text_confirmBtn = "";
            this.purchase = item;
            this.purchaseInfoDialog = true;
        },
        async onConfirm_purchaseDialog(item) {
            this.purchaseDialog = false;
            if (this.actionType == "post") {
                await this.postPurchase(item);
            } else if (this.actionType == "put") {
                await this.putPurchase(item);
            }
        },
        async onConfirm_confirmDialog(item) {
            this.confirmDialog = false;
            await this.deletePurchase(item);
        },
        async onFinish_purchaseInfoDialog(item) {
            this.purchaseInfoDialog = false;
            await this.putPurchase(item);
        },
        preSend(item) {
            item.ID = parseFloat(item.ID);
            item.Status = parseFloat(item.Status);
            item.PurchaseType = parseFloat(item.PurchaseType);
            item.QTY = parseFloat(item.QTY);
            item.ShippingAgentCutKrw = parseFloat(item.ShippingAgentCutKrw);
            item.ShippingAgentCutPercent = parseFloat(item.ShippingAgentCutPercent);
            item.Weight = parseFloat(item.Weight);
            item.ShippingFeeKr = parseFloat(item.ShippingFeeKr);
            item.ShippingFeeTw = parseFloat(item.ShippingFeeTw);
            item.ShippingFeeKokusaiKrw = parseFloat(item.ShippingFeeKokusaiKrw);
            item.ShippingFeeKokusaiPerKilo = parseFloat(item.ShippingFeeKokusaiPerKilo);
            item.ExchangeRateKrw = parseFloat(item.ExchangeRateKrw);
            item.TariffTwd = parseFloat(item.TariffTwd);
            item.TariffPerKilo = parseFloat(item.TariffPerKilo);
            item.TotalKrw = parseFloat(item.TotalKrw);
            item.TotalTwd = parseFloat(item.TotalTwd);
            item.Total = parseFloat(item.Total);
            item.DataOrder = parseFloat(item.DataOrder);
            return item;
        },
        async getPurchases() {
            await getPurchases()
                .then((response) => {
                    if (response.data.records != null) {
                        this.purchases = response.data.records;
                    }
                    else {
                        this.purchases = [];
                    }
                })
                .catch((error) => {
                });
        },
        async postPurchase(item) {
            item = this.preSend(item);
            item.ShippingArriveAt = item.ShippingArriveAt || null;
            item.ShippingCreateAt = item.ShippingCreateAt || null;
            item.ShippingEndAt = item.ShippingEndAt || null;
            await postPurchase(item)
                .then(async (response) => {
                    await this.getPurchases();
                    this.alert = true;
                    this.alertType = "Success";
                    this.alertText = "新增採購案成功";
                })
                .catch((error) => {
                    this.alert = true;
                    this.alertType = "Fail";
                    this.alertText = "新增採購案失敗";
                });
        },
        async putPurchase(item) {
            item = this.preSend(item);
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
        async deletePurchase(item) {
            await deletePurchase(item)
                .then(async (response) => {
                    await this.getPurchases();
                    this.alert = true;
                    this.alertType = "Success";
                    this.alertText = "刪除採購案成功";
                })
                .catch((error) => {
                    if (error.response.data.role == "model" && error.response.data.code == 1) {
                        this.alert = true;
                        this.alertType = "Fail";
                        this.alertText = "此採購案尚包含商品明細，請先移除相關商品明細後再移除此採購案";
                    }
                    else {
                        this.alert = true;
                        this.alertType = "Fail";
                        this.alertText = "刪除採購案失敗";
                    }
                });
        },
    },
}
</script>