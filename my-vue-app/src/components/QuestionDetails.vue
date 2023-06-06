<script setup>
import {useRoute, useRouter} from "vue-router";
import {onMounted, ref} from "vue";
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
})

const centerDialogVisible = ref(false)



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

</script>

<template>
  <div style="background:#F2F3F5;height: 100vh;font-size: 20px" v-if="question">
    <div class="common-layout" style="width: 100%;background:#909399;text-align: left">
     <div style="text-align: center"> <el-text class="mx-1" style="font-size: 25px">{{question.Title}}</el-text></div>

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