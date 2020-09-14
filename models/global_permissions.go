package models

// ZtcGlobalPermissions
type ZtcGlobalPermissions struct {
	ID           int64 //主键
	CreationTime int64
	UserId       string
	Authorize    bool
	Del          bool
	Modify       bool
	Read         bool
}

//
