<script setup>
import {onMounted, ref, watch,defineProps} from "vue";
import axios from "axios";


const props = defineProps(["beID"])

const comments =ref(null);

onMounted(async function getComments(){

  const result = await axios({
    url:"/api/getAllComments",
    method:"get",
    params:{
      beID:props.beID
    }
  }).catch(err => {
    console.log(err)
  })
  comments.value = result.data.comments

})

</script>

<template>
  <suspense>
    <template #default>
      <div>
        <div v-if="comments">
          <el-row v-for="i in comments" :key="comments.ID">
            <el-col :span="4">{{i.Username}} : </el-col>
            <el-col :span="12">{{i.Content}}</el-col>
            <el-col :span="8">{{i.Date}}</el-col>

          </el-row>
        </div>
      </div>
    </template>
    <template #fallback>
      <div>
        <Loading></Loading>
      </div>
    </template>
  </suspense>
</template>

<style scoped>

</style>