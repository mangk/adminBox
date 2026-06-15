// Package casbin provides a singleton Casbin enforcer for RBAC authorization.
//
// It integrates with GORM for policy storage and supports route-level
// authorization checks via sub/obj/act (subject, object, action) matching.
package casbin

import (
	"sync"

	casbinLib "github.com/casbin/casbin/v2"
	casbinModel "github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/mangk/adminBox/pkg/config"
	"github.com/mangk/adminBox/pkg/db"
	"github.com/mangk/adminBox/pkg/log"
)

var enforce *casbinLib.SyncedEnforcer
var once sync.Once

// Enforce returns the global Casbin enforcer, initializing it once.
func Enforce() *casbinLib.SyncedEnforcer {
	once.Do(func() {
		a, _ := gormadapter.NewAdapterByDBUseTableName(db.DB(), config.DBCfg()["default"].Prefix+"sys", "casbin_rule")
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
		m, err := casbinModel.NewModelFromString(text)
		if err != nil {
			log.Error("casbin model load failed!" + err.Error())
			return
		}
		enforce, _ = casbinLib.NewSyncedEnforcer(m, a)
	})
	return enforce
}
