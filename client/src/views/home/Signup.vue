<template>
    <div class="child signup">
        <center><h3>Anmälan</h3></center>
        <form>
            <label for="nameInput">Förnamn:</label>
            <input type="text" name="name" id="nameInput" v-model="name" />
            
            <label for="lastnameInput">Efternamn:</label>
            <input type="text" name="lastname" id="lastnameInput" v-model="lastname" />
            
            <label for="emailInput">Epost:</label>
            <input type="text" name="email" id="emailInput" v-model="email" />
            
            <label for="classSelect">Klass:</label>
            <select id="classSelect" name="klass" v-model="klass">
                <option v-for="(cls, key) in classes" :key="key" :value="cls">{{cls}}</option>
            </select>

            <label for="classroomSelect">Sal:</label>
            <select id="classroomSelect" name="classroom" v-model="classroom">
                <option v-for="(room, key) in rooms" :key="key" :value="room">{{room}}</option>
            </select>

            <center>
                <button type="submit" class="signupbtn" @click.prevent="signup">Anmäl</button>
            </center>
        </form>
    </div>
</template>
<script>
import VueNotifications from 'vue-notifications'
import API from '../../api'

export default {
    name: 'Signup',
    data () {
        return {
            classes: ['15El', '15Es', '15Te', '16Es', '16It', '16Te', '17Es', '17It', '17Te'],
            rooms: ['Ingen dator', '209', '210'],
            name: '',
            lastname: '',
            email: '',
            klass: '15El',
            classroom: 'Ingen dator'
        }
    },
    methods: {
        signup () {
            API.createSignup({
                name: this.name,
                lastname: this.lastname,
                email: this.email,
                class: this.klass,
                classroom: this.classroom
            }).then((response) => {
                VueNotifications.success({message: response.message})
                this.name = ''
                this.lastname = ''
                this.email = ''
                this.klass = ''
                this.classroom = ''
            }, (error) => {
                VueNotifications.error({message: error})
            })
        }
    }
}
</script>