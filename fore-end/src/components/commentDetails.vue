<script setup>

import {getRandomColor, handleResult, handleTime} from "../utils/index.js";
import {ref} from "vue";
import {useStore} from "vuex";
import axios from "axios";
const props = defineProps([
    "comment",
    "questionID",
])

const store = useStore()
const input = ref('')
const centerDialogVisible = ref(false)
const type = ref(0)
async function reply() {
switch (type.value){
  case 1:{
    centerDialogVisible.value = false
    let object = {
      content:input.value,
      beID:props.comment.ID,
      questionID: props.questionID,
      beUsername: props.comment.Username
    }
    let willReload = await store.dispatch("addComment",object)
    if (willReload) {
      input.value = ''
      setTimeout(()=>{
        location.reload()
      },500)


    }
    break
  }
  case 2:{
    centerDialogVisible.value = false
    if (input.value==""){
      ElMessage({
        message:"输入不可为空",
        type:"warning"
      })
      return
    }
    const formData = new FormData()
    formData.append("ID",props.comment.ID)
    formData.append("content",input.value)
    const result = await  axios({
      url:"/api/updateComment",
      method:"post",
      data:formData
    })

    handleResult(result)

    break
  }



}


}
function handle(){
  input.value=input.value.trim()

}
function modifyComment(){
  centerDialogVisible.value = true
  type.value= 2
  input.value=props.comment.Content

}

async function deleteComment(){

  const formData = new FormData()
  formData.append("ID",props.comment.ID)
  const result = await axios({
    url:'/api/deleteComment',
    method:"post",
    data : formData
  })


  handleResult(result)

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
 <div :style=" `background: ${getRandomColor()}`">
   <el-row>
     <el-col :span="6" style="text-align: left">{{props.comment.Username}} <span v-if="props.comment.BeUsername">->{{props.comment.BeUsername}}</span> : </el-col>
     <el-col :offset="14" :span="2" >
       <el-dropdown>
    <span class="el-dropdown-link">
       <el-icon><More /></el-icon>
    </span>
         <template #dropdown>
           <el-dropdown-menu>
             <el-dropdown-item @click="centerDialogVisible = true , type=1">回复</el-dropdown-item>
             <el-dropdown-item @click="modifyComment"
                               v-if="(store.state.userDetails.Username) && (store.state.userDetails.Username === props.comment.Username) "
             >修改</el-dropdown-item>
             <el-dropdown-item @click="deleteComment"
                               v-if="(store.state.userDetails.Username) && (store.state.userDetails.Username === props.comment.Username) "
             >删除</el-dropdown-item>
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
 </div>
  <div v-if="props.comment.Be">
    <CommentDetails v-for="comment in props.comment.Be " :comment="comment" :questionID="props.questionID"></CommentDetails>
  </div>
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