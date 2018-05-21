<template>
    <div class="child tournament-signup">
        <h3>{{name}}</h3>
        <div class="tournament-group" v-for="(team, index) of tournamentTeams" :key="index">
            <h4>{{index}}</h4>
            <table>
                <thead>
                    <tr>
                        <th>ID</th>
                        <th>Team</th>
                        <th>Namn</th>
                        <th>IGN</th>
                        <th>Email</th>
                        <th>Ledare</th>
                    </tr>
                </thead>
                <tbody>
                    <tr v-for="(member, imember) of team" :key="imember">
                        <td>{{member.id}}</td>
                        <td>{{member.team}}</td>
                        <td>{{member.name}}</td>
                        <td>{{member.ign}}</td>
                        <td>{{member.email}}</td>
                        <td>{{member.leader ? 'Ja' : 'Nej'}}</td>
                    </tr>
                </tbody>
            </table>
        </div>
    </div>
</template>
<script>
export default {
    name: 'tournaments',
    props: ['id', 'name', 'teams'],
    computed: {
        tournamentTeams () {
            let obj = {}
            for (let o of this.teams) {
                if (typeof obj[o.team] === 'undefined') obj[o.team] = []
                obj[o.team].push(o)
            }
            return obj
        }
    },
    methods: {
    }
}
</script>