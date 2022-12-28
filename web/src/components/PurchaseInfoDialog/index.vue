<template>
    <v-dialog v-model="purchaseInfoDialog" hide-overlay @click:outside="onClick_cancel" fullscreen
        transition="slide-x-reverse-transition">
        <v-card class="d-flex align-center flex-column justify-center">
            <v-btn icon dark absolute top left @click="purchaseInfoDialog = false">
                <v-icon>mdi-arrow-left</v-icon>
            </v-btn>
            <v-btn icon dark absolute top right @click="onClick_newDetailButton">
                <v-icon>mdi-plus</v-icon>
            </v-btn>
            <v-card-title class="text-h5">
                {{ text_cardTitle }}
            </v-card-title>
            <v-container class="pa-6">
                <v-row>
                    <v-col xs="12" sm="12" class="ml-auto mr-auto">

                        <c-data-table :prop_headers="purchaseDetailHeader" :prop_items="purchaseDetails" :prop_search="search"
                            @edit="onClick_editButton" @delete="onClick_deleteButton">
                            <template v-slot:item.Status="{ item }">
                                <span>{{ convertDisplayText(systemConfigs.PurchaseStatus, item.Status) }}</span>
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
        </v-card>
    </v-dialog>
</template>

<script>
import DatePicker from "../../components/Pickers/DatePicker.vue";
import BtnAdd from "../../components/Buttons/BtnAdd.vue";
import DataTable from "../../components/DataTables/DataTable.vue";

export default {
    name: 'purchaseInfoDialog',
    components: {
        "c-btn-add": BtnAdd,
        "c-date-picker": DatePicker,
        "c-data-table": DataTable,
    },
    data() {
        return {
            toggle_form: undefined,
            purchaseQty: null,
            validator: false,
            text_requiredRules: [
                v => !!v || '必填',
            ],
            text_requiredRules_isNumber: [
                v => !!v || '必填',
                v => {
                    return !!parseFloat(v) || '必須為數字'
                },
            ],
            purchaseDetails: [],
            purchaseDetailHeader: [
                { text: '採購名稱', value: 'Name' },
                { text: '採購狀態', value: 'Status' },
                { text: '採購種類', value: 'PurchaseType' },
                { text: "備註", value: "Remark", width: "10%" },
                { text: '', value: 'actions', sortable: false, width: "10%" },
            ],
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
    },
    methods: {
        onClick_confirm: function () { //有子元件的事件觸發 自定義事件childevent
            this.$refs.form.validate();
            if (this.validator == false) {
                return
            }
            this.$emit('confirm', this.purchaseItem);//觸發一個在子元件中宣告的事件 childEvnet
            // this.resetForm();
        },
        onClick_cancel() {
            this.purchaseInfoDialog = false;
            // this.resetForm();
        },
        resetForm() {
            this.$refs.form.reset();
            // this.productID = null;
            // this.productSku = null;
            // this.productName = null;
            // this.productStocks = null;
            // this.purchaseQty = null;
        },
    },
    // watch: {
    //     purchaseInfoDialog: function () {
    //         if (this.purchaseInfoDialog == false) {
    //             this.resetForm();
    //         }
    //     },
    // }
}
</script>