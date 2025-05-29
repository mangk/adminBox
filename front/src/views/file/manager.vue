<template>

    <div class="main-content">
        <el-row :gutter="10">
            <el-col :span="4">
                <el-tree class="file-tree-box" :data="pathTree" :props="{ label: 'name', children: 'children' }"
                    @node-click="handleNodeClick" default-expand-all :expand-on-click-node="false" draggable>
                    <template #default="{ node, data }">
                        <div :class="data.id == selectNodeId ? 'self-node node-active' : 'self-node'">
                            {{ node.label }}
                        </div>
                    </template>
                </el-tree>
            </el-col>
            <el-col :span="20">
                <div class="file-list-box">
                    <el-row style="display: flex;justify-content: space-between;">
                        <div>
                            <el-button type="primary" :icon="Top">
                                上传文件
                            </el-button>
                            <el-button :icon="RefreshRight">
                                刷新
                            </el-button>
                        </div>
                        <el-dropdown disabled>
                            <el-button disabled>
                                批量操作
                            </el-button>
                            <template #dropdown>
                                <el-dropdown-menu>
                                    <el-dropdown-item :icon="Plus">Action 1</el-dropdown-item>
                                    <el-dropdown-item :icon="CirclePlusFilled">
                                        Action 2
                                    </el-dropdown-item>
                                    <el-dropdown-item :icon="CirclePlus">Action 3</el-dropdown-item>
                                    <el-dropdown-item :icon="Check">Action 4</el-dropdown-item>
                                    <el-dropdown-item :icon="CircleCheck">Action 5</el-dropdown-item>
                                </el-dropdown-menu>
                            </template>
                        </el-dropdown>

                    </el-row>
                    <el-table :data="tableData" highlight-current-row show-overflow-tooltip stripe
                        height="var(--global-table)" style="width: 100%;margin-top: 10px;">

                        <el-table-column prop="date" label="Date" width="180" />
                        <el-table-column prop="name" label="Name" width="180" />
                        <el-table-column prop="address" label="Address" />
                    </el-table>
                    <el-pagination v-model:current-page="page" v-model:page-size="pageSize"
                        :page-sizes="[20, 50, 100, 200]" size="small" layout="total, sizes, prev, pager, next, jumper"
                        :total="total" @size-change="handleSizeChange" @current-change="handleCurrentChange" />
                </div>
            </el-col>
        </el-row>

    </div>
</template>
<script setup>
import { ref, watch } from 'vue'
import { Top, RefreshRight, MoreFilled } from '@element-plus/icons-vue'
import { fileDelete, fileUploadCfg, fileUploadPage, fileGroupTree } from '@/api/fileUpload'
import { useUserStore } from '@/pinia/useUserStore'
import { ElMessage, ElMessageBox } from 'element-plus'
import { serverHost } from '@/utils/requester'

const uploadConfig = ref({})
fileUploadCfg().then((res) => {
    uploadConfig.value = res.data
})

const pathTree = ref([])
fileGroupTree().then((res) => {
    const treeData = Array.isArray(res.data) ? res.data : [];
    pathTree.value = [{ id: '', name: '全部', children: treeData }]
})

const tableData = ref([])
const page = ref(1)
const pageSize = ref(20)
const total = ref(0)

const loadData = (query = {}) => {
    fileUploadPage(page.value, pageSize.value, query).then((res) => {
        tableData.value = res.data.list
        page.value = res.data.page
        pageSize.value = res.data.page_size
        total.value = res.data.total
    })
}
loadData()

const handleSizeChange = (size) => {
    pageSize.value = size
    loadData()
}

const handleCurrentChange = (changePage) => {
    page.value = changePage
    loadData()
}

const selectNodeId = ref('')
const handleNodeClick = (data) => {
    let query = {}
    if (data.id) {
        selectNodeId.value = data.id
        query = {
            group_id: data.id
        }
    } else {
        selectNodeId.value = ''
    }
    loadData(query)
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

.main-content {
    background-color: unset;
    padding: 0;
}

.el-col {
    padding: 0;
}

.file-tree-box {
    box-sizing: border-box;
    padding: 10px;
    width: 100%;
    height: calc(100vh - 70px);
    overflow: scroll;
}

.file-list-box {
    background-color: #fff;
    padding: 10px;
    box-sizing: border-box;
    height: calc(100vh - 70px);
    overflow-y: scroll;
}

.self-node {
    font-size: 14px;
    padding: 5px;
}
.node-active {
    color: var(--el-color-primary);
}
</style>
