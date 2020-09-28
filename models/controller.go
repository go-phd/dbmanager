package models

import "time"

// ZtcController
type ZtcController struct {
	Id                         int64 //主键
	Cluster_host               string
	LastAlive                  time.Time `orm:"auto_now;type(datetime)"`
	Public_identity            string
	Last_authorized_credential string
	V_major                    int32
	V_minor                    int32
	V_rev                      int32
	V_build                    int32
	Host_port                  int32
	Use_rabbitmq               bool
}

//
