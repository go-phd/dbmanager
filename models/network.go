package models

// ZtcNetwork 网络表
type ZtcNetwork struct {
	ID                int64 //主键
	CreationTime      int64
	OwnerId           string
	ControllerId      string
	Capabilities      string
	EnableBroadcast   bool
	LastModified      int64
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
