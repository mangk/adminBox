<template>
    <div class="upload-file-slot" @click="shoFileManagerBox">
        <slot>
            <el-button type="primary">
                <el-icon>
                    <UploadFilled />
                </el-icon>
                文件上传
            </el-button>
        </slot>
    </div>

    <el-dialog v-model="show" width="70vw" :show-close="false" :before-close="handleClose"
        style="height: 70vh;background-color: #f1f2f5;">
        <BaseManager style="width: calc(100% + 10px);height: calc(70vh - 50px);" />
        <!-- <template #footer>
            <div class="dialog-footer">
                <el-button @click="show = false">Cancel</el-button>
                <el-button type="primary" @click="show = false">
                    Confirm
                </el-button>
            </div>
        </template> -->
    </el-dialog>

</template>

<script setup>
import { ref } from "vue"
import BaseManager from './baseManager.vue'

const show = ref(true)


const shoFileManagerBox = () => {
    // fileUploadCfg().then((res) => {
    //     uploadConfig.value = res.data
    // })
    // loadData()
    show.value = true
}

const handleClose = () => {
    show.value = false
}
</script>
<style lang="css">
.el-dialog {
    padding: 10px;
}

.el-dialog__header {
    padding: 0;
}
</style>