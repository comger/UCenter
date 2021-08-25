package entity

import (
	"ucenter/domain/ov"
)

// Custom  在没有购买应用时, 它只是代表潜在客户
type Custom struct {
	Name     string
	CustomID ov.CustomID
}

type Group struct {
	CustomID ov.CustomID
	GroupID  ov.GroupID
	Name     string
	IsGroup  bool
	Members  []Group
}

func (g *Group) AppendMember(aMember *Group) {
	if g.IsGroup {
		g.Members = append(g.Members, *aMember)
	}

}

// Tenant 租户，代表着客户在应用程序ClientID 购买开通了Modules, 购买凭证为 TenantID
type Tenant struct {
	Custom
	ClientID     ov.ClientID
	TenantID     ov.TenantID
	OrderModuels ov.OrderModuels
	actived      bool
}

func (t *Tenant) Active() {
	t.actived = true
}

func (t *Tenant) Deactive() {
	t.actived = false
}

func (t *Tenant) IsActived() bool {
	return t.actived
}

// 初始化客户购买应用程序
func NewTenant(c *Custom, app *Application, aOrderModuels ov.OrderModuels) *Tenant {
	// 判断oOrderModuels 是否在当前app 的Modules keys 中
	if app.Modules.IsReady(aOrderModuels) {
		var t = &Tenant{
			Custom:       *c,
			ClientID:     app.ClientID,
			TenantID:     ov.TenantID(ov.NewOID()),
			OrderModuels: aOrderModuels,
		}
		return t
	} else {
		return nil
	}
}

// SetModules 客户重新购买，调整可用模块
func (aTenant *Tenant) SetModules(app *Application, aOrderModuels ov.OrderModuels) {
	// 判断oOrderModuels 是否在当前app 的Modules keys 中
	if app.Modules.IsReady(aOrderModuels) {
		aTenant.OrderModuels = aOrderModuels
	}
}

func (aTenant *Tenant) AddRole(aName string, aModules ov.OrderModuels) *Role {
	return NewRole(*aTenant, aName, aModules)
}
