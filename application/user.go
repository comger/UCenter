package application

import "ucenter/domain"

type UserAppInterface interface {
	domain.UserRepoInterface
	GetUserWithTenants(string, int64, int64) []domain.UserWithTenant
}

type UserApp struct {
	ur domain.UserRepoInterface
	tr domain.TenantRepoInterface
}

var _ UserAppInterface = &UserApp{}

func (ua *UserApp) CreateUser(uc *domain.UserCreate) (*domain.User, error) {
	return ua.ur.CreateUser(uc)
}

func (ua *UserApp) RegistUser(u *domain.UserRegist) error {
	return ua.ur.RegistUser(u)
}

func (ua *UserApp) UpdateUser(u *domain.User) error {
	return ua.ur.UpdateUser(u)
}

func (ua *UserApp) GetUsers(key string, offset int64, limit int64) []domain.User {
	return ua.ur.GetUsers(key, offset, limit)
}

// GetUserWithTenants is Aggreate for user and tenant
func (ua *UserApp) GetUserWithTenants(key string, offset int64, limit int64) []domain.UserWithTenant {
	users := ua.ur.GetUsers(key, offset, limit)
	tmap := ua.tr.GetTenantMapByIDs()

	uwts := []domain.UserWithTenant{}
	for _, u := range users {
		tenantName := tmap[u.TenantID]
		uwt := domain.UserWithTenant{User: u, TenantName: tenantName}
		uwts = append(uwts, uwt)
	}

	return uwts
}

func (ua *UserApp) GetUsersByTenantID(tenantID string, key string, offset int64, limit int64) []domain.User {
	return ua.ur.GetUsersByTenantID(tenantID, key, offset, limit)
}

func (ua *UserApp) GetUser(ID string) (*domain.User, error) {
	return ua.ur.GetUser(ID)
}
