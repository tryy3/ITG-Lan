<template>
    <div class="container">
        <center>
            <h1>Admin</h1>
        </center>

        <div class="grid frontpage">
            <SignupTable />
            <Configuration v-on:deleteTournament="deleteTournament" />
            <FAQList />

            <Tournaments
                v-for="(tour, index) of active_tournaments"
                :key="index"
                :id="tour.tournament.id"
                :name="tour.tournament.name"
                :teams="tour.teams"
            />
        </div>
    </div>
</template>

<script>
import SignupTable from './admin/SignupTable.vue'
import Configuration from './admin/Configuration.vue'
import FAQList from './admin/FAQList.vue'
import Tournaments from './admin/Tournaments.vue'
import api from '../api'

export default {
    name: 'admin',
    components: {
        SignupTable,
        Configuration,
        FAQList,
        Tournaments
    },
    data () {
        return {
            active_tournaments: []
        }
    },
    mounted () {
        api.findAllTournaments().then((response) => {
            this.active_tournaments = response
        }, (error) => {
            console.log(error)
        })
    },
    methods: {
        deleteTournament (id) {
            this.active_tournaments = this.active_tournaments.filter((t) => t.tournament.id !== id)
        }
    }
}
</script>