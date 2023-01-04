<template>
    <v-dialog v-model="confirmDialog" @click:outside="onClick_cancel" v-on:keydown.enter.prevent="onClick_confirm"
        fullscreen>
        <v-card class="d-flex align-center flex-column justify-center">
            <v-card-title class="text-h5">
                {{ text_cardTitle }}
            </v-card-title>
            <div class="text-h6" v-if="text_cardHint">
                {{ text_cardHint }}
            </div>
            <v-container class="pa-6">
                <v-row>
                    <v-col xs="12" sm="6" class="ml-auto mr-auto">
                        <div v-if="confirmTarget" class="d-flex flex-column justify-center align-center">
                            <h2 v-if="confirmTarget.Name">
                                {{ confirmTarget.Name }}
                            </h2>
                            <h4 v-if="confirmTarget.ID" class="font-weight-thin">
                                {{ confirmTarget.ID }}
                            </h4>
                        </div>
                    </v-col>
                </v-row>
            </v-container>

            <v-card-actions class="pa-6 pt-3">
                <v-spacer></v-spacer>
                <v-btn outlined rounded text @click.stop="onClick_cancel"> 取消 </v-btn>
                <slot name="actions" v-bind="{ item: confirmTarget }"></slot>
                <v-btn outlined rounded text @click.stop="onClick_confirm">
                    {{ text_confirmBtn }}
                </v-btn>
            </v-card-actions>
        </v-card>
    </v-dialog>
</template>

<script>
export default {
    name: "ConfirmDialog",
    components: {},
    data() {
        return {};
    },
    props: {
        prop_confirmDialog: {
            type: Boolean,
            required: true,
        },
        prop_text_cardTitle: {
            type: String,
            required: true,
        },
        prop_text_cardHint: {
            type: String,
            required: true,
        },
        prop_text_confirmBtn: {
            type: String,
            required: true,
        },
        prop_confirmTarget: {
            type: Object,
            required: false,
        },
    },
    computed: {
        confirmDialog: {
            get() {
                return this.prop_confirmDialog;
            },
            set(val) {
                this.$emit("update:prop_confirmDialog", val);
            },
        },
        text_cardTitle: {
            get() {
                return this.prop_text_cardTitle;
            },
            set(val) {
                this.$emit("update:prop_text_cardTitle", val);
            },
        },
        text_cardHint: {
            get() {
                return this.prop_text_cardHint;
            },
            set(val) {
                this.$emit("update:prop_text_cardHint", val);
            },
        },
        text_confirmBtn: {
            get() {
                return this.prop_text_confirmBtn;
            },
            set(val) {
                this.$emit("update:prop_text_confirmBtn", val);
            },
        },
        confirmTarget: {
            get() {
                return this.prop_confirmTarget;
            },
            set(val) {
                this.$emit("update:prop_confirmTarget", val);
            },
        },
    },
    methods: {
        onClick_confirm() {
            //有子元件的事件觸發 自定義事件childevent
            this.$emit("confirmClick", this.confirmTarget); //觸發一個在子元件中宣告的事件 childEvnet
            this.resetForm();
        },
        onClick_cancel() {
            this.confirmDialog = false;
            this.resetForm();
        },
        resetForm() {
            this.confirmTarget = null;
        },
    },
    watch: {
        confirmDialog: function () {
            if (this.confirmDialog == false) {
                this.resetForm();
            }
        },
    },
};
</script>
