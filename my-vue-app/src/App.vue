<script setup>
import MainPage from './components/MainPage.vue'
import {watch} from "vue";
import {useStore} from "vuex";
const state = useStore().state



</script>

<template>
  <router-view></router-view>
</template>

<style >
#app {
  width: 100vw;
  height: 100vh;
}

</style>
