import Vue from 'vue';
import Vuex from 'vuex';
import { getSystemConfigs } from "../apis/SystemConfigsAPI";

Vue.use(Vuex);

// 定義一個新的 Vue Store
const store = new Vuex.Store({
    state: {
      systemConfigs: {},
    },
    mutations: {
      // 將state設定為參數
      setSystemConfig(state, payload){
        state.systemConfigs = payload;
      }
    },
    actions:{
        async GetSystemConfig({ commit }){
            let systemConfig = {};
            await getSystemConfigs()
            .then((response)=>{
                response.data.map(function(x){
                    if(systemConfig[x.type] == undefined){
                        systemConfig[x.type] = [];
                    }
                    systemConfig[x.type].push({key: x.key, value: x.value});
                });
            });
            commit('setSystemConfig', systemConfig);
        },
    }

})
export default store;