package models

import "time"

// ZtcNetwork 网络表
type ZtcNetwork struct {
	Id                int64     //主键
	CreationTime      time.Time `orm:"auto_now_add;type(datetime)"` //第一次保存时才设置时间
	OwnerId           string
	ControllerId      string
	Capabilities      string
	EnableBroadcast   bool
	LastModified      time.Time `orm:"auto_now;type(datetime)"`
	Mtu               int32
	Name              string
	Private           bool
	RemoteTraceTarget string
	RemoteTraceLevel  int32
	Rules             string
	RulesSource       string
	Tags              string
	V4AssignMode      string
	V6AssignMode      string
	AuthTokens        string
	Routes            string
	IpAssignmentPools string
	Revision          int32
	Deleted           bool
}

//
