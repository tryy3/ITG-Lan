import Vue from 'vue'
import utils from './utils.js'
import auth from '../auth/index'

export default {
    createTournament () {
        return new Promise((resolve, reject) => {
            Vue.http.post(utils.CREATE_TOURNAMENT_URL, {}, { headers: auth.getAuthHeader() }).then(response => {
                if (response.body.success) resolve(response.body)
                else reject(response.body.message)
            }, response => {
                reject(response.statusText)
            })
        })
    },
    editTournament (payload) {
        return new Promise((resolve, reject) => {
            Vue.http.post(utils.EDIT_TOURNAMENT_URL, payload, { headers: auth.getAuthHeader() }).then(response => {
                if (response.body.success) resolve(response.body)
                else reject(response.body.message)
            }, response => {
                reject(response.statusText)
            })
        })
    },
    deleteTournament (payload) {
        return new Promise((resolve, reject) => {
            Vue.http.post(utils.DELETE_TOURNAMENT_URL, payload, { headers: auth.getAuthHeader() }).then(response => {
                if (response.body.success) resolve(response.body)
                else reject(response.body.message)
            }, response => {
                reject(response.statusText)
            })
        })
    },
    findTournaments () {
        return new Promise((resolve, reject) => {
            Vue.http.post(utils.FIND_TOURNAMENT_URL, {}, { headers: auth.getAuthHeader() }).then(response => {
                if (response.body.success === false) reject(response.body.message)
                else resolve(response.body)
            }, response => {
                reject(response.statusText)
            })
        })
    }
}
