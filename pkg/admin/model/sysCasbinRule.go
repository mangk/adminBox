package model

import (
	"errors"
	"strconv"

	"github.com/mangk/adminBox/pkg/casbin"
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
		e := casbin.Enforce()
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
	e := casbin.Enforce()
	success, _ := e.RemoveFilteredPolicy(v, p...)
	return success
}
