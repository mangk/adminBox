<template>

    <div class="main-content">
        <el-row :gutter="10">
            <el-col :span="4">
                <el-tree class="file-tree-box" :data="pathTree" :props="{ label: 'name', children: 'children' }"
                    @node-click="handleNodeClick" default-expand-all :expand-on-click-node="false" draggable>
                    <template #default="{ node, data }">
                        <div :class="data.id == searchGroupId ? 'self-node node-active' : 'self-node'">
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
                            <!-- <el-button :icon="RefreshRight" @click="loadData(true)">
                                刷新
                            </el-button> -->

                            <el-dropdown :disabled="!tableDataSelectIds.length" style="margin-left: 10px;">
                                <el-button :disabled="!tableDataSelectIds.length">
                                    批量操作
                                </el-button>
                                <template #dropdown>
                                    <el-dropdown-menu>
                                        <el-dropdown-item :icon="Plus">下载</el-dropdown-item>
                                        <el-dropdown-item :icon="Plus">移动到</el-dropdown-item>
                                        <el-dropdown-item :icon="CirclePlusFilled" type="primary">删除</el-dropdown-item>
                                    </el-dropdown-menu>
                                </template>
                            </el-dropdown>
                        </div>

                        <el-input v-model="searchName" style="max-width: 450px" placeholder="搜索文件名" clearable
                            class="input-with-select" @clear="loadData(true)">
                            <template #prepend>
                                <el-select v-model="searchTag" placeholder="筛选类型" clearable style="width: 115px"
                                    @clear="loadData(true)">
                                    <el-option label="Restaurant" value="1" />
                                    <el-option label="Order No." value="2" />
                                    <el-option label="Tel" value="3" />
                                </el-select>
                            </template>
                            <template #append>
                                <el-button :icon="Search" @click="loadData(true)" />
                            </template>
                        </el-input>


                    </el-row>
                    <el-table :data="tableData" row-key="id" highlight-current-row show-overflow-tooltip stripe
                        style="width: 100%;margin-top: 9px;border-top: 1px solid var(--el-table-border-color);height: var(--global-table);"
                        @selection-change="handleSelectionChange">
                        <el-table-column type="selection" width="55" fixed />
                        <el-table-column prop="name" label="文件名" width="400" />
                        <el-table-column prop="tag" label="类型" width="80" />
                        <el-table-column prop="group_info.name" label="分组" />
                        <el-table-column prop="ut" label="更新时间" width="160" />
                        <el-table-column prop="ct" label="创建时间" width="160" />
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
import { ref, reactive, watch } from 'vue'
import { Top, RefreshRight, Search, MoreFilled } from '@element-plus/icons-vue'
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

const page = ref(1)
const pageSize = ref(20)
const total = ref(0)
// 一些查询条件
const searchGroupId = ref('')
const searchName = ref('')
const searchTag = ref('')
// 数据列表
const tableData = ref([])
// 数据选中
const tableDataSelectIds = ref([])

const loadData = (resetPage = false) => {
    if (resetPage) {
        page.value = 1
        pageSize.value = 20
    }

    let query = {}
    if (searchGroupId.value) {
        query.group_id = searchGroupId.value
    }
    if (searchName.value) {
        query.name = searchName.value
    }
    if (searchTag.value) {
        query.tag = searchTag.value
    }


    fileUploadPage(page.value, pageSize.value, query).then((res) => {
        tableData.value = res.data.list
        page.value = res.data.page
        pageSize.value = res.data.page_size
        total.value = res.data.total
    })
    tableDataSelectIds.value = []
}

loadData()

const handleSelectionChange = (val) => {
    tableDataSelectIds.value = val.map(item => item.id)
}

const handleSizeChange = (size) => {
    pageSize.value = size
    loadData()
}

const handleCurrentChange = (changePage) => {
    page.value = changePage
    loadData()
}

const handleNodeClick = (data) => {
    if (data.id) {
        searchGroupId.value = data.id
    } else {
        searchGroupId.value = ''
    }
    loadData(true)
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
