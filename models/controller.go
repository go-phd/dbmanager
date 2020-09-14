package models

// ZtcController
type ZtcController struct {
	ID                         int64 //主键
	Cluster_host               string
	Last_alive                 int64
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
