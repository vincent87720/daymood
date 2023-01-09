<template>
    <div>
        <v-btn icon large @click="onClick_settings">
            <v-icon>mdi-cog</v-icon>
        </v-btn>
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