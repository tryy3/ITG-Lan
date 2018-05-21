import Vue from 'vue'
import Router from 'vue-router'
import auth from '../auth'

// Containers
import Full from '@/containers/Full'

// Views
import Home from '@/views/Home'
import Admin from '@/views/Admin'

// Views - Pages
import Page404 from '@/views/pages/Page404'
import Login from '@/views/pages/Login'
import Register from '@/views/pages/Register'

Vue.use(Router)

async function requireAuth (to, from, next) {
    try {
        await auth.isAuthenticated()
        next()
    } catch (e) {
        next('/login')
    }
}

export default new Router({
    mode: 'history', // Demo is living in GitHub.io, so required!
    linkActiveClass: 'open active',
    scrollBehavior: () => ({ y: 0 }),
    routes: [
        {
            path: '/',
            redirect: '/home',
            name: 'Index',
            component: Full,
            children: [
                {
                    path: 'home',
                    name: 'Home',
                    component: Home
                },
                {
                    path: 'admin',
                    name: 'Admin',
                    component: Admin,
                    beforeEnter: requireAuth
                }
            ]
        },
        {
            path: '/login',
            name: 'Login',
            component: Login
        },
        {
            path: '/register',
            name: 'Register',
            component: Register
        },
        {
            path: '/404',
            name: '404',
            component: Page404
        },
        {
            path: '*',
            redirect: '/404',
            name: 'NotFound'
        }
    ]
})
