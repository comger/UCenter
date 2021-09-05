package service

import (
	"ucenter/domain/entity"
	"ucenter/domain/ov"
	"ucenter/domain/repo"
)

type ICustomApp interface {
	CreateCustom(aCustom *entity.Custom)
	UpdateCustom(aCustom *entity.Custom)
	GetCustoms(name string) []*entity.Custom
	CreateGroup(aCustom *entity.Custom, aGroupName string)
	AppendGroupMember(aGroup *entity.Group, bGroup *entity.Group)
	GetGroupMembers(ov.GroupID) []*entity.Group
	RemoveGroup(*entity.Group)
}

type CustomApp struct {
	crepo repo.ICustomRepo
	grepo repo.IGroupRepo
}

var _ ICustomApp = &CustomApp{}

func (c *CustomApp) CreateCustom(aCustom *entity.Custom) {
	aCustom.CustomID = ov.CustomID(ov.NewOID())
	c.crepo.Create(aCustom)
}

func (c *CustomApp) UpdateCustom(aCustom *entity.Custom) {
	c.crepo.Update(aCustom)
}

func (c *CustomApp) GetCustoms(name string) []*entity.Custom {
	return c.crepo.Fetch(name)
}

func (c *CustomApp) CreateGroup(aCustom *entity.Custom, aGroupName string) {
	aGroup := &entity.Group{
		CustomID: aCustom.CustomID,
		GroupID:  ov.GroupID(aCustom.CustomID),
		Name:     aGroupName,
		IsGroup:  true,
	}
	c.grepo.Create(aGroup)
}

func (c *CustomApp) AppendGroupMember(_ *entity.Group, _ *entity.Group) {
	panic("not implemented") // TODO: Implement
}

func (c *CustomApp) GetGroupMembers(_ ov.GroupID) []*entity.Group {
	panic("not implemented") // TODO: Implement
}

func (c *CustomApp) RemoveGroup(_ *entity.Group) {
	panic("not implemented") // TODO: Implement
}

/**
### 应用程序的租户
* 修改租户的可用模块
* 为租户创建角色实例
* 设置角色实例的可用的功能模块，功能模块只能来源于应用开通给本租户的
* 将已存在客户用户，导入到应用的租户中
* 将已存在客户的用户，批量导入到应用的租户中
* 为租户实创建新用户
* 设租户用户的角色，角色只能来自于相同租户
* 为租户设置超管用户
* 删除租户设置超管用户
* 将用户从租户中删除
**/
type ITenantApp interface {
	UpdateModules(aTenant *entity.Tenant, orderModules ov.OrderModuels)
	CreateRole(aTenant *entity.Tenant, name string, modules ov.OrderModuels)
	UpdateRole(aRoel *entity.Role)
	FetchRoles(aTenant *entity.Tenant) []*entity.Role
	AppendUser(aTenant *entity.Tenant, aUID ov.UID, roleID ov.RoleID)
	FetchUsers(aTenant *entity.Tenant) []*entity.TenantUserInfo
	UpdateTenantUser(aTenantUser *entity.TenantUser)
	RemoveTenantUser(aTenantUser *entity.TenantUser)
}

type TenantApp struct {
	trepo  repo.ITenantRepo
	rrepo  repo.IRoleRepo
	turepo repo.ITenantUserRepo
}

var _ ITenantApp = &TenantApp{}

func (t *TenantApp) UpdateModules(aTenant *entity.Tenant, orderModules ov.OrderModuels) {
	aTenant.OrderModuels = orderModules
	t.trepo.Update(aTenant)
}

func (t *TenantApp) CreateRole(aTenant *entity.Tenant, name string, modules ov.OrderModuels) {
	role := &entity.Role{
		RoleID:   ov.RoleID(ov.NewOID()),
		Name:     name,
		Modules:  modules,
		TenantID: aTenant.TenantID,
	}
	t.rrepo.Create(role)

}

func (t *TenantApp) UpdateRole(aRoel *entity.Role) {
	t.rrepo.Update(aRoel)
}

func (t *TenantApp) FetchRoles(aTenant *entity.Tenant) []*entity.Role {
	return t.rrepo.Fetch(string(aTenant.TenantID))
}

func (t *TenantApp) AppendUser(aTenant *entity.Tenant, aUID ov.UID, roldID ov.RoleID) {
	tenantUser := &entity.TenantUser{
		UID:      aUID,
		TenantID: aTenant.TenantID,
		RoleID:   roldID,
	}
	t.turepo.Create(tenantUser)
}

func (t *TenantApp) FetchUsers(aTenant *entity.Tenant) []*entity.TenantUserInfo {
	tusers := t.turepo.Fetch(aTenant.TenantID)
	tuis := make([]*entity.TenantUserInfo, 0)
	for _, tuser := range tusers {
		tui := &entity.TenantUserInfo{
			TenantUser: *tuser,
		}

		tuis = append(tuis, tui)
	}

	return tuis
}

func (t *TenantApp) UpdateTenantUser(aTenantUser *entity.TenantUser) {
	t.turepo.Update(aTenantUser)
}

func (t *TenantApp) RemoveTenantUser(aTenantUser *entity.TenantUser) {
	//TODO
}
