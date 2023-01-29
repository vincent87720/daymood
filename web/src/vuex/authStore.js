const authStore = {
  state: {
    isLogin: false,
  },
  mutations: {
    // 將state設定為參數
    setLoginStatus(state, payload) {
      state.isLogin = payload;
    },
  },
  actions: {
    async SetLoginStatus({ commit },isLogin) {
      commit("setLoginStatus", isLogin);
    },
  },
};
export default authStore;
