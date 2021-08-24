package repo

import (
	"ucenter/domain/entity"
	"ucenter/domain/ov"
)

type IRepo interface {
	Save(*interface{})
	FindByID(*ov.OID) *interface{}
	Fetch(*interface{}) []interface{}
}

type IUserRepo interface {
	Save(*entity.User)
	FindByID(*ov.OID) *entity.Person
	FindByUserName(*ov.UserName) *entity.User
	Fetch(*entity.Person) entity.Persons
}

type ITenantRepo interface {
	IRepo
}

type IApplicationRepo interface {
	IRepo
}

type IRoleRepo interface {
	IRepo
}

type ITenantUser interface {
	IRepo
}

type IGroup interface {
	IRepo
}
