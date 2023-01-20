<template>
  <v-app>
    <v-app-bar app flat class="appbar d-flex justify-center">
      <div class="pa-3">
        <img src="@/assets/8467417.png" width="120" />
      </div>
    </v-app-bar>
    <v-main>
      <v-window v-model="onboarding">
        <v-window-item key="0">
          <Purchases />
        </v-window-item>
        <v-window-item key="1">
          <Products />
        </v-window-item>
        <v-window-item key="2">
          <Reports />
        </v-window-item>
        <v-window-item key="3">
          <Suppliers />
        </v-window-item>
        <v-window-item key="4">
          <DeliveryOrders />
        </v-window-item>
      </v-window>
    </v-main>
    <v-bottom-navigation v-model="onboarding" dark grow app>
      <v-btn x-large>
        <span>進貨</span>
        <v-icon>mdi-package-down</v-icon>
      </v-btn>
      <v-btn x-large>
        <span>商品管理</span>
        <v-icon>mdi-package-variant-closed</v-icon>
      </v-btn>
      <v-btn x-large>
        <span>報表</span>
        <v-icon>mdi-chart-bar</v-icon>
      </v-btn>
      <v-btn x-large>
        <span>廠商管理</span>
        <v-icon>mdi-store-outline</v-icon>
      </v-btn>
      <v-btn x-large>
        <span>出貨</span>
        <v-icon>mdi-package-up</v-icon>
      </v-btn>
    </v-bottom-navigation>
  </v-app>
</template>

<script>
import Products from './pages/Products/Products.vue'
import Suppliers from './pages/Suppliers/Suppliers.vue'
import Purchases from './pages/Purchases/Purchases.vue'
import DeliveryOrders from './pages/DeliveryOrders/DeliveryOrders.vue'
import ProductDialog from './pages/Products/ProductDialog.vue'
import Reports from './pages/Reports/index.vue'

export default {
  name: 'App',
  components: {
    Products,
    Suppliers,
    Purchases,
    DeliveryOrders,
    ProductDialog,
    Reports,
  },
  data: () => ({
    onboarding: 0,
  }),
  created() {
    this.$store.dispatch("GetSystemConfig");
    this.$store.dispatch("GetTradingSettings");
  },
  methods: {
  },
};
</script>
