<template>
    <v-dialog v-model="productDialog" @click:outside="onClick_cancel" fullscreen>
        <v-card class="d-flex align-center flex-column justify-center" width="100">
            <v-card-title class="text-h5 mt-4">
                {{ text_cardTitle }}
            </v-card-title>
            <v-container class="pa-6">
                <v-row>
                    <v-col xs="12" sm="6" class="ml-auto mr-auto">
                        <v-form ref="form" v-model="validator">
                            <v-text-field label="SKU商品編號" v-model="productItem.SKU" prepend-icon="mdi-identifier"
                                v-on:keydown.enter.prevent="onClick_confirm" autofocus></v-text-field>
                            <v-text-field label="商品名稱" v-model="productItem.Name" :rules="[rules.required]"
                                prepend-icon="mdi-text-short"
                                v-on:keydown.enter.prevent="onClick_confirm"></v-text-field>
                            <v-select label="商品種類" v-model="productItem.ProductType"
                                :rules="[rules.requiredIncludeZero, rules.number]" prepend-icon="mdi-tournament"
                                :items="systemConfigs.ProductType" item-text="value" item-value="key"></v-select>
                            <v-select label="進貨廠商" v-model="productItem.SupplierID" class="overflow-hidden"
                                :rules="[rules.number]" prepend-icon="mdi-store-outline" :items="supplierList"
                                item-text="value" item-value="key" clearable></v-select>
                            <v-text-field label="售價" v-model="productItem.RetailPrice" :rules="[rules.number]"
                                prepend-icon="mdi-currency-twd" v-on:keydown.enter.prevent="onClick_confirm"
                                :hint="hint_retailPrice" :persistent-hint="isMaterials == true"
                                :disabled="isMaterials == true"></v-text-field>
                            <v-text-field label="庫存" v-model="productItem.Stocks"
                                :rules="[rules.requiredIncludeZero, rules.number]" prepend-icon="mdi-package-variant"
                                v-on:keydown.enter.prevent="onClick_confirm"></v-text-field>
                            <v-text-field label="商品重量(g)" v-model="productItem.Weight" :rules="[rules.number]"
                                prepend-icon="mdi-weight-gram"
                                v-on:keydown.enter.prevent="onClick_confirm"></v-text-field>
                            <v-textarea label="備註" v-model="productItem.Remark" prepend-icon="mdi-text-long" auto-grow
                                rows="1" row-height="15"></v-textarea>
                        </v-form>
                    </v-col>
                </v-row>
            </v-container>
            <div class="pa-6 pt-3">

                <v-btn class="mx-2" outlined rounded text @click.stop="onClick_cancel">
                    取消
                </v-btn>
                <v-btn class="mx-2" outlined rounded text @click.stop="onClick_confirm">
                    {{ text_confirmBtn }}
                </v-btn>
            </div>
        </v-card>
    </v-dialog>
</template>

<script>

export default {
    name: 'ProductDialog',
    components: {

    },
    data() {
        return {
            imgFile: null,
            fooFile: null,
            validator: false,
            rules: {
                required: value => !!value || '必填',
                requiredIncludeZero: value => !(value === undefined || value == null || value == "") || '必填',
                number: value => !isNaN(value) || '必須為數字',
                email: value => {
                    const pattern = /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/
                    return pattern.test(value) || 'Invalid e-mail.'
                },
            },
            grossMargin: null,//毛利
            profitMargin: null,//毛利率

            hint_retailPrice: "",
        };
    },
    mounted() {
    },
    props: {
        prop_productDialog: {
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
        prop_actionType: {
            type: String,
            required: false
        },
        prop_productItem: {
            type: Object,
            required: false
        },
        prop_tradingSettings: {
            type: Object,
            required: true
        },
    },
    computed: {
        productDialog: {
            get() {
                return this.prop_productDialog
            },
            set(val) {
                this.$emit('update:prop_productDialog', val)

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
        actionType: {
            get() {
                return this.prop_actionType
            },
            set(val) {
                this.$emit('update:prop_actionType', val)
            }
        },
        productItem: {
            get() {
                return this.prop_productItem
            },
            set(val) {
                this.$emit('update:prop_productItem', val)
            }
        },
        tradingSettings: {
            get() {
                return this.prop_tradingSettings
            },
            set(val) {
                this.$emit('update:prop_tradingSettings', val)
            }
        },
        isMaterials() {
            let type_products = [1, 2, 3, 4, 5];//商品
            let type_materials = [6, 7, 8, 9, 10, 11, 12, 13];//耗材
            if (type_products.includes(this.productItem.ProductType)) {
                // if(this.actionType == "post") this.productItem.RetailPrice = undefined;
                this.hint_retailPrice = ""
                return false;
            }
            else if (type_materials.includes(this.productItem.ProductType)) {
                this.productItem.RetailPrice = 0;
                this.hint_retailPrice = "若商品種類為耗材，售價為零"
                return true
            }
        },
        systemConfigs() {
            return this.$store.state.systemConfigs;
        },
        supplierList() {
            return this.$store.state.suppliers;
        },
    },
    methods: {
        onClick_confirm() { //有子元件的事件觸發 自定義事件childevent
            this.$refs.form.validate();
            if (this.validator == false) {
                return
            }
            this.$emit('confirm', this.productItem);//觸發一個在子元件中宣告的事件 childEvnet
            // this.resetForm();
        },
        onClick_cancel() {
            this.productDialog = false;
            // this.resetForm();
        },
        resetForm() {
            this.productID = null;
            this.productSku = null;
            this.purchaseProductID = null;
            this.productName = null;
            this.productType = null;
            this.productImgName = null;
            this.productImgID = null;
            this.productWeight = null;
            this.krwWholesalePrice = null;
            this.ntdWholesalePrice = null;
            this.ntdListPrice = null;
            this.ntdSellingPrice = null;
            this.ntdCnf = null;
            this.ntdCost = null;
            this.imgFile = null;
            this.fooFile = null;
            this.$refs.form.reset();
        },
        clearFileInput() {
            this.$refs.imgFile.reset();
            this.imgFile = null;
        },
        calc_ntdWholesalePrice() {
            return (this.krwWholesalePrice / this.tradingSettings.ExchangeRate).toFixed(2).toString();
        },
        calc_listPrice() {
            var tempPrice = ((this.krwWholesalePrice / this.tradingSettings.ExchangeRate) * this.tradingSettings.Markup).toFixed();
            if (tempPrice % 10 == 0) {
                return (tempPrice).toString()
            }
            else if (tempPrice % 10 > 1) {
                return ((tempPrice - tempPrice % 10) + 10).toString()
            }
            else {
                return (tempPrice - tempPrice % 10).toString()
            }
        },
        calc_param() {
            var param = {};

            var gram = this.productWeight / 1000;
            param.ajeossiCalc = this.tradingSettings.Ajeossi / 100 + 1;//大哥抽成
            param.shippingFeeCalc = gram * this.tradingSettings.ShippingFee;
            param.tariffCalc = gram * this.tradingSettings.Tariff;
            return param
        },
        calc_cnfCost() {
            var param = this.calc_param();
            return ((this.krwWholesalePrice * param.ajeossiCalc + param.shippingFeeCalc) / this.tradingSettings.ExchangeRate + param.tariffCalc).toFixed(2).toString()
        },
        calc_cost() {
            var param = this.calc_param();
            return ((this.krwWholesalePrice * param.ajeossiCalc + param.shippingFeeCalc) / this.tradingSettings.ExchangeRate + param.tariffCalc + 25).toFixed(2).toString()
        },
        calc_grossMargin() {
            this.grossMargin = this.ntdSellingPrice - this.ntdCost;
            let percent = ((this.ntdSellingPrice - this.ntdCost) / this.ntdSellingPrice).toFixed(2).toString();

            if (isNaN(percent) == false) {
                this.profitMargin = percent;
            }
            else {
                this.profitMargin = 0;
            }
        }
    },
    watch: {
        productDialog: async function (newVal, oldVal) {
            if (newVal == true) {
                this.$store.dispatch("GetSuppliers");
            }
        },
        ntdSellingPrice: function () {
            this.calc_grossMargin();
        },
        ntdCost: function () {
            this.calc_grossMargin();
        },
    }
}
</script>
<style scoped src="./style.css">

</style>