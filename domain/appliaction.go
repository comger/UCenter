package domain

import (
	"github.com/gofrs/uuid"
	"gopkg.in/mgo.v2/bson"
)

type ApplicationCreate struct {
	Name string
}

type Application struct {
	ApplicationCreate
	ClientID     string
	ClientSecret string
}

func (a *Application) Prepare() {
	a.ClientID = bson.NewObjectId().String()
	hash, _ := uuid.NewV1()
	a.ClientSecret = hash.String()
}

type ApplicationRepoInterface interface {
	CreateApplication(*ApplicationCreate) (*Application, error)
	UpdateApplication(*Application) error
	GetApplications() []Application
	GetApplication(string) *Application
	GetAppMapByIDs() map[string]string
}
