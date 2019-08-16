<template>
  <div class="app-container">

    <div class="filter-container" style="margin-bottom: 10px">
      <el-input v-model="form.title" placeholder="Title" style="width: 70%;" class="filter-item" @keyup.enter.native="handleFilter" />
      <el-select v-model="form.importance" placeholder="Imp" clearable style="width: 90px" class="filter-item">
        <el-option v-for="item in importanceOptions" :key="item" :label="item" :value="item" />
      </el-select>
      <el-button v-waves class="filter-item" type="primary" icon="el-icon-search" @click="handleFilter" />
      <el-button v-waves class="filter-item" type="primary" icon="el-icon-plus" @click="handleFilter" />
    </div>

    <el-table
      v-loading="loading"
      :data="rows"
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
      <el-table-column label="Title">
        <template slot-scope="scope">
          {{ scope.row.title }}
        </template>
      </el-table-column>
      <el-table-column label="User" width="110" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.user }}</span>
        </template>
      </el-table-column>
      <el-table-column class-name="status-col" label="Status" width="110" align="center">
        <template slot-scope="scope">
          <el-tag :type="scope.row.status | statusFilter">{{ scope.row.status }}</el-tag>
        </template>
      </el-table-column>
      <el-table-column align="center" prop="created_at" label="CreateAt" width="160">
        <template slot-scope="scope">
          <el-tag>{{ scope.row.create_at }}</el-tag>
        </template>
      </el-table-column>

      <el-table-column label="Actions" align="center" width="200" class-name="small-padding fixed-width">
        <template slot-scope="{row}">
          <el-button type="primary" size="mini" icon="el-icon-edit" @click="dialog()" />
          <el-button v-if="row.status!='published'" size="mini" type="success" icon="el-icon-success" @click="handleModifyStatus(row,'published')" />
          <el-button v-if="row.status!='draft'" size="mini" icon="el-icon-question" @click="handleModifyStatus(row,'draft')" />
          <el-button v-if="row.status!='deleted'" size="mini" icon="el-icon-delete" type="danger" @click="handleModifyStatus(row,'deleted')" />
        </template>
      </el-table-column>
    </el-table>

    <pagination v-show="total>0" :total="total" :page.sync="page.page" :limit.sync="page.limit" @pagination="find" />

    <el-dialog :visible.sync="dialogPvVisible" title="Reading statistics">
      <el-table :data="pvData" border fit highlight-current-row style="width: 100%">
        <el-table-column prop="key" label="Channel" />
        <el-table-column prop="pv" label="Pv" />
      </el-table>
      <span slot="footer" class="dialog-footer">
        <el-button type="primary" @click="dialogPvVisible = false">Confirm</el-button>
      </span>
    </el-dialog>
  </div>
</template>

<script>
import Pagination from '@/components/Pagination'

export default {
  components: {
    Pagination
  },
  filters: {
    statusFilter(status) {
      const statusMap = {
        publish: 'success',
        draft: 'gray',
        deleted: 'danger'
      }
      return statusMap[status]
    }
  },
  data() {
    return {
      rows: [
        {
          id: 1,
          title: '文章标题',
          user: '作者',
          status: 'draft',
          create_at: '2019/08/11 11:12:12'
        },
        {
          id: 2,
          title: '文章标题2',
          user: '作者2',
          status: 'published',
          create_at: '2019/08/11 11:12:12'
        },
        {
          id: 3,
          title: '文章标题3',
          user: '作者2',
          status: 'deleted',
          create_at: '2019/08/11 11:12:12'
        }
      ],
      form: {
        title: ''
      },
      importanceOptions: [1, 2, 3],
      total: 200,
      page: {
        page: 1,
        limit: 20,
        importance: undefined,
        title: undefined,
        type: undefined,
        sort: '+id'
      },
      dialogPvVisible: false,
      loading: false
    }
  },
  created() {
    // this.fetchData()
  },
  methods: {
    fetchData() {
      this.loading = true
    },
    find() {
      console.info('find')
    },
    filter() {
      console.info('filter')
    },
    dialog() {
      this.dialogPvVisible = true
    },
    handleUpdate(row) {
      alert('handle Update')
    },
    handleModifyStatus(row, status) {
      alert('handle modify status' + JSON.stringify(row))
    }
  }
}
</script>
