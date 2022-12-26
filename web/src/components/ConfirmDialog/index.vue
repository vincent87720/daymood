<script>

    export default {
        name: 'ConfirmDialog',
        components:{
            
        },
        data(){
            return{
            };
        },
        props:{
            prop_confirmDialog:{
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
            prop_deleteTarget:{
                type:Object,
                required:false
            },
        },
        computed:{
            confirmDialog:{
                get(){
                    return this.prop_confirmDialog
                },
                set(val){
                    this.$emit('update:prop_confirmDialog',val)
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
            deleteTarget:{
                get(){
                    return this.prop_deleteTarget
                },
                set(val){
                    this.$emit('update:prop_deleteTarget',val)
                }
            },
        },
        methods:{
            onClick_confirm(){ //有子元件的事件觸發 自定義事件childevent
                this.$emit('confirmClick',this.deleteTarget);//觸發一個在子元件中宣告的事件 childEvnet
                this.resetForm();
            },
            onClick_cancel(){
                this.confirmDialog = false;
                this.resetForm();
            },
            resetForm () {
                this.deleteTarget = null;
            },
        },
        watch:  {
            confirmDialog: function(){
                if(this.confirmDialog == false){
                    this.resetForm();
                }
            },
        }
    }
</script>
<template src="./template.html"></template>