<script setup lang="ts">
import {ref} from "vue";
import {ElMessage} from "element-plus";
import {useStore} from "vuex";

const imageUrl = ref('')
const store = useStore()

const handleAvatarSuccess: UploadProps['onSuccess'] = (
    response,
    uploadFile
) => {
  imageUrl.value = URL.createObjectURL(uploadFile.raw!)

  if (response.status == 500){
    ElMessage({
      message:response.message,
      type:"warning"
    })
    return
  }
  ElMessage({
    message:response.message,
    type:"success"
  })

}

const beforeAvatarUpload: UploadProps['beforeUpload'] = (rawFile) => {
   if (rawFile.size / 1024 / 1024 > 2) {
    ElMessage.error('Avatar picture size can not exceed 2MB!')
    return false
  }
  return true
}
</script>

<template>
  <el-row justify="center">
    <el-upload
        class="avatar-uploader"
        action="/api/upload"
        :show-file-list="false"
        :on-success="handleAvatarSuccess"
        :before-upload="beforeAvatarUpload"
        headers="multipart/form-data"
        style="background: #CDD0D6;margin-left: 100px"
    >
      <!--    <img v-if="imageUrl" :src="imageUrl" class="avatar" />-->
      <el-icon  class="avatar-uploader-icon"><Plus /></el-icon>
    </el-upload>
  </el-row>
</template>

<style scoped>
.avatar-uploader .el-upload {
  border: 1px dashed var(--el-border-color);
  border-radius: 6px;
  cursor: pointer;
  position: relative;
  overflow: hidden;
  transition: var(--el-transition-duration-fast);
}

.avatar-uploader .el-upload:hover {
  border-color: var(--el-color-primary);
}

.el-icon.avatar-uploader-icon {
  font-size: 28px;
  color: #8c939d;
  width: 178px;
  height: 178px;
  text-align: center;
}
</style>