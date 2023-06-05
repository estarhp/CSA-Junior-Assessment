<script setup>
import {onMounted, ref, watch,defineProps} from "vue";
import axios from "axios";
import {handleTime} from "../utils/index.js";
import CommentDetails from "./commentDetails.vue";


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
      <div style="font-size: 16px;" v-if="comments">
        <div  v-for="i in comments" :key="comments.ID" style="background: #CDD0D6;margin-bottom: 20px" >
          <CommentDetails :comment="i"></CommentDetails>
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