<template>
    <v-dialog v-model="purchaseDetailDialog" @click:outside="onClick_cancel" fullscreen>
        <v-card class="d-flex align-center flex-column justify-center">
            <v-card-title class="text-h5">
                {{ text_cardTitle }}
            </v-card-title>
            <v-container class="pa-6" fluid>
                <v-row>
                    <v-col xs="12" sm="6" class="ml-auto mr-auto">
                        <v-form ref="form" v-model="validator">
                            <v-row class="ma-0 d-flex justify-center">
                                <v-col cols="12" md="6">
                                    <v-text-field label="採購商品編號" v-model="purchaseDetailItem.NamedID"
                                        prepend-icon="mdi-identifier" v-on:keydown.enter.prevent="onClick_confirm"
                                        dense></v-text-field>
                                </v-col>
                                <v-col cols="12" md="6">
                                    <v-text-field label="商品名稱" v-model="purchaseDetailItem.Name"
                                        :rules="[rules.required]" prepend-icon="mdi-text-short"
                                        v-on:keydown.enter.prevent="onClick_confirm" dense></v-text-field>
                                </v-col>
                            </v-row>
                            <v-row class="ma-0 d-flex justify-center">
                                <v-col cols="12" md="6">
                                    <v-select label="進貨廠商" v-model="purchaseDetailItem.SupplierID"
                                        class="overflow-hidden" prepend-icon="mdi-store-outline" :items="supplierList"
                                        item-text="value" item-value="key" clearable></v-select>
                                </v-col>
                                <v-col cols="12" md="6">
                                    <v-select label="商品" v-model="purchaseDetailItem.ProductID" class="overflow-hidden"
                                        prepend-icon="mdi-package-variant" :items="productList" item-text="value"
                                        item-value="key" clearable></v-select>
                                </v-col>
                            </v-row>
                            <v-row class="ma-0 d-flex justify-center">
                                <v-col>
                                    <v-select label="是否採用" v-model="purchaseDetailItem.Status" class="overflow-hidden"
                                        :rules="[rules.requiredIncludeZero, rules.number]"
                                        prepend-icon="mdi-archive-check" :items="systemConfigs.PurchaseDetailStatus"
                                        item-text="value" item-value="key"></v-select>
                                </v-col>
                            </v-row>
                            <v-row class="ma-0 d-flex justify-center">
                                <v-col cols="12" md="4">
                                    <v-text-field label="批價" v-model="purchaseDetailItem.WholesalePrice"
                                        :rules="[rules.number]" prepend-icon="mdi-currency-usd"
                                        v-on:keydown.enter.prevent="onClick_confirm" dense></v-text-field>
                                </v-col>
                                <v-col cols="12" md="4">
                                    <v-text-field label="數量" v-model="purchaseDetailItem.QTY"
                                        :rules="[rules.requiredIncludeZero, rules.number]" prepend-icon="mdi-numeric"
                                        v-on:keydown.enter.prevent="onClick_confirm" dense></v-text-field>
                                </v-col>
                                <v-col cols="12" md="4">
                                    <v-text-field label="小計" v-model="purchaseDetailItem.Subtotal"
                                        :rules="[rules.number]" prepend-icon="mdi-currency-usd"
                                        v-on:keydown.enter.prevent="onClick_confirm" dense :placeholder="calc_Subtotal"
                                        :hint="hint_Subtotal" :persistent-hint="true"></v-text-field>
                                </v-col>
                            </v-row>
                            <v-row class="ma-0 d-flex justify-center">
                                <v-col>
                                    <v-textarea label="備註" v-model="purchaseDetailItem.Remark"
                                        prepend-icon="mdi-text-long" auto-grow rows="1" row-height="15"
                                        dense></v-textarea>
                                </v-col>
                            </v-row>
                        </v-form>
                    </v-col>
                </v-row>
            </v-container>

            <v-card-actions class="pa-6 pt-3">
                <v-spacer></v-spacer>
                <v-btn outlined rounded text @click.stop="onClick_cancel">
                    取消
                </v-btn>
                <v-btn outlined rounded text :disabled="!validator" @click.stop="onClick_confirm">
                    {{ text_confirmBtn }}
                </v-btn>
            </v-card-actions>
        </v-card>
    </v-dialog>
</template>

<script>
import { getPurchaseDetails } from "../../apis/PurchaseDetailsAPI";

export default {
    name: 'purchaseDetailDialog',
    components: {
    },
    data() {
        return {
            validator: false,
            rules: {
                required: value => !!value || '必填',
                requiredIncludeZero: value => !(value === undefined || value == null || value == "") || '必填',
                number: value => !isNaN(value) || '必須為數字',
            },

            hint_Subtotal: "",
            hint_SubtotalTwd: "",
        };
    },
    mounted() {
    },
    props: {
        prop_purchaseDetailDialog: {
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
        prop_purchaseDetailItem: {
            type: Object,
            required: false
        },
    },
    computed: {
        purchaseDetailDialog: {
            get() {
                return this.prop_purchaseDetailDialog
            },
            set(val) {
                this.$emit('update:prop_purchaseDetailDialog', val)

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
        purchaseDetailItem: {
            get() {
                return this.prop_purchaseDetailItem
            },
            set(val) {
                this.$emit('update:prop_purchaseDetailItem', val)
            }
        },
        systemConfigs() {
            return this.$store.state.systemConfigs;
        },
        productList() {
            return this.$store.state.products;
        },
        supplierList() {
            return this.$store.state.suppliers;
        },
        calc_Subtotal() {
            this.purchaseDetailItem.Subtotal = "";
            this.hint_Subtotal = "";
            let result = 0;
            let wholesalePrice = parseFloat(this.purchaseDetailItem.WholesalePrice);
            let qty = parseFloat(this.purchaseDetailItem.QTY);
            let subtotal = wholesalePrice * qty
            if (isNaN(subtotal) == false) {
                this.purchaseDetailItem.Subtotal = subtotal;
                this.hint_Subtotal = `${this.purchaseDetailItem.WholesalePrice}(批價)＊${this.purchaseDetailItem.QTY}(數量)`
                result += subtotal
            }
            return result.toFixed(2);
        },
    },
    methods: {
        onClick_confirm: function () { //有子元件的事件觸發 自定義事件childevent
            this.$refs.form.validate();
            if (this.validator == false) {
                return
            }
            this.$emit('confirm', this.purchaseDetailItem);//觸發一個在子元件中宣告的事件 childEvnet
        },
        onClick_cancel() {
            this.purchaseDetailDialog = false;
        },
        async getPurchaseDetails(item) {
            await getPurchaseDetails(item)
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
    },
    watch: {
        purchaseDetailDialog: async function (newVal, oldVal) {
            if (newVal == true) {
                this.$store.dispatch("GetProducts");
                this.$store.dispatch("GetSuppliers");
            }
        },
    }
}
</script>