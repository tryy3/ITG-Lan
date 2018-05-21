<template>
    <div class="child tournament-signup">
        <img :src="image" :alt="name"><br>
        <h3>{{name}}</h3>
        <ul class="split-list">
            <li v-for="(team, index) of teams" :key="index">{{team}}</li>
        </ul>
        <template v-if="amount <= 1">
            <div class="input-group" v-if="inputs.members.length >= 1">
                <label>Namn:</label>
                <input type="text" v-model="inputs.members[0].name" />
                
                <label>Ingame namn:</label>
                <input type="text" v-model="inputs.members[0].ign" />
                
                <label>Email:</label>
                <input type="text" v-model="inputs.members[0].email" />
            </div>

            <center>
                <button type="submit" class="signupbtn" @click.prevent="signupSolo">Anmäl</button>
            </center>
        </template>
        <template v-else>
            <div class="input-group">
                <label>Team namn:</label>
                <input type="text" v-model="inputs.team" />
            </div>

            <div class="input-group" v-for="(member, index) of inputs.members" :key="index">
                <label v-if="index == 0">Ledar namn:</label>
                <label v-else>Namn:</label>
                <input type="text" v-model="inputs.members[index].name" />
                
                <label v-if="index == 0">Ledar ingame namn:</label>
                <label v-else>Ingame namn:</label>
                <input type="text" v-model="inputs.members[index].ign" />
                
                <label v-if="index == 0">Ledar email:</label>
                <label v-else>Email:</label>
                <input type="text" v-model="inputs.members[index].email" />
            </div>

            <center>
                <button type="submit" class="signupbtn" @click.prevent="signupTeam">Anmäl</button>
            </center>
        </template>
    </div>
</template>
<script>
import VueNotifications from 'vue-notifications'
import api from '../../api'

export default {
    name: 'tournaments',
    props: ['id', 'image', 'amount', 'name', 'teams'],
    data () {
        return {
            inputs: {
                team: '',
                members: []
            }
        }
    },
    mounted () { this.reset() },
    methods: {
        signupSolo () {
            api.createTeam([{
                team: this.inputs.members[0].ign,
                name: this.inputs.members[0].name,
                ign: this.inputs.members[0].ign,
                leader: true,
                email: this.inputs.members[0].email,
                id_tournament: this.id
            }]).then(response => {
                VueNotifications.success({message: response.message})
                this.reset()
            }, error => {
                VueNotifications.error({message: error})
            })
        },
        signupTeam () {
            let members = []
            for (let i = 0; i < this.inputs.members.length; i++) {
                members.push({
                    ...this.inputs.members[i],
                    team: this.inputs.team,
                    leader: (i === 0),
                    id_tournament: this.id
                })
            }
            api.createTeam(members).then(response => {
                VueNotifications.success({message: response.message})
                this.reset()
            }, error => {
                VueNotifications.error({message: error})
            })
        },
        reset () {
            this.inputs.members = []
            for (let i = 0; i < this.amount; i++) {
                this.inputs.members.push({name: '', ign: '', email: ''})
            }
        }
    }
}
</script>