package applications

import "ucenter/domain"

type AppAppInterface interface {
	domain.ApplicationRepoInterface
	GetAppTenants(string) []domain.AppTenantInfo
}

type AppApp struct {
	ar  domain.ApplicationRepoInterface
	atr domain.AppTenantRepoInterface
	tr  domain.TenantRepoInterface
}

var _ AppAppInterface = &AppApp{}

func (aa *AppApp) CreateApplication(app *domain.ApplicationCreate) (*domain.Application, error) {
	return aa.ar.CreateApplication(app)
}

func (aa *AppApp) UpdateApplication(_ *domain.Application) error {
	panic("not implemented") // TODO: Implement
}

func (aa *AppApp) GetApplications() []domain.Application {
	panic("not implemented") // TODO: Implement
}

func (aa *AppApp) GetApplication(_ string) *domain.Application {
	panic("not implemented") // TODO: Implement
}

func (aa *AppApp) GetAppTenants(ClientID string) []domain.AppTenantInfo {
	apps := aa.atr.GetAppTenantsByApp(ClientID)
	apptenantinfos := []domain.AppTenantInfo{}
	tmaps := aa.tr.GetTenantMapByIDs()
	amaps := aa.ar.GetAppMapByIDs()
	for _, app := range apps {
		tenantName := tmaps[app.TenantID]
		applicationName := amaps[app.ClientID]
		apptenantinfo := domain.AppTenantInfo{AppTenant: app,
			TenantName:      tenantName,
			ApplicationName: applicationName,
		}
		apptenantinfos = append(apptenantinfos, apptenantinfo)
	}
	return apptenantinfos
}

func (aa *AppApp) GetAppMapByIDs() map[string]string {
	return aa.ar.GetAppMapByIDs()
}
