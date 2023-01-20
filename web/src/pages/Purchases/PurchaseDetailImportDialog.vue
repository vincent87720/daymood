<template>
    <v-dialog v-model="purchaseDetailImportDialog" hide-overlay @click:outside="onClick_cancel" fullscreen
        transition="slide-x-reverse-transition">
        <v-card class="d-flex align-center flex-column justify-center">
            <v-btn icon dark absolute top left @click="purchaseDetailImportDialog = false">
                <v-icon>mdi-arrow-left</v-icon>
            </v-btn>
            <v-card-title class="text-h5">
                {{ text_cardTitle }}
            </v-card-title>
            <v-container class="pa-6 mb-8">
                <v-row class="d-flex justify-center">
                    <v-col cols="6">
                        <v-file-input truncate-length="15" v-model="files" label="選擇檔案" outlined
                            prepend-icon="mdi-table-large" @change="onUpload"
                            @click:clear="onClick_clear"></v-file-input>
                    </v-col>
                </v-row>
                <v-row class="ma-0 d-flex justify-center ">
                    <v-col cols="6" class="d-flex justify-center ma-0">
                        <span class="d-flex overflow-x-auto">
                            <v-btn v-for="worksheet in worksheets" class="mx-1"
                                @click="onSelect_worksheet(worksheet)">{{ worksheet }}</v-btn>
                        </span>
                    </v-col>
                </v-row>
                <v-row class="d-flex justify-center">
                    <v-col cols="10">
                        <v-card outlined v-if="items.length > 0">
                            <c-data-table :prop_headers="headers" :prop_items.sync="items">
                                <template v-slot:item.ProductID="props">
                                    <v-edit-dialog :return-value.sync="props.item.ProductID">
                                        {{ convertDisplayText_SKU(allProductsList, props.item.ProductID) }}
                                        <template v-slot:input>
                                            <v-select label="商品" v-model="props.item.ProductID"
                                                prepend-icon="mdi-tournament" :items="allProductsList" item-text="SKU"
                                                item-value="key"></v-select>
                                        </template>
                                    </v-edit-dialog>
                                </template>
                                <template v-slot:item.SupplierID="props">
                                    <v-edit-dialog :return-value.sync="props.item.SupplierID">
                                        {{ convertDisplayText(allSuppliersList, props.item.SupplierID) }}
                                        <template v-slot:input>
                                            <v-select label="廠商名稱" v-model="props.item.SupplierID"
                                                prepend-icon="mdi-tournament" :items="allSuppliersList"
                                                item-text="value" item-value="key"></v-select>
                                        </template>
                                    </v-edit-dialog>
                                </template>
                                <template v-slot:item.Status="props">
                                    <v-edit-dialog :return-value.sync="props.item.Status">
                                        {{ convertDisplayText(systemConfigs.PurchaseDetailStatus, props.item.Status) }}
                                        <template v-slot:input>
                                            <v-select label="是否採用" v-model="props.item.Status"
                                                prepend-icon="mdi-tournament"
                                                :items="systemConfigs.PurchaseDetailStatus" item-text="value"
                                                item-value="key"></v-select>
                                        </template>
                                    </v-edit-dialog>
                                </template>
                            </c-data-table>
                        </v-card>
                    </v-col>
                </v-row>
            </v-container>
            <v-card-actions class="pa-6 pt-3">
                <v-spacer></v-spacer>
                <v-btn outlined rounded text @click.stop="onClick_cancel">
                    取消
                </v-btn>
                <v-btn outlined rounded text @click.stop="onClick_confirm">
                    {{ text_confirmBtn }}
                </v-btn>
            </v-card-actions>
        </v-card>
        <Alert :prop_alert.sync="alert" :prop_alertType="alertType" :prop_alertText="alertText"></Alert>
    </v-dialog>
</template>

<script>
import Alert from '@/components/Alert/index.vue'
import DatePicker from "@/components/Pickers/DatePicker.vue";
import DataTable from "@/components/DataTables/DataTable.vue";
const Excel = require('exceljs');

export default {
    name: 'PurchaseDetailImportDialog',
    components: {
        Alert,
        "c-date-picker": DatePicker,
        "c-data-table": DataTable,
    },
    data() {
        return {
            //Alert
            alert: false,
            alertType: "",
            alertText: "",

            files: undefined,
            worksheets: [],
            items: [],
            headers: [
                { text: '是否採用', value: 'Status' },
                { text: '廠商名稱', value: 'SupplierID' },
                { text: 'SKU', value: 'ProductID' },
                { text: '採購商品編號', value: 'NamedID' },
                { text: '商品名稱', value: 'Name' },
                { text: '批價', value: 'WholesalePrice' },
                { text: '數量', value: 'QTY' },
                { text: '小計', value: 'Subtotal' },
                { text: '備註', value: 'Remark' },
            ],
        };
    },
    props: {
        prop_purchaseDetailImportDialog: {
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
    },
    mounted() {
    },
    computed: {
        purchaseDetailImportDialog: {
            get() {
                return this.prop_purchaseDetailImportDialog
            },
            set(val) {
                this.$emit('update:prop_purchaseDetailImportDialog', val)

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
        systemConfigs() {
            return this.$store.state.systemConfigs;
        },
        allSuppliersList() {
            return this.$store.state.allSuppliers;
        },
        allProductsList() {
            return this.$store.state.allProducts;
        },
        PurchaseDetailStatusMaps() {
            let map = {};
            this.$store.state.systemConfigs.PurchaseDetailStatus.map(function (x) {
                map[x.value] = x.key;
            })
            return map;
        },
        SupplierMaps() {
            let map = {};
            this.$store.state.allSuppliers.map(function (x) {
                map[x.value] = x.key;
            })
            return map;
        },
        ProductMaps() {
            let map = {};
            this.$store.state.allProducts.map(function (x) {
                map[x.SKU] = x.key;
            })
            return map;
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
        convertDisplayText_SKU(list, key) {
            let result = list.find(x => x.key == key);
            if (result) {
                return result.SKU
            }
            return "";
        },
        onClick_confirm: function () {
            this.$emit('confirm', this.items);
        },
        onClick_cancel() {
            this.purchaseDetailImportDialog = false;
        },
        onClick_clear() {
            this.files = undefined;
            this.worksheets = [];
            this.items = [];
        },
        onUpload(item) {
            var workbook = new Excel.Workbook();
            let vthis = this;
            workbook.xlsx.load(item)
                .then(function () {
                    workbook.eachSheet(function (worksheet, sheetId) {
                        vthis.worksheets.push(worksheet.name);
                    });
                });
        },
        onSelect_worksheet(worksheet) {
            var workbook = new Excel.Workbook();
            let vworksheet = worksheet;
            let vthis = this;
            workbook.xlsx.load(this.files)
                .then(function () {
                    var worksheet = workbook.getWorksheet(vworksheet);
                    let getHeaders = function (index) {
                        let result = {}

                        let row = worksheet.getRow(index);

                        if (row === null || !row.values || !row.values.length) return [];

                        for (let i = 1; i < row.values.length; i++) {
                            let cell = row.getCell(i);
                            result[cell.text] = i;
                        }
                        return result;
                    }
                    let headers = getHeaders(1);
                    let dataAry = [];
                    worksheet.eachRow({ includeEmpty: true }, function (row, rowNumber) {
                        if (rowNumber > 1) {
                            let obj = {
                                Status: vthis.PurchaseDetailStatusMaps[row.getCell(headers["是否採用"]).value],
                                SupplierID: vthis.SupplierMaps[row.getCell(headers["廠商"]).value],
                                ProductID: vthis.ProductMaps[row.getCell(headers["SKU"]).value],
                                NamedID: row.getCell(headers["採購編號"]).value,
                                Name: row.getCell(headers["商品名稱"]).value,
                                WholesalePrice: row.getCell(headers["韓幣批價"]).value,
                                QTY: row.getCell(headers["數量"]).value,
                                Subtotal: row.getCell(headers["批價總額(WON)"]).value,
                                Remark: row.getCell(headers["備註"]).value,
                            }
                            dataAry.push(obj)
                        }
                    });
                    vthis.items = dataAry;
                })
                .catch(function () {
                    vthis.alert = true;
                    vthis.alertType = "Fail";
                    vthis.alertText = "格式錯誤，請確認資料表包含「廠商名稱」、「SKU」、「採購商品編號」、「商品名稱」、「批價」、「數量」、「小計」及「是否採用」欄位";
                });
        },
    },
    watch: {
    }
}
</script>