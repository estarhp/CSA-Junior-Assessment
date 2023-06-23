<script setup>

import {useStore} from "vuex";
import {computed, onMounted, ref} from "vue";
import QandCDetails from "./QandCDetails.vue";
import Upload from "./Upload.vue";
import {ElMessage} from "element-plus";
import axios from "axios";
import {useRoute} from "vue-router";
import {getRandomColor, handleResult} from "../utils/index.js";


const store = useStore()
const telephone = ref("")
const address = ref("")
const regExp = /^1[3-9]\d{9}$/;

const input = ref(true)
const route = useRoute()
const userDetails = ref("")

const IsFollowed = ref(false)
const isOther = computed(()=>{
  if (userDetails.value.Username === store.state.userDetails.Username){
    return true
  }
  if (route.params.isOther ==="false"){
    return true
  }
  return false
})
onMounted(async ()=>{
    userDetails.value = JSON.parse(route.params.userDetails)
    if (isOther.value === false) {

      await IsFollow()
    }
})

async function IsFollow(){
  const result = await axios({
    url:"/api/IsFollowed",
    params:{
      BeUsername:userDetails.value.Username
    }
  })

  IsFollowed.value =  result.data.IsFollowed
}

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

async function follow(){
  const result = await axios({
    url:"/api/follow",
    method:"get",
    params:{
      followed : userDetails.value.Username
    }
  })
  handleResult(result,false)
  await IsFollow()
}

async function unfollow() {
  const result = await axios({
    url:"/api/unfollow",
    method:"get",
    params:{
      followed : userDetails.value.Username
    }
  })
  handleResult(result,false)
  await IsFollow()
}
</script>

<template>
  <div style="background: #B2DFDB">
    <el-descriptions

        :column="4"
        size="large"
        border
    >
      <el-descriptions-item label="Username" v-if="isOther">{{store.state.userDetails.Username}}</el-descriptions-item>
      <el-descriptions-item label="Username" v-else>{{userDetails.Username}}</el-descriptions-item>
      <el-descriptions-item label="Telephone">
        <el-input
            v-model="store.state.userDetails.Telephone"
            @blur="isPhoneNumber"
            @focus="input = true"
            placeholder="你可以输入你的电话号码"
            v-if="isOther"
        >

        </el-input>
        <el-input
            v-model="userDetails.Telephone"
            @blur="isPhoneNumber"
            @focus="input = true"
            placeholder=""
            v-else
            disabled
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
            v-if="isOther"
        >

        </el-input>
        <el-input
            v-model="userDetails.Address"
            placeholder=""
            v-else
            disabled
        >

        </el-input>
      </el-descriptions-item>

    </el-descriptions>
    <div style="margin-bottom: 20px"></div>
    <el-row justify="center" v-if="isOther">
      <el-image style="width: 250px;display: inline-block;" :src="store.state.userDetails.AvatarImage" :fit="fit" />
      <Upload v-if="isOther"></Upload>
    </el-row>
    <el-row justify="center"  v-else>
      <el-image style="width: 250px;display: inline-block;" :src="userDetails.AvatarImage" :fit="fit" />
      <Upload v-if="isOther"></Upload>
    </el-row>
    <div style="margin-bottom: 20px"></div>
    <div  v-if="isOther">
      <el-button type="primary" round @click="handSave" v-if="alreadyModify" >保存修改</el-button>
      <el-button type="primary" round disabled v-else >保存修改</el-button>
    </div>
    <div v-else>
      <el-button type="success" plain @click="follow" v-if="!IsFollowed">关注</el-button>
      <el-button type="success" plain @click="unfollow" v-else>取消关注</el-button>
    </div>
    <QandCDetails v-if="isOther">
    </QandCDetails>

 <div v-if="isOther">
   <h3 style="text-align: left" v-if="store.state.userDetails">你的关注 ({{store.state.userDetails.Follows.length}}) </h3>
   <el-row><el-col :span="4" v-for="i in store.state.userDetails.Follows" :style=" `background: ${getRandomColor()}`">{{i}}</el-col></el-row>
 </div>
    <div v-else>
      <h3 style="text-align: left" v-if="store.state.userDetails.Follow">他/她的关注 ({{store.state.userDetails.Follows.length}}) </h3>
    </div>

  </div>
</template>

<style scoped>

</style>