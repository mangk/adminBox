<template>
  <div>
    <!--    搜索-->
    <div class="gva-search-box">
      <el-form ref="searchForm" :inline="true" :model="searchInfo">
        {[{ range $f := .field }]}
        {[{ formatElement $f }]}
        {[{ end }]}
        <el-form-item>
          <el-button type="primary" icon="search" @click="__list(true)">查询</el-button>
          <el-button icon="refresh" @click="__list(false)">重置</el-button>
        </el-form-item>
      </el-form>
    </div>

    <!--    表格-->
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button type="primary" icon="plus" @click="_create('add')" v-show="!{[{.createBtnHide}]}">新增</el-button>
        <el-popover v-model="deleteVisible" placement="top" width="160">
          <p>确定要删除吗？</p>
          <div style="text-align: right; margin-top: 8px;">
            <el-button type="primary" link @click="deleteVisible = false">取消</el-button>
            <el-button type="primary" @click="_delete">确定</el-button>
          </div>
          <template #reference>
            <el-button icon="delete" :disabled="false" style="margin-left: 10px;" @click="deleteVisible = true"
                       v-show="!{[{.deleteBtnHide}]}">
              删除
            </el-button>
          </template>
        </el-popover>
      </div>
      <el-table :data="tableData.list" @sort-change="_sortChange" @selection-change="_selectionChange">
        <el-table-column type="selection" width="55"/>
        {[{ range .field }]}
        <el-table-column align="left" label="{[{.Name}]}" min-width="160" prop="{[{.Column}]}"
                         :sortable="{[{ if .SortAble }]}'custom'{[{ else }]}false{[{end}]}"/>
        {[{ end }]}
        <el-table-column align="left" fixed="right" label="操作" width="120">
          <template #default="scope">
            <el-tooltip class="box-item" effect="dark" content="查看详情" placement="top">
              <el-button icon="view" type="primary" link @click="_detail(scope.row)"></el-button>
            </el-tooltip>
            <el-tooltip class="box-item" effect="dark" content="编辑" placement="top">
              <el-button icon="edit" type="primary" link @click="_edit(scope.row)"
                         v-show="!{[{.editBtnHide}]}"></el-button>
            </el-tooltip>
            <el-tooltip class="box-item" effect="dark" content="删除" placement="top">
              <el-button icon="delete" type="primary" link @click="_delete(scope.row,false)"
                         v-show="!{[{.deleteBtnHide}]}"></el-button>
            </el-tooltip>
          </template>
        </el-table-column>
      </el-table>
      <div class="gva-pagination">
        <el-pagination
            :current-page="tableData.page"
            :page-size="tableData.pageSize"
            :page-sizes="[20, 50, 100]"
            :total="tableData.count"
            layout="total, sizes, prev, pager, next, jumper"
            @current-change="_changeCurrentPage"
            @size-change="_changePageSize"
        />
      </div>

    </div>

    <!--    弹窗：新建、编辑、详情-->
    <el-dialog v-model="dialog.formVisible" :before-close="closeDialog" :title="dialog.title">
      <el-form ref="apiForm" :model="dialog.form" :rules="rules" label-width="80px">
        {[{ range .field }]}
        <el-form-item label="{[{.Name}]}" prop="{[{.Column}]}">
          <el-input v-model="dialog.form.{[{.Column}]}" :disabled="dialog.disabled || !{[{.EditAble}]}"
                    autocomplete="off"/>
        </el-form-item>
        {[{ end }]}
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="_dialogClose">取 消</el-button>
          <el-button type="primary" @click="_dialogEnter">确 定</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script>

export default {
  name: "a",
  data() {
    return {
      // 数据列表
      tableData: {
        list: [],
        page: 1,
        pageSize: 20,
        count: 0,
      },
      // 数据多选
      selected: [],
      // 数据筛选
      search: {},
      sort: "",
      dialog: {
        // 对话框标题
        title: "",
        // 对话框是否展示
        formVisible: false,
        // 表单数据
        form: {},
        disabled: false,
      },
      deleteVisible: false,
      // @formatter:off 这里的变量来自 golang 对模版的渲染
      _pk: {[{.pk}]},
      _ap: {[{.ap}]},
      // @formatter:on
    }
  },
  methods: {
    __list(search = false) {
      let data = {
        page: this.tableData.page,
        pageSize: this.tableData.pageSize,
        query: []
      }
      if (!search) {
        this.search = {}
        data.query = []
      }

      for (const formKey in this.search) {
        let d = this.search[formKey]
        if (d.length != 0) {
          if (Array.isArray(d)) {
            if (d[0] instanceof Date) {
              d[0] = this.$dayjs(d[0]).format("YYYY-MM-DD HH:mm:ss")
              d[1] = this.$dayjs(d[1]).format("YYYY-MM-DD HH:mm:ss")
              data.query.push({c: formKey, v: d.join(',')})
            } else {
              data.query.push({c: formKey, v: d.join(',')})
            }
          } else {
            data.query.push({c: formKey, v: d})
          }
        }
      }

      if (this.sort) {
        data.sort = this.sort
      }

      this.$http({
        method: 'post',
        url: this._ap + '/p',
        data
      }).then(res => {
        this.tableData = res.data
      });
    },
    __getDetail(id) {
      console.log(id)
      let data = {id: typeof (id) === 'string' ? id : id.toString()}
      this.$http({
        method: "post",
        url: this._ap + '/r',
        data,
      }).then(res => {
        this.dialog.form = res.data
      })
    },
    _create(key) {
      this.dialog.title = "创建"
      this.dialog.formVisible = true
      this.dialog.disabled = false
      this.dialog.form = {}
    },
    _edit(row) {
      this.dialog.title = "编辑"
      this.dialog.formVisible = true
      this.dialog.disabled = false
      this.__getDetail(row[this._pk].toString())
    },
    _detail(row) {
      this.dialog.title = "查看详情"
      this.dialog.formVisible = true
      this.dialog.disabled = true
      this.__getDetail(row[this._pk])
    },
    _delete(row, many = true) {
      let data = {
        ids: []
      }
      if (!many) {
        data.ids = [row[this._pk].toString()]
      } else {
        this.selected.forEach((value) => {
          data.ids.push(value[this._pk].toString())
        })
      }

      this.$messageBox.confirm('此操作将永久删除选中的数据, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(async () => {
        this.$http({
          method: 'delete',
          url: this._ap + '/d',
          data
        }).then(res => {
          if (res.code === 0) {
            this.$message({
              type: 'success',
              message: '删除成功!'
            })
            this.__list()
          }
        })
      })
    },
    _changePageSize(val) {
      this.tableData.page = 1
      this.tableData.pageSize = val
      this.__list(true)
    },
    _changeCurrentPage(val) {
      this.tableData.page = val
      this.__list(true)
    },
    _dialogClose() {
      this.dialog.formVisible = false
    },
    _dialogEnter() {
      let url = this._ap
      let method = "post"
      let data = {query: []}
      switch (this.dialog.title) {
        case "编辑":
          url += '/u'
          method = "patch"
          break
        case "创建":
          url += '/c'
          data.query.push({c: "order_id", v: ""})
          break
      }
      for (const formKey in this.dialog.form) {
        if (formKey !== this._pk) {
          data.query.push({c: formKey, v: this.dialog.form[formKey]})
        } else {
          data.ids = [this.dialog.form[formKey].toString()]
        }
      }
      this.$http({
        method: method,
        url: url,
        data
      }).then(res => {
        if (res.code === 0) {
          this.$message({
            type: 'success',
            message: res.msg,
            showClose: true
          })
          this.dialog.formVisible = false
          this.dialog.form = {}
          if (this.dialog.title === "创建") {
            this.tableData.page = 1
          }
          this.__list()
        }
      })
    },
    _selectionChange(val) {
      this.selected = val
    },
    _sortChange({prop, order}) {
      if (order === "ascending") {
        this.sort = prop + " asc"
      } else if (order === "descending") {
        this.sort = prop + " desc"
      } else {
        this.sort = ""
      }
      this.__list(true)
    }
  },
  mounted() {
    this.__list()
  }
}
</script>

<style scoped>

</style>