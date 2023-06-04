<script setup>
import {handleTime} from "../utils/index.js";

defineProps([
    "question"
    ]
)

function handleDetails(text){
  if (text.length > 100){
    return text.substring(0,100)+"....."
  }
  return text
}
</script>

<template>
  <el-descriptions :title="question.Title" style="background: gray;width: 100%">
    <el-descriptions-item></el-descriptions-item>
    <el-descriptions-item label="发布者:">{{question.Username}}</el-descriptions-item>
    <el-descriptions-item></el-descriptions-item>
    <el-descriptions-item></el-descriptions-item>
    <el-descriptions-item label="内容:">{{handleDetails(question.Details)}}</el-descriptions-item>
    <el-descriptions-item></el-descriptions-item>
    <el-descriptions-item label="发布时间:">
      {{handleTime(question.Date)}}
    </el-descriptions-item>


  </el-descriptions>
</template>

<style scoped>

</style>