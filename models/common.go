package models

import uuid "github.com/satori/go.uuid"

// DbSyncNotify 数据库同步广播标识
const DbSyncNotify = "dbSyncNotify"

const (
	UserTable = 1
)

const (
	// InsertOp insert 操作
	InsertOp = 1
	// UpdateOp update 操作
	UpdateOp = 2
	// DeleteOp delete 操作
	DeleteOp = 3
)

// DbSyncDes 数据库同步数据描述符
type DbSyncDes struct {
	Type  int       // 1 insert, 2 update, 3 delete
	Table int       // 1 user
	UUID  uuid.UUID // 本实例uuid
	Data  []byte    // 表数据
}
