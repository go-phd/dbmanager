package models

// ZtcNetworkRoute 网络路由表
type ZtcNetworkRoute struct {
	ID           int64 //主键
	CreationTime int64
	NetworkId    string
	Address      string
	Bits         string
	via          string
}
