<template>
  <div class="app-container">
    <el-button type="text" @click="dialogFormVisible = true">添加分类</el-button>

    <el-dialog title="添加分类" :visible.sync="dialogFormVisible">
      <el-form :model="form">
        <el-form-item label="分类名称" :label-width="formLabelWidth">
          <el-input v-model="form.name" autocomplete="off" />
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="dialogFormVisible = false">取 消</el-button>
        <el-button type="primary" @click="addCategory">确 定</el-button>
      </div>
    </el-dialog>
    <el-dialog title="编辑分类" :visible.sync="dialogEditFormVisible">
      <el-form :model="editForm">
        <el-form-item label="分类名称" :label-width="formLabelWidth">
          <el-input v-model="editForm.name" autocomplete="off" />
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="dialogFormVisible = false">取 消</el-button>
        <el-button type="primary" @click="handleDialog">确 定</el-button>
      </div>
    </el-dialog>
    <el-table
      v-loading="listLoading"
      :data="list"
      element-loading-text="Loading"
      border
      fit
      highlight-current-row
    >
      <el-table-column align="center" label="ID" width="50">
        <template slot-scope="scope">
          {{ scope.row.id }}
        </template>
      </el-table-column>
      <el-table-column label="名称">
        <template slot-scope="scope">
          {{ scope.row.name }}
        </template>
      </el-table-column>
      <el-table-column label="博客数量" width="80" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.count }}</span>
        </template>
      </el-table-column>
      <el-table-column label="创建时间" width="200" align="center">
        <template slot-scope="scope">
          {{ scope.row.create_time }}
        </template>
      </el-table-column>
      <el-table-column label="是否启用" width="100" align="center">
        <template slot-scope="scope">
          <el-button
            size="mini"
            @click="handleEnable(scope.row.id)"
          >
            <div v-if="scope.row.enable">禁用</div>
            <div v-else>启用</div>
          </el-button>
        </template>
      </el-table-column>
      <el-table-column label="操作" width="200">
        <template slot-scope="scope">

          <el-button
            size="mini"
            type="danger"
            @click="openEditDialog(scope.row.id, scope.row.name)"
          >编辑</el-button>
          <el-button
            size="mini"
            type="danger"
            @click="handleDelete(scope.row.id)"
          >删除</el-button>
        </template>
      </el-table-column>
    </el-table>
    <el-pagination
      layout="prev, pager, next"
      :page-size="pageSize"
      :total="count"
      @current-change="pageChange"
    />
  </div>
</template>

<script>
import { getList } from '@/api/category'
import { deleteCategory, editCategory, enableCategory, saveCategory } from '../../api/category'

export default {
  filters: {
    statusFilter(status) {
      const statusMap = {
        published: 'success',
        draft: 'gray',
        deleted: 'danger'
      }
      return statusMap[status]
    }
  },
  data() {
    return {
      // list: null,
      count: 0,
      pageSize: 10,
      page: 1,
      listLoading: false,
      form: {
        name: ''
      },
      editForm: {
        id: 0,
        name: ''
      },
      dialogFormVisible: false,
      dialogEditFormVisible: false,
      formLabelWidth: '120px'
    }
  },
  created() {
    this.fetchData()
  },
  methods: {
    fetchData() {
      this.listLoading = true
      getList({ page: this.page, page_size: this.pageSize }).then(response => {
        this.list = response.data
        this.listLoading = false
        this.count = response.count
      })
    },
    addDialog() {
      this.form.name = ''
    },
    addCategory() {
      saveCategory({ name: this.form.name }).then(res => {
        this.fetchData()
        this.dialogFormVisible = false
      })
    },
    handleEnable(id) {
      enableCategory({ id: id }).then(res => {
        this.fetchData()
      })
    },
    handleDelete(id) {
      deleteCategory({ id: id }).then(res => {
        this.fetchData()
      })
    },
    openEditDialog(id, name) {
      console.log('id = ', id, 'name = ', name)
      this.editForm.id = id
      this.editForm.name = name
      this.dialogEditFormVisible = true
    },
    handleDialog() {
      editCategory({ id: this.editForm.id, name: this.editForm.name }).then(res => {
        this.fetchData()
        this.dialogEditFormVisible = false
      })
    },
    pageChange(page) {
      this.page = page
      this.fetchData()
    }
  }
}
</script>
