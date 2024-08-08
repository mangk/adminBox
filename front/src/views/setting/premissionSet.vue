<template>
  <el-drawer
    v-model="isVisible"
    direction="rtl"
    :before-close="closeDrawer"
    :with-header="false"
    size="90%"
  >
    <el-table
      :header-cell-class-name="'global-el-table-header'"
      :data="permissionMap"
      height="calc(100vh - 62px - 40px )"
      style="width: 100%"
      row-key="id"
      border
      default-expand-all
    >
      <el-table-column label="菜单 · API · 动作" width="190px" fixed>
        <template #default="scope">
          <el-icon>
            <Menu v-if="scope.row.type == 'menu'" />
            <Link v-if="scope.row.type == 'api'" />
            <Pointer v-if="scope.row.type == 'action'" />
          </el-icon>
          {{ scope.row.name }}
        </template>
      </el-table-column>

      <el-table-column label="结果展示" fixed width="80px">
        <template #default="scope">
          <el-icon style="color: #67c23a" v-if="permissionResult[scope.row.id] > 0">
            <Check />
          </el-icon>
          <el-icon style="color: #f56c6c" v-if="permissionResult[scope.row.id] < 0">
            <Close />
          </el-icon>
        </template>
      </el-table-column>

      <el-table-column label="当前设置" fixed width="220px">
        <template #default="scope">
          <el-radio-group v-model="permissionSet[scope.row.id]">
            <el-radio size="small" :value="1">允许</el-radio>
            <el-radio size="small" :value="-1">拒绝</el-radio>
            <el-radio size="small" :value="0">不限</el-radio>
          </el-radio-group>
        </template>
      </el-table-column>

      <el-table-column
        v-for="(item, index) in permissionOther"
        width="120px"
        :key="index"
        :label="item.record_name"
      >
        <template #default="scope">
          <el-icon
            style="color: #67c23a"
            v-if="item.list && item.list[scope.row.id] && item.list[scope.row.id] > 0"
          >
            <Check />
          </el-icon>
          <el-icon
            style="color: #f56c6c"
            v-if="item.list && item.list[scope.row.id] && item.list[scope.row.id] < 0"
          >
            <Close />
          </el-icon>
        </template>
      </el-table-column>
    </el-table>

    <template #footer>
      <div style="display: flex; flex-flow: row nowrap; justify-content: space-between">
        <div>
          <el-button @click="save" type="primary">保存</el-button>
          <el-button @click="closeDrawer">关闭</el-button>
        </div>
        <div style="text-align: left">
          <div style="color: #f56c6c; font-size: 12px">
            <el-icon>
              <WarningFilled />
            </el-icon>
            权限分级说明: 个人设置 > 角色设置 > 部门设置
          </div>
          <div style="color: #e6a23c; font-size: 12px">
            <el-icon>
              <WarningFilled />
            </el-icon>
            权限同级说明: 拒绝 > 允许;
          </div>
        </div>
      </div>
    </template>
  </el-drawer>
</template>

<script setup>
import { defineModel, reactive, ref, watch } from 'vue'
import { permissionAll, permissionGetByIdAndModule, permissionSave } from '@/api/auth'

const id = defineModel({ type: Number })
const type = defineModel('type', { type: String })

const isVisible = ref(false)
const permissionMap = ref([]) // 权限树
const permissionSet = reactive({}) // 手动设置权限
const permissionOther = ref([]) // 从角色或部门继承权限
const permissionResult = reactive({}) // 结果展示

// 页面打开，监听value变化拉取接口数据
watch(id, (newId) => {
  if (newId) {
    isVisible.value = true
    permissionAll().then((res) => {
      permissionMap.value = res.data
    })
    permissionGetByIdAndModule(id.value, type.value).then((res) => {
      permissionOther.value = res.data?.ohter_set_list
      Object.assign(permissionSet, res.data?.cur_set?.list)
      Object.assign(permissionResult, res.data?.result_set?.list)
    })
  }
})

watch(permissionSet, (newVal) => {
  for (const key in newVal) {
    if (Object.hasOwnProperty.call(newVal, key)) {
      permissionResult[key] = checkResule(key)
    }
  }
})

// 计算具体是允许权限，还是禁止权限
const checkResule = (key) => {
  var fromSet = 0
  var fromRole = []
  var fromDepartment = []
  if (permissionSet[key]) {
    fromSet = permissionSet[key]
  }
  for (const i in permissionOther.value) {
    if (Object.hasOwnProperty.call(permissionOther.value, i)) {
      const element = permissionOther.value[i]
      if (element.module == 'sys_role' && element.list && element.list[key]) {
        fromRole.push(element.list[key])
      }
      if (element.module == 'sys_department' && element.list && element.list[key]) {
        fromDepartment.push(element.list[key])
      }
    }
  }

  var role = getPremission(fromRole)
  var department = getPremission(fromDepartment)

  var resule = 0
  if (department !== 0) {
    resule = department
  }
  if (role !== 0) {
    // 角色权限覆盖部门权限
    resule = role
  }
  if (fromSet !== 0) {
    // 对于用户，手动设置权限最高，覆盖角色，覆盖部门
    resule = fromSet
  }

  return resule
}

// 同类型数据计算权限允许还是禁止
const getPremission = (data) => {
  // 同类型权限内部，拒绝 > 允许
  var hasOne = 0
  for (let i = 0; i < data.length; i++) {
    const e = data[i]
    if (e == -1) {
      return -1
    }
    if (e == 1) {
      hasOne++
    }
  }
  if (hasOne > 0) {
    return 1
  }
  return 0
}

// 保存
const save = () => {
  permissionSave({ id: id.value, module: type.value, list: permissionSet }).then(() => {
    closeDrawer()
  })
}

// 页面关闭
const closeDrawer = () => {
  isVisible.value = false
  id.value = 0
  permissionOther.value = []
  for (const key in permissionSet) {
    if (Object.hasOwnProperty.call(permissionSet, key)) {
      delete permissionSet[key]
    }
  }
  for (const key in permissionResult) {
    if (Object.hasOwnProperty.call(permissionResult, key)) {
      delete permissionResult[key]
    }
  }
}
</script>

<style scoped>
:deep(.cell) {
  padding: 0 10px;
}

:deep(.el-table__cell) {
  padding: 0;
}

.select-item-box {
  border-top: 1px solid #eceef5;
  text-align: right;
  padding: 0 var(--global-padding);
  height: 30px;
  line-height: 30px;
  box-sizing: border-box;
}

.with-out-border-top {
  border-top: 0;
}
</style>
