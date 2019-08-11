<template>
  <div class="app-container">
    <el-form ref="form" :model="form" label-width="120px">
      <el-form-item label="Title">
        <el-input v-model="form.title" />
      </el-form-item>
      <el-form-item label="Category">
        <el-select v-model="form.category" placeholder="please select category">
          <el-option label="Tech" value="Tech" />
          <el-option label="Life" value="life" />
        </el-select>
      </el-form-item>
      <el-form-item label="Content">
        <el-input v-model="form.content" :rows="17" type="textarea" />
      </el-form-item>
      <el-form-item label="MDContent">
        <markdown-editor v-model="content" />
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="onSubmit">Create</el-button>
        <el-button @click="onCancel">Cancel</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

import MarkdownEditor from '@/components/MarkdownEditor'

<script>
export default {
  data() {
    return {
      form: {
        title: '',
        category: '',
        mkcontent: '',
        content: ''
      }
    }
  },
  methods: {
    onSubmit() {
      this.$message('submit!')
    },
    onCancel() {
      this.$message({
        message: 'cancel!',
        type: 'warning'
      })
    }
  }
}
</script>

<style scoped>
.line{
  text-align: center;
}
</style>

