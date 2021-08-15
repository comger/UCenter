package domain

// UserCreate is
type UserCreate struct {
	UserName string
	Password string
	NickName string
	Email    string
	Mobile   string
	TenantID string
}

type UserRegist struct {
	UserCreate
	PasswordCheck string
}

type UserPasswordSet struct {
	Password string
	ID       string
}

type UserIsSuperdSet struct {
	IsSuper bool
	ID      string
}

type User struct {
	UserCreate
	ID      string
	IsSuper bool
}

type UserWithTenant struct {
	User
	TenantName string
}

type UserRepoInterface interface {
	CreateUser(*UserCreate) (*User, error)
	RegistUser(*UserRegist) error
	UpdateUser(*User) error
	GetUsers(string, int64, int64) []User
	GetUsersByTenantID(string, string, int64, int64) []User
	GetUser(string) (*User, error)
}
