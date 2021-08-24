package entity

import "ucenter/domain/ov"

type EventLog struct {
	UID       ov.UID      // 谁
	CreatedAt int64       // 在什么时间
	TenantID  ov.TenantID // 在哪里
	Event     interface{} // 做了什么事
}
