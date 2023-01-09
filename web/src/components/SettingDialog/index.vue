<template>
    <div>
        <v-dialog v-model="settingDialog" @click:outside="onClick_cancel" fullscreen>
            <v-card class="d-flex align-center flex-column justify-center">
                <v-card-title class="text-h5"> {{ text_cardTitle }} </v-card-title>
                <v-container class="pa-6">
                    <v-row>
                        <v-col cols="12" sm="6" class="ml-auto mr-auto">
                            <v-form ref="form" v-model="valid_settingForm">
                                <v-text-field label="貨運行抽成(%)" v-model="tradings.Ajeossi"
                                    :rules="text_requiredRules_isNumber" prepend-icon="mdi-truck"
                                    v-on:keydown.enter.prevent="onClick_confirm" autofocus required></v-text-field>
                                <v-text-field label="國際運費" v-model="tradings.ShippingFee"
                                    :rules="text_requiredRules_isNumber" prepend-icon="mdi-ferry"
                                    v-on:keydown.enter.prevent="onClick_confirm" required></v-text-field>
                                <v-text-field label="韓圓匯率" v-model="tradings.ExchangeRate"
                                    :rules="text_requiredRules_isNumber" prepend-icon="mdi-cash-marker"
                                    v-on:keydown.enter.prevent="onClick_confirm" required></v-text-field>
                                <v-text-field label="關稅" v-model="tradings.Tariff" :rules="text_requiredRules_isNumber"
                                    prepend-icon="mdi-cash-lock" v-on:keydown.enter.prevent="onClick_confirm"
                                    required></v-text-field>
                                <v-text-field label="定價倍率" v-model="tradings.Markup"
                                    :rules="text_requiredRules_isNumber" prepend-icon="mdi-cash-plus"
                                    v-on:keydown.enter.prevent="onClick_confirm" required></v-text-field>
                                <v-combobox v-model="tradings.Costs" :items="items" :item-text="convertCostText"
                                    clearable multiple small-chips append-outer-icon="mdi-plus"
                                    prepend-icon="mdi-ribbon" label="其他成本"
                                    @click:append-outer.stop="onClick_addCosts"></v-combobox>
                            </v-form>
                        </v-col>
                    </v-row>
                    <v-dialog v-model="dialog" max-width="290">
                        <v-card>
                            <v-container class="px-4">
                                <v-row>
                                    <v-col cols="6">
                                        <v-text-field v-model="cost.Key" label="項目"></v-text-field>
                                    </v-col>
                                    <v-col cols="6">
                                        <v-text-field v-model="cost.Value" label="金額"></v-text-field>
                                    </v-col>
                                </v-row>
                            </v-container>

                            <v-card-actions>
                                <v-spacer></v-spacer>
                                <v-btn text @click="dialog = false">
                                    取消
                                </v-btn>
                                <v-btn text @click="onConfirm_addCost(cost)">
                                    新增
                                </v-btn>
                            </v-card-actions>
                        </v-card>
                    </v-dialog>
                </v-container>

                <v-card-actions class="pa-6 pt-3">
                    <v-spacer></v-spacer>
                    <v-btn outlined rounded text @click.stop="onClick_cancel"> 取消 </v-btn>
                    <v-btn outlined rounded text :disabled="!valid_settingForm" @click.stop="onClick_confirm">
                        {{ text_confirmBtn }}
                    </v-btn>
                </v-card-actions>
            </v-card>
        </v-dialog>
        <Alert :prop_alert.sync="alert" :prop_alertType="alertType" :prop_alertText="alertText"></Alert>
    </div>
</template>

<script>
import Alert from '../../components/Alert/index.vue'
import { getTradingSettings, putTradingSettings } from "@/apis/TradingSettingsAPI.js";
export default {
    name: 'SettingDialog',
    components: {
        Alert,
    },
    data() {
        return {
            //Alert
            alert: false,
            alertType: "",
            alertText: "",

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
            tradings: {},
            items: new Array(),
            dialog: false,
            cost: {
                Key: "",
                Value: undefined,
            },
        };
    },
    async mounted() {
        await this.getTradingSettings();
    },
    props: {
        prop_settingDialog: {
            type: Boolean,
            required: true
        },
        prop_text_cardTitle: {
            type: String,
            required: false,
            default: "設定",
        },
        prop_text_confirmBtn: {
            type: String,
            required: false,
            default: "確定",
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
    },
    methods: {
        convertCostText(item) {
            return `${item.Key} ${item.Value}`;
        },
        onClick_addCosts() {
            this.dialog = true;
        },
        onClick_confirm: async function () {
            this.$refs.form.validate();
            if (this.valid_settingForm == false) {
                return
            }
            await this.putTradingSettings(this.tradings);
        },
        onClick_cancel() {
            this.settingDialog = false;
        },
        onConfirm_addCost(item) {
            if (this.items == null) {
                this.items = [item];
            }
            else {
                this.items.push(item)
            }
            this.dialog = false;
        },
        async getTradingSettings() {
            await getTradingSettings().then((response) => {
                if (response.data.trading) {
                    this.tradings = response.data.trading;

                    //deep copy
                    this.items = JSON.parse(JSON.stringify(response.data.trading.Costs));
                }
                else {
                    this.tradings = {};
                    this.items = [];
                }
            })
        },
        preSend(item) {
            item.Ajeossi = parseFloat(item.Ajeossi);
            item.ShippingFee = parseFloat(item.ShippingFee);
            item.ExchangeRate = parseFloat(item.ExchangeRate);
            item.Tariff = parseFloat(item.Tariff);
            item.Markup = parseFloat(item.Markup);
            if (item.Costs) {
                item.Costs.map((x) => { x.Value = parseFloat(x.Value) })
            }
            return item
        },
        async putTradingSettings(item) {
            item = this.preSend(item);
            await putTradingSettings(item)
                .then(async (response) => {
                    this.tradings = response.data.trading;
                    this.alert = true;
                    this.alertType = "Success";
                    this.alertText = "編輯設定成功";
                    this.settingDialog = false;
                })
                .catch((error) => {
                    this.alert = true;
                    this.alertType = "Fail";
                    this.alertText = "編輯設定失敗";
                });
        },
    },
}
</script>