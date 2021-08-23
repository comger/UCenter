package applications

import "ucenter/domain"

type TenantAppInterface interface {
	domain.TenantRepoInterface
	GetAppTenants(string) []domain.AppTenantInfo
}

type TenantApp struct {
	tr  domain.TenantRepoInterface
	ar  domain.ApplicationRepoInterface
	atr domain.AppTenantRepoInterface
}

var _ TenantAppInterface = &TenantApp{}

func (ta *TenantApp) CreateTenant(tc *domain.TenantCreate) (*domain.Tenant, error) {
	return ta.tr.CreateTenant(tc)
}

func (ta *TenantApp) UpdateTenant(t *domain.Tenant) error {
	return ta.tr.UpdateTenant(t)
}

func (ta *TenantApp) GetTenants() []domain.Tenant {
	return ta.tr.GetTenants()
}

func (ta *TenantApp) GetTenant(ID string) (*domain.Tenant, error) {
	return ta.tr.GetTenant(ID)
}

func (ta *TenantApp) GetTenantMapByIDs() map[string]string {
	return ta.tr.GetTenantMapByIDs()
}

func (ta *TenantApp) GetAppTenants(tenantID string) []domain.AppTenantInfo {
	apps := ta.atr.GetAppTenantsByTenant(tenantID)
	apptenantinfos := []domain.AppTenantInfo{}
	tmaps := ta.tr.GetTenantMapByIDs()
	amaps := ta.ar.GetAppMapByIDs()
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
