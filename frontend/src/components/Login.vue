<script setup>
import axios from "axios";
import { reactive} from 'vue'
import { useStore } from 'vuex'
import {reload} from "../utils/index.js";
import {encrypt} from "../cryption/index.js";


const store = useStore()

const formInline = reactive({
  username: '',
  password: '',
})


const app = document.querySelector(".common-layout")




const onSubmit = async() => {
  if (formInline.username.trim() == ""||
    formInline.password.trim() == ""){
    ElMessage({
      message:"输入不合法",
      type:'warning',
      customClass:"messages",
    })
    return
  }


  const result = await axios({
    url:"/api/login",
    method:"GET",
    params:{
      username:encrypt(formInline.username),
      password:encrypt(formInline.password)
    }
  })
  ElMessage({
    message:result.data.message
  })
 reload(result)

}

const register = async ()=> {
  if (formInline.username.trim() == ""||
      formInline.password.trim() == ""){
    ElMessage({
      message:"错误的输入！！",
      type:"warning"
    })
    return
  }

  let formData = new FormData();

  formData.append("username",encrypt(formInline.username))
  formData.append("password",encrypt(formInline.password))
  const result = await axios({
    url:"/api/register",
    method:"POST",
    data:formData
  })

  ElMessage({
    message:result.data.message,
  })

}
const handleInput = ()=>{
  formInline.password = formInline.password.trim()
  formInline.username = formInline.username.trim()
}
</script>

<template>
  <div
      ref="messages"><el-form :inline="true" :model="formInline" class="demo-form-inline">
    <el-form-item label="用户名">
      <el-input v-model="formInline.username" placeholder="请输入用户名" maxlength="20" minlength="1" @blur="handleInput"/>
    </el-form-item>
    <el-form-item label="密码">
      <el-input v-model="formInline.password" placeholder="请输入密码" maxlength="20" minlength="4" @blur="handleInput"/>
    </el-form-item>
    <el-form-item>
      <el-button type="primary" @click="onSubmit">登录</el-button>
      <el-button type="primary" @click="register">注册</el-button>
    </el-form-item>
  </el-form>
  </div>
</template>

<style>

</style>