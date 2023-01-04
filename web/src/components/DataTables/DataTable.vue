<template>
    <v-data-table :headers="headers" :items="items" :items-per-page="7" :search="search" dark
        class="page__table">
        <template v-slot:item.actions="{ item }">
            <v-icon small class="mx-1" @click.stop="onClick_editButton(item)">
                mdi-pencil
            </v-icon>
            <v-icon small class="mx-1" @click.stop="onClick_deleteButton(item)">
                mdi-delete
            </v-icon>
            <slot name="item.actions.plus" v-bind="{ item }"></slot>
        </template>
        <!-- Code to pass on the $slots: -->
        <!-- <template v-for="(index, name) in $slots" v-slot:[name]>
            <slot :name="name" />
        </template> -->
        <!-- Code to pass on the $slots: -->
        <!-- Code to pass on the $scopedSlots -->
        <template v-for="(index, name) in $scopedSlots" v-slot:[name]="data">
            <slot :name="name" v-bind="data"></slot>
        </template>
        <!-- Code to pass on the $scopedSlots -->
        <!-- <template v-slot:body="data">
            <draggable :list="data.items" tag="tbody">
                <tr v-for="(item, index) in data.items" :key="`row.${index}`">
                    <td v-if="headers.find(x => x.value == idx)" v-for="(col, idx) in item" :key="`col.${idx}`">
                        <slot :name="`item.${idx}`" v-bind="{'item':item}">{{ col }}</slot>
                    </td>
                    <td>
                        <v-icon small class="mr-2" @click.stop="onClick_editButton(item)">
                            mdi-pencil
                        </v-icon>
                        <v-icon small @click.stop="onClick_deleteButton(item)">
                            mdi-delete
                        </v-icon>
                    </td>
                </tr>
            </draggable>
        </template> -->
    </v-data-table>
</template>

<script>
import Draggable from 'vuedraggable';
export default {
    name: 'DataTable',
    data() {
        return {
        };
    },
    components: {
        Draggable,
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

<style lang="scss">
.page--table {
    .page {
        &__table {
            margin-top: 20px;
        }

        &__grab-icon {
            cursor: move;
        }
    }
}
</style>