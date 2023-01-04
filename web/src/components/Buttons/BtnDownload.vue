<template>
    <v-btn icon large @click="onClick">
        <v-icon>mdi-download</v-icon>
    </v-btn>
</template>
<script>
export default {
    name: 'BtnDownload',
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