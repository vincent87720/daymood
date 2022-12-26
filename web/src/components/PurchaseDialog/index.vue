<script>

    export default {
        name: 'PurchaseDialog',
        components:{
            
        },
        data(){
            return{
                purchaseQty: null,
                valid_purchaseForm: false,
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
        props:{
            prop_purchaseDialog:{
                type:Boolean,
                required:true
            },
            prop_text_cardTitle:{
                type:String,
                required:true
            },
            prop_text_confirmBtn:{
                type:String,
                required:true
            },
            prop_productID:{
                type:Number,
                required:false
            },
            prop_productSku:{
                type:String,
                required:false
            },
            prop_productName:{
                type:String,
                required:false
            },
            prop_productStocks:{
                type:Number,
                required:false
            },
        },
        computed:{
            purchaseDialog:{
                get(){
                    return this.prop_purchaseDialog
                },
                set(val){
                    this.$emit('update:prop_purchaseDialog',val)
                    
                }
            },
            text_cardTitle:{
                get(){
                    return this.prop_text_cardTitle
                },
                set(val){
                    this.$emit('update:prop_text_cardTitle',val)
                }
            },
            text_confirmBtn:{
                get(){
                    return this.prop_text_confirmBtn
                },
                set(val){
                    this.$emit('update:prop_text_confirmBtn',val)
                }
            },
            productID:{
                get(){
                    return this.prop_productID
                },
                set(val){
                    this.$emit('update:prop_productID',val)
                }
            },
            productSku:{
                get(){
                    return this.prop_productSku
                },
                set(val){
                    this.$emit('update:prop_productSku',val)
                }
            },
            productName:{
                get(){
                    return this.prop_productName
                },
                set(val){
                    this.$emit('update:prop_productName',val)
                }
            },
            productStocks:{
                get(){
                    return this.prop_productStocks
                },
                set(val){
                    this.$emit('update:prop_productStocks',val)
                }
            },
        },
        methods:{
            onClick_confirm:function(){ //有子元件的事件觸發 自定義事件childevent
                this.$refs.form.validate();
                if(this.valid_purchaseForm == false){
                    return
                }
                var purchaseInfo = {};
                purchaseInfo.productID = this.productID;
                purchaseInfo.productSku = this.productSku;
                purchaseInfo.productName = this.productName;
                purchaseInfo.stocks = this.productStocks;
                purchaseInfo.quantity = this.purchaseQty;
                this.$emit('confirmClick',purchaseInfo);//觸發一個在子元件中宣告的事件 childEvnet
                this.resetForm();
            },
            onClick_cancel(){
                this.purchaseDialog = false;
                this.resetForm();
            },
            resetForm () {
                this.$refs.form.reset();
                this.productID = null;
                this.productSku = null;
                this.productName = null;
                this.productStocks = null;
                this.purchaseQty = null;
            },
        },
        watch:  {
            purchaseDialog: function(){
                if(this.purchaseDialog == false){
                    this.resetForm();
                }
            },
        }
    }
</script>
<template src="./template.html"></template>