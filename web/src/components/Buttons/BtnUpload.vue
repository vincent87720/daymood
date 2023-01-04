<template>
    <v-btn icon large @click="onClick">
        <v-icon>mdi-upload</v-icon>
    </v-btn>
</template>
<script>
export default {
    name: 'BtnUpload',
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