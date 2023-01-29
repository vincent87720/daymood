import Vue from "vue";
import Vuex from "vuex";
import authStore from "@/vuex/authStore";
import dataStore from "@/vuex/dataStore";
import configStore from "@/vuex/configStore";
import createPersistedState from "vuex-persistedstate";

Vue.use(Vuex);

const store = new Vuex.Store({
  modules: {
    auth: authStore,
    data: dataStore,
    conf: configStore,
  },
  plugins: [
    createPersistedState({
      reducer(val) {
        return {
          auth: val.auth,
        };
      },
    }),
  ],
});

export default store;
