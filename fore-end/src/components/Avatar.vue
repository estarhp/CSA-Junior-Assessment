<script setup>
import {useRoute, useRouter} from "vue-router";
import {onMounted, ref} from "vue";
import axios from "axios";

const props = defineProps(["username"])
const router = useRouter()
function toDetails(){
  router.push({
    name:"user.details",
    params:{
      userDetails: JSON.stringify(userDetails.value),
      isOther:true
    }
  })
}
const userDetails = ref("")
onMounted(async ()=>{
  const result = await axios({
    url: "/api/otherUserDetails",
    method: "get",
    params:{
      username:props.username
    }
  })

  userDetails.value = result.data.userDetails


})
</script>

<template>
  <el-avatar v-if="userDetails.AvatarImage" :src="userDetails.AvatarImage" @click="toDetails"></el-avatar>
  <el-avatar v-else> user </el-avatar>
</template>

<style scoped>

</style>