<script setup>

import {handleTime} from "../utils/index.js";
import {ref} from "vue";
import axios from "axios";

const props = defineProps([
    "comment"
])
const input = ref('')
const centerDialogVisible = ref(false)
async function reply() {
  centerDialogVisible.value = false
  if (input.value=""){
    ElMessage({
      message:"输入不可为空",
      type:"warning"
    })
    return
  }

  const formData = new FormData()
  formData.append("beID",props.beID)
  formData.append("content",input.value)
  formData.append("questionID",)
  axios({
    url:"/api/addComment",
    method:"post"
  })
}
function handle(){
  input.value=input.value.trim()
}
</script>

<template>
  <el-dialog v-model="centerDialogVisible"
             width="30%"
             :modal="false"
             top="75vh"
             :lock-scroll="true"
  >
    <el-input v-model="input" placeholder="Please input" @blur="handle" />
    <template #footer>
      <span class="dialog-footer">
        <el-button @click="centerDialogVisible = false">Cancel</el-button>
        <el-button type="primary" @click="reply">
          Confirm
        </el-button>
      </span>
    </template>
  </el-dialog>
  <el-row >
    <el-col :span="2">{{props.comment.Username}} : </el-col>
    <el-col :offset="20" :span="2" >
      <el-dropdown>
    <span class="el-dropdown-link">
       <el-icon><More /></el-icon>
    </span>
        <template #dropdown>
          <el-dropdown-menu>
            <el-dropdown-item @click="centerDialogVisible = true">回复</el-dropdown-item>
            <el-dropdown-item>Action 2</el-dropdown-item>
            <el-dropdown-item>Action 3</el-dropdown-item>
          </el-dropdown-menu>
        </template>
      </el-dropdown>
    </el-col>
  </el-row>
  <el-row >
    <el-col :span="16">{{props.comment.Content}}</el-col>
  </el-row>
  <el-row>
    <el-col :span="8" :offset="17">{{handleTime(props.comment.Date)}}</el-col>
  </el-row>
</template>

<style scoped>
.example-showcase .el-dropdown-link {
  cursor: pointer;
  color: var(--el-color-primary);
  display: flex;
  align-items: center;
}
.dialog-footer button:first-child {
  margin-right: 10px;
}
</style>