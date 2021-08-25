package repo

import (
	"ucenter/domain/entity"
	"ucenter/domain/ov"
)

type IUserRepo interface {
	Create(*entity.User)
	Update(*entity.User)
	FindByID(ov.UID) *entity.Person
	FindByUserName(ov.UserName) *entity.User
	Fetch(*entity.Person) entity.Persons
}

type ICustomRepo interface {
	Create(*entity.Custom)
	Update(*entity.Custom)
	FindByID(ov.CustomID) *entity.Custom
	Fetch(string) []*entity.Custom
}

type IGroup interface {
	Create(*entity.Group)
	Update(*entity.Group)
	FindByID(ov.GroupID) []*entity.Group
}

type ITenantRepo interface {
	Create(*entity.Tenant)
	Update(*entity.Tenant)
	FindByID(ov.TenantID) *entity.Tenant
	FindByClientID(ov.ClientID) []*entity.Tenant
	FindByCustomID(ov.CustomID) []*entity.Tenant
}

type IApplicationRepo interface {
	Create(*entity.Application)
	Update(*entity.Application)
	FindByID(ov.ClientID) *entity.Application
	Fetch(string) []*entity.Application
}

type IRoleRepo interface {
	Create(*entity.Role)
	Update(*entity.Role)
	FindByID(ov.RoleID) *entity.Role
	Fetch(string) []*entity.Role
}

type ITenantUser interface {
	Create(*entity.TenantUser)
	Update(*entity.TenantUser)
	FindByID(ov.TenantID, ov.UID) *entity.TenantUser
	Fetch() []*entity.TenantUser
}

type IEventLog interface {
	Create(*entity.EventLog)
	FindByID(ov.TenantID, ov.UID) []*entity.EventLog
	FindByEventID(ov.OID) []*entity.EventLog
}
