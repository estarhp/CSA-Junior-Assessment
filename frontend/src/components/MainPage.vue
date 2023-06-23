<script setup>

import Navigator from "./Navigator.vue";
import Main from "./Main.vue";
import {useStore} from "vuex";

const store = useStore()


</script>

<template>
  <div class="common-layout">
    <el-container>
      <el-header style="background: #00BCD4;height: auto" ><Navigator></Navigator></el-header>
      <el-main style="background: url('../../public/kwp97MvPYV.jpg');height: 100%" ><Main></Main></el-main>
    </el-container>
  </div>
</template>

<style scoped>

</style>
