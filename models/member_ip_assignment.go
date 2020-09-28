package models

import "time"

// ZtcMemberIpAssignment
type ZtcMemberIpAssignment struct {
	Id           int64     //主键
	CreationTime time.Time `orm:"auto_now_add;type(datetime)"` //第一次保存时才设置时间
	NetworkId    string    `orm:"size(16)"`
	MemberId     string    `orm:"size(16)"`
	Address      string
}

//
