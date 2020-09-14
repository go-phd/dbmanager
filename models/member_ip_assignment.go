package models

// ZtcMemberIpAssignment
type ZtcMemberIpAssignment struct {
	ID           int64 //主键
	CreationTime int64
	NetworkId    string
	MemberId     string
	Address      string
}

//
