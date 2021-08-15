package domain

// TenantCreate is input when Tenant Create
type TenantCreate struct {
	Name string
}

// Tenant is
type Tenant struct {
	TenantCreate
	ID string
}

// TenantRepoInterface Repo Interface
type TenantRepoInterface interface {
	CreateTenant(*TenantCreate) (*Tenant, error)
	UpdateTenant(*Tenant) error
	GetTenants() []Tenant
	GetTenantMapByIDs() map[string]string
	GetTenant(string) (*Tenant, error)
}
