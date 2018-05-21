<template>
    <div class='navbar'>
        <div class='logo'>
            <template v-for="(img, key) in images">
                <span class='sep' v-if="img.sep" :key="key" />
                <img v-else :src="'/static/img/' + img.image + '.png'" :alt="img.alt" :ref="img.image" :style="img.style" :key="key"  />
            </template>
        </div>
    </div>
</template>
<script>
export default {
    name: 'header',
    methods: {
        toggleImage (i, show) {
            if (show) {
                this.images[i].style.height = '50px'
                this.images[i].style.opacity = 1
            } else {
                this.images[i].style.height = '0px'
                this.images[i].style.opacity = 0
            }
        }
    },
    data () {
        return {
            images: [
                {image: 'w', alt: 'w', style: {backgroundColor: '#000'}},
                {image: 'e-1', alt: 'e', style: {backgroundColor: '#000'}},
                {sep: true},
                {image: 'l', alt: 'l', style: {backgroundColor: '#0098C3'}},
                {image: 'o', alt: 'o', style: {backgroundColor: '#CF0072'}},
                {image: 'v', alt: 'v', style: {backgroundColor: '#FFA02F'}},
                {image: 'e-2', alt: 'e', style: {backgroundColor: '#000'}},
                {sep: true},
                {image: 'i', alt: 'i', style: {backgroundColor: '#000'}},
                {image: 't', alt: 't', style: {backgroundColor: '#000'}},
                {image: 'l-2', alt: 'l', style: {backgroundColor: '#000', opacity: 0, height: 0}},
                {image: 'a', alt: 'a', style: {backgroundColor: '#000', opacity: 0, height: 0}},
                {image: 'n', alt: 'n', style: {backgroundColor: '#000', opacity: 0, height: 0}}
            ],
            defaultColor: '#000',
            colors: ['#0098C3', '#CF0072', '#FFA02F'],
            loops: [
                [0, 1, 3, 4, 5, 6, 8, 9],
                [0, 1, 3, 4, 5, 6, 10, 11, 12]
            ],
            currentLoop: 0,
            colorLoop: [3, 4, 5],
            steps: {
                current: 0,
                switch: 3
            }
        }
    },
    mounted () {
        setInterval(() => {
            let colors = this.loops[this.currentLoop]

            // Switch (it/lan)
            this.steps.current++
            if (this.steps.current >= colors.length) {
                this.steps.current = 0
                this.currentLoop = (this.currentLoop === 0) ? 1 : 0
                colors = this.loops[this.currentLoop]
                switch (this.currentLoop) {
                case 0:
                    this.toggleImage(8, true)
                    this.toggleImage(9, true)
                    this.toggleImage(10, false)
                    this.toggleImage(11, false)
                    this.toggleImage(12, false)
                    break
                case 1:
                    this.toggleImage(8, false)
                    this.toggleImage(9, false)
                    this.toggleImage(10, true)
                    this.toggleImage(11, true)
                    this.toggleImage(12, true)
                    break
                }
            }

            for (let i in this.images) {
                if (this.images[i].sep) continue
                let randomRotate = Math.floor(Math.random() * 30) - 15
                let randomHeight = Math.floor(Math.random() * 6) - 3
                this.images[i].style.transform = 'rotate(' + randomRotate + 'deg) translateY(' + randomHeight + 'px)'
                this.images[i].style.backgroundColor = this.defaultColor
            }

            for (let i in this.colorLoop) {
                this.images[colors[this.colorLoop[i]]].style.backgroundColor = this.colors[i]
                this.colorLoop[i]++
                if (this.colorLoop[i] >= colors.length) this.colorLoop[i] = 0
            }
        }, 1000)
    }
}
</script>