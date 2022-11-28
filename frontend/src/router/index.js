import { createRouter, createWebHashHistory } from "vue-router";

const router =createRouter({
    history: createWebHashHistory(),
    routes: [{
        name: 'login',
        path: '/',
        component: () => import('../components/Login.vue')
    },
        {
            name: 'register',
            path: '/register',
            component: () => import('../components/HelloWorld.vue')
        }
    ]
})

export  {
    router
}
