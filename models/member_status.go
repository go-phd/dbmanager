package models

import "time"

// ZtcMemberStatus
type ZtcMemberStatus struct {
	Id           int64     //主键
	CreationTime time.Time `orm:"auto_now_add;type(datetime)"` //第一次保存时才设置时间
	NetworkId    string    `orm:"size(16)"`
	Member_id    string    `orm:"size(16)"`
	Address      string
	LastUpdated  time.Time `orm:"auto_now;type(datetime)"` //每次 model 保存时都会对时间自动更新
}

//
