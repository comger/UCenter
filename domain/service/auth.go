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

type IUserApp interface {
	Auth(ov.TenantID, ov.UserName, ov.Password) *AccessToken
	Me(AccessToken) *entity.TenantUserInfo
	IsSuperUser(ov.UID) bool
	CreatePerson(aPerson *entity.Person)
	UpdatePerson(aPerson *entity.Person)
	CreateUser(aUser *entity.User, aPassword ov.Password)
	ResetPassword(aUID ov.UID, aPassword ov.Password)
}

type UserApp struct {
	urepo  repo.IUserRepo
	turepo repo.ITenantUserRepo
	trepo  repo.ITenantRepo
	rrepo  repo.IRoleRepo
}

var _ IUserApp = &UserApp{}

func (ua *UserApp) Auth(aTenantID ov.TenantID, aUserName ov.UserName, aPassword ov.Password) *AccessToken {
	user := ua.urepo.FindByUserName(aUserName)
	// 用户是否激活
	if user.IsActived() {
		// 用户名密码是否正确
		if user.VerifyPassword(string(aPassword)) {
			// 用户是否存在租户中
			if ov.IsNull(aTenantID) && ua.turepo.FindByID(aTenantID, user.UID) != nil {
				return &AccessToken{
					UID:      user.UID,
					TenantID: aTenantID,
				}
			}
		}
	}
	return nil
}

func (ua *UserApp) Me(aAccessToken AccessToken) *entity.TenantUserInfo {
	tUser := ua.turepo.FindByID(aAccessToken.TenantID, aAccessToken.UID)
	tUserInfo := &entity.TenantUserInfo{
		TenantUser: *tUser,
		Person:     *ua.urepo.FindByID(tUser.UID),
		Tenant:     *ua.trepo.FindByID(tUser.TenantID),
		Role:       *ua.rrepo.FindByID(tUser.RoleID),
	}

	return tUserInfo
}

func (ua *UserApp) IsSuperUser(aUID ov.UID) bool {
	return ua.urepo.IsSuperUser(aUID)
}

func (ua *UserApp) CreatePerson(aPerson *entity.Person) {
	aPerson.UID = ov.UID(ov.NewOID())
	user := &entity.User{
		Person: *aPerson,
	}
	ua.urepo.Create(user)
}

func (ua *UserApp) UpdatePerson(aPerson *entity.Person) {
	ua.urepo.Update(&entity.User{
		Person: *aPerson,
	})
}

func (ua *UserApp) CreateUser(aUser *entity.User, aPassword ov.Password) {
	aUser.SetPassword(string(aPassword))
	ua.urepo.Create(aUser)
}

func (ua *UserApp) ResetPassword(aUID ov.UID, aPassword ov.Password) {
	ua.urepo.SetPassword(aUID, aPassword)
}
