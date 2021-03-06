import Vue from 'vue'
import Router from 'vue-router'
import Cookie from 'vue-cookie'
import Dashboard from '@/components/dashboard/Layout'
import Doc from '@/components/dashboard/Document'
import AuthLayount from '@/components/auth/Login'

Vue.use(Router);
Vue.use(Cookie);

const routes =  [
    {
        path: '/',
        name: 'Dashboard',
        component: Dashboard,
        children: [
            {
                path: '/doc/:fid/:did',
                component: Doc
            },
        ],
        meta: {
            requireAuth: true
        }
    },
    {
        path: '/login',
        name: 'Auth',
        component: AuthLayount,
    }
]

const router = new Router({
    routes,
    // mode: 'history'
});

router.beforeEach((to, from, next) => {
    const authObj = getAuth()
    const now = Math.round((new Date()).getTime() / 1000);
    if (to.matched.some(r => r.meta.requireAuth)) {
        if (authObj == null || parseInt(authObj.expired_at) < now) {
            next({
                path: "/login",
                query: {redirect: to.fullPath}
            })
        } else {
            next()
        }
    } else {
        next()
    }
});

function getAuth() {
    const auth = Vue.cookie.get('rest-doc')
    const authObj = JSON.parse(auth)
    return authObj
}

export default router
