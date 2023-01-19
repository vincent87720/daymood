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
                    <c-data-table :prop_headers="isSmalldevice ? supplierHeaderLess : supplierHeader" :prop_items="suppliers" :prop_search="search"
                        @edit="onClick_editButton" @delete="onClick_deleteButton"></c-data-table>
                </v-col>
            </v-row>
        </v-container>
        <SupplierDialog :prop_supplierDialog.sync="supplierDialog" :prop_text_cardTitle="text_cardTitle"
            :prop_text_confirmBtn="text_confirmBtn" :prop_supplierItem="supplier" @confirm="onConfirm_supplierDialog" />
        <ConfirmDialog :prop_confirmDialog.sync="confirmDialog" :prop_text_cardTitle="text_cardTitle"
            :prop_text_cardHint="text_cardHint" :prop_text_confirmBtn="text_confirmBtn"
            :prop_confirmTarget.sync="confirmTarget" v-on:confirmClick="onConfirm_confirmDialog">
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
import key from "keymaster";
import SupplierDialog from "../../components/SupplierDialog/index.vue";
import Alert from "../../components/Alert/index.vue";
import ConfirmDialog from "../../components/ConfirmDialog/index.vue";
import BtnAdd from "../../components/Buttons/BtnAdd.vue";
import BtnDownload from "../../components/Buttons/BtnDownload.vue";
import BtnSetting from "../../components/Buttons/BtnSetting.vue";
import DataTable from "../../components/DataTables/DataTable.vue";
import {
    getSuppliers,
    postSupplier,
    putSupplier,
    deleteSupplier,
} from "../../apis/SuppliersAPI";

class Supplier {
    ID = undefined;
    Name = "";
    Address = "";
    Remark = "";
    DataOrder = undefined;
    CreateAt = "";
    UpdateAt = "";
}

const FileDownload = require("js-file-download");

export default {
    name: "Suppliers",
    components: {
        SupplierDialog,
        Alert,
        ConfirmDialog,
        "c-btn-add": BtnAdd,
        "c-btn-download": BtnDownload,
        "c-btn-setting": BtnSetting,
        "c-data-table": DataTable,
    },
    data() {
        return {
            search: "",
            text_cardTitle: "新增",
            text_cardHint: "",
            text_confirmBtn: "新增",

            supplier: new Supplier(),
            supplierDialog: false,
            supplierHeader: [
                { text: "廠商名稱", value: "Name" },
                { text: "廠商地址", value: "Address" },
                { text: "備註", value: "Remark" },
                { text: "", value: "actions", sortable: false, width: "10%" },
            ],
            supplierHeaderLess: [
                { text: "廠商名稱", value: "Name" },
                { text: "", value: "actions", sortable: false, width: "10%" },
            ],
            suppliers: [],

            //Alert
            alert: false,
            alertType: "",
            alertText: "",

            confirmDialog: false,
            confirmTarget: null,
        };
    },
    async mounted() {
        key("command+/", this.onFocus_searchFields);
        key("ctrl+/", this.onFocus_searchFields);
        await this.getSuppliers();
        // this.dumpFirms();
    },
    computed: {
        isSmalldevice() {
            if(this.$vuetify.breakpoint.name == "xs"){
                return true;
            }
            return false;
        },
    },
    methods: {
        onFocus_searchFields() {
            // this.$refs.searchField.focus();
        },
        onClick_newButton() {
            this.text_cardTitle = "新增廠商";
            this.text_confirmBtn = "新增";
            this.actionType = "post";
            this.supplier = new Supplier();
            this.supplierDialog = true;
        },
        onClick_editButton(item) {
            this.text_cardTitle = "編輯廠商";
            this.text_confirmBtn = "修改";
            this.actionType = "put";
            this.supplier = item;
            this.supplierDialog = true;
        },
        onClick_deleteButton(item) {
            this.text_cardTitle = "確認刪除";
            this.text_confirmBtn = "刪除";
            this.confirmDialog = true;
            this.confirmTarget = item;
        },
        onClick_download() {
            // this.dumpFirms();
        },
        async onClick_changeDataStatus(item) {
            this.confirmDialog = false;
            item.DataStatus = 0;
            await this.putSupplier(item);
        },
        onConfirm_supplierDialog(val) {
            this.supplierDialog = false;
            if (this.actionType == "post") {
                this.postSupplier(val);
            } else if (this.actionType == "put") {
                this.putSupplier(val);
            }
        },
        onConfirm_confirmDialog(item) {
            this.confirmDialog = false;
            this.deleteSupplier(item);
        },
        filterEnableData(item) {
            return item.filter(x => x.DataStatus == 1);
        },
        async getSuppliers() {
            await getSuppliers()
                .then((response) => {
                    if (response.data.records != null) {
                        this.suppliers = this.filterEnableData(response.data.records);
                    }
                    else {
                        this.suppliers = [];
                    }
                })
                .catch((error) => {
                    // console.log(error.response);
                });
        },
        async postSupplier(item) {
            await postSupplier(item)
                .then(async (response) => {
                    await this.getSuppliers();
                    this.alert = true;
                    this.alertType = "Success";
                    this.alertText = "新增廠商成功";
                })
                .catch((error) => {
                    // console.log(error.response);
                    this.alert = true;
                    this.alertType = "Fail";
                    this.alertText = "新增廠商失敗";
                });
        },
        async putSupplier(item) {
            await putSupplier(item)
                .then(async (response) => {
                    await this.getSuppliers();
                    this.alert = true;
                    this.alertType = "Success";
                    this.alertText = "編輯廠商成功";
                })
                .catch((error) => {
                    // console.log(error.response);
                    this.alert = true;
                    this.alertType = "Fail";
                    this.alertText = "編輯廠商失敗";
                });
        },
        async deleteSupplier(item) {
            await deleteSupplier(item)
                .then(async (response) => {
                    await this.getSuppliers();
                    this.alert = true;
                    this.alertType = "Success";
                    this.alertText = "刪除廠商成功";
                })
                .catch((error) => {
                    if (error.response.data.role == "model" && error.response.data.code == 1) {
                        this.alert = true;
                        this.alertType = "Fail";
                        this.alertText = "尚有商品屬於此廠商，請先移除相關商品後再移除此廠商";
                    }
                });
        },
        // dumpFirms() {
        //     let vthis = this;

        //     this.axios({
        //         url: process.env.VUE_APP_BASE_URL + "firm/dumping",
        //         method: "GET",
        //         responseType: "blob", // Important
        //     })
        //         .then(function (response) {
        //             if (response.data != null) {
        //                 FileDownload(response.data, "firms.csv");
        //             }
        //             // console.log(response.data);
        //         })
        //         .catch((error) => {
        //             // console.log(error.response);
        //         });
        // },
    },
};
</script>
