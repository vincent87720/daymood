<script>

    export default {
        name: 'ProductInfoDialog',
        components:{
            
        },
        data(){
            return{
                valid_productForm: false,
                select_type: null,
                typeList:[
                    { name: '耳環', val: 'earrings' },
                    { name: '耳骨夾', val: 'earcuff' },
                    { name: '項鍊', val: 'necklace' },
                    { name: '戒指', val: 'rings' },
                    { name: '手鍊', val: 'bracelet' },
                ],
                tradingSettings: null,
                text_requiredRules: [
                    v => !!v || '必填',
                ],
                text_requiredRules_isNumber: [
                    v => !!v || '必填',
                    v => {
                        return !!parseFloat(v) || '必須為數字'
                    },
                ],
                firmName: null,
                imgSrc: null,

                ajeossi: 0,
                shippingFee: 0,
                exchangeRate: 0,
                tariff: 0,
                markup: 0
            };
        },
        mounted(){
        },
        props:{
            prop_productInfoDialog:{
                type:Boolean,
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
        },
        computed:{
            productInfoDialog:{
                get(){
                    return this.prop_productInfoDialog
                },
                set(val){
                    this.$emit('update:prop_productInfoDialog',val)
                    
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
        },
        methods:{
            onClick_cancel(){
                this.productInfoDialog = false;
                this.resetForm();
            },
            resetForm () {
                this.productID = null;
                this.productSku = null;
                this.purchaseProductID = null;
                this.productName = null;
                this.productType = null;
                this.productImgID = null;
                this.productWeight = null;
                this.krwWholesalePrice = null;
                this.ntdWholesalePrice = null;
                this.ntdListPrice = null;
                this.ntdSellingPrice = null;
                this.ntdCnf = null;
                this.ntdCost = null;
                this.imgSrc = null;
            },
            calc_listPrice(){
                var tempPrice = ((this.krwWholesalePrice/this.exchangeRate)*this.markup).toFixed();
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
            calc_cnfCost(){
                var gram = this.productWeight/1000;
                var ajeossiCalc = this.ajeossi/100+1;//大哥抽成
                var shippingFeeCalc = gram*this.shippingFee;
                var tariffCalc = gram*this.tariff;
                return ((this.krwWholesalePrice*ajeossiCalc+shippingFeeCalc)/this.exchangeRate+tariffCalc).toFixed(2).toString()
            },
            calc_cost(){
                var gram = this.productWeight/1000;
                var ajeossiCalc = this.ajeossi/100+1;//大哥抽成
                var shippingFeeCalc = gram*this.shippingFee;
                var tariffCalc = gram*this.tariff;
                return ((this.krwWholesalePrice*ajeossiCalc+shippingFeeCalc)/this.exchangeRate+tariffCalc+25).toFixed(2).toString()
            },
            getProductType(){
                this.typeList.forEach((element) => {
                    if (element.val == this.productType){
                        this.select_type = element.name;
                    }
                });
            },
            getFirmName(){
                this.firmList.forEach((element) => {
                    if (element.ID == this.firmID){
                        this.firmName =  element.Name;
                    }
                });
            },
        },
        watch:  {
            productInfoDialog: function(){
                if(this.productInfoDialog == false){
                    this.resetForm();
                }
                else{
                    this.getProductType();
                    this.getFirmName();
                }
            },
            productImgID: function(){
                if(this.productImgID){
                    this.imgSrc = process.env.VUE_APP_BASE_URL+"products/images/"+this.productImgID;
                }
            }
        }
    }
</script>
<template src="./template.html"></template>