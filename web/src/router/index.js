import Vue from "vue";
import VueRouter from "vue-router";
import store from "@/vuex/store";

Vue.use(VueRouter);

const routes = [
  {
    path: "/login",
    name: "Login",
    component: () => import("@/pages/Authentication/Login.vue"),
  },
  {
    path: "/main",
    name: "Main",
    component: () => import("@/pages/Main/Main.vue"),
    meta: { requireAuth: true },
  },
];

const router = new VueRouter({
  mode: "history",
  base: process.env.BASE_URL,
  routes,
});

router.beforeEach(async (to, from, next) => {
  // console.log(from)

  if (to.meta.requireAuth) {
    if (store.state.auth.isLogin == true) {
      next();
    } else {
      next({ name: "Login" });
    }
  } else {
    if(to.name == "Login" && store.state.auth.isLogin == true){
      next({ name: "Main" });
    }
    else{
      next();
    }
  }
});

export default router;
