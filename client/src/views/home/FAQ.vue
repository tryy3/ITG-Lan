<template>
    <div class="child questions">
        <center><h3>Frågor</h3></center>
        <form>
            <label for="questionInput" style="height: 150px">Fråga:</label>
            <textarea type="text" name="question" id="questionInput" v-model="question" />

            <center>
                <button type="submit" class="btn" @click.prevent="sendQuestion">Fråga</button>
            </center>
        </form>
    </div>
</template>
<script>
import VueNotifications from 'vue-notifications'
import API from '../../api'

export default {
    name: 'FAQ',
    data () {
        return {
            question: ''
        }
    },
    methods: {
        sendQuestion () {
            API.createFAQ({question: this.question}).then(response => {
                this.question = ''
                VueNotifications.success({message: response.message})
            }, (error) => {
                VueNotifications.error({message: error})
            })
        }
    }
}
</script>