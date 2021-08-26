package service

import (
	"ucenter/domain/entity"
	"ucenter/domain/ov"
)

/**
### 应用程序
* 生成应用安全信息
* 重置应用安全信息
* 上线应用的可用功能模块
* 获取可用功能模块列表
* 下线没有被分配过的功能模块
* 修改可用模块的显示名称
* 创建应用程序的租户
* 取消租户授权
**/
type IApplication interface {
	CreateApplication(*entity.Application) *entity.Application
	ResetApplicationScrect(*entity.Application) ov.ClientSecret
	UpdateApplication(*entity.Application)
	GetApplication(ov.ClientID) *entity.Application
	GetApplications() []*entity.Application
	GetAppUsedModules(*entity.Application) ov.Modules
	PubAppModule(*entity.Application, string, string)
	RMAppModule(*entity.Application, string) error
	CreateTenant(*entity.Custom) *entity.Tenant
	DisableTenant(*entity.Custom)
}
