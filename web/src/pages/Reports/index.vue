<template>
    <div>
        <v-container fluid class="d-none d-lg-block">
                <v-row class="ma-0">
                    <v-col class="pa-0 d-flex justify-end">
                        <c-btn-setting></c-btn-setting>
                        <c-btn-logout></c-btn-logout>
                    </v-col>
                </v-row>
            </v-container>
        <v-container class="d-flex align-center justify-center">
            <v-row class="pa-3">
                <v-col cols="12" sm="4">
                    <c-card-rounded class="pa-5 d-flex flex-column align-center">
                        <h1>總支出</h1>
                        <h2>$ {{ balance.PurchaseTotal }}</h2>
                    </c-card-rounded>
    
                </v-col>
                <v-col cols="12" sm="4">
                    <c-card-rounded class="pa-5 d-flex flex-column align-center">
                        <h1>收支差</h1>
                        <h2 :class="[isNegative ? 'text-color-red' : 'text-color-green']">$ {{ calcBalance(balance.PurchaseTotal,balance.DeliveryTotal) }}</h2>
                    </c-card-rounded>
    
                </v-col>
                <v-col cols="12" sm="4">
                    <c-card-rounded class="pa-5 d-flex flex-column align-center">
                        <h1>總收入</h1>
                        <h2>$ {{ balance.DeliveryTotal }}</h2>
                    </c-card-rounded>
    
                </v-col>
            </v-row>
        </v-container>
    </div>
</template>

<script>
import BtnSetting from "@/components/Buttons/BtnSetting.vue";
import BtnLogout from "@/components/Buttons/BtnLogout.vue";
import CardRounded from "@/components/Cards/CardRounded.vue";
import { getBalancesReports } from "@/apis/ReportsAPI";


class Balance {
    PurchaseTotal = undefined;
    DeliveryTotal = undefined;
}

export default {
    name: 'Reports',
    components: {
        "c-btn-setting": BtnSetting,
        "c-btn-logout": BtnLogout,
        "c-card-rounded": CardRounded,
    },
    data() {
        return {
            balance: new Balance(),
        };
    },
    async mounted() {
        await this.getBalancesReports();
    },
    props: {
    },
    computed: {
        isNegative(){
            if(this.calcBalance(this.balance.PurchaseTotal,this.balance.DeliveryTotal) < 0){
                return true;
            }
            return false;
        }
    },
    methods: {
        calcBalance(purchaseTotal, deliveryTotal){
            return parseFloat(deliveryTotal) - parseFloat(purchaseTotal);
        },
        async getBalancesReports() {
            await getBalancesReports()
                .then((response) => {
                    if (response.data.records != null) {
                        this.balance = response.data.records[0];
                    }
                    else {
                        this.balance = {};
                    }
                })
                .catch((error) => {
                });
        },
    },
}
</script>

<style scoped>
.text-color-red{
    color: #965455;
}
.text-color-green{
    color: #849A8F;
}
</style>