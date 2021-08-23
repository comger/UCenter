package domain

type AppTenant struct {
	ClientID string
	TenantID string
	Modules  []string
}

type AppTenantInfo struct {
	AppTenant
	ApplicationName string
	TenantName      string
}

type AppTenantRepoInterface interface {
	CreateAppTenant(*AppTenant) (*AppTenant, error)
	UpdateAppTenant(*AppTenant) error
	DeleteAppTenant(*AppTenant) error
	GetAppTenantsByApp(string) []AppTenant
	GetAppTenantsByTenant(string) []AppTenant
	GetAppTenant(string, string) (*AppTenant, error)
}
