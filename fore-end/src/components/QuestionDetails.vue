<script setup>
import {useRoute, useRouter} from "vue-router";
import {onMounted, ref, watch} from "vue";
import axios from "axios";
import Comment from "./Comment.vue";
import {useStore} from "vuex";
import {handleResult} from "../utils/index.js";



const textarea = ref('')

const store = useStore()
const route =useRoute()
const router = useRouter()
const question = ref(null)

const text = ref(null)
const title = ref(null)

const islike = ref(false)
const likesNumber = ref(0)

onMounted(async ()=>{
  const result  =await  axios({
    url:"/api/getQuestion",
    method:"get",
    params:{
      ID:route.params.id
    }
  })

  question.value = result.data.question
  text.value = question.value.Details
  title.value = question.value.Title

  islike.value = await isLike(question.value.ID,islike)
  likesNumber.value = await likeNumber(question.value.ID)


})

const centerDialogVisible = ref(route.params.edit === "true")



async function Submit(){
  let object = {
    content:textarea.value,
    beID:question.value.ID,
    questionID:question.value.ID
  }

  const willReload = await store.dispatch("addComment",object)
  if (willReload){
    textarea.value = ""
    location.reload()
  }

}



function handle(){
  textarea.value = textarea.value.trim()
}

async function deleteQuestion() {
  const formData = new FormData()

  formData.append("ID",question.value.ID)
  const result = await axios({
    url:"/api/deleteQuestion",
    method:"post",
    data:formData
  }).catch(err => {
    console.error(err)
  })

  if (result.data.status == 200) {
    router.back({
      path:"/",
      replace:true
    })
    ElMessage({
      message:result.data.message,
      type:"warning"
    })
  }


}

async function editQuestion(){
   centerDialogVisible.value = false
   if (title.value == "" || text.value == ""){
     ElMessage({
       message:"输入不可为空",
       type:"warning"
     })
     return
   }

   const  formData = new FormData()
  formData.append("title",title.value)
  formData.append("details",text.value)
  formData.append("ID",question.value.ID)
    const result = await axios({
      url:"/api/editQuestion",
      method:"post",
      data:formData
    })

  handleResult(result)

}

 async function like(questionID) {
  const result = await axios({
    url:"/api/like",
    method:"get",
    params:{
      questionID
    }
  })
   handleResult(result,false)
   islike.value = await isLike(question.value.ID,islike)
   likesNumber.value = await likeNumber(question.value.ID)

}
 async function unlike(questionID) {
  const result = await axios({
    url:"/api/unlike",
    method:"get",
    params:{
      questionID
    }
  })
   handleResult(result,false)
   islike.value = await isLike(question.value.ID,islike)
   likesNumber.value = await likeNumber(question.value.ID)
}

 async function likeNumber(questionID) {
  const result = await axios({
    url:"/api/likeNumber",
    method:"get",
    params:{
      questionID
    }
  })

   return result.data.likes
}

async function isLike(questionID){
  const result = await axios({
    url:"/api/isLike",
    method:"get",
    params:{
      questionID:questionID
    }
  })
  return  result.data.isLike

}





</script>

<template>
  <div style="background:#F2F3F5;height: 100vh;font-size: 20px" v-if="question">
    <div class="common-layout" style="width: 100%;background:#BBDEFB;text-align: left">
     <div style="text-align: center"> <el-text class="mx-1" style="font-size: 25px">{{question.Title}}</el-text></div>

      <el-text class="mx-1" type="success" size="large">{{question.Username}} : </el-text>
      <el-text class="mx-1">{{question.Details}}</el-text>
    </div>
    <div style="margin: 20px 0" />
    <div style="text-align: right;">
      <el-input
          v-model="textarea"
          type="textarea"
          placeholder="Please input"
          :autosize="{ minRows: 5, maxRows: 10 }"
          @blur="handle"
      />
      <el-button type="success" style="margin-right: 20px" @click="Submit">Submit</el-button>
      <span style="margin-right: 10px;color: gray">{{likesNumber}}</span>
      <el-icon style="margin-right: 10px;cursor: pointer;height: ;: 30px" @click="like(question.ID) && isLike(question.ID,islike)" v-if="!islike"><Star /></el-icon>
      <el-icon style="margin-right: 10px;cursor: pointer;color: red;height: 30px " @click="unlike(question.ID) && isLike(question.ID,islike)"  v-else><Star /></el-icon>
    </div>
    <div>
      <Comment :beID="question.ID"></Comment>
    </div>
<div >


  <el-button
      type="danger"
      round icon="delete"
      v-if="store.state.userDetails.Username && store.state.userDetails.Username === question.Username"
      @click="deleteQuestion"
  >delete</el-button>

  <el-button
      type="primary"
      icon="edit"
      v-if="store.state.userDetails.Username && store.state.userDetails.Username === question.Username"
      @click="centerDialogVisible = true"
      round >
    Edit
  </el-button>

</div>
    <el-dialog v-model="centerDialogVisible"
               width="30%"
               center
               :fullscreen="true"
    >
      title : <el-input
          v-model="title"
          autosize
          type="textarea"
          placeholder="Please input"
          @blur="handle"
      />
      <div style="margin-bottom: 20px"></div>
      content :
      <el-input
          v-model="text"
          autosize
          type="textarea"
          placeholder="Please input"
          @blur="handle"
      />
      <template #footer>
      <span class="dialog-footer">
        <el-button @click="centerDialogVisible = false">Cancel</el-button>
        <el-button type="primary" @click="editQuestion">
          Confirm
        </el-button>
      </span>
      </template>
    </el-dialog>


  </div>
</template>

<style scoped>

</style>