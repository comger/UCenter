package entity

import (
	"ucenter/domain/ov"
)

type Role struct {
	Name     string
	TenantID ov.TenantID
	RoleID   ov.RoleID
	Modules  ov.OrderModuels
}

func NewRole(aTenant Tenant, aName string, aModules ov.OrderModuels) *Role {
	if aTenant.OrderModuels.IsReady(aModules) {
		var aRole = &Role{
			Name:     aName,
			TenantID: aTenant.TenantID,
			RoleID:   ov.RoleID(ov.NewOID()),
			Modules:  aModules,
		}
		return aRole
	} else {
		return nil
	}
}

func (r *Role) SetModules(aTenant Tenant, aModules ov.OrderModuels) {
	if aTenant.OrderModuels.IsReady(aModules) {
		r.Modules = aModules
	}
}
