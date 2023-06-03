<script setup>
import {useRoute, useRouter} from "vue-router";
import {computed, onMounted, ref} from "vue";
import axios from "axios";
import Comment from "./Comment.vue";


const textarea = ref('')

defineProps(["id"])
const router = useRouter()
const route =useRoute()
const question = computed(()=>{
  return JSON.parse(route.params.id)
})



async function Submit(){

  if (textarea.value == ""){
    ElMessage({
      message:"输入不合法",
      type:"warning"
    })
    return
  }

  const formData = new FormData()
  formData.append("content",textarea.value)
  formData.append("beID",question.value.ID)
  const result = await axios({
    url:"/api/addComment",
    method:"POST",
    data: formData
  })

  ElMessage({
    message:result.data.message
  })
}

function handle(){
  textarea.value = textarea.value.trim()
}

</script>

<template>
  <div style="background:#F2F3F5;height: 100vh;font-size: 20px">
    <div class="common-layout" style="width: 100%;background:#909399;text-align: left">
      <el-text class="mx-1" type="success" size="large">{{question.Username}} : </el-text>
      <el-text class="mx-1">{{question.Details}}</el-text>
    </div>
    <div style="margin: 20px 0" />
    <div style="text-align: right">
      <el-input
          v-model="textarea"
          type="textarea"
          placeholder="Please input"
          :autosize="{ minRows: 5, maxRows: 10 }"
          @blur="handle"
      />
      <el-button type="success" style="margin-right: 20px" @click="Submit">Submit</el-button>
    </div>
    <div>
      <Comment :beID="question.ID"></Comment>
    </div>

  </div>
</template>

<style scoped>

</style>