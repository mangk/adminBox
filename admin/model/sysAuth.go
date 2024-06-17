package model

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mangk/adminX/db"
)

type SysAuth struct {
	Model
	TableId     int    `json:"table_id" gorm:"type:int(11);comment:表ID"`
	TableModule string `json:"table_module" gorm:"size:100;comment:表模块"`
	Type        string `json:"type" gorm:"size:100;comment:类型: menu,api,action"`
	Key         string `json:"key" gorm:"size:50;comment:数据标记: menu: id, api: id, action: mapKey"`
	SetValue    int    `json:"set_value" gorm:"type:int(11);comment:设置值"`
}

func (SysAuth) TableName() string {
	return "sys_auth"
}

func (SysAuth) SplitStr() string {
	return "|"
}

func (s SysAuth) LoadDetail(id int, module string) (curSet, resultSet AuthDetail, otherSetList []AuthDetail, err error) {

	list := []SysAuth{}
	if err = db.DB().Where("table_id = ? AND table_module = ?", id, module).Find(&list).Error; err != nil {
		return
	}

	curSet.Id = id
	curSet.Module = module
	curSet.RecordName = "正在编辑"
	curSet.List = make(map[string]int)
	if len(list) > 0 {
		for _, auth := range list {
			curSet.List[fmt.Sprintf("%s%s%s", auth.Type, s.SplitStr(), auth.Key)] = auth.SetValue
		}
	}

	if module == (SysUser{}).TableName() {
		var user SysUser
		user, err = SysUser{}.Detail(id)
		if err != nil {
			return
		}

		for _, record := range user.RoleList {
			otherSetList = append(otherSetList, AuthDetail{
				Id:         record.ID,
				Module:     record.TableName(),
				RecordName: "来自角色:" + record.Name,
			})
		}

		for _, record := range user.DepartmentList {
			otherSetList = append(otherSetList, AuthDetail{
				Id:         record.ID,
				Module:     record.TableName(),
				RecordName: "来自部门:" + record.Name,
			})
		}
	}

	for i, queryCondition := range otherSetList {
		list := []SysAuth{}
		if err = db.DB().Where("table_id = ? AND table_module = ?", queryCondition.Id, queryCondition.Module).Find(&list).Error; err != nil {
			return
		}
		if len(list) > 0 {
			otherSetList[i].List = make(map[string]int)
			for _, auth := range list {
				otherSetList[i].List[fmt.Sprintf("%s%s%s", auth.Type, s.SplitStr(), auth.Key)] = auth.SetValue
			}
		}
	}

	resultSet.List = s.result(append(otherSetList, curSet))

	return
}

func (s SysAuth) LoadMenuIdListByUserId(id int) (list []int, err error) {
	list = []int{}
	_, result, _, err := s.LoadDetail(id, SysUser{}.TableName())
	if err != nil {
		return
	}

	for key,value := range result.List {
		st := strings.Split(key, s.SplitStr())
		if len(st) >= 2 && st[0] == "menu" && value == 1 {
			menuId, err := strconv.Atoi(st[1])
			if err != nil {
				continue
			}

			list = append(list, menuId)
		}
	}
	return
}

func (s SysAuth) LoadApiDetailByUserId(id int) (list [][]string, err error) {
	list = [][]string{}
	_, result, _, err := s.LoadDetail(id, SysUser{}.TableName())
	if err != nil {
		return
	}

	allApi := SysApi{}.All(true)
	for key := range result.List {
		st := strings.Split(key, s.SplitStr())
		if len(st) >= 2 && st[0] == "api" {
			apiId, err := strconv.Atoi(st[1])
			if err != nil {
				continue
			}

			for _, v := range allApi {
				if apiId == v.ID {
					list = append(list, []string{strconv.Itoa(id), v.Path, v.Method})
				}
			}
		}
	}

	return
}

func (SysAuth) result(data []AuthDetail) (result map[string]int) {
	result = make(map[string]int)
	var (
		fromUser       = make(map[string][]int)
		fromRole       = make(map[string][]int)
		fromDepartment = make(map[string][]int)
	)
	if len(data) > 0 {
		for _, list := range data {
			switch list.Module {
			case (SysUser{}).TableName():
				for k, v := range list.List {
					if _, has := fromUser[k]; !has {
						fromUser[k] = make([]int, 0)
					}
					fromUser[k] = append(fromUser[k], v)
				}
			case (SysRole{}).TableName():
				for k, v := range list.List {
					if _, has := fromRole[k]; !has {
						fromRole[k] = make([]int, 0)
					}
					fromRole[k] = append(fromRole[k], v)
				}
			case (SysDepartment{}).TableName():
				for k, v := range list.List {
					if _, has := fromDepartment[k]; !has {
						fromDepartment[k] = make([]int, 0)
					}
					fromDepartment[k] = append(fromDepartment[k], v)
				}
			}
		}

		keyMap := []string{}
		for key, value := range fromUser {
			keyMap = append(keyMap, key)
			fromUser[key] = []int{checkSlice(value)}
		}
		for key, value := range fromRole {
			keyMap = append(keyMap, key)
			fromRole[key] = []int{checkSlice(value)}
		}
		for key, value := range fromDepartment {
			keyMap = append(keyMap, key)
			fromDepartment[key] = []int{checkSlice(value)}
		}

		for _, key := range keyMap {
			if v, has := fromDepartment[key]; has {
				result[key] = v[0]
			}
			if v, has := fromRole[key]; has {
				result[key] = v[0]
			}
			if v, has := fromUser[key]; has {
				result[key] = v[0]
			}
		}
	}

	return
}

func checkSlice(slice []int) int {
	hasOne := false

	for _, value := range slice {
		if value == -1 {
			return -1
		}
		if value == 1 {
			hasOne = true
		}
	}

	if hasOne {
		return 1
	}

	return 0
}

type AuthDetail struct {
	Id         int            `json:"id"`
	Module     string         `json:"module"`
	RecordName string         `json:"record_name"`
	List       map[string]int `json:"list"`
}
