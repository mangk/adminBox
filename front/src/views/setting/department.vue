<template>
  <div class="main-content">
    <el-form-item label="">
      <el-button type="primary" @click="edit()">
        <el-icon>
          <Plus />
        </el-icon>
        新建部门
      </el-button>
    </el-form-item>

    <el-table
      :data="tableData"
      row-key="id"
      height="var(--global-table)"
      border
      highlight-current-row
      show-overflow-tooltip
      default-expand-all
    >
      <el-table-column prop="id" label="ID" sortable fixed />
      <el-table-column prop="name" label="部门名称" />
      <el-table-column prop="description" label="部门简介" />
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
            :title="'删除后不可恢复，确定删除菜单【' + scope.row.name + '】?'"
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
      :title="form.id ? '编辑部门' : '新建部门'"
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
        inline
      >
        <el-form-item label="父级部门" prop="pid" style="width: 40%">
          <el-cascader
            v-model="form.pid"
            :options="departmentOption"
            :props="{
              expandTrigger: 'hover',
              checkStrictly: true,
              value: 'id',
              label: 'name',
              emitPath: false
            }"
            style="width: 100%"
          />
        </el-form-item>
        <el-form-item label="部门名称" prop="name" style="width: 40%">
          <el-input v-model="form.name" />
        </el-form-item>
        <el-form-item label="部门简介" prop="description" style="width: 40%">
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

    <PremissionSet v-model="authoritySetId" type="sys_department" />
  </div>
</template>

<script setup>
import { getCurrentInstance, onMounted, reactive, ref } from 'vue'
import * as ElementPlusIconsVue from '@element-plus/icons-vue'
import {
  departmentCreate,
  departmentDelete,
  departmentDetail,
  departmentPage,
  departmentUpdate
} from '@/api/department.js'
import PremissionSet from '@/views/setting/premissionSet.vue'

const { proxy } = getCurrentInstance()

const authoritySetId = ref(0)

const dialogVisible = ref(false)
const tableData = ref([])

const iconList = reactive([])
for (const [component] of Object.entries(ElementPlusIconsVue)) {
  iconList.push({
    key: component.__name,
    label: component.__name
  })
}

const formRef = ref()
const form = reactive({
  id: 0,
  pid: '',
  name: '',
  description: ''
})

const rules = reactive({
  pid: [
    {
      required: true,
      trigger: 'blur',
      validator: (rule, value, callback) => {
        console.log(rule, value, callback)
        if (value === '') {
          callback(new Error('父级部门 不能为空'))
          return
        }
        if (value === form.id) {
          callback(new Error('所选部门不能为当前编辑部门'))
          return
        }
        callback()
      }
    }
  ],
  name: [{ required: true, message: '路由名称 不能为空', trigger: 'blur' }]
})

const submitForm = async (formEl) => {
  if (!formEl) return
  await formEl.validate((valid, fields) => {
    if (!valid) {
      console.log('error submit!', fields)
    } else {
      let req = {}
      console.log(form.id)
      if (form.id) {
        req = departmentUpdate(form)
      } else {
        req = departmentCreate(form)
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

const edit = (id = false) => {
  if (id) {
    form.id = id
    departmentDetail(id).then((res) => {
      if (res.code === 0) {
        form.pid = res.data.pid
        form.name = res.data.name
        form.description = res.data.description
      }
    })
  } else {
    form.id = 0
  }
  setOptions()
  dialogVisible.value = true
}

const del = (id) => {
  if (!id) {
    proxy.$message.error('请选择数据ID')
  } else {
    departmentDelete(id).then((res) => {
      if (res.code === 0) {
        proxy.$message.success(res.msg)
        loadData()
      }
    })
  }
  console.log('del', id)
}

const loadData = () => {
  departmentPage().then((res) => {
    tableData.value = res.data
  })
}

const departmentOption = ref([
  {
    id: 0,
    name: '根目录'
  }
])

const setOptions = () => {
  departmentOption.value = [
    {
      id: 0,
      name: '根目录'
    }
  ]
  setMenuOptions(tableData.value, departmentOption.value, false)
}

const setMenuOptions = (menuData, optionsData, disabled) => {
  menuData &&
    menuData.forEach((item) => {
      if (item.children && item.children.length) {
        const option = {
          name: item.name,
          id: item.id,
          disabled: disabled || item.id === form.id,
          children: []
        }
        setMenuOptions(item.children, option.children, disabled || item.id === form.id)
        optionsData.push(option)
      } else {
        const option = {
          name: item.name,
          id: item.id,
          disabled: disabled || item.id === form.id
        }
        optionsData.push(option)
      }
    })
}

onMounted(() => {
  loadData()
})
</script>

<style lang="scss" scoped>

.icon-select {
  :deep(.el-select__placeholder) {
    padding-left: 20px;
  }
}

.el-form-item {
  margin-bottom: var(--global-padding);
}

.el-dialog {
  .el-form-item {
    margin: 0 16px calc(var(--global-padding) * 2);
  }
}
</style>
