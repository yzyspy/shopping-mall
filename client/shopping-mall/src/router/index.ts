import {createRouter, createWebHashHistory} from "vue-router";


const routes = [
    {
        path: '/',
        name: 'Login',
        meta: {title:"admin"},
        component: () => import('../views/Login.vue')

    }
]

const router = createRouter({
    history: createWebHashHistory(),
    routes: routes
})

export default router