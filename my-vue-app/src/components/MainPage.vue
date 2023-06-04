<script setup>

import Navigator from "./Navigator.vue";
import Main from "./Main.vue";
import {onMounted} from "vue";
import {useStore} from "vuex";

const store = useStore()


</script>

<template>
  <div class="common-layout">
    <el-container>
      <el-header style="background: #606266;height: auto" ><Navigator></Navigator></el-header>
      <el-main style="background: #CDD0D6;height: 100%"><Main></Main></el-main>
    </el-container>
  </div>
</template>

<style scoped>

</style>
