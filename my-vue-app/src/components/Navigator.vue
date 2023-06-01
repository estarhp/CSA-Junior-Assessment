<script setup>

import {ref} from "vue";
import {useStore} from "vuex";
const state = useStore().state


const outerVisible = ref(false)





</script>

<template>

  <el-row>
    <el-col :span="8">
    <div class="grid-content ep-bg-purple" >

    <el-button text @click="outerVisible = true" v-if="state.isLogin =='false'" style="height: 100%">
    <span style="font-size: 2em;color: black"  >登录/注册
      <el-icon><CaretBottom /></el-icon>
    </span>
  </el-button>
    <div v-else style="font-size: 2em;line-height: 2em;height: 100%"> 你好！ 亲爱的用户</div>
  </div>
  </el-col>


  </el-row>


  <el-dialog v-model="outerVisible" title="请输入用户名 和 密码" draggable>
    <template #default>
    <Login></Login>

    </template>
    <template #footer>
      <div class="dialog-footer">
        <el-button @click="outerVisible = false">Cancel</el-button>


      </div>
    </template>
  </el-dialog>
</template>

<style scoped>

.dialog-footer button:first-child {
  margin-right: 10px;
}


.el-col {
  border-radius: 4px;
}

.grid-content {
  border-radius: 4px;
  min-height: 36px;
}

</style>