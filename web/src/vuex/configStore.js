import { getSystemConfigs } from "@/apis/SystemConfigsAPI";
import { getTradingSettings } from "@/apis/TradingSettingsAPI";

const configStore = {
  state: {
    systemConfigs: {},
    tradingSettings: {},
  },
  mutations: {
    // 將state設定為參數
    setSystemConfig(state, payload) {
      state.systemConfigs = payload;
    },
    setTradingSettings(state, payload) {
      state.tradingSettings = payload;
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
  },
};
export default configStore;
