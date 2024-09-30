package model

import "github.com/mangk/adminBox/db"

type SysDepartment struct {
	Model
	Pid         int             `json:"pid" gorm:"type:int(11);comment:父级Id"`
	Name        string          `json:"name" gorm:"size:50;comment:角色名称"`
	Description string          `json:"description" gorm:"size:255;comment:角色简介"`
	Children    []SysDepartment `json:"children" gorm:"-"`
	UserList    []*SysUser      `json:"user_list,omitempty" gorm:"many2many:sys_user_department"`
}

func (SysDepartment) TableName() string {
	return "sys_department"
}

func (s SysDepartment) All() ([]SysDepartment, error) {
	list := []SysDepartment{}
	db.DB().Find(&list)
	data, _ := s.buildTree(list)

	return data, nil
}

func (s SysDepartment) buildTree(menuList []SysDepartment) ([]SysDepartment, error) {
	var err error
	treeMap := make(map[int][]SysDepartment)
	for _, v := range menuList {
		treeMap[v.Pid] = append(treeMap[v.Pid], v)
	}
	resMenuList := treeMap[0]
	for i := 0; i < len(resMenuList); i++ {
		err = s.getBaseChildrenList(&resMenuList[i], treeMap)
	}
	return resMenuList, err
}

func (s SysDepartment) getBaseChildrenList(menu *SysDepartment, treeMap map[int][]SysDepartment) (err error) {
	menu.Children = treeMap[menu.ID]
	for i := 0; i < len(menu.Children); i++ {
		err = s.getBaseChildrenList(&menu.Children[i], treeMap)
	}
	return err
}
