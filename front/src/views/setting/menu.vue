<template>
  <div class="main-content">
    <el-form-item label="">
      <el-button type="primary" @click="edit()">
        <el-icon>
          <Plus />
        </el-icon>
        新建目录
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
      <el-table-column label="菜单名称" min-width="120">
        <template #default="scope">
          <span>{{ scope.row.meta.title }}</span>
        </template>
      </el-table-column>
      <el-table-column label="图标" min-width="140">
        <template #default="scope">
          <div v-if="scope.row.meta.icon" class="icon-column">
            <el-icon>
              <component :is="scope.row.meta.icon" />
            </el-icon>
            <span>{{ scope.row.meta.icon }}</span>
          </div>
        </template>
      </el-table-column>
      <el-table-column prop="name" label="路由名称" width="120" />
      <el-table-column prop="path" label="路由地址" width="120" />
      <el-table-column prop="hidden" label="是否隐藏" />
      <el-table-column prop="pid" label="父节点" />
      <el-table-column prop="sort" label="排序" width="120" sortable />
      <el-table-column prop="component" label="模版路经" width="260" />
      <el-table-column label="服务端模版" width="260">
        <template #default="scope">
          {{ scope.row.meta ? scope.row.meta.sc_path : '' }}
        </template>
      </el-table-column>
      <el-table-column prop="action_list" label="动作列表" width="260" />
      <el-table-column fixed="right" label="操作" width="210">
        <template #default="scope">
          <el-button link type="primary" size="small" @click="edit(scope.row.id)">
            <el-icon>
              <Edit />
            </el-icon>
            编辑
          </el-button>
          <el-button link type="primary" size="small">
            <el-icon>
              <Cherry />
            </el-icon>
            编辑动作
          </el-button>
          <el-popconfirm
            v-if="!scope.row.children"
            :title="'删除后不可恢复，确定删除菜单【' + scope.row.meta.title + '】?'"
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
      :title="form.id ? '编辑菜单' : '新建菜单'"
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
        inline
      >
        <el-form-item label="父级菜单" prop="pid" style="width: 40%">
          <el-cascader
            v-model="form.pid"
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
        <el-form-item label="菜单名称" prop="meta.title" style="width: 40%">
          <el-input v-model="form.meta.title" />
        </el-form-item>
        <el-form-item label="路由名称(name)" prop="name" style="width: 40%">
          <el-input v-model="form.name" />
        </el-form-item>
        <el-form-item label="路由地址(path)" prop="path" style="width: 40%">
          <el-input v-model="form.path" />
        </el-form-item>
        <el-form-item label="图标" prop="meta.icon" style="width: 40%">
          <span
            style="position: absolute; z-index: 9999; padding: 4px 10px 0"
            v-if="form.meta.icon"
          >
            <el-icon>
              <component :is="form.meta.icon" />
            </el-icon>
          </span>
          <el-select v-model="form.meta.icon" clearable class="icon-select">
            <el-option
              v-for="item in iconList"
              :key="item.key"
              :label="item.key"
              :value="item.label"
            >
              <span class="gva-icon" style="padding: 3px 0 0" :class="item.label">
                <el-icon>
                  <component :is="item.label" />
                </el-icon>
              </span>
              <span style="text-align: left; margin: 0 0 0 10px">{{ item.key }}</span>
            </el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="排序" prop="sort" style="width: 40%">
          <el-input type="number" v-model.number="form.sort" />
        </el-form-item>
        <el-form-item label="模版地址" prop="meta.component" style="width: 40%">
          <el-input v-model="form.component" placeholder="views/util/serverComponent.vue" />
        </el-form-item>
        <el-form-item label="服务端模版地址" prop="meta.sc_path" style="width: 40%">
          <el-input v-model="form.meta.sc_path" />
        </el-form-item>
        <el-form-item label="是否隐藏" prop="hidden" style="width: 40%">
          <el-select v-model="form.hidden">
            <el-option label="否" :value="false" />
            <el-option label="是" :value="true" />
          </el-select>
        </el-form-item>
        <el-form-item label="KeepAlive" prop="meta.keep_alive" style="width: 40%">
          <el-select v-model="form.meta.keep_alive">
            <el-option label="否" :value="false" />
            <el-option label="是" :value="true" />
          </el-select>
        </el-form-item>
        <el-form-item label="基础页面" prop="meta.default_menu" style="width: 40%">
          <el-select v-model="form.meta.default_menu">
            <el-option label="否" :value="false" />
            <el-option label="是" :value="true" />
          </el-select>
        </el-form-item>
        <el-form-item label="自动关闭标签" prop="meta.auto_close" style="width: 40%">
          <el-select v-model="form.meta.auto_close">
            <el-option label="否" :value="false" />
            <el-option label="是" :value="true" />
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
  </div>
</template>

<script setup>
import { menuCreate, menuDelete, menuDetail, menuPage, menuUpdate } from '@/api/menu.js'
import { getCurrentInstance, onMounted, reactive, ref } from 'vue'
import * as ElementPlusIconsVue from '@element-plus/icons-vue'

const { proxy } = getCurrentInstance()

const dialogVisible = ref(false)
const tableData = ref([])
const xxx = ref('')

const iconList = reactive([])
for (const [component] of Object.entries(ElementPlusIconsVue)) {
  iconList.push({
    key: component,
    label: component
  })
}

const formRef = ref()
const form = reactive({
  id: 0,
  pid: '',
  name: '',
  path: '',
  hidden: false,
  component: '',
  sort: 0,
  meta: {
    title: '',
    keep_alive: false,
    default_menu: false,
    icon: '',
    auto_close: false,
    sc_path: '',
    action_list: null
  }
})

const validateFields = (rule, value, callback) => {
  if (!form.component && !form.meta.sc_path) {
    callback(new Error('模版地址 和 服务端模版地址 不能同时为空'))
  } else {
    callback()
  }
}

const rules = reactive({
  pid: [
    {
      required: true,
      trigger: 'blur',
      validator: (rule, value, callback) => {
        if (value === '') {
          callback(new Error('父级菜单 不能为空'))
          return
        }
        if (value === form.id) {
          callback(new Error('所选菜单不能为当前编辑菜单'))
          return
        }
        callback()
      }
    }
  ],
  name: [{ required: true, message: '路由名称 不能为空', trigger: 'blur' }],
  path: [{ required: true, message: '路由地址 不能为空', trigger: 'blur' }],
  hidden: [{ required: true, message: '请选择 是否隐藏', trigger: 'blur' }],
  'meta.title': [{ required: true, message: '菜单名称 不能为空', trigger: 'blur' }],
  'meta.keep_alive': [{ required: true, message: '请选择 是否KeepAlive', trigger: 'blur' }],
  'meta.default_menu': [{ required: true, message: '请选择 是否为基础页面', trigger: 'blur' }],
  'meta.auto_close': [{ required: true, message: '请选择 是否自动关闭标签', trigger: 'blur' }]
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
        req = menuUpdate(form)
      } else {
        req = menuCreate(form)
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
    menuDetail(id).then((res) => {
      if (res.code === 0) {
        form.pid = res.data.pid
        form.meta.title = res.data.meta.title
        form.name = res.data.name
        form.path = res.data.path
        form.hidden = res.data.hidden
        form.component = res.data.component
        form.sort = res.data.sort
        form.meta = res.data.meta
      }
    })
  } else {
    form.id = 0
  }

  setDisableOption()
  dialogVisible.value = true
}

const del = (id) => {
  if (!id) {
    proxy.$message.error('请选择数据ID')
  } else {
    menuDelete(id).then((res) => {
      if (res.code === 0) {
        proxy.$message.success(res.msg)
        loadData()
      }
    })
  }
  console.log('del', id)
}

const loadData = () => {
  menuPage().then((res) => {
    tableData.value = res.data[0].children
    menuOption.value = res.data
  })
}

const menuOption = ref([])

const setDisableOption = () => {
  recursionSetDisable(menuOption.value, false)
}

const recursionSetDisable = (menu, disabled) => {
  menu &&
    menu.forEach((item) => {
      if (item.children && item.children.length) {
        item.title = item.meta.title
        item.disabled = (disabled || item.id === form.id) && form.id
        recursionSetDisable(item.children, item.disabled)
      } else {
        item.title = item.meta.title
        item.disabled = (disabled || item.id === form.id) && form.id
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
