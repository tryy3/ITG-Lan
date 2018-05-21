<template>
    <div class="child questions">
        <center><h3>Frågor</h3></center>
        <span class="questionChild" v-for="(faq, index) of questions" :key="index">
            <div>{{faq.question}}</div>
            <input type="text" v-model="questions[index].answer">
            <button @click.prevent="editQuestion(faq.id)">Svara</button>
        </span>
    </div>
</template>
<script>
import VueNotifications from 'vue-notifications'
import API from '../../api'

export default {
    name: 'FAQ',
    data () {
        return {
            questions: []
        }
    },
    mounted () {
        API.findAllFAQ().then((response) => {
            this.questions = response
        }, (error) => {
            // VueNotifications.error({message: error})
            console.log(error)
        })
    },
    methods: {
        editQuestion (id) {
            let answer = this.questions.find((q) => q.id === id).answer
            API.answerFAQ({id, answer}).then((response) => {
                VueNotifications.success({message: response.message})
            }, (error) => {
                VueNotifications.error({message: error})
            })
        }
    }
}
</script>