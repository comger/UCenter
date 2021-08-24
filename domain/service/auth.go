package service

import (
	"ucenter/domain/entity"
	"ucenter/domain/ov"
	"ucenter/domain/repo"
)

type AuthService struct {
}

type AccessToken struct {
	UID      ov.UID
	TenantID ov.TenantID
}

func Auth(UserName string, Password string) *AccessToken {
	return nil
}

type IUserApp interface {
	Auth(ov.UserName, ov.Password) *AccessToken
	Me(ov.UID) *entity.TenantUser
	IsSuperUser(ov.UID) bool
	CreatePerson(aPerson *entity.Person)
	UpdatePerson(aPerson *entity.Person)
	CreateUser(aUser *entity.User)
	ResetPassword(ov.UID, ov.Password)
}

type UserApp struct {
	Urepo repo.IUserRepo
	Trepo repo.ITenantRepo
}

var _ IUserApp = &UserApp{}
