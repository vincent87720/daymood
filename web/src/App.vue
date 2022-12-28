<template>
  <v-app>
    <v-app-bar
    app
    flat
    class="appbar d-flex justify-center"
    >
      <div class="pa-3">
        <img src="@/assets/8467417.png" width="120"/>
      </div>
    </v-app-bar>
    <v-main>
      <!-- <ProductsOld
      v-if="activeType != 2 && activeType != 0&& activeType != 1"
      ref="ProductsComponent"
      :prop_activeType.sync="activeType"
      /> -->
      <Products
      v-if="activeType == 1"
      :prop_activeType.sync="activeType"
      />
      <Purchases
      ref="Purchases"
      v-if="activeType == 0"
      />
      <Suppliers 
      v-if="activeType == 2"
      :prop_activeType.sync="activeType"
      />
    </v-main>
    <v-bottom-navigation 
    v-model="activeType"
    dark
    grow
    app
    >
      <v-btn
      x-large
      activeType="productPurchase"
      >
        <span>進貨</span>

        <v-icon>mdi-package-down</v-icon>
      </v-btn>

      <v-btn
      x-large
      activeType="productManagement"
      >
        <span>商品管理</span>

        <v-icon>mdi-package-variant-closed</v-icon>
      </v-btn>

      <v-btn
      x-large
      activeType="firmManagement"
      >
        <span>廠商管理</span>

        <v-icon>mdi-store-outline</v-icon>
      </v-btn>

      <v-btn
      x-large
      activeType="productDelivery"
      >
        <span>出貨</span>

        <v-icon>mdi-package-up</v-icon>
      </v-btn>
    </v-bottom-navigation>
  </v-app>
</template>

<script>
import ProductsOld from './components/ProductsOld/index.vue'
import Products from './components/Products/index.vue'
import Suppliers from './components/Suppliers/index.vue'
import Purchases from './components/Purchases/index.vue'
import ProductDialog from './components/ProductDialog/index.vue'

export default {
  name: 'App',
  components:{
    Products,
    ProductsOld,
    Suppliers,
    Purchases,
    ProductDialog,
  },
  data: () => ({
    activeType: -1,
  }),
  created(){
    this.$store.dispatch("GetSystemConfig");
  },
  methods:{
  },
};
</script>
