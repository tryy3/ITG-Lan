<template>
    <div class="app">
        <AppHeader/>
        <router-view></router-view>
    </div>
</template>

<script>
import AppHeader from '../components/Header.vue'

export default {
    name: 'full',
    components: {
        AppHeader
    }
}
</script>
