import Vue from "vue";
import Vuex from "vuex";
import { getSystemConfigs } from "../apis/SystemConfigsAPI";
import { getTradingSettings } from "../apis/TradingSettingsAPI";
import { getProducts } from "../apis/ProductsAPI";
import { getSuppliers } from "../apis/SuppliersAPI";

Vue.use(Vuex);

// 定義一個新的 Vue Store
const store = new Vuex.Store({
  state: {
    systemConfigs: {},
    tradingSettings: {},
    products: [],
    suppliers: [],
  },
  mutations: {
    // 將state設定為參數
    setSystemConfig(state, payload) {
      state.systemConfigs = payload;
    },
    setTradingSettings(state, payload) {
      state.tradingSettings = payload;
    },
    setProducts(state, payload) {
      state.products = payload;
    },
    setSuppliers(state, payload) {
      state.suppliers = payload;
    },
  },
  actions: {
    async GetSystemConfig({ commit }) {
      let systemConfig = {};
      await getSystemConfigs().then((response) => {
        response.data.map(function (x) {
          if (systemConfig[x.type] == undefined) {
            systemConfig[x.type] = [];
          }
          systemConfig[x.type].push({ key: x.key, value: x.value });
        });
      });
      commit("setSystemConfig", systemConfig);
    },
    async GetTradingSettings({ commit }) {
      await getTradingSettings().then((response) => {
        if (response.data.trading != null) {
          commit("setTradingSettings", response.data.trading);
        }
      });
    },
    async GetProducts({ commit }) {
      let products = [];
      await getProducts()
        .then((response) => {
          if (response.data.records != null) {
            response.data.records.map(function (x) {
              products.push({
                key: x.ID,
                value: x.SKU + " " + x.Name,
              });
            });
          }
        })
        .catch((error) => {});
      commit("setProducts", products);
    },
    async GetSuppliers({ commit }) {
      let suppliers = [];
      await getSuppliers()
        .then((response) => {
          if (response.data.records != null) {
            response.data.records.map(function (x) {
              suppliers.push({
                key: x.ID,
                value: x.Name,
              });
            });
          }
        })
        .catch((error) => {
        });
      commit("setSuppliers", suppliers);
    },
  },
});
export default store;
