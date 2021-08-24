package entity

import (
	"ucenter/domain/ov"
)

type Application struct {
	Name         string
	Modules      ov.Modules
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
