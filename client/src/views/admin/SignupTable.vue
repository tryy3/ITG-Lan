<template>
    <div class="child information">
        <h3>Anmälda</h3>
        <table>
            <thead>
                <tr>
                    <th>#</th>
                    <th>Namn</th>
                    <th>Email</th>
                    <th>Klass</th>
                    <th>Rum</th>
                </tr>
            </thead>
            <tbody>
                <tr v-for="(s, index) of rows" :key="index">
                    <td>{{index + 1}}</td>
                    <td>{{s.Namn}}</td>
                    <td>{{s.Email}}</td>
                    <td>{{s.Klass}}</td>
                    <td>{{s.Room}}</td>
                </tr>
            </tbody>
        </table>
    </div>
</template>
<script>
import API from '../../api'

export default {
    name: 'SignupTable',
    data () {
        return {
            rows: []
        }
    },
    mounted () {
        API.findSignup().then((response) => {
            this.rows = []
            for (let s of response) {
                this.rows.push({
                    Namn: s.name + ' ' + s.lastname,
                    Email: s.email,
                    Klass: s.class,
                    Room: s.classroom
                })
            }
        }, (error) => {
            console.log(error)
        })
    }
}
</script>