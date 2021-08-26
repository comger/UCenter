package service

import (
	"ucenter/domain/entity"
	"ucenter/domain/ov"
)

type ICustomApp interface {
	CreateCustom(*entity.Custom) *entity.Custom
	SetCustom(*entity.Custom)
	GetCustoms() []*entity.Custom
	CreateGroup(*entity.Custom, *entity.Group)
	AppendGroupMember(*entity.Group, *entity.Group)
	GetGroupMembers(ov.GroupID) []*entity.Group
	RemoveGroup(*entity.Group)
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
	UpdateModules(*entity.Tenant, ov.OrderModuels)
	CreateRole(*entity.Tenant, string, ov.OrderModuels) *entity.Role
	UpdateRole(*entity.Role)
	AppendUser(*entity.Tenant, ov.UID) *entity.TenantUser
	SetTenantUser(*entity.TenantUser)
	RemoveTenantUser(*entity.TenantUser)
}
