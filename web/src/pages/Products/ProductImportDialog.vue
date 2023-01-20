<template>
    <v-dialog v-model="productImportDialog" hide-overlay @click:outside="onClick_cancel" fullscreen
        transition="slide-x-reverse-transition">
        <v-card class="d-flex align-center flex-column justify-center">
            <v-btn icon dark absolute top left @click="productImportDialog = false">
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
                        <div class="d-flex flex-row justify-center mb-4">
                            <v-btn v-for="worksheet in worksheets" class="mx-1"
                                @click="onSelect_worksheet(worksheet)">{{ worksheet }}</v-btn>
                        </div>
                    </v-col>
                </v-row>
                <v-row class="d-flex justify-center">
                    <v-col cols="10">
                        <v-card outlined v-if="items.length > 0">
                            <c-data-table :prop_headers="headers" :prop_items.sync="items">
                                <template v-slot:item.ProductType="props">
                                    <v-edit-dialog :return-value.sync="props.item.ProductType">
                                        {{ convertDisplayText(systemConfigs.ProductType, props.item.ProductType) }}
                                        <template v-slot:input>
                                            <v-select label="商品種類" v-model="props.item.ProductType"
                                                prepend-icon="mdi-tournament" :items="systemConfigs.ProductType"
                                                item-text="value" item-value="key"></v-select>
                                        </template>
                                    </v-edit-dialog>
                                </template>
                                <template v-slot:item.SupplierID="props">
                                    <v-edit-dialog :return-value.sync="props.item.SupplierID">
                                        {{ convertDisplayText(allSuppliersList, props.item.SupplierID) }}
                                        <template v-slot:input>
                                            <v-select label="商品種類" v-model="props.item.SupplierID"
                                                prepend-icon="mdi-tournament" :items="systemConfigs.ProductType"
                                                item-text="value" item-value="key"></v-select>
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
    name: 'ProductImportDialog',
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
                { text: '廠商名稱', value: 'SupplierID' },
                { text: 'SKU', value: 'SKU' },
                { text: '商品名稱', value: 'Name' },
                { text: '商品種類', value: 'ProductType' },
                { text: '庫存(個)', value: 'Stocks' },
                { text: '售價', value: 'RetailPrice' },
                { text: '重量(g)', value: 'Weight' },
            ],
        };
    },
    props: {
        prop_productImportDialog: {
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
        productImportDialog: {
            get() {
                return this.prop_productImportDialog
            },
            set(val) {
                this.$emit('update:prop_productImportDialog', val)

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
        ProductTypeMaps() {
            let map = {};
            this.$store.state.systemConfigs.ProductType.map(function (x) {
                map[x.value] = x.key;
            })
            return map;
        },
        allSuppliersList() {
            return this.$store.state.allSuppliers;
        },
        SupplierMaps() {
            let map = {};
            this.$store.state.allSuppliers.map(function (x) {
                map[x.value] = x.key;
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
        onClick_confirm: function () { //有子元件的事件觸發 自定義事件childevent
            this.$emit('confirm', this.items);//觸發一個在子元件中宣告的事件 childEvnet
        },
        onClick_cancel() {
            this.productImportDialog = false;
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
                                SupplierID: vthis.SupplierMaps[row.getCell(headers["廠商名稱"]).value],
                                SKU: row.getCell(headers["SKU"]).value,
                                Name: row.getCell(headers["商品名稱"]).value,
                                ProductType: vthis.ProductTypeMaps[row.getCell(headers["商品種類"]).value],
                                Stocks: row.getCell(headers["庫存"]).value,
                                RetailPrice: row.getCell(headers["售價"]).value,
                                Weight: row.getCell(headers["重量"]).value,
                                DataStatus: 1,
                            }
                            dataAry.push(obj)
                        }
                    });
                    vthis.items = dataAry;
                })
                .catch(function () {
                    vthis.alert = true;
                    vthis.alertType = "Fail";
                    vthis.alertText = "格式錯誤，請確認資料表包含「廠商名稱」、「SKU」、「商品名稱」、「商品種類」、「庫存」、「售價」及「重量」欄位";
                });
        },
    },
    watch: {
    }
}
</script>