<script setup>
import MainPage from './components/MainPage.vue'
import {onMounted, watch} from "vue";
import {useStore} from "vuex";
const store = useStore()



onMounted(async()=>{
  await store.dispatch("AlreadyLogin")
})

onMounted(async ()=>{

    await store.dispatch("getAllQuestions")



})



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
