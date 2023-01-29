<template>
    <div>
        <v-tooltip bottom open-delay="500">
            <template v-slot:activator="{ on, attrs }">
                <v-btn icon large v-bind="attrs" v-on="on" @click="onClick_settings">
                    <v-icon>mdi-cog</v-icon>
                </v-btn>
            </template>
            <span>設定</span>
        </v-tooltip>
        <SettingDialog :prop_settingDialog.sync="settingDialog" />
    </div>
</template>
<script>
import SettingDialog from '../../components/SettingDialog/index.vue'
export default {
    name: 'BtnSetting',
    components: {
        SettingDialog,
    },
    data() {
        return {
            settingDialog: false,
        };
    },
    methods: {
        onClick_settings() {
            this.settingDialog = true;
        },
    },
}   
</script>