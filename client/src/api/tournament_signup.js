import Vue from 'vue'
import utils from './utils.js'
import auth from '../auth/index'

export default {
    createTeam (payload) {
        return new Promise((resolve, reject) => {
            Vue.http.post(utils.CREATE_TOURNAMENT_SIGNUP_URL, payload).then(response => {
                if (response.body.success === false) reject(response.body.message)
                else resolve(response.body)
            }, response => {
                reject(response.statusText)
            })
        })
    },
    findActiveTournaments () {
        return new Promise((resolve, reject) => {
            Vue.http.post(utils.FIND_ACTIVE_TOURNAMENT_SIGNUP_URL).then(response => {
                if (response.body.success === false) reject(response.body.message)
                else resolve(response.body)
            }, response => {
                reject(response.statusText)
            })
        })
    },
    findAllTournaments () {
        return new Promise((resolve, reject) => {
            Vue.http.post(utils.FIND_ALL_TOURNAMENT_SIGNUP_URL, {}, { headers: auth.getAuthHeader() }).then(response => {
                if (response.body.success === false) reject(response.body.message)
                else resolve(response.body)
            }, response => {
                reject(response.statusText)
            })
        })
    }
}
