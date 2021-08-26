package pgsql

import (
	"time"
	"ucenter/domain"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
)

type UserORM struct {
	ID        string     `gorm:"primary_key;auto_increment" json:"id"`
	UserName  string     `gorm:"size:35;not null;unique" json:"username"`
	Email     string     `gorm:"size:100;not null;unique" json:"email"`
	Mobile    string     `gorm:"size:15;not null;unique" json:"mobile"`
	TanentID  uint64     `gorm:"not null" json:"tanent_id"`
	Password  string     `gorm:"text;not null;" json:"password"`
	NickName  string     `gorm:"size:45;not null;" json:"nickname"`
	IsSuper   bool       `json:"issuper"`
	CreatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

func NewUserORM(uc *domain.UserCreate) *UserORM {
	// var uo UserORM
	// for key, val :=range(uc){
	// 	uo[key] = val
	// }
	id := bson.NewObjectId().String()
	hash, _ := bcrypt.GenerateFromPassword([]byte(uc.Password), bcrypt.MinCost)
	return &UserORM{ID: id,
		UserName: uc.UserName,
		NickName: uc.NickName,
		Email:    uc.Email,
		Mobile:   uc.Mobile,
		Password: string(hash),
	}
}

type UserRepo struct {
	db *gorm.DB
}

var _ domain.UserRepoInterface = &UserRepo{}

func (ur *UserRepo) CreateUser(_ *domain.UserCreate) (*domain.User, error) {
	panic("not implemented") // TODO: Implement
}

func (ur *UserRepo) RegistUser(_ *domain.UserRegist) error {
	panic("not implemented") // TODO: Implement
}

func (ur *UserRepo) UpdateUser(_ *domain.User) error {
	panic("not implemented") // TODO: Implement
}

func (ur *UserRepo) GetUsers(_ string, _ int64, _ int64) []domain.User {
	panic("not implemented") // TODO: Implement
}

func (ur *UserRepo) GetUsersByTenantID(_ string, _ string, _ int64, _ int64) []domain.User {
	panic("not implemented") // TODO: Implement
}

func (ur *UserRepo) GetUser(_ string) (*domain.User, error) {
	panic("not implemented") // TODO: Implement
}
