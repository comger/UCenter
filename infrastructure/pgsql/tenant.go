package pgsql

import (
	"time"
	"ucenter/domain"

	"github.com/jinzhu/gorm"
	"gopkg.in/mgo.v2/bson"
)

type TenantORM struct {
	ID        string     `gorm:"primary_key;not null;unique" json:"id"`
	Name      string     `gorm:"size:100;not null;unique" json:"name"`
	CreatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

func (TenantORM) TableName() string {
	return "Tenant"
}

func NewTenantORM(tc *domain.TenantCreate) *TenantORM {
	return &TenantORM{Name: tc.Name, ID: bson.NewObjectId().String()}
}

func InitTenantORM(t *domain.Tenant) *TenantORM {
	return &TenantORM{Name: t.Name, ID: t.ID}
}

func (to *TenantORM) toModel() *domain.Tenant {
	t := &domain.Tenant{}
	t.Name = to.Name
	t.ID = to.ID
	return t
}

type TenantRepo struct {
	db *gorm.DB
}

var _ domain.TenantRepoInterface = &TenantRepo{}

func NewTenantRepo(db *gorm.DB) *TenantRepo {
	return &TenantRepo{db}
}

func (tr *TenantRepo) CreateTenant(tc *domain.TenantCreate) (*domain.Tenant, error) {
	to := NewTenantORM(tc)
	err := tr.db.Debug().Create(to).Error
	if err != nil {
		return nil, err
	}

	return to.toModel(), nil
}

func (tr *TenantRepo) UpdateTenant(t *domain.Tenant) error {
	to := InitTenantORM(t)
	err := tr.db.Debug().Save(to).Error
	if err != nil {
		return err
	}
	return nil
}

func (tr *TenantRepo) GetTenants() []domain.Tenant {
	var tos []TenantORM
	err := tr.db.Order("created_at desc").Find(&tos).Error
	if err != nil {
		return nil
	}

	var ts []domain.Tenant
	for _, t := range tos {
		ts = append(ts, *t.toModel())
	}
	return ts

}

func (tr *TenantRepo) GetTenantMapByIDs() map[string]string {
	panic("not implemented") // TODO: Implement
}

func (tr *TenantRepo) GetTenant(ID string) (*domain.Tenant, error) {
	var t domain.Tenant
	err := tr.db.Where("ID=?", ID).Take(t).Error
	if err != nil {
		return nil, err
	}
	return &t, nil
}
