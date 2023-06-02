<script setup>

import {reactive, ref} from 'vue'
import axios from "axios";


const question = reactive({
  input : "",
  textarea : ""
})

let dialogVisible = ref(false)

async function handler() {
  dialogVisible = false

  if (question.input == '' || question.textarea == "") {
    ElMessage({
      message:"输入不合法",
      type:'warning'
    })
    return
  }

  let formData = new FormData();

  formData.append("title", question.input)
  formData.append("details", question.textarea)

  const result = await axios({
    url: "/api/addQuestion",
    method: "POST",
    data: formData
  })

  if (result.data.status == 200 ){
    ElMessage({
      message:result.data.message,
      center:true,
      type:"success"
    })

    return
  }

  ElMessage({
    message:result.data.message,
    center:true,
    type:"error"
  })


}

const handleInput = ()=>{
  question.input = question.input.trim()
  question.textarea = question.textarea.trim()
}

</script>

<template>
  <el-col :span="3" >
    <el-button text @click="dialogVisible = true" color="white" >
      <el-icon :size="20" color="#303133"><Plus /></el-icon>
    </el-button>
    <el-dialog v-model="dialogVisible" title="Add" width="30%" draggable>

      <el-input v-model="question.input" placeholder="title" @blur="handleInput" />
      <div style="margin: 20px 0"  />
      <el-input
          v-model="question.textarea"
          autosize
          type="textarea"
          placeholder="Please input"
          @blur="handleInput"
      />
      <template #footer>
      <span class="dialog-footer">
        <el-button @click="dialogVisible = false">Cancel</el-button>
        <el-button type="primary" @click="handler">
          Confirm
        </el-button>
      </span>
      </template>
    </el-dialog>
  </el-col>



</template>

<style scoped>

</style>