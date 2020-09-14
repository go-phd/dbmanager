package models

// ZtcNetwork 网络表
type ZtcMember struct {
	ID                              int64 //主键
	creation_time                   int64
	NetworkId                       string
	Active_bridge                   bool
	Authorized                      bool
	Capabilities                    string
	Identity                        string
	Last_authorized_time            int64
	Last_deauthorized_time          int64
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
