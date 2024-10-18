package model

import (
	"encoding/json"
	"time"

	"github.com/mangk/adminBox/cache"
	"github.com/mangk/adminBox/db"
	"gorm.io/gorm"
)

type SysMenu struct {
	Model
	Pid       int       `json:"pid" gorm:"type:int(11);comment:父级Id"`
	Name      string    `json:"name" gorm:"unique;size:50;comment:菜单名称"`
	Path      string    `json:"path" gorm:"size:50;comment:路由地址"`
	Hidden    bool      `json:"hidden" gorm:"comment:是否隐藏"`
	Component string    `json:"component" gorm:"size:255;comment:模版地址"`
	Sort      int       `json:"sort" gorm:"type:int(11);comment:排序"`
	Children  []SysMenu `json:"children" gorm:"-"`
	Meta      `json:"meta" gorm:"embedded;comment:附加属性"`

	Premiss int `json:"premiss,omitempty" gorm:"-"` // 权限列表展示使用
}

type Meta struct {
	Title       string                 `json:"title" gorm:"size:50;comment:菜单名"`
	KeepAlive   bool                   `json:"keep_alive" gorm:"comment:是否缓存"`
	DefaultMenu bool                   `json:"default_menu" gorm:"comment:是否是基础路由"`
	Icon        string                 `json:"icon" gorm:"size:255;comment:菜单图标"`
	AutoClose   bool                   `json:"auto_close" gorm:"comment:自动关闭tab"`
	SCPath      string                 `json:"sc_path" gorm:"size:255;comment:服务端模版地址"`
	ActionList  map[string]interface{} `json:"action_list" gorm:"serializer:json;comment:动作列表"` // 列举了当前模块下所有按钮是否展示，对应前端应该从当前路由的 meta 数据中读取当前用户权限下的列表，并在页面的按钮上使用 v-if="$router.currenRouter.actionList['edit'] != false"
	ApiList     []SysApi               `json:"api_list" gorm:"-"`                               // Api 列表
}

func (s SysMenu) TableName() string {
	return "sys_menu"
}

func (s SysMenu) All(loadSystem bool) []SysMenu {
	list := []SysMenu{}
	db.DB().Find(&list)

	if loadSystem {
		list = append(list, s.SystemMenu()...)
	}
	return list
}

func (s SysMenu) Tree(loadSystem, withApi, buildTree, loadHidden bool, sort string, userid ...int) ([]SysMenu, error) {
	menus := []SysMenu{}
	q := db.DB().Order(sort)
	if !loadHidden {
		q = q.Where("hidden = false")
	}
	q.Find(&menus)

	if loadSystem {
		menus = append(menus, s.SystemMenu()...)
	}

	if len(userid) > 0 {
		tmpMeun := []SysMenu{}
		list, err := SysAuth{}.LoadMenuIdListByUserId(userid[0])
		if err != nil {
			return []SysMenu{}, err
		}

		for _, v := range menus {
			for _, menuId := range list {
				if menuId == v.ID {
					tmpMeun = append(tmpMeun, v)
				}
			}
		}
		menus = tmpMeun
	}

	root := SysMenu{
		Name: "根目录",
		Meta: Meta{Title: "根目录"},
	}

	if withApi {
		apis := SysApi{}.All(true)
		for _, api := range apis {
			var find bool
			for i, menu := range menus {
				if api.MenuId == menu.ID {
					find = true
					menus[i].Meta.ApiList = append(menus[i].Meta.ApiList, api)
				}
			}
			if !find {
				root.Meta.ApiList = append(root.Meta.ApiList, api)
			}
		}

	}

	if buildTree {
		data, _ := s.buildTree(menus)
		root.Children = data

		withRoot := []SysMenu{}
		withRoot = append(withRoot, root)
		return withRoot, nil
	} else {
		return append(menus, root), nil
	}
}

func (s SysMenu) buildTree(menuList []SysMenu) ([]SysMenu, error) {
	var err error
	treeMap := make(map[int][]SysMenu)
	for _, v := range menuList {
		treeMap[v.Pid] = append(treeMap[v.Pid], v)
	}
	resMenuList := treeMap[0]
	for i := 0; i < len(resMenuList); i++ {
		err = s.getBaseChildrenList(&resMenuList[i], treeMap)
	}
	return resMenuList, err
}

func (s SysMenu) getBaseChildrenList(menu *SysMenu, treeMap map[int][]SysMenu) (err error) {
	menu.Children = treeMap[menu.ID]
	for i := 0; i < len(menu.Children); i++ {
		err = s.getBaseChildrenList(&menu.Children[i], treeMap)
	}
	return err
}

const sysMenuTranMapKey = "model:sys_menu:TranMap"

func (s *SysMenu) BeforeSave(tx *gorm.DB) error {
	if s.Component == "" {
		s.Component = "views/util/serverComponent.vue"
	}

	cache.RedisDel(sysMenuTranMapKey)

	return nil
}

func (s SysMenu) TranMap() map[int]string {
	dat := cache.RedisHasOrQuery(sysMenuTranMapKey, func() (data string, exp time.Duration) {
		d := make(map[int]string)

		list := []SysMenu{}
		db.DB().Find(&list)
		list = append(list, s.SystemMenu()...)
		for _, menu := range list {
			d[menu.ID] = menu.Title
		}

		str, _ := json.Marshal(d)
		return string(str), 4 * time.Hour
	})

	data := make(map[int]string)
	json.Unmarshal([]byte(dat), &data)
	return data
}

func (s SysMenu) SystemMenu() []SysMenu {
	return []SysMenu{
		{Model: Model{ID: -100}, Pid: 0, Name: "welcome", Path: "welcome", Hidden: true, Component: "views/welcome.vue", Sort: 100, Meta: Meta{Title: "欢迎", KeepAlive: true, Icon: "sugar"}},
		{Model: Model{ID: -200}, Pid: 0, Name: "fileUpload", Path: "fileUpload", Component: "views/setting/index.vue", Sort: 100, Meta: Meta{Title: "文件管理", KeepAlive: true, Icon: "upload-filled"}},
		{Model: Model{ID: -300}, Pid: 0, Name: "setting", Path: "setting", Component: "views/setting/index.vue", Sort: 100, Meta: Meta{Title: "系统设置", KeepAlive: true, Icon: "setting"}},
		{Model: Model{ID: -301}, Pid: -300, Name: "auth", Path: "auth", Hidden: true, Component: "", Sort: 110, Meta: Meta{Title: "授权分组"}},
		{Model: Model{ID: -302}, Pid: -300, Name: "menu", Path: "menu", Component: "views/setting/menu.vue", Sort: 120, Meta: Meta{Title: "菜单管理", KeepAlive: true, Icon: "menu"}},
		{Model: Model{ID: -303}, Pid: -300, Name: "api", Path: "api", Component: "views/setting/api.vue", Sort: 130, Meta: Meta{Title: "API管理", KeepAlive: true, Icon: "link"}},
		{Model: Model{ID: -304}, Pid: -300, Name: "user", Path: "user", Component: "views/setting/user.vue", Sort: 140, Meta: Meta{Title: "用户管理", KeepAlive: true, Icon: "user"}},
		{Model: Model{ID: -305}, Pid: -300, Name: "role", Path: "role", Component: "views/setting/role.vue", Sort: 150, Meta: Meta{Title: "角色管理", KeepAlive: true, Icon: "filter"}},
		{Model: Model{ID: -306}, Pid: -300, Name: "department", Path: "department", Component: "views/setting/department.vue", Sort: 160, Meta: Meta{Title: "部门管理", KeepAlive: true, Icon: "office-building"}},
	}
}
