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

    <el-dialog v-model="show" title="Tips" width="500" :before-close="handleClose">
        <span>This is a message</span>
        <template #footer>
            <div class="dialog-footer">
                <el-button @click="show = false">Cancel</el-button>
                <el-button type="primary" @click="show = false">
                    Confirm
                </el-button>
            </div>
        </template>
    </el-dialog>

</template>

<script setup>
import { ref } from "vue"

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
<style lang="css" module></style>