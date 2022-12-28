<template>
    <v-menu v-model="menu" :close-on-content-click="true" :nudge-right="40" transition="scale-transition" offset-y
        min-width="auto">
        <template v-slot:activator="{ on, attrs }">
            <v-text-field v-model="date" :label="prop_label" prepend-icon="mdi-calendar" readonly
                v-bind="attrs" v-on="on"></v-text-field>
        </template>
        <v-date-picker v-model="date" @input="menu = false"></v-date-picker>
    </v-menu>
</template>

<script>
export default {
    name: 'DatePicker',
    data() {
        return {
            menu: false,
        };
    },
    props: {
        prop_date: {
            type: String,
            required: false
        },
        prop_label: {
            type: String,
            required: false
        },
    },
    computed: {
        date: {
            get() {
                if(this.prop_date){
                    return this.prop_date.substring(0, 10)
                }
                else{
                    return ""
                }
            },
            set(val) {
                this.$emit('update:prop_date', val)

            }
        },
        label: {
            get() {
                return this.prop_label
            },
            set(val) {
                this.$emit('update:prop_label', val)

            }
        },
    },
}   
</script>