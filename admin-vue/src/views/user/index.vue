<template>
  <div class="app-container">
    <el-table
      v-loading="loading"
      :data="users"
      element-loading-text="Loading"
      border
      fit
      highlight-current-row
    >
      <el-table-column align="center" label="ID" width="95">
        <template slot-scope="scope">
          {{ scope.$index }}
        </template>
      </el-table-column>
      <el-table-column label="Username" width="200">
        <template slot-scope="scope">
          {{ scope.row.username }}
        </template>
      </el-table-column>
      <el-table-column label="Password" width="110" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.password }}</span>
        </template>
      </el-table-column>
      <el-table-column label="Email" width="200" align="center">
        <template slot-scope="scope">
          {{ scope.row.email }}
        </template>
      </el-table-column>
      <el-table-column class-name="status-col" label="Status" width="110" align="center">
        <template slot-scope="scope">
          <el-tag :type="scope.row.status | statusFilter">{{ scope.row.status }}</el-tag>
        </template>
      </el-table-column>
      <el-table-column align="center" prop="created_at" label="CreateAt" width="200">
        <template slot-scope="scope">
          <i class="el-icon-time" />
          <span>{{ scope.row.create_at }}</span>
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>

<script>
import { findUsers } from '@/api/user'

export default {
  filters: {
    statusFilter(status) {
      const statusMap = {
        active: 'success',
        // inactive: 'gray',
        inactive: 'danger'
      }
      return statusMap[status]
    }
  },
  data() {
    return {
      users: [
          {
              id:1,
              username:"username",
              password:"password",
              email:"tantao700@163.com",
              status:"inactive",
              create_at:'2019-08-11 11:12:12'
          }
      ],
      loading: false
    }
  },
  created() {
    // this.fetchData()
  },
  methods: {
    fetchData() {
      this.loading = true
      findUsers().then(response => {
        this.users = response.data.items
        this.loading = false
      })
    }
  }
}
</script>
