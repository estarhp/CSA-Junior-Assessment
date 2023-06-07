<script setup>
import {onMounted, reactive, ref} from "vue";
import axios from "axios";
import Question from "./Question.vue";
import {useRouter} from "vue-router";
import {useStore} from "vuex";


const router = useRouter()
const store = useStore()

const questions =ref(null)
onMounted(async ()=>{
  if (!store.state.allQuestions){
    await store.dispatch("getAllQuestions")
  }

  questions.value = store.state.allQuestions
})


function goDetail(i){
  router.push({
    name:"question.details",
    params:{
      id:i.ID,
      edit:false
    }
  })
}



</script>

<template>

  <Suspense>
    <template #default>
     <div v-if="questions">
       <el-row v-for="i in questions" :key="i.ID">
         <Question :question="i" @click="goDetail(i)"></Question>
       </el-row>
     </div>
    </template>

    <template #fallback>
      <!-- 这里是等待异步组件加载时显示的内容 -->
      <div><Loading></Loading></div>
    </template>
  </Suspense>

</template>

<style scoped>
.example-pagination-block + .example-pagination-block {
  margin-top: 10px;
}
.example-pagination-block .example-demonstration {
  margin-bottom: 16px;
}

.el-row {
  margin-bottom: 20px;
}
.el-row:last-child {
  margin-bottom: 0;
}
</style>