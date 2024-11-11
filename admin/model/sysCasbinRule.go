package model

import (
	"errors"
	"strconv"
	"sync"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/mangk/adminBox/config"
	"github.com/mangk/adminBox/db"
	"github.com/mangk/adminBox/log"
)

type SysCasbinRole struct{}

func (s SysCasbinRole) UpdateCasbin(authorityId int) error {
	id := strconv.Itoa(authorityId)
	s.ClearCasbin(0, id)

	rules, err := SysAuth{}.LoadApiDetailByUserId(authorityId)
	if err != nil {
		return err
	}

	if len(rules) > 0 {
		e := LoadEnforce()
		success, _ := e.AddPolicies(rules)
		if !success {
			return errors.New("存在相同api,添加失败,请联系管理员")
		}

		err = e.LoadPolicy()
		if err != nil {
			return err
		}
	}
	return nil
}

func (s SysCasbinRole) ClearCasbin(v int, p ...string) bool {
	e := LoadEnforce()
	success, _ := e.RemoveFilteredPolicy(v, p...)
	return success
}

var enforce *casbin.SyncedEnforcer
var once sync.Once

func LoadEnforce() *casbin.SyncedEnforcer {
	once.Do(func() {
		a, _ := gormadapter.NewAdapterByDBUseTableName(db.DB(), config.DBCfg()["default"].Prefix + "sys", "casbin_rule")
		text := `
		[request_definition]
		r = sub, obj, act
		
		[policy_definition]
		p = sub, obj, act
		
		[role_definition]
		g = _, _
		
		[policy_effect]
		e = some(where (p.eft == allow))
		
		[matchers]
		m = r.sub == p.sub && keyMatch2(r.obj,p.obj) && r.act == p.act
		`
		m, err := model.NewModelFromString(text)
		if err != nil {
			log.Error("字符串加载模型失败!" + err.Error())
			return
		}
		enforce, _ = casbin.NewSyncedEnforcer(m, a)
		// enforce.StartAutoLoadPolicy(time.Second * 10)
	})
	return enforce
}
