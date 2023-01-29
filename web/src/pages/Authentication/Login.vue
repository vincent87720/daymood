<template>
    <v-app>
        <v-progress-linear :indeterminate="loading" :color="progressLinearColor"></v-progress-linear>
        <div class="c-login-page d-flex align-center justify-center">
            <v-container fluid>
                <v-row class="d-flex justify-center">
                    <v-col col="12" class="d-flex justify-center c-bg-color-black">
                        <v-card class="pa-8 d-flex flex-column justify-center rounded-xl" elevation="5" max-width="340">
                            <div class="d-flex justify-center mb-6">
                                <img src="@/assets/8467417.png" width="240" />
                            </div>
                            <v-text-field v-model="user.Username" color="#796956" label="帳號" filled outlined
                                dense v-on:keydown.enter.prevent="onClick_login()"></v-text-field>
                            <v-text-field v-model="user.Password" color="#796956" label="密碼" filled outlined dense
                                type="password" v-on:keydown.enter.prevent="onClick_login()"></v-text-field>
                            <v-btn block outlined color="#796956" class="c-btn-login mt-2"
                                @click="onClick_login()">登入</v-btn>
                        </v-card>
                    </v-col>
                </v-row>
            </v-container>
            <Alert :prop_alert.sync="alert" :prop_alertType="alertType" :prop_alertText="alertText"></Alert>
        </div>
    </v-app>
</template>

<script>
import Alert from "@/components/Alert/index.vue";
import CardRounded from "@/components/Cards/CardRounded.vue";
import { login } from "@/apis/AuthenticationAPI";

class User {
    ID = undefined;
    Username = "";
    Password = "";
    Name = "";
    Email = "";
    CreateAt = "";
    UpdateAt = "";
}

export default {
    name: 'Authentication',
    components: {
        Alert,
        "c-card-rounded": CardRounded,
    },
    data() {
        return {
            //Alert
            alert: false,
            alertType: "",
            alertText: "",

            loading: false,

            user: new User(),
        };
    },
    mounted() {
    },
    props: {
    },
    computed: {
        progressLinearColor(){
            if(this.loading == true){
                return "#796956";
            }
            return "#121212";
        }
    },
    methods: {
        async onClick_login() {
            await this.login(this.user);
        },
        async login(item) {
            this.loading = true;
            await login(item)
                .then(async (response) => {
                    await this.$store.dispatch("SetLoginStatus",response.data.isLogin);
                    this.$router.push({ name: 'Main' }).catch(() => { })
                })
                .catch((error) => {
                    // console.log(error)
                    this.alert = true;
                    this.alertType = "Fail";
                    this.alertText = "登入失敗";
                })
                .finally(() => {
                    this.loading = false;
                });
        },
    },
    watch: {
    }
}
</script>

<style scoped>
.c-login-page {
    height: 100%;
}

.c-bg-color-black {
    background-color: #121212;
}

.c-btn-login {
    color: #796956;
    background-color: #E2DDD3;
}
</style>