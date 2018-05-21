<template>
    <div class='container'>
        <center>
            <h1>Höstlovs-LAN</h1>
            <h2>27 Oktober 2017</h2>
        </center>

        <div class='grid frontpage'>
            <Signup />
            <FAQ />
            <Information />

            <Tournaments
                v-for="(tour, index) of active_tournaments"
                :key="index"
                :id="tour.tournament.id"
                :image="tour.tournament.image"
                :amount="tour.tournament.amount"
                :name="tour.tournament.name"
                :teams="tour.teams"
            />
        </div>

        <div class='footer'>
            <p @click='easterOne'>Gjord av Dennis Planting</p>
        </div>
    </div>
</template>

<script>
import Signup from './home/Signup.vue'
import Information from './home/Information.vue'
import FAQ from './home/FAQ.vue'
import Tournaments from './home/Tournaments.vue'
import api from '../api'

function sleep (ms) {
    return new Promise(resolve => setTimeout(resolve, ms))
}

export default {
    name: 'home',
    components: {
        Signup,
        Information,
        FAQ,
        Tournaments
    },
    data () {
        return {
            easterOneRun: false,
            active_tournaments: []
        }
    },
    methods: {
        async easterOne () {
            if (this.easterOneRun) return
            this.easterOneRun = true
            for (let i = 0; i < 3; i++) {
                document.getElementsByClassName('logo')[0].style = 'position: relative; animation-name: logo; animation-duration: 2s;'
                setTimeout(() => {
                    document.getElementsByClassName('logo')[0].style = ''
                }, 2000)
                await sleep(2200)
            }
            this.easterOneRun = false
        }
    },
    mounted () {
        api.findActiveTournaments().then((response) => {
            this.active_tournaments = response
        }, (error) => {
            console.log(error)
        })
    }
}
</script>