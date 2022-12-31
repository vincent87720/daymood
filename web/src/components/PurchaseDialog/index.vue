<template>
    <v-dialog v-model="purchaseDialog" @click:outside="onClick_cancel" fullscreen>
        <v-card class="d-flex align-center flex-column justify-center">
            <v-card-title class="text-h5">
                {{ text_cardTitle }}
            </v-card-title>
            <v-container class="pa-6">
                <v-row>
                    <v-col xs="12" sm="6" class="ml-auto mr-auto">
                        <v-btn-toggle v-model="toggle_form" rounded mandatory dense class="d-flex justify-sm-center">
                            <v-btn> 基本資訊 </v-btn>
                            <v-btn> 貨運資訊 </v-btn>
                        </v-btn-toggle>
                        <v-form ref="form" v-model="validator">
                            <div v-if="toggle_form == 0">
                                <v-text-field label="採購名稱" v-model="purchaseItem.Name" :rules="text_requiredRules"
                                    required prepend-icon="mdi-text-short"></v-text-field>
                                <v-select class="overflow-hidden" label="採購狀態" v-model="purchaseItem.Status"
                                    :rules="text_requiredRules" prepend-icon="mdi-package-variant"
                                    :items="systemConfigs.PurchaseStatus" item-text="value" item-value="key" required
                                    dense full-width></v-select>
                                <v-select class="overflow-hidden" label="採購類型" v-model="purchaseItem.PurchaseType"
                                    :rules="text_requiredRules" prepend-icon="mdi-package-variant"
                                    :items="systemConfigs.PurchaseType" item-text="value" item-value="key" required
                                    dense full-width></v-select>
                            </div>
                            <div v-if="toggle_form == 1">
                                <v-text-field label="貨運行" v-model="purchaseItem.ShippingAgent" required
                                    prepend-icon="mdi-text-short"></v-text-field>
                                <v-text-field label="貨運團主" v-model="purchaseItem.ShippingInitiator" required
                                    prepend-icon="mdi-text-short"></v-text-field>
                                <c-date-picker prop_label="開團日期"
                                    :prop_date.sync="purchaseItem.ShippingCreateAt"></c-date-picker>
                                <c-date-picker prop_label="結單日期"
                                    :prop_date.sync="purchaseItem.ShippingEndAt"></c-date-picker>
                                <c-date-picker prop_label="送達日期"
                                    :prop_date.sync="purchaseItem.ShippingArriveAt"></c-date-picker>
                                <v-text-field label="貨運總重" v-model="purchaseItem.Weight" required
                                    prepend-icon="mdi-text-short"></v-text-field>
                                <v-text-field label="國內運費 韓國" v-model="purchaseItem.ShippingFeeKr" required
                                    prepend-icon="mdi-text-short"></v-text-field>
                                <v-text-field label="國內運費 台灣" v-model="purchaseItem.ShippingFeeTw" required
                                    prepend-icon="mdi-text-short"></v-text-field>
                                <v-text-field label="國際運費" v-model="purchaseItem.ShippingFeeKokusai" required
                                    prepend-icon="mdi-text-short"></v-text-field>
                            </div>
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
import DatePicker from "../../components/Pickers/DatePicker.vue";

export default {
    name: 'PurchaseDialog',
    components: {
        "c-date-picker": DatePicker,
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
        };
    },
    props: {
        prop_purchaseDialog: {
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
        prop_purchaseItem: {
            type: Object,
            required: false
        },
    },
    computed: {
        purchaseDialog: {
            get() {
                return this.prop_purchaseDialog
            },
            set(val) {
                this.$emit('update:prop_purchaseDialog', val)

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
            this.purchaseDialog = false;
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
    //     purchaseDialog: function () {
    //         if (this.purchaseDialog == false) {
    //             this.resetForm();
    //         }
    //     },
    // }
}
</script>