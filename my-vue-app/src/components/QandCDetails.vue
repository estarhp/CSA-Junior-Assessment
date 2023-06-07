<template>
  <h3 style="text-align: left">your questions</h3>
  <el-table :data="tableData" style="width: 100%">
    <el-table-column label="Date" width="180">
      <template #default="scope">
        <div style="display: flex; align-items: center">
          <el-icon><timer /></el-icon>
          <span style="margin-left: 10px">{{ handleTime(scope.row.Date) }}</span>
        </div>
      </template>
    </el-table-column>
    <el-table-column label="Title" width="180">
      <template #default="scope">
        <div style="display: flex; align-items: center">
          <span style="margin-left: 10px">{{ scope.row.Title }}</span>
        </div>
      </template>
    </el-table-column>

      <el-table-column label="Details" width="180">
        <template #default="scope">
          <div style="display: flex; align-items: center">
            <span style="margin-left: 10px">{{ handleDetails(scope.row.Details) }}</span>
          </div>
        </template>
      </el-table-column>

      <el-table-column label="Operations">
      <template #default="scope">
        <el-button size="small" @click="handleEdit(scope.$index, scope.row)"
        >Edit</el-button
        >
        <el-button
            size="small"
            type="danger"
            @click="handleDelete(scope.$index, scope.row)"
        >Delete</el-button
        >
      </template>
    </el-table-column>
  </el-table>
   <div style="margin-bottom: 20px"></div>
  <h3 style="text-align: left">your comments</h3>
  
</template>

<script lang="ts" setup>

import {useStore} from "vuex";


import { Timer } from '@element-plus/icons-vue'
import {handleTime} from '../utils/index.js'
import {onMounted, ref} from "vue";
import axios from "axios";
import {useRouter} from "vue-router";
interface Question {
  Date: string
  Title: string,
  Details:string
}

const store = useStore()
const router = useRouter()

async function deleteQuestion(ID : string ) {
  const formData = new FormData()

  formData.append("ID",ID)
  const result = await axios({
    url:"/api/deleteQuestion",
    method:"post",
    data:formData
  }).catch(err => {
    console.error(err)
  })

  if (result.data.status == 200) {
    ElMessage({
      message:result.data.message,
      type:"warning"
    })
    setTimeout(()=>{
      location.reload()
    },1000)
  }


}

const handleEdit = (index: number, row: Question) => {
  router.push({
    name:"question.details",
    params:{
      id:row.ID,
      edit:true,
    }
  })
  console.log(index, row.ID)
}
const handleDelete = async (index: number, row: Question) => {
   await deleteQuestion(row.ID)
}

const comments = ref(null)

const tableData = ref([])
onMounted(async ()=>{
  const  result1  = await  axios({
    url:"/api/getUserQuestion",
    method:"get",
  })
  const  result2  = await  axios({
    url:"/api/getUserComments",
    method:"get"
  })

  comments.value = result2.data.comments
  tableData.value = result1.data.questions
})

function handleDetails(detials : string){
  if (detials.length>20){
    return detials.substring(0,20) + "..."
  }
  return detials
}



</script>
