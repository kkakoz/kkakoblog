<template>
  <div class="app-container">
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
      <el-table-column label="博客名称">
        <template slot-scope="scope">
          {{ scope.row.title }}
        </template>
      </el-table-column>
      <el-table-column label="博客作者" width="80" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.username }}</span>
        </template>
      </el-table-column>
      <el-table-column label="博客分类" width="200" align="center">
        <template slot-scope="scope">
          {{ scope.row.category_name }}
        </template>
      </el-table-column>
      <el-table-column label="博客点赞" width="200" align="center">
        <template slot-scope="scope">
          {{ scope.row.star }}
        </template>
      </el-table-column>
      <el-table-column label="博客阅读量" width="200" align="center">
        <template slot-scope="scope">
          {{ scope.row.create_time }}
        </template>
      </el-table-column>
      <el-table-column label="博客创建时间" width="200" align="center">
        <template slot-scope="scope">
          {{ scope.row.view }}
        </template>
      </el-table-column>
      <el-table-column label="是否展示" width="100" align="center">
        <template slot-scope="scope">
          <el-button
            size="mini"
            @click="handleEnable(scope.row.id)"
          >
            <div v-if="scope.row.enable">启用</div>
            <div v-else>未启用</div>
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
      list: [
        {
          name: 'spring',
          count: '10',
          create_time: '1010'
        }
      ],
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
      getList().then(response => {
        console.log(response)
        this.list = response.data
        this.listLoading = false
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
    }
  }
}
</script>
