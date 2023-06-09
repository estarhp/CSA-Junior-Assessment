<script setup>

import {useStore} from "vuex";
import {ref} from "vue";
import QandCDetails from "./QandCDetails.vue";
import Upload from "./Upload.vue";

const store = useStore()
const telephone = ref("")
const address = ref("")
const regExp = /^1[3-9]\d{9}$/;

const input = ref(true)
function isPhoneNumber() {
  telephone.value = telephone.value.trim()

   if (!regExp.test(telephone.value)&&!telephone.value == ""){
     input.value = false
   }


}
</script>

<template>
  <el-descriptions
      title="Vertical list with border"
      :column="4"
      size="large"
      border
  >
    <el-descriptions-item label="Username">{{store.state.userDetails.Username}}</el-descriptions-item>
    <el-descriptions-item label="Telephone">
      <el-input
          v-model="telephone"
          @blur="isPhoneNumber"
          @focus="input = true"
          placeholder="你可以输入你的电话号码"
      >

      </el-input>
      <div>
        <el-text class="mx-1" type="danger" v-if="!input">请输入正确的手机号</el-text>
      </div>
    </el-descriptions-item>
    <el-descriptions-item label="Adress" :span="2">
      <el-input
      v-model="address"
      placeholder="你可以输入自己的住址"
      >

      </el-input>
    </el-descriptions-item>

  </el-descriptions>
  <Upload></Upload>
  <QandCDetails>
  </QandCDetails>
</template>

<style scoped>

</style>