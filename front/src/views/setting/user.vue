<template>
  <div class="main-content">
    <el-form-item label="">
      <el-button type="primary" @click="edit()">
        <el-icon>
          <Plus />
        </el-icon>
        新增用户
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
      <el-table-column prop="id" label="ID" sortable fixed />
      <el-table-column prop="uuid" label="UUID" width="300" />
      <el-table-column prop="nick_name" label="昵称" width="120" />
      <el-table-column prop="username" label="用户名" />
      <el-table-column prop="phone" label="手机号" width="160" />
      <el-table-column prop="email" label="Email" width="220" />
      <el-table-column prop="avatar" label="头像" width="120" />
      <el-table-column prop="enable" label="是否启用" width="120" />
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
            :title="
              '删除后不可恢复，确定删除API【' +
              scope.row.nickName +
              '(' +
              scope.row.userName +
              ')】?'
            "
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
      :title="form.id ? '编辑用户' : '新建用户'"
      width="80vw"
      append-to-body
      @close="cancel(formRef)"
    >
      <el-form
        ref="formRef"
        :model="form"
        :rules="rules"
        status-icon
        inline
        label-position="top"
        :show-all-levels="false"
      >
        <el-form-item label="用户名" prop="username" style="width: 40%">
          <el-input v-model="form.username" />
        </el-form-item>
        <el-form-item label="手机号" prop="phone" style="width: 40%">
          <el-input v-model="form.phone" />
        </el-form-item>
        <el-form-item label="Email" prop="email" style="width: 40%">
          <el-input v-model="form.email" />
        </el-form-item>
        <el-form-item label="昵称" prop="nick_name" style="width: 40%">
          <el-input v-model="form.nick_name" />
        </el-form-item>
        <el-form-item label="所属部门" prop="department_ids" style="width: 40%">
          <el-cascader
            v-model="form.department_ids"
            :options="departmentOption"
            placeholder="设置用户部门"
            :props="{
              checkStrictly: true,
              expandTrigger: 'hover',
              value: 'id',
              label: 'name',
              multiple: true,
              emitPath: false
            }"
            style="width: 100%"
          />
        </el-form-item>
        <el-form-item label="用户角色" prop="role_ids" style="width: 40%">
          <el-select v-model="form.role_ids" multiple placeholder="设置用户角色">
            <el-option
              v-for="item in roleOption"
              :key="item.id"
              :label="item.name"
              :value="item.id"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="头像" prop="avatar" style="width: 40%">
          <el-image
            style="
              width: 75px;
              height: 75px;
              border: 1px dashed var(--el-border-color);
              border-radius: 5px;
            "
            :src="form.avatar"
            :fit="'contain'"
          />
          <FileUpload v-model="form.avatar" :onlyPath="true">
            <el-icon
              style="
                width: 75px;
                height: 75px;
                border: 1px dashed var(--el-border-color);
                border-radius: 5px;
                margin-left: 10px;
              "
            >
              <Plus />
            </el-icon>
          </FileUpload>
        </el-form-item>
        <el-form-item label="是否启用" prop="enable" style="width: 40%">
          <el-switch
            v-model="form.enable"
            inline-prompt
            active-icon="Check"
            inactive-icon="Close"
          />
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

    <PermissionSet v-model="authoritySetId" type="sys_user" />
  </div>
</template>

<script setup>
import { getCurrentInstance, onMounted, reactive, ref } from 'vue'
import { userCreate, userDelete, userDetail, userPage, userUpdate } from '@/api/user.js'
import PermissionSet from '@/views/setting/premissionSet.vue'
import { departmentPage } from '@/api/department.js'
import { roleAll } from '@/api/role'
import FileUpload from '../util/fileUpload.vue'

const { proxy } = getCurrentInstance()

const tableData = ref([])
const page = ref(1)
const pageSize = ref(20)
const total = ref(0)

const authoritySetId = ref(0)

const dialogVisible = ref(false)
const formRef = ref()
const form = reactive({
  id: 0,
  username: '',
  phone: '',
  email: '',
  nick_name: '',
  avatar: '',
  department_ids: [],
  role_ids: [],
  enable: true
})

const rules = reactive({
  name: [{ required: true, message: 'API名称 不能为空', trigger: 'blur' }],
  path: [{ required: true, message: 'API地址 不能为空', trigger: 'blur' }],
  method: [{ required: true, message: 'API方法 不能为空', trigger: 'blur' }]
})

const departmentOption = ref([])
const roleOption = ref([])

const loadData = () => {
  userPage(page.value, pageSize.value).then((res) => {
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
  departmentPage().then((res) => {
    departmentOption.value = res.data
  })
  roleAll().then((res) => {
    roleOption.value = res.data
  })

  if (id) {
    form.id = id
    userDetail(id).then((res) => {
      if (res.code === 0) {
        form.id = res.data.id
        form.username = res.data.username
        form.phone = res.data.phone
        form.email = res.data.email
        form.nick_name = res.data.nick_name
        form.avatar = res.data.avatar
        form.enable = res.data.enable
        form.department_ids = res.data.department_ids
        form.role_ids = res.data.role_ids
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
    userDelete(id).then((res) => {
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

      if (form.id) {
        req = userUpdate(form)
      } else {
        req = userCreate(form)
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
