package ov

import (
	"github.com/gofrs/uuid"
	"gopkg.in/mgo.v2/bson"
)

type OID bson.ObjectId

type ClientID OID // 应用程序标记
type CustomID OID // 客户标记
type TenantID OID // 租户标记
type RoleID OID   // 角色标记
type UID OID      // 用户标记
type GroupID OID  // 分组标记

type Email string    // 邮件地址 格式只能是 xx@xx
type Mobile string   // 手机号码 长度只能是 11位
type Password string // 密码, 不可以明文存储, 不允许返回
type UserName string // 用户名

type Modules map[string]string // 可售卖的功能模块
type OrderModuels []string     // 已购买的模块列表
type UUIDString string         // 全局不重复UUID 字符串
type ClientSecret UUIDString   // 应用的安全秘钥
type CreatedAt int64           // 实例创建时间

func (cs *ClientSecret) IsNull() bool {
	return cs == nil
}

func NewUUIDString() UUIDString {
	hash, _ := uuid.NewV1()
	return UUIDString(hash.String())
}

func NewClientSecret() ClientSecret {
	return ClientSecret(NewUUIDString())
}

func NewOID() OID {
	return OID(bson.NewObjectId())
}

// 判断aOrderModuels 是否在当前Modules keys 中
func (ms *Modules) IsReady(aOrderModuels OrderModuels) bool {

	keys := make([]string, 0, len(*ms))
	for k := range *ms {
		keys = append(keys, k)
	}

	if len(keys) == 0 || len(keys) < len(aOrderModuels) {
		return false
	}

	return true
}

func (oms OrderModuels) IsReady(aOrderModuels OrderModuels) bool {
	keys := make([]string, 0, len(oms))

	if len(keys) == 0 || len(keys) < len(aOrderModuels) {
		return false
	}

	return true
}
