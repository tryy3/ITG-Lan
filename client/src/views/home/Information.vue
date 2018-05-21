<template>
    <div class='child information'>
        <center><h3>Information</h3></center>
        <ul>
            <li>Lanet börjar klockan 17:00 den 27:e Oktober.</li>
            <li>Lanet avslutas 10:00 den 28:e Oktober.</li>
            <li>Sal 204a kommer ha spel så som Cod4, Quake.</li>
            <li>Sal 209/210 kommer att användas för dem med datorer.</li>
            <li>Under lanet gäller skolans vanliga ordningsregler.</li>
            <li>Tänk på att du ur brandskyddssynpunkt inte kan ta med dig för stora datastolar.</li>
        </ul>
        <center><h3>Frågor</h3></center>
        <span class='faq' v-for="(q, key) of faq" :key="key">
            <span class='question'><b>Q: </b>{{q.question}}</span>
            <span class='svar'><b>A: </b>{{q.answer}}</span>
        </span>
    </div>
</template>
<script>
import API from '../../api'

export default {
    name: 'Information',
    data () {
        return {
            faq: []
        }
    },
    mounted () {
        this.faq = []
        API.findAnsweredFAQ().then((response) => {
            this.faq = response
        }, (error) => {
            console.log(error)
        })
    }
}
</script>