package crud

import (
	"bytes"
	"fmt"
	"html/template"
	"math"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mangk/adminBox/admin/model"
	"github.com/mangk/adminBox/db"
	"github.com/mangk/adminBox/front"
	"github.com/mangk/adminBox/http"
	"github.com/mangk/adminBox/http/middleware"
	"github.com/mangk/adminBox/http/request"
	"github.com/mangk/adminBox/http/response"
	"github.com/mangk/adminBox/log"
)

// 数据筛选自动sql条件
const (
	WhereEQ   uint8 = iota + 1 // =
	WhereLike                  // LIKE
	WhereIn                    // IN
	WhereGt                    // >
	WhereLt                    // <
	WhereGe                    // >=
	WhereLe                    // <=
	WhereTst                   // > AND <
	WhereEst                   // >= AND <
	WhereTse                   // > AND <=
	WhereEse                   // >= AND <=
	WhereFs                    // FIND_IN_SET
)

// Element 组件
const (
	ElementComponentInput uint8 = iota + 1 // input
	ElementComponentDatePicker
	ElementComponentDateRangePicker
	ElementComponentDateTimePicker
	ElementComponentDateTimeRangePicker
	ElementComponentSelect
	ElementComponentCascader
)

type ElOption struct {
	Type       uint8         // 对应 Element 组件类型
	SearchType uint8         // 后台搜索的条件类型
	Options    func() string // 筛选框 item
	Props      string        // 筛选框属性设置
}

// 字段
type Field struct {
	Name             string                                           // 展示名
	Column           string                                           // 字段名
	Type             string                                           // 字段类型
	ElOption         ElOption                                         // 对应前端 Element 类型 TODO 实现这里的功能，通过模版渲染switch来完成 表单，列表内容的输出
	TranslateFunc    func() map[string]string                         // 字段翻译函数：返回用来翻译字段的map
	FormatFunc       func(value, originValue interface{}) interface{} // 字段格式化函数
	SortAble         bool                                             // 是否排序字段
	Hide             bool                                             // 是否对外展示
	EditAble         bool                                             // 是否可更新
	DefaultValueFunc func() string                                    // 数据创建时的默认值
}

// 实体
type Engine struct {
	mutex      sync.Mutex
	ap         string                     // absolute path
	opt        Options                    // 配置选项
	field      []Field                    // 字段
	middleware []gin.HandlerFunc          // 应用的中间件
	userFunc   map[string]gin.HandlerFunc // 用户自己注册的方法
}

// 选项参数
type Options struct {
	DbName           string // 指定数据库连接
	TableName        string // 指定表名
	PK               string // 指定主键字段
	SoftDelete       string // 指定软删除字段
	OrderBy          string // 指定默认排序
	HideCreateBtn    bool   // 隐藏创建按钮
	HideDeleteBtn    bool   // 隐藏删除按钮
	HideEditBtn      bool   // 隐藏编辑按钮
	CustomDataOrigin DataOrigin
}

type DataOrigin interface {
	Page(req request.CRUDRequest) (data []map[string]interface{}, count int64, pageSize int64, err error)
	GetById(req request.CRUDRequest) (data map[string]interface{}, err error)
	UpdateById(req request.CRUDRequest) (id interface{}, err error)
	Create(req request.CRUDRequest) (id interface{}, err error)
	Delete(req request.CRUDRequest) (id interface{}, err error)
}

func NewEngine(absolutePath string, opt Options) *Engine {
	if opt.DbName == "" {
		opt.DbName = "default"
	}
	if opt.PK == "" {
		opt.PK = "id"
	}

	e := &Engine{
		ap:         strings.Trim(absolutePath, "/"),
		mutex:      sync.Mutex{},
		field:      nil,
		middleware: []gin.HandlerFunc{},
		userFunc:   make(map[string]gin.HandlerFunc),
		opt:        opt,
	}
	return e
}

func (e *Engine) AddMiddleware(m ...gin.HandlerFunc) *Engine {
	e.middleware = append(e.middleware, m...)
	return e
}

func (e *Engine) AddField(filed ...Field) *Engine {
	e.field = append(e.field, filed...)
	return e
}

func (e *Engine) Register() {
	g := http.HttpEngine()
	for _, handlerFunc := range e.middleware {
		g.Use(handlerFunc)
	}

	// 注册模版
	e.register("tmp", "GET", e.tmp)

	// 注册系统方法 TODO 系统默认鉴权 middleware 的加入
	for _, key := range []string{"p|POST", "r|POST", "c|POST", "u|PATCH", "d|DELETE"} {
		st := strings.Split(key, "|")
		var f []gin.HandlerFunc
		if _, ok := e.userFunc[key]; ok {
			f = []gin.HandlerFunc{e.userFunc[key]}
			delete(e.userFunc, key)
		} else {
			switch st[0] {
			case "p":
				f = []gin.HandlerFunc{middleware.PublicRequestCrud(), e.page}
			case "r":
				f = []gin.HandlerFunc{middleware.PublicRequestCrud(), e.getById}
			case "c":
				f = []gin.HandlerFunc{middleware.PublicRequestCrud(), e.create}
			case "u":
				f = []gin.HandlerFunc{middleware.PublicRequestCrud(), e.updateById}
			case "d":
				f = []gin.HandlerFunc{middleware.PublicRequestCrud(), e.delete}
			}
		}
		e.register(st[0], st[1], f...)
	}

	// 注册用户自定义方法
	for key, f := range e.userFunc {
		st := strings.Split(key, "|")
		e.register(st[0], st[1], f)
	}
}

func (e *Engine) RegisterGetHandler(relativePath string, handlerFunc gin.HandlerFunc) *Engine {
	relativePath = strings.Trim(relativePath, "/")
	key := fmt.Sprintf("%s|%s", relativePath, "GET")
	if _, ok := e.userFunc[key]; ok {
		log.Error("Unsupported methods")
		panic("path as registered")
	}
	e.userFunc[key] = handlerFunc
	return e
}

func (e *Engine) RegisterPostHandler(relativePath string, handlerFunc gin.HandlerFunc) *Engine {
	relativePath = strings.Trim(relativePath, "/")
	key := fmt.Sprintf("%s|%s", relativePath, "POST")
	if _, ok := e.userFunc[key]; ok {
		log.Error("Unsupported methods")
		panic("path as registered")
	}
	e.userFunc[key] = handlerFunc
	return e
}

func (e *Engine) RegisterPatchHandler(relativePath string, handlerFunc gin.HandlerFunc) *Engine {
	relativePath = strings.Trim(relativePath, "/")
	key := fmt.Sprintf("%s|%s", relativePath, "PATCH")
	if _, ok := e.userFunc[key]; ok {
		log.Error("Unsupported methods")
		panic("path as registered")
	}
	e.userFunc[key] = handlerFunc
	return e
}

func (e *Engine) RegisterDeleteHandler(relativePath string, handlerFunc gin.HandlerFunc) *Engine {
	relativePath = strings.Trim(relativePath, "/")
	key := fmt.Sprintf("%s|%s", relativePath, "DELETE")
	if _, ok := e.userFunc[key]; ok {
		log.Error("Unsupported methods")
		panic("path as registered")
	}
	e.userFunc[key] = handlerFunc
	return e
}

func (e *Engine) register(relativePath, method string, handlerFunc ...gin.HandlerFunc) {
	g := http.HttpEngine()
	switch method {
	case "POST":
		g.POST(e.ap+"/"+relativePath, handlerFunc...)
	case "GET":
		g.GET(e.ap+"/"+relativePath, handlerFunc...)
	case "PATCH":
		g.PATCH(e.ap+"/"+relativePath, handlerFunc...)
	case "DELETE":
		g.DELETE(e.ap+"/"+relativePath, handlerFunc...)
	default:
		log.Error("Unsupported methods")
		panic("Unsupported methods")
	}
}

func (e *Engine) tmp(ctx *gin.Context) {
	data := map[string]interface{}{
		"pk":            e.opt.PK,
		"field":         e.field,
		"ap":            "/" + e.ap,
		"createBtnHide": e.opt.HideCreateBtn,
		"deleteBtnHide": e.opt.HideDeleteBtn,
		"editBtnHide":   e.opt.HideEditBtn,
	}

	var buf bytes.Buffer
	t, _ := template.New("convert").Delims("{[{", "}]}").Funcs(template.FuncMap{
		"formatElement": formatElement,
	}).Parse(front.CrudTemplate)
	_ = t.Execute(&buf, data)
	ctx.String(200, "%s", buf.String())
}

func (e *Engine) page(ctx *gin.Context) {
	req := request.PublicRequestCrud(ctx)

	var count int64
	var data []map[string]interface{}
	var err error

	if e.opt.CustomDataOrigin != nil {
		data, count, req.PageSize, err = e.opt.CustomDataOrigin.Page(req)
	} else {
		query := db.DB(e.opt.DbName).Table(e.opt.TableName)
		for _, condition := range req.Query {
			for _, field := range e.field {
				if condition.Column == field.Column {
					switch field.ElOption.SearchType {
					case WhereEQ:
						query = query.Where(fmt.Sprintf("%s = ?", condition.Column), condition.Value)
					case WhereLike:
						query = query.Where(fmt.Sprintf("%s like ?", condition.Column), "%"+condition.Value+"%")
					case WhereIn:
						query = query.Where(fmt.Sprintf("%s in ?", condition.Column), strings.Split(condition.Value, ","))
					case WhereGt:
						query = query.Where(fmt.Sprintf("%s > ?", condition.Column), condition.Value)
					case WhereLt:
						query = query.Where(fmt.Sprintf("%s < ?", condition.Column), condition.Value)
					case WhereGe:
						query = query.Where(fmt.Sprintf("%s >= ?", condition.Column), condition.Value)
					case WhereLe:
						query = query.Where(fmt.Sprintf("%s <= ?", condition.Column), condition.Value)
					case WhereTst:
						val := strings.Split(condition.Value, ",")
						if len(val) == 2 {
							query = query.Where(fmt.Sprintf("%s > ?", condition.Column), val[0]).Where(fmt.Sprintf("%s < ?", condition.Column), val[1])
						} else {
							response.FailWithDetail(ctx, "范围查询参数提交错误", err)
							return
						}
					case WhereEst:
						val := strings.Split(condition.Value, ",")
						if len(val) == 2 {
							query = query.Where(fmt.Sprintf("%s >= ?", condition.Column), val[0]).Where(fmt.Sprintf("%s < ?", condition.Column), val[1])
						} else {
							response.FailWithDetail(ctx, "范围查询参数提交错误", err)
							return
						}
					case WhereTse:
						val := strings.Split(condition.Value, ",")
						if len(val) == 2 {
							query = query.Where(fmt.Sprintf("%s > ?", condition.Column), val[0]).Where(fmt.Sprintf("%s <= ?", condition.Column), val[1])
						} else {
							response.FailWithDetail(ctx, "范围查询参数提交错误", err)
							return
						}
					case WhereEse:
						val := strings.Split(condition.Value, ",")
						if len(val) == 2 {
							query = query.Where(fmt.Sprintf("%s >= ?", condition.Column), val[0]).Where(fmt.Sprintf("%s <= ?", condition.Column), val[1])
						} else {
							response.FailWithDetail(ctx, "范围查询参数提交错误", err)
							return
						}
					case WhereFs:
					default:
					}
				}
			}
		}
		if e.opt.SoftDelete != "" {
			query = query.Where(fmt.Sprintf("%s IS NULL", e.opt.SoftDelete))
		}
		err = query.Count(&count).Error
		if err != nil {
			response.FailWithDetail(ctx, "数据库查询错误", err)
			return
		}
		if count > 0 {
			query = query.Select(e.selectColumns()).Limit(int(req.PageSize)).Offset(int((req.Page - 1) * req.PageSize))
			if req.Sort != "" {
				query = query.Order(req.Sort)
			} else {
				if e.opt.OrderBy != "" {
					query = query.Order(e.opt.OrderBy)
				} else {
					query = query.Order(fmt.Sprintf("%s desc", e.opt.PK))
				}
			}
			data, err = model.Find(query)
			// 对每个字段的数据进行翻译
			for _, field := range e.field {
				var tMap map[string]string
				if field.TranslateFunc != nil {
					tMap = field.TranslateFunc()
				}
				for i := range data {
					// 先翻译
					if field.TranslateFunc != nil {
						if val, ok := data[i][field.Column].(string); ok {
							if v, has := tMap[val]; has {
								data[i][field.Column+"_origin"] = data[i][field.Column]
								data[i][field.Column] = v
							}
						}
					}
					// 再格式化
					if field.FormatFunc != nil {
						o, ok := data[i][field.Column+"_origin"]
						if !ok {
							o = ""
						}
						data[i][field.Column] = field.FormatFunc(data[i][field.Column], o)
					}
				}
			}
			if err != nil {
				response.FailWithDetail(ctx, "数据库查询错误", err)
				return
			}
		}
	}

	if err != nil {
		response.FailWithDetail(ctx, "查询错误", err)
		return
	}

	response.OkWithDetail(ctx, "success", response.CRUDResponse{
		Count:     count,
		Page:      req.Page,
		PageSize:  req.PageSize,
		PageCount: int(math.Ceil(float64(count) / float64(req.PageSize))),
		List:      data,
	})
}

func (e *Engine) getById(ctx *gin.Context) {
	req := request.PublicRequestCrud(ctx)

	var data map[string]interface{}
	var err error

	if e.opt.CustomDataOrigin != nil {
		data, err = e.opt.CustomDataOrigin.GetById(req)
	} else {
		if len(req.Ids) == 0 {
			response.FailWithMsg(ctx, "未提交数据查询 ID")
			return
		}

		data, err = e.queryById(req.Ids[0])
		if err != nil {
			response.FailWithDetail(ctx, "数据查询错误", err.Error())
			return
		}
		if data == nil {
			response.FailWithMsg(ctx, "未找到数据")
			return
		}
	}
	if err != nil {
		response.FailWithDetail(ctx, "查询错误", err.Error())
		return
	}

	response.OkWithDetail(ctx, "ok", data)
}

func (e *Engine) updateById(ctx *gin.Context) {
	req := request.PublicRequestCrud(ctx)

	var id interface{}
	var err error

	if e.opt.CustomDataOrigin != nil {
		id, err = e.opt.CustomDataOrigin.UpdateById(req)
	} else {
		if len(req.Ids) == 0 {
			response.FailWithMsg(ctx, "未提交数据查询 ID")
			return
		}

		data, err := e.queryById(req.Ids[0])
		if err != nil {
			response.FailWithDetail(ctx, "数据查询错误", err.Error())
			return
		}
		if data == nil {
			response.FailWithMsg(ctx, "未找到数据")
			return
		}

		update := make(map[string]interface{})
		for _, condition := range req.Query {
			for _, field := range e.field {
				if condition.Column == field.Column && field.EditAble && condition.Value != "" {
					update[condition.Column] = condition.Value
				}
			}
		}

		err = db.DB(e.opt.DbName).Table(e.opt.TableName).Where(fmt.Sprintf("%s = ?", e.opt.PK), req.Ids[0]).Updates(update).Error
		if err != nil {
			response.FailWithDetail(ctx, "数据更新错误", err.Error())
			return
		}
		id = req.Ids[0]
	}
	if err != nil {
		response.FailWithDetail(ctx, "数据更新错误", err.Error())
		return
	}

	response.OkWithDetail(ctx, "ok", id)
}

func (e *Engine) create(ctx *gin.Context) {
	req := request.PublicRequestCrud(ctx)

	var id interface{}
	var err error
	if e.opt.CustomDataOrigin != nil {
		id, err = e.opt.CustomDataOrigin.Create(req)
	} else {
		create := make(map[string]interface{})
		for _, condition := range req.Query {
			for _, field := range e.field {
				if condition.Column == field.Column {
					var v interface{}
					if condition.Value == "" {
						if field.DefaultValueFunc == nil {
							goto WITHOUT_DEFAULT
						}
						v = field.DefaultValueFunc()
					} else {
						v = condition.Value
					}
					create[condition.Column] = v
				}
			WITHOUT_DEFAULT:
			}
		}
		err := db.DB(e.opt.DbName).Table(e.opt.TableName).Create(create).Error
		if err != nil {
			response.FailWithDetail(ctx, "数据创建错误", err.Error())
			return
		}
	}
	if err != nil {
		response.FailWithDetail(ctx, "数据创建错误", err.Error())
		return
	}

	response.OkWithDetail(ctx, "ok", id)
}

func (e *Engine) delete(ctx *gin.Context) {
	req := request.PublicRequestCrud(ctx)

	var id interface{}
	var err error
	if e.opt.CustomDataOrigin != nil {
		id, err = e.opt.CustomDataOrigin.Create(req)
	} else {
		if len(req.Ids) == 0 {
			response.FailWithMsg(ctx, "未提交数据查询 ID")
			return
		}

		query := db.DB(e.opt.DbName).Table(e.opt.TableName).Where(fmt.Sprintf("%s in ?", e.opt.PK), req.Ids)
		var err error
		if e.opt.SoftDelete != "" {
			err = query.Updates(map[string]interface{}{e.opt.SoftDelete: time.Now().Format("2006-01-02 15:05:05")}).Error
		} else {
			err = query.Delete(nil).Error
		}
		if err != nil {
			response.FailWithDetail(ctx, "数据删除错误", err.Error())
			return
		}
		id = req.Ids
	}
	if err != nil {
		response.FailWithDetail(ctx, "数据删除错误", err.Error())
		return
	}

	response.OkWithDetail(ctx, "ok", id)
}

func (e *Engine) selectColumns() string {
	var selectColumns []string
	for _, field := range e.field {
		if !field.Hide && field.Column != e.opt.SoftDelete {
			selectColumns = append(selectColumns, field.Column)
		}
	}
	return "`" + strings.Join(selectColumns, "`,`") + "`"
}

func (e *Engine) queryById(id string) (map[string]interface{}, error) {
	query := db.DB(e.opt.DbName).Table(e.opt.TableName).Select(e.selectColumns()).Where(fmt.Sprintf("%s = ?", e.opt.PK), id)
	if e.opt.SoftDelete != "" {
		query = query.Where(fmt.Sprintf("%s IS NULL", e.opt.SoftDelete))
	}

	data, err := model.Find(query.Limit(1))

	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return nil, nil
	}
	return data[0], nil
}

func formatElement(field Field) (html interface{}) {
	if field.ElOption.SearchType == 0 {
		return
	}
	switch field.ElOption.Type {
	case ElementComponentInput:
		return template.HTML(`
        <el-form-item label="` + field.Name + `">
          <el-input v-model="search.` + field.Column + `" placeholder="` + field.Name + `" clearable/>
        </el-form-item>
`)
	case ElementComponentDatePicker, ElementComponentDateRangePicker, ElementComponentDateTimePicker, ElementComponentDateTimeRangePicker:
		t := ""
		switch field.ElOption.Type {
		case ElementComponentDatePicker:
			t = "date"
		case ElementComponentDateRangePicker:
			t = "daterange"
		case ElementComponentDateTimePicker:
			t = "datetime"
		case ElementComponentDateTimeRangePicker:
			t = "datetimerange"
		}
		return template.HTML(`
        <el-form-item label="` + field.Name + `">
			<el-date-picker v-model="search.` + field.Column + `" type="` + t + `" placeholder="` + field.Name + `" range-separator="～" start-placeholder="开始时间" end-placeholder="结束时间" clearable/>
        </el-form-item>`)
	case ElementComponentSelect:
		options := field.ElOption.Options
		if options != nil {
			return template.HTML(`
        <el-form-item label="` + field.Name + `">
			<el-select v-model="search.` + field.Column + `" placeholder="` + field.Name + `" filterable clearable>
				<el-option v-for='item in ` + options() + `' :key="item.k" :label="item.v" :value="item.k" />
			</el-select>
        </el-form-item>`)
		}
		return template.HTML(`
        <el-form-item label="` + field.Name + `">
			<el-select v-model="search.` + field.Column + `" placeholder="` + field.Name + `" filterable clearable>
				<el-option v-for="item in []" :key="item.v" :label="item.k" :value="item.v" />
			</el-select>
        </el-form-item>`)
	case ElementComponentCascader:
		options := field.ElOption.Options
		if options != nil {
			return template.HTML(`
        <el-form-item label="` + field.Name + `">
			<el-cascader v-model="search.` + field.Column + `" :options='` + options() + `' :props="` + field.ElOption.Props + `" clearable :show-all-levels="false" collapse-tags collapse-tags-tooltip filterable />
        </el-form-item>`)
		}
		return template.HTML(`
        <el-form-item label="` + field.Name + `">
			<el-cascader v-model="search.` + field.Column + `" :options="[]" :props="` + field.ElOption.Props + `" clearable :show-all-levels="false" filterable />
        </el-form-item>`)
	}
	return

}
