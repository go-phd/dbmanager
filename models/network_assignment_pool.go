package models

import "time"

// ZtcNetworkAssignmentPool
type ZtcNetworkAssignmentPool struct {
	Id           int64     //主键
	CreationTime time.Time `orm:"auto_now_add;type(datetime)"` //第一次保存时才设置时间
	NetworkId    string    `orm:"size(16)"`
	IpRangeStart string
	IpRangeEnd   string
}

//
