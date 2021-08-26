package entity

import (
	"ucenter/domain/ov"

	"golang.org/x/crypto/bcrypt"
)

type Person struct {
	Name   string // 姓名
	Email  ov.Email
	Mobile ov.Mobile
	UID    ov.UID
}

type Persons []Person

type User struct {
	Person
	UserName ov.UserName
	password ov.Password
	actived  bool
}

type TenantUser struct {
	UID            ov.UID
	TenantID       ov.TenantID
	RoleID         ov.RoleID
	IsTenantMaster bool
}

type TenantUserInfo struct {
	TenantUser
	Person Person
	Tenant Tenant
	Role   Role
}

type AdminUser struct {
	User
	IsSuperUser bool
}

func (u *User) Active() {
	u.actived = true
}

func (u *User) Deactive() {
	u.actived = false
}

func (u *User) IsActived() bool {
	return u.actived
}

func (u *User) SetPassword(aPassword string) {
	hash, err := bcrypt.GenerateFromPassword([]byte(aPassword), bcrypt.MinCost)
	if err != nil {
		u.password = ""
	} else {
		u.password = ov.Password(string(hash))
	}
}

func (u *User) VerifyPassword(aPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.password), []byte(aPassword))
	return err == nil
}
