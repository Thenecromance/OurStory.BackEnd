import { createRouter, createWebHistory } from "vue-router";
import Dashboard from "../views/Dashboard.vue";
import Tables from "../views/Tables.vue";
import Billing from "../views/Billing.vue";
import Profile from "../views/Profile.vue";
import Signup from "../views/Signup.vue";
import Signin from "../views/Signin.vue";

import Home from "../ourStories/Home.vue";
import Login from "../ourStories/signin/Signin.vue";
import Register from "../ourStories/signup/signup.vue";
const routes = [
  {
    path: "/",
    name: "/",
    component: Home,
  },
  {
    path: "/dashboard-default",
    name: "Dashboard",
    component: Dashboard,
  },
  {
    path: "/tables",
    name: "Tables",
    component: Tables,
  },
  {
    path: "/billing",
    name: "Billing",
    component: Billing,
  },
  {
    path: "/profile",
    name: "Profile",
    component: Profile,
  },
  {
    path: "/signin",
    name: "Signin",
    component: Signin,
  },
  {
    path: "/signup",
    name: "Signup",
    component: Signup,
  },
  {
    path:"/login",
    name:"Login",
    component:Login,
  },
  {
    path:"/register",
    name:"Register",
    component:Register,
  },
  {
    path:"/:pathMatch(.*)",
    name:"404",
    // redirect:"/404",
    component:()=>import("../ourStories/404/404.vue"),
    hidden:true,
  },
  {
    path:"/test",
    name :"Test",
    component:()=>import("../ourStories/TestComponent/Test.vue"),
  }
];

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes,
  linkActiveClass: "active",
});

export default router;
