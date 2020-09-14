package models

// ZtcNetworkAssignmentPool
type ZtcNetworkAssignmentPool struct {
	ID           int64 //主键
	CreationTime int64
	NetworkId    string
	IpRangeStart string
	IpRangeEnd   string
}

//
