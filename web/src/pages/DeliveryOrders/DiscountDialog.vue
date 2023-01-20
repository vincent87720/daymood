<template>
    <v-dialog v-model="discountDialog" @click:outside="onClick_cancel" fullscreen>
        <v-card class="d-flex align-center flex-column justify-center">
            <v-card-title class="text-h5">
                {{ text_cardTitle }}
            </v-card-title>
            <v-container class="pa-6" fluid>
                <v-row>
                    <v-col xs="12" sm="6" class="ml-auto mr-auto">
                        <v-form ref="form" v-model="validator">
                            <v-row class="ma-0 d-flex justify-center">
                                <v-col>
                                    <v-select label="折扣方式" v-model="discountItem.DiscountType"
                                        class="overflow-hidden" prepend-icon="mdi-package-variant" :items="systemConfigs.DiscountType"
                                        item-text="value" item-value="key" clearable dense></v-select>
                                </v-col>
                            </v-row>
                            <v-row class="ma-0 d-flex justify-center">
                                <v-col>
                                    <v-text-field label="折扣名稱" v-model="discountItem.Name"
                                        prepend-icon="mdi-text-short" v-on:keydown.enter.prevent="onClick_confirm"
                                        dense></v-text-field>
                                </v-col>
                            </v-row>
                            <v-row class="ma-0 d-flex justify-center">
                                <v-col>
                                    <v-text-field label="折扣金額" v-model="discountItem.Price"
                                        prepend-icon="mdi-currency-twd" v-on:keydown.enter.prevent="onClick_confirm"
                                        dense></v-text-field>
                                </v-col>
                            </v-row>
                            <v-row class="ma-0 d-flex justify-center">
                                <v-col>
                                    <v-textarea label="備註" v-model="discountItem.Remark"
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
    name: 'DiscountDialog',
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
        prop_discountDialog: {
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
        prop_deliveryOrderItem: {
            type: Object,
            required: false
        },
        prop_discountItem: {
            type: Object,
            required: false
        },
    },
    computed: {
        discountDialog: {
            get() {
                return this.prop_discountDialog
            },
            set(val) {
                this.$emit('update:prop_discountDialog', val)

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
        deliveryOrderItem: {
            get() {
                return this.prop_deliveryOrderItem
            },
            set(val) {
                this.$emit('update:prop_deliveryOrderItem', val)
            }
        },
        discountItem: {
            get() {
                return this.prop_discountItem
            },
            set(val) {
                this.$emit('update:prop_discountItem', val)
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
            this.$emit('confirm', this.discountItem);//觸發一個在子元件中宣告的事件 childEvnet
        },
        onClick_cancel() {
            this.discountDialog = false;
        },
    },
    watch: {
        discountDialog: async function (newVal, oldVal) {
            if (newVal == true) {
                this.$store.dispatch("GetProducts");
            }
        },
    }
}
</script>