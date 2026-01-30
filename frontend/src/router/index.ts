import { createRouter, createWebHistory } from "vue-router";

const routes = [
  {
    path: "/",
    name: "Home",
    component: () => import("../pages/Home.vue"),
  },
  {
    path: "/ai-writing",
    name: "AIWriting",
    component: () => import("../pages/AISciWriting.vue"),
  },
  {
    path: "/documents",
    name: "Documents",
    component: () => import("../pages/Documents.vue"),
  },
  {
    path: "/login",
    name: "Login",
    component: () => import("../pages/Login.vue"),
  },
  {
    path: "/register",
    name: "Register",
    component: () => import("../pages/Register.vue"),
  },
  {
    path: "/subscription",
    name: "Subscription",
    component: () => import("../pages/SubscriptionPage.vue"),
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
