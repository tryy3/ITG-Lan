// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import router from './router/index'
import VueResource from 'vue-resource'
import VueNotifications from 'vue-notifications'
import { options as VueNotificationsOptions } from './notifications'

Vue.use(VueResource)
Vue.use(VueNotifications, VueNotificationsOptions)

/* eslint-disable no-new */
new Vue({
    el: '#app',
    router,
    template: '<App/>',
    components: {
        App
    }
})
