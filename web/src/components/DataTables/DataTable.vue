<template>
    <v-data-table :headers="headers" :items="items" :items-per-page="7" :search="search" dark
        class="elevation-1 mx-4 mb-4">
        <template v-slot:item.actions="{ item }">
            <v-icon small class="mr-2" @click.stop="onClick_editButton(item)">
                mdi-pencil
            </v-icon>
            <v-icon small @click.stop="onClick_deleteButton(item)">
                mdi-delete
            </v-icon>
        </template>
    </v-data-table>
</template>

<script>
export default {
    name: 'DataTable',
    data() {
        return {
        };
    },
    props: {
        prop_headers: {
            type: Array,
            required: false
        },
        prop_items: {
            type: Array,
            required: false
        },
        prop_search: {
            type: String,
            required: false
        },
    },
    computed: {
        headers: {
            get() {
                return this.prop_headers
            },
            set(val) {
                this.$emit('update:prop_headers', val)

            }
        },
        items: {
            get() {
                return this.prop_items
            },
            set(val) {
                this.$emit('update:prop_items', val)

            }
        },
        search: {
            get() {
                return this.prop_search
            },
            set(val) {
                this.$emit('update:prop_search', val)

            }
        },
    },
    methods: {
        onClick_editButton(item) {
            this.$emit('edit', item);
        },
        onClick_deleteButton(item) {
            this.$emit('delete', item);
        },
    },
}   
</script>