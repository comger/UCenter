package service

import (
	"ucenter/domain/entity"
	"ucenter/domain/ov"
	"ucenter/domain/repo"
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
	CreateApplication(appCreate *entity.ApplicationCreate) *entity.Application
	ResetApplicationScrect(app *entity.Application) ov.ClientSecret
	UpdateApplication(app *entity.Application)
	GetApplication(aClientID ov.ClientID) *entity.Application
	GetApplications() []*entity.Application
	GetAppUsedModules(app *entity.Application) ov.Modules
	PubAppModule(app *entity.Application, key string, val string)
	RMAppModule(app *entity.Application, key string) error
	CreateTenant(app *entity.Application, aCustom *entity.Custom) *entity.Tenant
	DisableTenant(app *entity.Application, aCustom *entity.Custom)
}

type Application struct {
	arepo repo.IApplicationRepo
	trepo repo.ITenantRepo
}

var _ IApplication = &Application{}

func (a *Application) CreateApplication(appCreate *entity.ApplicationCreate) *entity.Application {
	app := &entity.Application{
		ApplicationCreate: *appCreate,
		ClientID:          ov.ClientID(ov.NewOID()),
	}
	app.ResetClientScrect()
	a.arepo.Create(app)
	return app
}

func (a *Application) ResetApplicationScrect(app *entity.Application) ov.ClientSecret {
	app.ResetClientScrect()
	a.arepo.Update(app)
	return app.GetClientScrect()
}

func (a *Application) UpdateApplication(app *entity.Application) {
	a.arepo.Update(app)
}

func (a *Application) GetApplication(aClientID ov.ClientID) *entity.Application {
	return a.arepo.FindByID(aClientID)
}

func (a *Application) GetApplications() []*entity.Application {
	return a.arepo.Fetch()
}

func (a *Application) GetAppUsedModules(app *entity.Application) ov.Modules {
	panic("not implemented") // TODO: Implement
}

func (a *Application) PubAppModule(app *entity.Application, key string, val string) {
	app.Modules[key] = val
	a.arepo.Update(app)
}

func (a *Application) RMAppModule(app *entity.Application, key string) error {
	panic("not implemented") // TODO: Implement
}

func (a *Application) CreateTenant(app *entity.Application, aCustom *entity.Custom) *entity.Tenant {
	tenant := &entity.Tenant{
		Custom:   *aCustom,
		TenantID: ov.TenantID(ov.NewOID()),
		ClientID: app.ClientID,
	}

	return tenant
}

func (a *Application) DisableTenant(app *entity.Application, aCustom *entity.Custom) {
	tenant := a.trepo.FindByAppCustom(app.ClientID, aCustom.CustomID)
	tenant.Deactive()
	a.trepo.Update(tenant)
}
