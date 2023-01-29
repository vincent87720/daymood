<template>
    <div>
        <v-tooltip bottom open-delay="500">
            <template v-slot:activator="{ on, attrs }">
                <v-btn icon large v-bind="attrs" v-on="on" @click="onClick()">
                    <v-icon>mdi-logout</v-icon>
                </v-btn>
            </template>
            <span>登出</span>
        </v-tooltip>
        <Alert :prop_alert.sync="alert" :prop_alertType="alertType" :prop_alertText="alertText"></Alert>
    </div>
</template>
<script>
import Alert from "@/components/Alert/index.vue";
import { logout } from "@/apis/AuthenticationAPI"
export default {
    name: 'BtnLogout',
    components: {
        Alert,
    },
    data() {
        return {
            //Alert
            alert: false,
            alertType: "",
            alertText: "",
        };
    },
    methods: {
        async onClick() {
            await logout()
                .then(async (response) => {
                    await this.$store.dispatch("SetLoginStatus", response.data.isLogin);
                    this.$router.push({ name: 'Login' }).catch(() => { });
                })
                .catch(() => {
                    this.alert = true;
                    this.alertType = "Fail";
                    this.alertText = "登出功能異常";
                });
        },
    },
}   
</script>