<template>
  <div class="upload-file-slot" @click="showDrawer">
    <slot>
      <el-button type="primary">
        <el-icon>
          <UploadFilled />
        </el-icon>
        文件上传
      </el-button>
    </slot>
  </div>

  <el-drawer
    v-model="show"
    direction="rtl"
    :before-close="handleClose"
    :show-close="true"
    append-to-body
    size="50%"
    style="max-width: 600px"
  >
    <template #header>
      <div style="padding: 0px var(--global-padding)">
        <el-button type="primary" @click="selectOK">确定</el-button>
        <div v-if="Object.keys(uploadConfig).length > 1">
          <el-radio-group v-model="driver" size="small">
            <el-radio-button
              :label="item.name"
              :value="item.driver"
              v-for="(item, index) in uploadConfig"
              :key="item.driver"
            />
          </el-radio-group>
        </div>
      </div>
    </template>

    <div style="display: flex; flex-flow: row wrap; padding: var(--global-padding)">
      <div
        :class="['img-item', checkSelect(item)]"
        v-for="(item, index) in tableData"
        @mouseleave="mouseLeave(item.id)"
        @mouseover="mouseOver(item.id)"
        @click="selectFile(item)"
      >
        <el-icon class="img-del" v-if="curMouseOnId == item.id" @click="delImage(item)">
          <CircleClose />
        </el-icon>
        <el-image style="width: 100%" :src="item.url" fit="scale-down" lazy>
          <template #error>
            <div
              style="
                font-size: 30px;
                text-align: center;
                margin: var(--global-padding);
                margin-bottom: 0;
              "
            >
              <el-icon style="font-size: 28px">
                <Document />
              </el-icon>
            </div>
            <div style="font-size: 12px; line-height: 12px; padding: 3px; text-align: center">
              {{ item.name.split('.')[0] }}
            </div>
            <div
              style="
                font-size: 12px;
                line-height: 12px;
                padding: 3px;
                position: absolute;
                left: 0;
                top: 0;
                background-color: #409eff;
                color: #fff;
                border-radius: 0 0 5px 0;
              "
            >
              {{ item.name.split('.')[1] }}
            </div>
          </template>
        </el-image>
      </div>
    </div>

    <template #footer>
      <el-pagination
        v-model:current-page="page"
        v-model:page-size="pageSize"
        :page-sizes="[20, 50, 100, 200]"
        :size="'small'"
        layout="total, prev, pager, next"
        :total="total"
        @size-change="handleSizeChange"
        @current-change="handleCurrentChange"
      />
      <el-upload
        drag
        :action="`${serverHost()}sys/fileUpload/upload?driver=${driver}`"
        :headers="{ Authorization: 'Bearer ' + userStore.userAuth().token }"
        :show-file-list="false"
        multiple
        :on-success="uploadSuccess"
      >
        <div
          class="el-upload__text"
          style="
            height: 60px;
            line-height: 60px;
            text-align: center;
            overflow: hidden;
            display: flex;
            flex-flow: row nowrap;
            justify-content: center;
          "
        >
          <el-icon
            class="el-icon--upload"
            style="font-size: 35px; margin: 10px var(--global-padding) 0 0"
          >
            <upload-filled />
          </el-icon>
          将文件拖到此处或<em>点击上传</em>
        </div>
      </el-upload>
    </template>
  </el-drawer>
</template>
<script setup>
import { ref, watch } from 'vue'
import { fileDelete, fileUploadCfg, fileUploadPage } from '@/api/fileUpload'
import { useUserStore } from '@/pinia/useUserStore'
import { ElMessage, ElMessageBox } from 'element-plus'
import { serverHost } from '@/utils/requester'

const uploadConfig = ref({})
const driver = ref('default')
const userStore = useUserStore()

const props = defineProps({
  modelValue: {
    type: [String, Array], // 修正为数组格式，支持多种类型
    required: true,
    default: '' // 默认值为字符串
  },
  selected: {
    type: Function, // 修正为 Function 类型
    required: false,
    default: () => {} // 默认值为一个空函数
  },
  multiple: {
    type: Number,
    default: 1
  },
  // fullInfo: {
  //   type: Boolean,
  //   default: false,
  // },
  suffixList: {
    type: Array, // 修正 Array 的默认值使用函数返回
    default: () => [] // 返回一个空数组
  }
})

// 定义 emit 事件
const emit = defineEmits(['update:modelValue'])

// 创建内部状态，用来处理双向绑定的数据
const internalValue = ref(props.modelValue)
if (!Array.isArray(internalValue.value)) {
  if (internalValue.value) {
    internalValue.value = [internalValue.value]
  } else {
    internalValue.value = []
  }
} else {
  let v = []
  for (const key in internalValue.value) {
    if (Object.prototype.hasOwnProperty.call(internalValue.value, key)) {
      const element = internalValue.value[key]
      if (element) {
        v.push(element)
      }
    }
  }
  internalValue.value = v
}

// 监听 modelValue 的变化，确保内部状态保持同步
// watch(
//   () => props.modelValue,
//   (newVal) => {
//     console.log(999, newVal);

//     internalValue.value = newVal
//   }
// )

const selectOK = () => {
  let value = []
  for (const key in internalValue.value) {
    if (Object.prototype.hasOwnProperty.call(internalValue.value, key)) {
      const element = internalValue.value[key]
      value.push(element)
    }
  }

  if (props.multiple == 1) {
    value = value[0]
  }

  // 通过 emit 通知父组件更新 modelValue
  emit('update:modelValue', value)
  // 如果传入了 selected 回调，执行该回调
  if (props.selected) {
    props.selected(value)
  }
  show.value = false
}

const show = ref(false)

const selectFile = (item) => {
  let delFlag = false
  for (let index = 0; index < internalValue.value.length; index++) {
    const element = internalValue.value[index]
    if (element && element == item.url) {
      internalValue.value.splice(index, 1)
      delFlag = true
    }
  }
  if (internalValue.value.length >= props.multiple) {
    ElMessage({ type: 'warning', message: `最多选择 ${props.multiple} 个` })
    return
  }
  if (!delFlag) {
    internalValue.value.push(item.url)
  }
}

const checkSelect = (item) => {
  for (const key in internalValue.value) {
    if (Object.prototype.hasOwnProperty.call(internalValue.value, key)) {
      const element = internalValue.value[key]
      if (element && element == item.url) {
        return 'img-item-select'
      }
    }
  }
}

const uploadSuccess = (response, uploadFile, uploadFiles) => {
  loadData()
}

const delImage = (item) => {
  ElMessageBox.confirm('删除文件?', '删除', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    fileDelete(item.id).then((res) => {
      loadData()
      for (let index = 0; index < internalValue.value.length; index++) {
        const element = internalValue.value[index]
        if (element == item.url) {
          internalValue.value.splice(index, 1)
        }
      }
    })
  })
}

const tableData = ref([])
const page = ref(1)
const pageSize = ref(25)
const total = ref(0)

const loadData = () => {
  fileUploadPage(page.value, pageSize.value, { tag: props.suffixList }).then((res) => {
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

const curMouseOnId = ref(0)
const mouseLeave = (id) => {
  curMouseOnId.value = 0
}
const mouseOver = (id) => {
  curMouseOnId.value = id
}

const showDrawer = () => {
  fileUploadCfg().then((res) => {
    uploadConfig.value = res.data
  })
  loadData()
  show.value = true
}

const handleClose = () => {
  show.value = false
}
</script>
<style lang="scss">
.el-upload-dragger {
  padding: var(--global-padding);
}

.el-upload-dragger.is-dragover {
  padding: calc(var(--global-padding) - 1px);
}

.el-drawer__header {
  padding: var(--global-padding);
  padding-bottom: 0;
  margin-bottom: var(--global-padding);
}

.el-drawer__body {
  padding: 0 var(--global-padding);
}

.img-item {
  position: relative;
  width: calc(20% - 1px * 2 - var(--global-padding));
  height: 0;
  padding-bottom: calc(20% - 1px * 2 - var(--global-padding));
  border: 1px dashed var(--el-border-color);
  margin: calc(var(--global-padding) / 2);
  border-radius: 5px;
  overflow: hidden;
}

.img-item-select {
  border: 2px dashed var(--el-color-primary);
  width: calc(20% - 2px * 2 - var(--global-padding));
  padding-bottom: calc(20% - 2px * 2 - var(--global-padding));
}

.upload-file-slot {
  margin: 0;
  padding: 0;
  line-height: 0;
}

.img-del {
  position: absolute;
  z-index: 20;
  font-size: 16px;
  top: 5px;
  right: 5px;
  color: #fff;
  background-color: var(--el-color-error);
  border-radius: 100px;
}
</style>
