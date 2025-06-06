<template>

    <div class="main-content">
        <el-row :gutter="10">
            <el-col :span="4" style="position: relative;">
                <el-tree class="file-tree-box" :data="pathTree" :props="{ label: 'name', children: 'children' }"
                    @node-click="handleNodeClick" default-expand-all :expand-on-click-node="false" draggable>
                    <template #default="{ node, data }">
                        <div :class="data.id == searchGroupId ? 'self-node node-active' : 'self-node'">
                            {{ node.label }}
                        </div>
                    </template>
                </el-tree>
                <div style="position: absolute;bottom: 5px;left: 10px;">
                    <el-popover placement="top-start" :width="260" trigger="hover">
                        <template #reference>
                            <el-button :icon="Plus" size="small" style="border-right: 0px;" link />
                        </template>
                        <el-input v-model="fileGroupCreateInput" style="width: 200px" size="small"
                            placeholder="输入分组名称" />
                        <el-button :icon="Select" type="success" style="margin-left: 10px;" @click="handlePathCreate"
                            link />
                    </el-popover>
                    <el-button :icon="Minus" size="small" @click="handlePathDelete" link />
                </div>
            </el-col>
            <el-col :span="20">
                <div class="file-list-box">
                    <el-row style="display: flex;justify-content: space-between;">
                        <div>
                            <el-popover placement="bottom-start" :width="400" :visible="uploadVisible">
                                <template #reference>
                                    <el-button type="primary" :icon="Top" @click="uploadVisible = !uploadVisible">
                                        上传文件
                                    </el-button>
                                </template>


                                <el-upload class="upload-demo" drag :auto-upload="false" ref="fileUploadRef"
                                    :http-request="handleFileUpload" v-model:file-list="fileUploadList"
                                    :on-change="handleFileListChange" multiple>

                                    <template #trigger>
                                        <div class="el-upload__text"
                                            style="height: 30px;line-height: 30px;text-align: center;overflow: hidden;display: flex;flex-flow: row nowrap;justify-content: center;">
                                            <el-icon class="el-icon--upload"
                                                style="font-size: 25px; margin: 0px var(--global-padding) 0 0">
                                                <upload-filled />
                                            </el-icon>
                                            将文件拖到此处或<em>点击选择文件</em>
                                        </div>
                                    </template>

                                    <template #tip>
                                        <div class="el-upload__tip">
                                            <el-button v-if="fileUploadList.length" type="success" size="small"
                                                :icon="Check" @click="handleFileUploadStart">
                                                开始上传
                                            </el-button>
                                        </div>
                                    </template>
                                </el-upload>
                            </el-popover>

                            <!-- <el-button :icon="RefreshRight" @click="loadData(true)">
                                刷新
                            </el-button> -->

                            <el-dropdown :disabled="!tableDataSelectIds.length" style="margin-left: 10px;"
                                @command="handleBatchCommand">
                                <el-button :disabled="!tableDataSelectIds.length">
                                    批量操作
                                </el-button>
                                <template #dropdown>
                                    <el-dropdown-menu>
                                        <el-dropdown-item command="downlaod" :icon="Bottom">下载文件</el-dropdown-item>
                                        <el-dropdown-item command="copyUrl" :icon="CopyDocument">复制链接</el-dropdown-item>
                                        <el-dropdown-item command="move" :icon="Right">
                                            <template #default>
                                                <el-popover placement="right" :width="300" trigger="hover">
                                                    <template #reference>
                                                        移动到
                                                    </template>
                                                    <el-tree :data="pathTree"
                                                        :props="{ label: 'name', children: 'children' }"
                                                        @node-click="handleMoveNodeClick" default-expand-all
                                                        :expand-on-click-node="false" draggable>
                                                        <template #default="{ node, data }">
                                                            {{ node.label }}
                                                        </template>
                                                    </el-tree>
                                                </el-popover>
                                            </template>
                                        </el-dropdown-item>
                                        <el-dropdown-item command="delete" :icon="Close" divided>删除</el-dropdown-item>
                                    </el-dropdown-menu>
                                </template>
                            </el-dropdown>
                        </div>

                        <el-input v-model="searchName" style="max-width: 450px" placeholder="搜索文件名" clearable
                            class="input-with-select" @clear="loadData(true)">
                            <!-- <template #prepend>
                                <el-select v-model="searchTag" placeholder="筛选类型" clearable style="width: 115px"
                                    @clear="loadData(true)">
                                    <el-option label="Restaurant" value="1" />
                                    <el-option label="Order No." value="2" />
                                    <el-option label="Tel" value="3" />
                                </el-select>
                            </template> -->
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
                        <el-table-column prop="url" label="链接" width="80" />
                        <el-table-column prop="tag" label="类型" width="80" />
                        <el-table-column prop="group_info.name" label="分组" />
                        <!-- <el-table-column prop="ut" label="更新时间" width="160" /> -->
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
import { Top, Search, Bottom, Right, Close, CopyDocument, Check, Warning, Plus, Minus, Select } from '@element-plus/icons-vue'
import { fileDelete, fileMove, fileUploadCfg, fileUploadPage, fileGroupTree, fileGroupCreate, fileGroupDelete, fileUploadToken } from '@/api/fileUpload'
import { useUserStore } from '@/pinia/useUserStore'
import { ElMessage, ElMessageBox } from 'element-plus'
import { serverHost } from '@/utils/requester'
import * as qiniu from 'qiniu-js'

const uploadConfig = ref({})
fileUploadCfg().then((res) => {
    uploadConfig.value = res.data
})

const pathTree = ref([])
const loadFileGroupTree = () => {
    fileGroupTree().then((res) => {
        const treeData = Array.isArray(res.data) ? res.data : [];
        pathTree.value = treeData
    })
}
loadFileGroupTree()

const fileGroupCreateInput = ref("")

const uploadVisible = ref(false)
const fileUploadRef = ref()
const fileUploadStatus = ref(0) // 0:未上传 1:上传中 2:暂停上传
const fileUploadList = ref([])

const page = ref(1)
const pageSize = ref(20)
const total = ref(0)
// 一些查询条件
const searchGroupId = ref(0)
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

const handlePathCreate = () => {
    if (!fileGroupCreateInput.value) {
        ElMessage.warning('请输入分组名称')
        return
    }
    fileGroupCreate(fileGroupCreateInput.value, searchGroupId.value).then((res) => {
        loadFileGroupTree()
        fileGroupCreateInput.value = ''
    })
}

const handlePathDelete = () => {
    if (searchGroupId.value) {
        fileGroupDelete(searchGroupId.value).then((res) => {
            loadFileGroupTree()
        })
    } else {
        ElMessage.warning('请选择分组')
    }
}

const handleFileUpload = async (options) => {
    // TODO 实现上传进度
    // console.log("点击上传", options);
    let tokenRes = await fileUploadToken(options.file.name)
    // console.log("获取token", tokenRes);

    const observable = qiniu.upload(options.file, tokenRes.data.key, tokenRes.data.token)
    const subscription = observable.subscribe((next) => {
        options.onProgress(next.total)
    }, (error) => {
        options.onError(error)
        setTimeout(() => {
            searchGroupId.value = 0
            loadData(true)
        }, 1500);
    }, (complete) => {
        options.onSuccess(complete)
        setTimeout(() => {
            searchGroupId.value = 0
            loadData(true)
        }, 1500);
    })
}

const handleFileUploadStart = () => {
    fileUploadRef.value.submit()
}

const handleFileListChange = (f, fs) => {
    // console.log(1234,f, fs);
    if (fs.findIndex(fi => fi.name == f.name) != fs.findLastIndex(fi => fi.name == f.name)) {
        ElMessage.info(f.name + " 已存在于列表中")
        fs.pop()
    }
    fileUploadList.value = fs
}

const handleFileUploadListClean = () => {
    fileUploadList.value = []
}

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

const handleMoveNodeClick = (data) => {
    fileMove(tableDataSelectIds.value, data.id).then((res) => {
        if (res.code != 0) {
            ElMessage.error(res.msg)
        } else {
            ElMessage.success('移动成功')
        }
        loadData()
    })
}

const handleBatchCommand = async (command) => {
    switch (command) {
        case "downlaod":
            tableData.value.forEach(async item => {
                if (tableDataSelectIds.value.includes(item.id)) {
                    const iframe = document.createElement('iframe');
                    iframe.style.display = 'none';
                    iframe.src = item.url + "?attname="; // 如果是 Content-Disposition: attachment 的话就会自动下载并显示浏览器进度
                    document.body.appendChild(iframe);

                    setTimeout(() => {
                        document.body.removeChild(iframe);
                    }, 5000);
                }
            })
            break;
        case "copyUrl":
            try {
                let urls = ""
                tableData.value.forEach(item => {
                    if (tableDataSelectIds.value.includes(item.id)) {
                        urls += item.url + "\n"
                    }
                })

                // 使用 Clipboard API 复制文本到剪贴板
                await navigator.clipboard.writeText(urls);
                ElMessage({
                    message: '已复制',
                    type: 'success'
                })
            } catch (err) {
                ElMessage({
                    message: '未授权读取剪贴板',
                    type: 'error'
                })
            }
            break;
        case "delete":
            ElMessageBox.confirm('删除后文件无法找回！', '提示', {
                confirmButtonText: '确定',
                cancelButtonText: '取消',
                type: 'warning'
            }).then(() => {
                fileDelete(tableDataSelectIds.value).then((res) => {
                    if (res.code == 0) {
                        ElMessage({
                            message: '已删除',
                            type: 'success'
                        })
                    }
                    loadData()
                })
            })
            break;
        default:
            break;
    }
}

loadData()
// fileUploadToken().then((res) => {
//     console.log(333, res);
// })
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
    padding-bottom: 30px;
}

.file-list-box {
    background-color: #fff;
    padding: 10px;
    box-sizing: border-box;
    height: calc(100vh - 70px);
    overflow-y: scroll;
}

.el-tree-node__content>.el-tree-node__expand-icon {
    padding: 0px;
}

.self-node {
    font-size: 14px;
    padding: 5px;
    padding-left: 2px;
}

.node-active {
    color: var(--el-color-primary);
}
</style>
