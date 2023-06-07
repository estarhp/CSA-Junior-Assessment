<template>
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
</template>

<script lang="ts" setup>

import {useStore} from "vuex";


import { Timer } from '@element-plus/icons-vue'
import {handleTime} from '../utils/index.js'
import {onMounted, ref} from "vue";
interface Question {
  Date: string
  Title: string,
  Details:string
}

const store = useStore()

const handleEdit = (index: number, row: Question) => {
  console.log(index, row)
}
const handleDelete = (index: number, row: Question) => {
  console.log(index, row)
}

const tableData = ref([])
onMounted(async ()=>{
  if (!store.state.allQuestions){
    await store.dispatch("getAllQuestions")
  }

  tableData.value  = store.state.allQuestions
})

function handleDetails(detials : string){
  if (detials.length>20){
    return detials.substring(0,20) + "..."
  }
  return detials
}

</script>
