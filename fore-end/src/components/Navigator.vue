<script setup>

import {ref} from "vue";
import {useStore} from "vuex";
import AddQuestion from "./AddQuestion.vue";
import axios from "axios";
import {useRouter} from "vue-router";
const state = useStore().state
const router = useRouter()

const outerVisible = ref(false)



async function logoff() {
  const result = await axios({
    url:"/api/logoff",
    method:"post"
  }).catch(err => {
    console.error(err)
    return
  })

  ElMessage({
    message:result.data.message,
    type:"success"
  })
  location.reload()
}

function userDetails() {
  router.push({
   name:"user.details",
    params:{
     userDetails: JSON.stringify(""),
      isOther:false
    }
  })
}



</script>

<template>

  <el-row >
    <AddQuestion></AddQuestion>
    <el-col :span="8" :offset="12">
    <div class="grid-content ep-bg-purple" >
    <el-button text @click="outerVisible = true" v-if="state.isLogin =='false'" style="height: 100%">
    <span style="font-size: 1em;color: black;line-height: 2em"  >登录/注册
      <el-icon><CaretBottom /></el-icon>
    </span>
    </el-button>

      <el-row v-if="state.isLogin === 'true'">
       <el-col :span="8">
        <span style="line-height: 3em">
           你好！<span>{{state.userDetails.Username}}</span>
        </span>
       </el-col>

        <el-col :span="8">
          <el-avatar class="center" v-if="!state.userDetails.AvatarImage"
                     src="https://cube.elemecdn.com/0/88/03b0d39583f48206768a7534e55bcpng.png"
                     @click="userDetails"
          />
          <el-avatar class="center" v-else
                     :src="state.userDetails.AvatarImage"
                     @click="userDetails"
          />
        </el-col>
        <el-col :span="8">
          <el-button type="warning" style="margin-left: 100px;height: 2.5em;margin-top: 0.5em" @click="logoff" v-if="state.isLogin === 'true'" >注销</el-button>
        </el-col>
      </el-row>



        </div>
  </el-col>


  </el-row>


  <el-dialog v-model="outerVisible" title="请输入用户名 和 密码" draggable :modal="false">
    <template #default>
    <Login></Login>

    </template>
    <template #footer>
      <div class="dialog-footer">
        <el-button @click="outerVisible = false">Cancel</el-button>


      </div>
    </template>
  </el-dialog>
</template>

<style scoped>

.dialog-footer button:first-child {
  margin-right: 10px;
}


.el-col {
  border-radius: 4px;
}

.grid-content {
  border-radius: 4px;
}

.center {
 margin-bottom: -20px;
}

</style>