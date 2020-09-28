package models

import "time"

// ZtcGlobalPermissions
type ZtcGlobalPermissions struct {
	Id           int64     //主键
	CreationTime time.Time `orm:"auto_now_add;type(datetime)"`
	UserId       string
	Authorize    bool
	Del          bool
	Modify       bool
	Read         bool
}

//
