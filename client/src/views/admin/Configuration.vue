<template>
    <div class="child tournaments">
        <h3>General</h3>
        <h3>Turneringar</h3>
        <button @click.prevent="addTournament">Lägg till ny turnering</button>
        <div class="tournament" v-for="(tour, index) of tournaments" :key="index">
            <label>Namn:</label>
            <input type="text" v-model="tournaments[index].name" />
            
            <label>Antal:</label>
            <input type="text" v-model="tournaments[index].amount" />
            
            <label>Bild URL:</label>
            <input type="text" v-model="tournaments[index].image" />

            <label class="switch-btn" ref="switch">
                <input class="checked-switch" type="checkbox" v-model="tournaments[index].active" />
                <span class="text-switch" data-yes="aktiv" data-no="inaktiv"></span>
                <span class="toggle-btn"></span>
            </label>
        
            <center>
                <button type="submit" @click.prevent="editTournament(tour.id)">Ändra turneringen</button>
                <button type="submit" @click.prevent="deleteTournament(tour.id)">Ta bort turneringen</button>
            </center>
        </div>
        <hr />
    </div>
</template>
<script>
import VueNotifications from 'vue-notifications'
import API from '../../api'

export default {
    name: 'Configuration',
    data () {
        return {
            tournaments: []
        }
    },
    mounted () {
        API.findTournaments().then((response) => {
            this.tournaments = response
        }, (error) => {
            // VueNotifications.error({message: error})
            console.log(error)
        })
    },
    methods: {
        addTournament () {
            API.createTournament().then((response) => {
                VueNotifications.success({message: response.message})
                this.tournaments.push(response.data)
            }, (error) => {
                VueNotifications.error({message: error})
            })
        },
        editTournament (id) {
            let tour = this.tournaments.find((t) => t.id === id)
            tour.amount = parseInt(tour.amount)
            API.editTournament(tour).then((response) => {
                VueNotifications.success({message: response.message})
            }, (error) => {
                VueNotifications.error({message: error})
            })
        },
        deleteTournament (id) {
            API.deleteTournament({id}).then((response) => {
                VueNotifications.success({message: response.message})
                this.tournaments = this.tournaments.filter((t) => t.id !== id)
                this.$emit('deleteTournament', id)
            }, (error) => {
                VueNotifications.error({message: error})
            })
        }
    }
}
</script>