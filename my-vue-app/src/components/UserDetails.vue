<script setup>

import {useStore} from "vuex";
import {computed, onMounted, ref} from "vue";
import QandCDetails from "./QandCDetails.vue";
import Upload from "./Upload.vue";
import {createAlova, useRequest} from "alova";
import VueHook from 'alova/vue';
import GlobalFetch from 'alova/GlobalFetch';
import {ElMessage} from "element-plus";
import axios from "axios";


const store = useStore()
const telephone = ref("")
const address = ref("")
const regExp = /^1[3-9]\d{9}$/;

const input = ref(true)




function isPhoneNumber() {


  store.state.userDetails.Telephone = store.state.userDetails.Telephone.trim()

   if (!regExp.test( store.state.userDetails.Telephone)&&!store.state.userDetails.Telephone == ""){
     input.value = false
   }


}

let alreadyModify = computed(()=>{

  const initData = store.state.initData
  const {Telephone,Address} = store.state.userDetails
  if (Telephone === initData.telephone && Address === initData.address ){
    return false
  }

  if (!input.value){
    return false
  }
  return true
})

async function handSave(){
   if (!input.value) {
     return
   }

  const {Telephone,Address} = store.state.userDetails
 const result  = await axios({
   url:"/api/saveUserDetails",
   data:{
     telephone:Telephone,
     address:Address
   },
   method:"post"

 })

  let data = result.data
  if (data.status == 500){
    ElMessage({
      message:data.message,
      type:"warning"
    })
    return;
  }else {
    ElMessage({
      message:data.message,
      type:"success"
    })
    setTimeout(()=>{
      location.reload()
    },1000)
  }

}
</script>

<template>
  <div style="background: #B2DFDB">
    <el-descriptions

        :column="4"
        size="large"
        border
    >
      <el-descriptions-item label="Username">{{store.state.userDetails.Username}}</el-descriptions-item>
      <el-descriptions-item label="Telephone">
        <el-input
            v-model="store.state.userDetails.Telephone"
            @blur="isPhoneNumber"
            @focus="input = true"
            placeholder="你可以输入你的电话号码"
        >

        </el-input>
        <div>
          <el-text class="mx-1" type="danger" v-if="!input">请输入正确的手机号</el-text>
        </div>
      </el-descriptions-item>
      <el-descriptions-item label="Adress" :span="2">
        <el-input
            v-model="store.state.userDetails.Address"
            placeholder="你可以输入自己的住址"
        >

        </el-input>
      </el-descriptions-item>

    </el-descriptions>
    <div style="margin-bottom: 20px"></div>
    <Upload></Upload>
    <div style="margin-bottom: 20px"></div>
    <el-button type="primary" round @click="handSave" v-if="alreadyModify" >保存修改</el-button>
    <el-button type="primary" round disabled v-else >保存修改</el-button>
    <QandCDetails>
    </QandCDetails>
  </div>
</template>

<style scoped>

</style>