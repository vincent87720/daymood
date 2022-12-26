<template>
    <v-dialog v-model="supplierDialog" @click:outside="onClick_cancel" fullscreen>
        <v-card class="d-flex align-center flex-column justify-center">
            <v-card-title class="text-h5">
                {{ text_cardTitle }}
            </v-card-title>
            <v-container class="pa-6">
                <v-row>
                    <v-col cols="12" sm="6" class="ml-auto mr-auto">
                        <v-form ref="form" v-model="validator">
                            <v-text-field label="廠商名稱" v-model="supplierItem.Name" :rules="text_requiredRules"
                                prepend-icon="mdi-store-outline" v-on:keydown.enter.prevent="onClick_confirm" autofocus
                                required></v-text-field>
                            <v-text-field label="廠商地址" v-model="supplierItem.Address" :rules="text_requiredRules"
                                prepend-icon="mdi-map-marker-outline" v-on:keydown.enter.prevent="onClick_confirm"
                                required></v-text-field>
                            <v-text-field label="備註" v-model="supplierItem.Remark" prepend-icon="mdi-text-long"
                                v-on:keydown.enter.prevent="onClick_confirm"></v-text-field>
                        </v-form>
                    </v-col>
                </v-row>
            </v-container>

            <div class="pa-6 pt-3">
                <v-btn class="mx-2" outlined rounded text @click.stop="onClick_cancel">
                    取消
                </v-btn>
                <v-btn class="mx-2" outlined rounded text :disabled="!validator" @click.stop="onClick_confirm">
                    {{ text_confirmBtn }}
                </v-btn>
            </div>
        </v-card>
    </v-dialog>
</template>

<script>

export default {
    name: 'SupplierDialog',
    components: {

    },
    data() {
        return {
            supplierForm: {
                name: "",
                address: "",
                remark: "",
            },
            validator: false,
            text_requiredRules: [
                v => !!v || '必填',
            ],
        };
    },
    props: {
        prop_supplierDialog: {
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
        prop_supplierItem: {
            type: Object,
            required: false
        },
    },
    computed: {
        supplierDialog: {
            get() {
                return this.prop_supplierDialog
            },
            set(val) {
                this.$emit('update:prop_supplierDialog', val)

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
        supplierItem: {
            get() {
                return this.prop_supplierItem
            },
            set(val) {
                this.$emit('update:prop_supplierItem', val)
            }
        },
    },
    methods: {
        onClick_confirm: function () { //有子元件的事件觸發 自定義事件childevent
            this.$refs.form.validate();
            if (this.validator == false) {
                return
            }
            this.$emit('confirm', this.supplierItem);//觸發一個在子元件中宣告的事件 childEvnet
            // this.resetForm();
        },
        onClick_cancel() {
            this.supplierDialog = false;
            // this.resetForm();
        },
        resetForm() {
            this.supplierItem = {
                name: "",
                address: "",
                remark: "",
            };
            this.$refs.form.reset();
        },
    },
    // watch:  {
    //     supplierDialog: function(){
    //         if(this.supplierDialog == false){
    //             this.resetForm();
    //         }
    //     },
    // }
}
</script>