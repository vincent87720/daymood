import { getProducts } from "@/apis/ProductsAPI";
import { getSuppliers } from "@/apis/SuppliersAPI";

const dataStore = {
  state: {
    products: [],
    allProducts: [],
    suppliers: [],
    allSuppliers: [],
  },
  mutations: {
    setProducts(state, payload) {
      state.products = payload;
    },
    setAllProducts(state, payload) {
      state.allProducts = payload;
    },
    setSuppliers(state, payload) {
      state.suppliers = payload;
    },
    setAllSuppliers(state, payload) {
      state.allSuppliers = payload;
    },
  },
  actions: {
    async GetProducts({ commit }) {
      let products = [];
      let allProducts = [];
      await getProducts()
        .then((response) => {
          if (response.data.records != null) {
            response.data.records.map(function (x) {
              allProducts.push({
                key: x.ID,
                value: x.Name,
                SKU: x.SKU,
              });
              if (x.DataStatus == 1) {
                products.push({
                  key: x.ID,
                  value: x.SKU + " " + x.Name,
                  RetailPrice: x.RetailPrice,
                });
              }
            });
          }
        })
        .catch((error) => { });
      commit("setProducts", products);
      commit("setAllProducts", allProducts);
    },
    async GetSuppliers({ commit }) {
      let suppliers = [];
      let allSuppliers = [];
      await getSuppliers()
        .then((response) => {
          if (response.data.records != null) {
            response.data.records.map(function (x) {
              allSuppliers.push({
                key: x.ID,
                value: x.Name,
              });
              if (x.DataStatus == 1) {
                suppliers.push({
                  key: x.ID,
                  value: x.Name,
                });
              }
            });
          }
        })
        .catch((error) => { });
      commit("setSuppliers", suppliers);
      commit("setAllSuppliers", allSuppliers);
    },
  },
};
export default dataStore;
