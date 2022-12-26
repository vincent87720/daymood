<template>
    <v-btn icon large @click="onClick">
        <v-icon>mdi-plus</v-icon>
    </v-btn>
</template>
<script>
    export default {
    name: 'BtnAdd',
    data(){
        return{
        };
    },
    methods: {
        onClick(){
            this.$emit('click',true);
        },
    },
}   
</script>