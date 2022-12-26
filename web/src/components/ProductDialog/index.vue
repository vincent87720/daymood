<script>

    export default {
        name: 'ProductDialog',
        components:{
            
        },
        data(){
            return{
                imgFile: null,
                fooFile: null,
                valid_productForm: false,
                // select_type: null,
                typeList:[
                    { name: '耳環', val: 'earrings' },
                    { name: '耳骨夾', val: 'earcuff' },
                    { name: '項鍊', val: 'necklace' },
                    { name: '戒指', val: 'rings' },
                    { name: '手鍊', val: 'bracelet' },
                ],
                text_requiredRules: [
                    v => !!v || '必填',
                ],
                text_requiredRules_isNumber: [
                    v => !!v || '必填',
                    v => {
                        return !!parseFloat(v) || '必須為數字'
                    },
                ],
                grossMargin: null,//毛利
                profitMargin: null,//毛利率
            };
        },
        mounted(){
        },
        props:{
            prop_productDialog:{
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
            prop_actionType:{
                type:String,
                required:false
            },
            prop_productID:{
                type:Number,
                required:false
            },
            prop_productSku:{
                type:String,
                required:false
            },
            prop_purchaseProductID:{
                type:String,
                required:false
            },
            prop_productName:{
                type:String,
                required:false
            },
            prop_productType:{
                type:String,
                required:false
            },
            prop_productImgName:{
                type:String,
                required:false
            },
            prop_productImgID:{
                type:String,
                required:false
            },
            prop_productWeight:{
                type:String,
                required:false
            },
            prop_krwWholesalePrice:{
                type:String,
                required:false
            },
            prop_ntdWholesalePrice:{
                type:String,
                required:false
            },
            prop_ntdListPrice:{
                type:String,
                required:false
            },
            prop_ntdSellingPrice:{
                type:String,
                required:false
            },
            prop_ntdCnf:{
                type:String,
                required:false
            },
            prop_ntdCost:{
                type:String,
                required:false
            },
            prop_firmID:{
                type:Number,
                required:false
            },
            prop_firmList:{
                type:Array,
                required:false
            },
            prop_tradingSettings:{
                type:Object,
                required:true
            },
        },
        computed:{
            productDialog:{
                get(){
                    return this.prop_productDialog
                },
                set(val){
                    this.$emit('update:prop_productDialog',val)
                    
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
            actionType:{
                get(){
                    return this.prop_actionType
                },
                set(val){
                    this.$emit('update:prop_actionType',val)
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
            purchaseProductID:{
                get(){
                    return this.prop_purchaseProductID
                },
                set(val){
                    this.$emit('update:prop_purchaseProductID',val)
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
            productType:{
                get(){
                    return this.prop_productType
                },
                set(val){
                    this.$emit('update:prop_productType',val)
                }
            },
            productImgName:{
                get(){
                    return this.prop_productImgName
                },
                set(val){
                    this.$emit('update:prop_productImgName',val)
                }
            },
            productImgID:{
                get(){
                    return this.prop_productImgID
                },
                set(val){
                    this.$emit('update:prop_productImgID',val)
                }
            },
            productWeight:{
                get(){
                    return this.prop_productWeight
                },
                set(val){
                    this.$emit('update:prop_productWeight',val)
                }
            },
            krwWholesalePrice:{
                get(){
                    return this.prop_krwWholesalePrice
                },
                set(val){
                    this.$emit('update:prop_krwWholesalePrice',val)
                }
            },
            ntdWholesalePrice:{
                get(){
                    return this.prop_ntdWholesalePrice
                },
                set(val){
                    this.$emit('update:prop_ntdWholesalePrice',val)
                }
            },
            ntdListPrice:{
                get(){
                    return this.prop_ntdListPrice
                },
                set(val){
                    this.$emit('update:prop_ntdListPrice',val)
                }
            },
            ntdSellingPrice:{
                get(){
                    return this.prop_ntdSellingPrice
                },
                set(val){
                    this.$emit('update:prop_ntdSellingPrice',val)
                }
            },
            ntdCnf:{
                get(){
                    return this.prop_ntdCnf
                },
                set(val){
                    this.$emit('update:prop_ntdCnf',val)
                }
            },
            ntdCost:{
                get(){
                    return this.prop_ntdCost
                },
                set(val){
                    this.$emit('update:prop_ntdCost',val)
                }
            },
            firmID:{
                get(){
                    return this.prop_firmID
                },
                set(val){
                    this.$emit('update:prop_firmID',val)
                }
            },
            firmList:{
                get(){
                    return this.prop_firmList
                },
                set(val){
                    this.$emit('update:prop_firmList',val)
                }
            },
            tradingSettings:{
                get(){
                    return this.prop_tradingSettings
                },
                set(val){
                    this.$emit('update:prop_tradingSettings',val)
                }
            },
        },
        methods:{
            onClick_confirm:function(){ //有子元件的事件觸發 自定義事件childevent
                this.$refs.form.validate();
                if(this.valid_productForm == false){
                    return
                }
                var productInfo = {};
                productInfo.actionType = this.actionType;
                productInfo.productID = this.productID;
                productInfo.productSku = this.productSku;
                productInfo.purchaseProductID = this.purchaseProductID;
                productInfo.productName = this.productName;
                productInfo.productType = this.productType;
                productInfo.productWeight = this.productWeight;
                productInfo.krwWholesalePrice = this.krwWholesalePrice;
                productInfo.ntdWholesalePrice = this.ntdWholesalePrice;
                productInfo.ntdListPrice = this.ntdListPrice;
                productInfo.ntdSellingPrice = this.ntdSellingPrice;
                productInfo.ntdCnf = this.ntdCnf;
                productInfo.ntdCost = this.ntdCost;
                productInfo.firmID = this.firmID;
                productInfo.fooFile = this.fooFile;
                productInfo.imgFile = this.imgFile;
                productInfo.productImgName = this.productImgName;
                productInfo.productImgID = this.productImgID;
                console.log(this.fooFile);
                console.log(this.imgFile);
                this.$emit('confirmClick',productInfo);//觸發一個在子元件中宣告的事件 childEvnet
                this.resetForm();
            },
            onClick_cancel(){
                this.productDialog = false;
                this.resetForm();
            },
            resetForm () {
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
            clearFileInput(){
                this.$refs.imgFile.reset();
                this.imgFile = null;
            },
            calc_ntdWholesalePrice(){
                return (this.krwWholesalePrice/this.tradingSettings.ExchangeRate).toFixed(2).toString();
            },
            calc_listPrice(){
                var tempPrice = ((this.krwWholesalePrice/this.tradingSettings.ExchangeRate)*this.tradingSettings.Markup).toFixed();
                if(tempPrice%10 == 0){
                    return (tempPrice).toString()
                }
                else if(tempPrice%10 > 1){
                    return ((tempPrice-tempPrice%10)+10).toString()
                }
                else{
                    return (tempPrice-tempPrice%10).toString()
                }
            },
            calc_param(){
                var param = {};

                var gram = this.productWeight/1000;
                param.ajeossiCalc = this.tradingSettings.Ajeossi/100+1;//大哥抽成
                param.shippingFeeCalc = gram*this.tradingSettings.ShippingFee;
                param.tariffCalc = gram*this.tradingSettings.Tariff;
                return param
            },
            calc_cnfCost(){
                var param = this.calc_param();
                return ((this.krwWholesalePrice * param.ajeossiCalc + param.shippingFeeCalc)/this.tradingSettings.ExchangeRate + param.tariffCalc).toFixed(2).toString()
            },
            calc_cost(){
                var param = this.calc_param();
                return ((this.krwWholesalePrice * param.ajeossiCalc + param.shippingFeeCalc)/this.tradingSettings.ExchangeRate + param.tariffCalc+25).toFixed(2).toString()
            },
            calc_grossMargin(){
                this.grossMargin = this.ntdSellingPrice-this.ntdCost;
                let percent = ((this.ntdSellingPrice-this.ntdCost)/this.ntdSellingPrice).toFixed(2).toString();

                if(isNaN(percent) == false){
                    this.profitMargin = percent;
                }
                else{
                    this.profitMargin = 0;
                }
            }
        },
        watch:  {
            productDialog: function(){
                if(this.productDialog == false){
                    this.resetForm();
                }
                else{
                    if(this.actionType == "put" && this.productImgName){
                        const fooFile = new File(["foo"], this.productImgName, {
                            type: "text/plain",
                            name: this.productImgName,
                        });
                        this.imgFile = fooFile;
                        this.fooFile = fooFile;
                    }
                }
            },
            ntdSellingPrice: function(){
                this.calc_grossMargin();
            },
            ntdCost: function(){
                this.calc_grossMargin();
            },
        }
    }
</script>
<style scoped src="./style.css"></style>
<template src="./template.html"></template>