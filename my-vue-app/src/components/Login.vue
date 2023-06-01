<script setup>
import axios from "axios";
import {onMounted, reactive} from 'vue'
import { useStore } from 'vuex'

const store = useStore()

const formInline = reactive({
  username: '',
  password: '',
})



const onSubmit = async() => {
  if (formInline.username == ""||
    formInline.password == ""){
    return
  }
  axios({
    url:"/api/login",
    method:"GET",
    params:{
      username:formInline.username,
      password:formInline.password
    }
  })
}

const register = async ()=> {
  let formData = new FormData();

  formData.append("username",formInline.username)
  formData.append("password",formInline.password)
  axios({
    url:"/api/register",
    method:"POST",
    data:formData
  })
}
</script>

<template>
  <el-form :inline="true" :model="formInline" class="demo-form-inline">
    <el-form-item label="用户名">
      <el-input v-model="formInline.username" placeholder="请输入用户名" />
    </el-form-item>
    <el-form-item label="密码">
      <el-input v-model="formInline.password" placeholder="请输入密码" />
    </el-form-item>
    <el-form-item>
      <el-button type="primary" @click="onSubmit">登录</el-button>
      <el-button type="primary" @click="register">注册</el-button>
    </el-form-item>
  </el-form>
</template>

<style scoped>

</style>