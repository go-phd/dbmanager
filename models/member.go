package models

import "time"

// ZtcNetwork 网络表
type ZtcMember struct {
	Id                              int64     //主键
	CreationTime                    time.Time `orm:"auto_now_add;type(datetime)"` //第一次保存时才设置时间
	NetworkId                       string    `orm:"size(16)"`
	Active_bridge                   bool
	Authorized                      bool
	Capabilities                    string
	Identity                        string
	LastAuthorizedTime              time.Time `orm:"auto_now;type(datetime)"`
	LastDeauthorizedTime            time.Time `orm:"auto_now;type(datetime)"`
	No_auto_assign_ips              string
	Remote_trace_target             bool
	Remote_trace_level              string
	Revision                        int32
	Tags                            string
	V_major                         int32
	V_minor                         int32
	V_rev                           int32
	V_proto                         int32
	Ip_assignments                  string
	Last_authorized_credential_type string
	Last_authorized_credential      string
	Deleted                         bool
	Address                         string
	Objtype                         string
	Hidden                          bool
}

//
