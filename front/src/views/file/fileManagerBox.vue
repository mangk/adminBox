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
        <BaseManager v-model="internalValue" :select-key="'url'" style="width: calc(100% + 10px);height: calc(70vh - 65px);" />
        <div style="padding: 10px 0px;display: flex;justify-content: space-between;">
            <span :style="{
                'font-size': '12px',
                'color': internalValue.length > multiple ? 'red' : ''
            }">{{ `已选 ` + internalValue.length + ` 条 / 可选 ` + multiple + ` 条` }}</span>

            <div>
                <el-button @click="internalValue = []">清空选择</el-button>
                <el-button type="primary" @click="selectedOK" :disabled="internalValue.length > multiple">确定</el-button>
                <el-button type="danger" @click="show = false">关闭</el-button>
            </div>
        </div>

    </el-dialog>

</template>

<script setup>
import { ref, watch } from "vue"
import BaseManager from './baseManager.vue'

const show = ref(false)
const props = defineProps({
    modelValue: {
        type: Array, // 修正为数组格式，支持多种类型
        default: () => []
    },
    multiple: {
        type: Number,
        default: 1
    },
    selected: {
        type: Function, // 修正为 Function 类型
        required: false,
        default: (value) => { } // 默认值为一个空函数
    },
})

// 创建内部状态，用来处理双向绑定的数据
const internalValue = ref(props.modelValue)
// 定义 emit 事件
const modelValueEmit = defineEmits(['update:modelValue'])
watch(internalValue, (newVal) => {
    modelValueEmit('update:modelValue', newVal)
})


const shoFileManagerBox = () => {
    show.value = true
}

const handleClose = () => {
    show.value = false
}

const selectedOK = () => {
    props.selected(internalValue.value)
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