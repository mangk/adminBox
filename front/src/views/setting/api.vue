<template>
  <div class="main-content">
    <el-form-item label="">
      <el-button type="primary" @click="edit()">
        <el-icon>
          <Plus />
        </el-icon>
        新建API
      </el-button>
    </el-form-item>

    <el-table
      :data="tableData"
      row-key="id"
      height="var(--global-table)"
      :header-cell-class-name="'global-el-table-header'"
      border
      highlight-current-row
      show-overflow-tooltip
    >
      <el-table-column prop="id" label="ID" sortable fixed width="80" />
      <el-table-column prop="name" label="api名称" />
      <el-table-column prop="description" label="api描述" width="160" />
      <el-table-column prop="path" label="api路径" width="260" />
      <el-table-column label="方法">
        <template #default="scope">
          <el-tag type="info" v-if="scope.row.method === 'GET'">{{ scope.row.method }}</el-tag>
          <el-tag type="success" v-if="scope.row.method === 'POST'">{{ scope.row.method }}</el-tag>
          <el-tag type="warning" v-if="scope.row.method === 'PUT'">{{ scope.row.method }}</el-tag>
          <el-tag type="danger" v-if="scope.row.method === 'DELETE'">{{ scope.row.method }}</el-tag>
        </template>
      </el-table-column>
      <el-table-column prop="menu_name" label="所属菜单" width="120" />
      <el-table-column fixed="right" label="操作" width="125">
        <template #default="scope">
          <el-button link type="primary" size="small" @click="edit(scope.row.id)">
            <el-icon>
              <Edit />
            </el-icon>
            编辑
          </el-button>
          <el-popconfirm
            v-if="!scope.row.children"
            :title="'删除后不可恢复，确定删除API【' + scope.row.name + '】?'"
            @confirm="del(scope.row.id)"
            width="200"
          >
            <template #reference>
              <el-button link type="primary" size="small">
                <el-icon>
                  <Delete />
                </el-icon>
                删除
              </el-button>
            </template>
          </el-popconfirm>
        </template>
      </el-table-column>
    </el-table>

    <el-dialog
      v-model="dialogVisible"
      :title="form.id ? '编辑API' : '新建API'"
      width="800"
      @close="cancel(formRef)"
    >
      <el-form
        ref="formRef"
        :model="form"
        :rules="rules"
        status-icon
        label-position="top"
        :show-all-levels="false"
      >
        <el-form-item label="父级菜单" prop="menuId">
          <el-cascader
            v-model="form.menu_id"
            :options="menuOption"
            :props="{
              expandTrigger: 'hover',
              checkStrictly: true,
              value: 'id',
              label: 'title',
              emitPath: false
            }"
            style="width: 100%"
          />
        </el-form-item>
        <el-form-item label="API名称" prop="name">
          <el-input v-model="form.name" />
        </el-form-item>
        <el-form-item label="API描述" prop="description">
          <el-input v-model="form.description" />
        </el-form-item>
        <el-form-item label="API地址" prop="path">
          <el-input v-model="form.path" />
        </el-form-item>
        <el-form-item label="API方法" prop="method">
          <el-select v-model="form.method">
            <el-option label="GET" value="GET" />
            <el-option label="POST" value="POST" />
            <el-option label="PUT" value="PUT" />
            <el-option label="DELETE" value="DELETE" />
          </el-select>
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="resetForm(formRef)">重置</el-button>
          <el-button @click="cancel(formRef)">取消</el-button>
          <el-button type="primary" @click="submitForm(formRef)">保存</el-button>
        </div>
      </template>
    </el-dialog>

    <el-pagination
      v-model:current-page="page"
      v-model:page-size="pageSize"
      :page-sizes="[20, 50, 100, 200]"
      :size="'small'"
      layout="total, sizes, prev, pager, next, jumper"
      :total="total"
      @size-change="handleSizeChange"
      @current-change="handleCurrentChange"
    />
  </div>
</template>

<script setup>
import { getCurrentInstance, onMounted, reactive, ref } from 'vue'
import { apiCreate, apiDelete, apiDetail, apiPage, apiUpdate } from '@/api/api.js'
import { menuPage } from '@/api/menu.js'

const { proxy } = getCurrentInstance()

const tableData = ref([])
const page = ref(1)
const pageSize = ref(20)
const total = ref(0)

const dialogVisible = ref(false)
const formRef = ref()
const form = reactive({
  id: 0,
  menu_id: 0,
  name: '',
  description: '',
  path: '',
  method: ''
})

const rules = reactive({
  name: [{ required: true, message: 'API名称 不能为空', trigger: 'blur' }],
  path: [{ required: true, message: 'API地址 不能为空', trigger: 'blur' }],
  method: [{ required: true, message: 'API方法 不能为空', trigger: 'blur' }]
})

const loadData = () => {
  apiPage(page.value, pageSize.value).then((res) => {
    tableData.value = res.data.list
    page.value = res.data.page
    pageSize.value = res.data.page_size
    total.value = res.data.total
  })
}

const handleSizeChange = (size) => {
  pageSize.value = size
  loadData()
}

const handleCurrentChange = (changePage) => {
  page.value = changePage
  loadData()
}

const edit = (id = false) => {
  if (id) {
    form.id = id
    apiDetail(id).then((res) => {
      if (res.code === 0) {
        form.id = res.data.id
        form.menu_id = res.data.menu_id
        form.name = res.data.name
        form.description = res.data.description
        form.path = res.data.path
        form.method = res.data.method
      }
    })
  } else {
    form.id = 0
  }

  loadMenuOptions()
  dialogVisible.value = true
}

const del = (id) => {
  if (!id) {
    proxy.$message.error('请选择数据ID')
  } else {
    apiDelete(id).then((res) => {
      if (res.code === 0) {
        proxy.$message.success(res.msg)
        loadData()
      }
    })
  }
  console.log('del', id)
}

const submitForm = async (formEl) => {
  if (!formEl) return
  await formEl.validate((valid, fields) => {
    if (!valid) {
      console.log('error submit!', fields)
    } else {
      let req = {}
      console.log(form.id)
      if (form.id) {
        req = apiUpdate(form)
      } else {
        req = apiCreate(form)
      }
      req.then((res) => {
        if (res.code === 0) {
          proxy.$message.success(res.msg)
          loadData()
          dialogVisible.value = false
          resetForm(formEl)
        }
      })
    }
  })
}

const cancel = (formEl) => {
  dialogVisible.value = false
  resetForm(formEl)
}

const resetForm = (formEl) => {
  if (!formEl) return
  formEl.resetFields()
}

const menuOption = ref()

const loadMenuOptions = async () => {
  menuPage(true).then((res) => {
    menuOption.value = res.data
    recursionSetDisable(menuOption.value, false)
  })
}

const recursionSetDisable = (menu, disabled) => {
  menu &&
    menu.forEach((item) => {
      if (item.children && item.children.length) {
        item.title = item.meta.title
        // item.disabled = (disabled || item.id === form.id) && form.id
        recursionSetDisable(item.children, item.disabled)
      } else {
        item.title = item.meta.title
        // item.disabled = (disabled || item.id === form.id) && form.id
      }
    })
}

onMounted(() => {
  loadData()
})
</script>

<style scoped></style>
