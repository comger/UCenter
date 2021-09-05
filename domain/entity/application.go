package entity

import (
	"ucenter/domain/ov"
)

type ApplicationCreate struct {
	Name    string
	Modules ov.Modules
}

type Application struct {
	ApplicationCreate
	ClientID     ov.ClientID
	clientSecret ov.ClientSecret
}

func (s *Application) GetClientScrect() ov.ClientSecret {
	if s.clientSecret.IsNull() {
		s.clientSecret = ov.NewClientSecret()
	}
	return s.clientSecret
}

func (s *Application) ResetClientScrect() ov.ClientSecret {
	s.clientSecret = ov.NewClientSecret()
	return s.clientSecret
}
