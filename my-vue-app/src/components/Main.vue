<script setup>
import {onMounted, reactive, ref} from "vue";
import axios from "axios";
import Question from "./Question.vue";
import {useRouter} from "vue-router";


const router = useRouter()


const questions =ref(null)
onMounted(async ()=>{


  const result = await axios({
      url:"/api/getAllQuestions",
      method:"get",
  })

  questions.value = result.data.questions



})


function goDetail(i){

  console.log(i)
  router.push({name:"question.details",params:{id:JSON.stringify(i)}})
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