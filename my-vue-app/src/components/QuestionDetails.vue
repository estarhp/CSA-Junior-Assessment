<script setup>
import {useRoute, useRouter} from "vue-router";
import {computed, onMounted, ref} from "vue";
import axios from "axios";
import Comment from "./Comment.vue";
import {useStore} from "vuex";


const textarea = ref('')

const store = useStore()
const route =useRoute()
const router = useRouter()
const question = computed(()=>{
  return JSON.parse(route.params.id)
})

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

}

</script>

<template>
  <div style="background:#F2F3F5;height: 100vh;font-size: 20px" v-if="question">
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
        @click="editQuestion"
        round >
      Edit
    </el-button>



  </div>
</template>

<style scoped>

</style>