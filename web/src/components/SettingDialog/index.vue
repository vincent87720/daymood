<script>

export default {
    name: 'SettingDialog',
    components: {

    },
    data() {
        return {
            valid_settingForm: false,
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
        prop_settingDialog: {
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
        prop_ajeossi: {
            type: String,
            required: false
        },
        prop_shippingFee: {
            type: String,
            required: false
        },
        prop_exchangeRate: {
            type: String,
            required: false
        },
        prop_tariff: {
            type: String,
            required: false
        },
        prop_markup: {
            type: String,
            required: false
        },
    },
    computed: {
        settingDialog: {
            get() {
                return this.prop_settingDialog
            },
            set(val) {
                this.$emit('update:prop_settingDialog', val)

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
        ajeossi: {
            get() {
                return this.prop_ajeossi
            },
            set(val) {
                this.$emit('update:prop_ajeossi', val)
            }
        },
        shippingFee: {
            get() {
                return this.prop_shippingFee
            },
            set(val) {
                this.$emit('update:prop_shippingFee', val)
            }
        },
        exchangeRate: {
            get() {
                return this.prop_exchangeRate
            },
            set(val) {
                this.$emit('update:prop_exchangeRate', val)
            }
        },
        tariff: {
            get() {
                return this.prop_tariff
            },
            set(val) {
                this.$emit('update:prop_tariff', val)
            }
        },
        markup: {
            get() {
                return this.prop_markup
            },
            set(val) {
                this.$emit('update:prop_markup', val)
            }
        },
    },
    methods: {
        onClick_confirm: function () { //有子元件的事件觸發 自定義事件childevent
            this.$refs.form.validate();
            if (this.valid_settingForm == false) {
                return
            }
            var settingInfo = {};
            settingInfo.ajeossi = this.ajeossi;
            settingInfo.shippingFee = this.shippingFee;
            settingInfo.exchangeRate = this.exchangeRate;
            settingInfo.tariff = this.tariff;
            settingInfo.markup = this.markup;

            this.$emit('confirmClick', settingInfo);//觸發一個在子元件中宣告的事件 childEvnet
            this.resetForm();
        },
        onClick_cancel() {
            this.settingDialog = false;
            this.resetForm();
        },
        resetForm() {
            this.ajeossi = null;
            this.shippingFee = null;
            this.exchangeRate = null;
            this.tariff = null;
            this.markup = null;
            this.$refs.form.reset();
        },
    },
    watch: {
        settingDialog: function () {
            if (this.settingDialog == false) {
                this.resetForm();
            }
        },
    }
}
</script>
<template src="./template.html"></template>