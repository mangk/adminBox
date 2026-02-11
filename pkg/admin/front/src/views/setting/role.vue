<template>
  <div class="main-content">
    <el-form-item label="">
      <el-button type="primary" @click="edit()">
        <el-icon>
          <Plus />
        </el-icon>
        新建角色
      </el-button>
    </el-form-item>

    <el-table
      :data="tableData"
      row-key="id"
      height="var(--global-table)"
      border
      highlight-current-row
      show-overflow-tooltip
    >
      <el-table-column prop="id" label="ID" sortable fixed width="80" />
      <el-table-column prop="name" label="角色名称" width="260" />
      <el-table-column prop="description" label="角色描述" />
      <el-table-column fixed="right" label="操作" width="200">
        <template #default="scope">
          <el-button link type="primary" size="small" @click="edit(scope.row.id)">
            <el-icon>
              <Edit />
            </el-icon>
            编辑
          </el-button>
          <el-button link type="primary" size="small" @click="authoritySetId = scope.row.id">
            <el-icon>
              <Filter />
            </el-icon>
            权限设置
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

    <el-dialog
      v-model="dialogVisible"
      :title="form.id ? '编辑角色' : '新建角色'"
      width="80vw"
      append-to-body
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
        <el-form-item label="角色名称" prop="name">
          <el-input v-model="form.name" />
        </el-form-item>
        <el-form-item label="角色描述" prop="description">
          <el-input v-model="form.description" />
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

    <PremissionSet v-model="authoritySetId" type="sys_role" />
  </div>
</template>

<script setup>
import { getCurrentInstance, onMounted, reactive, ref } from 'vue'
import { roleCreate, roleDelete, roleDetail, rolePage, roleUpdate } from '@/api/role.js'
import PremissionSet from '@/views/setting/premissionSet.vue'

const authoritySetId = ref(0)

const { proxy } = getCurrentInstance()

const tableData = ref([])
const page = ref(1)
const pageSize = ref(20)
const total = ref(0)

const dialogVisible = ref(false)
const formRef = ref()
const form = reactive({
  id: 0,
  name: '',
  description: ''
})

const rules = reactive({
  name: [{ required: true, message: '角色名称 不能为空', trigger: 'blur' }]
})

const loadData = () => {
  rolePage(page.value, pageSize.value).then((res) => {
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
    roleDetail(id).then((res) => {
      if (res.code === 0) {
        form.id = res.data.id
        form.name = res.data.name
        form.description = res.data.description
      }
    })
  } else {
    form.id = 0
  }
  dialogVisible.value = true
}

const del = (id) => {
  if (!id) {
    proxy.$message.error('请选择数据ID')
  } else {
    roleDelete(id).then((res) => {
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
        req = roleUpdate(form)
      } else {
        req = roleCreate(form)
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

onMounted(() => {
  loadData()
})
</script>

<style scoped></style>
