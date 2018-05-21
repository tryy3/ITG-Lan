import Vue from 'vue'
import utils from './utils.js'
import auth from '../auth/index'

export default {
    createFAQ (payload) {
        return new Promise((resolve, reject) => {
            Vue.http.post(utils.CREATE_FAQ_URL, payload).then(response => {
                if (response.body.success === false) reject(response.body.message)
                else resolve(response.body)
            }, response => {
                reject(response.statusText)
            })
        })
    },
    findAllFAQ () {
        return new Promise((resolve, reject) => {
            Vue.http.post(utils.FIND_ALL_FAQ_URL, {}, { headers: auth.getAuthHeader() }).then(response => {
                if (response.body.success === false) reject(response.body.message)
                else resolve(response.body)
            }, response => {
                reject(response.statusText)
            })
        })
    },
    findAnsweredFAQ () {
        return new Promise((resolve, reject) => {
            Vue.http.post(utils.FIND_ANSWERED_FAQ_URL).then(response => {
                if (response.body.success === false) reject(response.body.message)
                else resolve(response.body)
            }, response => {
                reject(response.statusText)
            })
        })
    },
    answerFAQ (payload) {
        return new Promise((resolve, reject) => {
            Vue.http.post(utils.ANSWER_FAQ_URL, payload, { headers: auth.getAuthHeader() }).then(response => {
                if (response.body.success) resolve(response.body)
                else reject(response.body.message)
            }, response => {
                reject(response.statusText)
            })
        })
    }
}
