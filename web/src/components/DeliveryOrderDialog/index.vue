<template>
    <v-dialog v-model="deliveryOrderDialog" @click:outside="onClick_cancel" fullscreen>
        <v-card class="d-flex align-center flex-column justify-center">
            <v-card-title class="text-h5">
                {{ text_cardTitle }}
            </v-card-title>
            <v-container class="pa-6">
                <v-row>
                    <v-col xs="12" sm="6" class="ml-auto mr-auto">
                        <v-btn-toggle v-model="toggle_form" rounded mandatory dense
                            class="d-flex justify-sm-center mb-3">
                            <v-btn> 貨運資訊 </v-btn>
                            <v-btn> 費用資訊 </v-btn>
                        </v-btn-toggle>
                        <v-form ref="form" v-model="validator">
                            <div v-show="toggle_form == 0">
                                <v-select class="overflow-hidden" label="出貨單狀態" v-model="deliveryOrderItem.Status"
                                    prepend-icon="mdi-toggle-switch" :items="systemConfigs.DeliveryOrderStatus"
                                    item-text="value" item-value="key"></v-select>
                                <v-select class="overflow-hidden" label="出貨方式" v-model="deliveryOrderItem.DeliveryType"
                                    prepend-icon="mdi-truck" :items="systemConfigs.DeliveryType" item-text="value"
                                    item-value="key"></v-select>
                                <v-select class="overflow-hidden" label="出貨狀態"
                                    v-model="deliveryOrderItem.DeliveryStatus" prepend-icon="mdi-package-variant"
                                    :items="systemConfigs.DeliveryStatus" item-text="value" item-value="key"></v-select>
                                <c-date-picker prop_label="下訂日期"
                                    :prop_date.sync="deliveryOrderItem.OrderAt"></c-date-picker>
                                <c-date-picker prop_label="出貨日期"
                                    :prop_date.sync="deliveryOrderItem.SendAt"></c-date-picker>
                                <c-date-picker prop_label="送達日期"
                                    :prop_date.sync="deliveryOrderItem.ArriveAt"></c-date-picker>
                                <v-textarea label="備註" v-model="deliveryOrderItem.Remark" prepend-icon="mdi-text-long"
                                    auto-grow rows="1" row-height="15"></v-textarea>
                            </div>
                            <div v-show="toggle_form == 1">
                                <v-select class="overflow-hidden" label="運費狀態"
                                    v-model="deliveryOrderItem.DeliveryFeeStatus" prepend-icon="mdi-truck"
                                    :items="systemConfigs.DeliveryFeeStatus" item-text="value"
                                    item-value="key"></v-select>
                                <v-select class="overflow-hidden" label="付款方式" v-model="deliveryOrderItem.PaymentType"
                                    prepend-icon="mdi-cash-register" :items="systemConfigs.PaymentType"
                                    item-text="value" item-value="key"></v-select>
                                <v-select class="overflow-hidden" label="付款狀態" v-model="deliveryOrderItem.PaymentStatus"
                                    prepend-icon="mdi-cash-check" :items="systemConfigs.PaymentStatus" item-text="value"
                                    item-value="key"></v-select>
                                <v-text-field label="原價" v-model="deliveryOrderItem.TotalOriginal"
                                    prepend-icon="mdi-currency-usd"></v-text-field>
                                <v-text-field label="折扣" v-model="deliveryOrderItem.Discount"
                                    prepend-icon="mdi-currency-usd"></v-text-field>
                                <v-text-field label="總價" v-model="deliveryOrderItem.TotalDiscounted"
                                    prepend-icon="mdi-currency-usd"></v-text-field>
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
    name: 'DeliveryOrderDialog',
    components: {
        "c-date-picker": DatePicker,
    },
    data() {
        return {
            toggle_form: undefined,
            validator: false,
            rules: {
                required: value => !!value || '必填',
                requiredIncludeZero: value => !(value === undefined || value == null || value == "") || '必填',
                number: value => !isNaN(value) || '必須為數字',
            },
        };
    },
    props: {
        prop_deliveryOrderDialog: {
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
        prop_deliveryOrderItem: {
            type: Object,
            required: false
        },
    },
    computed: {
        deliveryOrderDialog: {
            get() {
                return this.prop_deliveryOrderDialog
            },
            set(val) {
                this.$emit('update:prop_deliveryOrderDialog', val)

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
        deliveryOrderItem: {
            get() {
                return this.prop_deliveryOrderItem
            },
            set(val) {
                this.$emit('update:prop_deliveryOrderItem', val)
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
            this.$emit('confirm', this.deliveryOrderItem);//觸發一個在子元件中宣告的事件 childEvnet
        },
        onClick_cancel() {
            this.deliveryOrderDialog = false;
        },
    },
}
</script>