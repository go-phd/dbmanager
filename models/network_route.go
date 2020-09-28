package models

import "time"

// ZtcNetworkRoute 网络路由表
type ZtcNetworkRoute struct {
	Id           int64     //主键
	CreationTime time.Time `orm:"auto_now_add;type(datetime)"` //第一次保存时才设置时间
	NetworkId    string    `orm:"size(16)"`
	Address      string
	Bits         string
	via          string
}

//
