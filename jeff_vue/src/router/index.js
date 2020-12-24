import Vue from 'vue'
import VueRouter from 'vue-router'
import shop from "../views/shop";
import shopinfo from "../views/shopinfo";

Vue.use(VueRouter)

const routes = [
    {
        path: '/home',
        name: 'home',
        // route level code-splitting
        // this generates a separate chunk (about.[hash].js) for this route
        // which is lazy-loaded when the route is visited.
        component: () => import(/* webpackChunkName: "about" */ '../views/home.vue')
    },
    {
        path: '/',
        redirect:'/home'
    },
    {
        path: '/:tag/:page',
        name: 'shop',
        component: shop,
    },
    {
        path: '/:tag',
        redirect:'/:tag/1'
    },
    {
        path: '/shop/info/:pk',
        name: 'shopinfo',
        component: shopinfo,
    }
]

const router = new VueRouter({
    mode: 'history',
    base: process.env.BASE_URL,
    routes
})

export default router
