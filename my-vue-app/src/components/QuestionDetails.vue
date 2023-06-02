<script setup>
import {useRoute, useRouter} from "vue-router";
import {computed, onMounted, ref} from "vue";
import axios from "axios";


const textarea = ref('')

defineProps(["id"])
const router = useRouter()
const route =useRoute()
const question = computed(()=>{
  return JSON.parse(route.params.id)
})

async function Submit(){
  const result = await axios({
    url:"getAllComment",
    method:"get",
    data:{
      beID:question.value.ID
    }
  })
}

</script>

<template>
  <div style="background:#F2F3F5;height: 100vh;font-size: 20px">
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
      />
      <el-button type="success" style="margin-right: 20px" @click="Submit">Submit</el-button>
    </div>

  </div>
</template>

<style scoped>

</style>