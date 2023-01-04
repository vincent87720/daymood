<template>
    <v-btn icon large @click="onClick">
        <v-icon>mdi-cog</v-icon>
    </v-btn>
</template>
<script>
export default {
    name: 'BtnSetting',
    data() {
        return {
        };
    },
    methods: {
        onClick() {
            this.$emit('click', true);
        },
    },
}   
</script>