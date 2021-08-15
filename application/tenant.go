package application

import "ucenter/domain"

type TenantAppInterface interface {
	domain.TenantRepoInterface
}

type TenantApp struct {
	tr domain.TenantRepoInterface
}

var _ TenantAppInterface = &TenantApp{}

func (ta TenantApp) CreateTenant(tc *domain.TenantCreate) (*domain.Tenant, error) {
	return ta.tr.CreateTenant(tc)
}

func (ta TenantApp) UpdateTenant(t *domain.Tenant) error {
	return ta.tr.UpdateTenant(t)
}

func (ta TenantApp) GetTenants() []domain.Tenant {
	return ta.tr.GetTenants()
}

func (ta TenantApp) GetTenant(ID string) (*domain.Tenant, error) {
	return ta.tr.GetTenant(ID)
}

func (ta TenantApp) GetTenantMapByIDs() map[string]string {
	return ta.tr.GetTenantMapByIDs()
}
