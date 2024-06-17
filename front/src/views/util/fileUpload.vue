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
    :show-close="false"
    size="50%;"
  >
    <template #header>
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
    </template>

    <div style="display: flex; flex-flow: row wrap; padding: var(--global-padding)">
      <div
        :class="['img-item', checkSelect(item.id)]"
        v-for="(item, index) in tableData"
        @mouseleave="mouseLeave(item.id)"
        @mouseover="mouseOver(item.id)"
      >
        <el-icon class="img-del" v-if="curMouseOnId == item.id" @click="delImage(item.id)">
          <CircleClose />
        </el-icon>
        <el-image
          style="width: 100%"
          :src="item.url"
          fit="scale-down"
          lazy
          @click="selectFile(item)"
        />
      </div>
    </div>

    <template #footer>
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
      <el-upload
        drag
        :action="`${serverHost}/sys/fileUpload/upload?driver=${driver}`"
        :headers="{ Authorization: 'Bearer ' + userStore.token() }"
        :show-file-list="false"
        multiple
        :on-success="uploadSuccess"
      >
        <div
          class="el-upload__text"
          style="
            height: 40px;
            line-height: 40px;
            text-align: center;
            overflow: hidden;
            display: flex;
            flex-flow: row nowrap;
            justify-content: center;
          "
        >
          <el-icon
            class="el-icon--upload"
            style="font-size: 35px; margin: 0 var(--global-padding) 0 0"
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
import { defineModel, ref } from 'vue'
import { fileDelete, fileUploadCfg, fileUploadPage } from '@/api/fileUpload'
import { useUserStore } from '@/pinia/useUserStore'
import { ElMessage } from 'element-plus'

const userStore = useUserStore()

const model = defineModel()
const multiple = defineModel('multiple', { default: 1, type: Number })
const multipleSelectModel = ref([])

const uploadConfig = ref({})
const driver = ref('default')

const show = ref(false)

const adminx = window.adminX ? window.adminX : {}
const runAt = adminx.RunAt ? adminx.RunAt : ''
const serverHost = runAt + (adminx.BackendPrefix ? '/' + adminx.BackendPrefix : '')

const tableData = ref([])
const page = ref(1)
const pageSize = ref(25)
const total = ref(0)

const loadData = () => {
  fileUploadPage(page.value, pageSize.value).then((res) => {
    tableData.value = res.data.list
    page.value = res.data.page
    pageSize.value = res.data.page_size
    total.value = res.data.total
  })
}

const selectFile = (item) => {
  if (multiple.value == 1) {
    if (multipleSelectModel.value.length && multipleSelectModel.value[0].id == item.id) {
      multipleSelectModel.value = []
    } else {
      multipleSelectModel.value = [item]
    }
    model.value = multipleSelectModel.value
    return
  }

  for (let index = 0; index < multipleSelectModel.value.length; index++) {
    const element = multipleSelectModel.value[index]
    if (element.id == item.id) {
      multipleSelectModel.value.splice(index, 1)
      model.value = multipleSelectModel.value
      return
    }
  }

  if (multipleSelectModel.value.length >= multiple.value) {
    ElMessage({ type: 'warning', message: `最多选择 ${multiple.value} 个` })
    return
  }

  multipleSelectModel.value.push(item)
  model.value = multipleSelectModel.value
}

const checkSelect = (id) => {
  for (const key in multipleSelectModel.value) {
    if (Object.hasOwnProperty.call(multipleSelectModel.value, key)) {
      const element = multipleSelectModel.value[key]
      if (element.id == id) {
        return 'img-item-select'
      }
    }
  }
}

const uploadSuccess = (response, uploadFile, uploadFiles) => {
  // console.log(response, uploadFile, uploadFiles);
  loadData()
}

const curMouseOnId = ref(0)
const mouseLeave = (id) => {
  curMouseOnId.value = 0
}
const mouseOver = (id) => {
  curMouseOnId.value = id
}

const delImage = (id) => {
  fileDelete(id).then((res) => {
    for (let index = 0; index < multipleSelectModel.value.length; index++) {
      const element = multipleSelectModel.value[index]
      if (element.id == id) {
        multipleSelectModel.value.splice(index, 1)
        model.value = multipleSelectModel.value
        return
      }
    }
    console.log('删除', id)
    loadData()
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
