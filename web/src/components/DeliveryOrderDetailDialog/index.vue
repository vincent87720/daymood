<template>
    <v-dialog v-model="deliveryOrderDetailDialog" @click:outside="onClick_cancel" fullscreen>
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
                                    <v-select label="商品" v-model="deliveryOrderDetailItem.ProductID"
                                        class="overflow-hidden" prepend-icon="mdi-package-variant" :items="productList"
                                        item-text="value" item-value="key" clearable dense
                                        @change="onChange_product"></v-select>
                                </v-col>
                                <v-col cols="12" md="6">
                                    <v-text-field label="售價" v-model="deliveryOrderDetailItem.RetailPrice"
                                        prepend-icon="mdi-identifier" v-on:keydown.enter.prevent="onClick_confirm"
                                        dense></v-text-field>
                                </v-col>
                            </v-row>
                            <v-row class="ma-0 d-flex justify-center">
                                <v-col cols="12" md="6">
                                    <v-text-field label="數量" v-model="deliveryOrderDetailItem.QTY"
                                        prepend-icon="mdi-text-short" v-on:keydown.enter.prevent="onClick_confirm"
                                        dense></v-text-field>
                                </v-col>
                                <v-col cols="12" md="6">
                                    <v-text-field label="小計" v-model="deliveryOrderDetailItem.Subtotal"
                                        prepend-icon="mdi-text-short" v-on:keydown.enter.prevent="onClick_confirm" dense
                                        :placeholder="calc_Subtotal" :hint="hint_Subtotal"
                                        :persistent-hint="true"></v-text-field>
                                </v-col>
                            </v-row>
                            <v-row class="ma-0 d-flex justify-center">
                                <v-col>
                                    <v-textarea label="備註" v-model="deliveryOrderDetailItem.Remark"
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

export default {
    name: 'deliveryOrderDetailDialog',
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
        };
    },
    mounted() {
    },
    props: {
        prop_deliveryOrderDetailDialog: {
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
        prop_deliveryOrderDetailItem: {
            type: Object,
            required: false
        },
    },
    computed: {
        deliveryOrderDetailDialog: {
            get() {
                return this.prop_deliveryOrderDetailDialog
            },
            set(val) {
                this.$emit('update:prop_deliveryOrderDetailDialog', val)

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
        deliveryOrderDetailItem: {
            get() {
                return this.prop_deliveryOrderDetailItem
            },
            set(val) {
                this.$emit('update:prop_deliveryOrderDetailItem', val)
            }
        },
        productList() {
            return this.$store.state.products;
        },
        calc_Subtotal() {
            this.deliveryOrderDetailItem.Subtotal = "";
            this.hint_Subtotal = "";
            let result = 0;
            let wholesalePrice = parseFloat(this.deliveryOrderDetailItem.RetailPrice);
            let qty = parseFloat(this.deliveryOrderDetailItem.QTY);
            let subtotal = wholesalePrice * qty
            if (isNaN(subtotal) == false) {
                this.deliveryOrderDetailItem.Subtotal = subtotal;
                this.hint_Subtotal = `${this.deliveryOrderDetailItem.RetailPrice}(售價)＊${this.deliveryOrderDetailItem.QTY}(數量)`
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
            this.$emit('confirm', this.deliveryOrderDetailItem);//觸發一個在子元件中宣告的事件 childEvnet
        },
        onClick_cancel() {
            this.deliveryOrderDetailDialog = false;
        },
        onChange_product(productID) {
            let product = this.productList.find(x => x.key == productID);
            if (product != null) {
                this.deliveryOrderDetailItem.RetailPrice = product.RetailPrice;
            }
            else {
                this.deliveryOrderDetailItem.RetailPrice = "";
            }
        },
    },
    watch: {
        deliveryOrderDetailDialog: async function (newVal, oldVal) {
            if (newVal == true) {
                this.$store.dispatch("GetProducts");
            }
        },
    }
}
</script>