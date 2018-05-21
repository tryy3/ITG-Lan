import Vue from 'vue'
import utils from './utils.js'
import auth from '../auth/index'

export default {
    createSignup (payload) {
        return new Promise((resolve, reject) => {
            Vue.http.post(utils.CREATE_SIGNUP_URL, payload).then(response => {
                if (response.body.success) resolve(response.body)
                else reject(response.body.message)
            }, response => {
                reject(response.statusText)
            })
        })
    },
    findSignup () {
        return new Promise((resolve, reject) => {
            Vue.http.post(utils.FIND_SIGNUP_URL, {}, { headers: auth.getAuthHeader() }).then(response => {
                if (response.body.success === false) reject(response.body.message)
                else resolve(response.body)
            }, response => {
                reject(response.statusText)
            })
        })
    }
}
