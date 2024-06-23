package gormdb

import (
	"github.com/yohanesmario/online-book-store/internal/domain/repository/base"
	"gorm.io/gorm"
)

type dbConnectionImpl struct {
	db *gorm.DB
}

func NewDBConnection(db *gorm.DB) base.IDBConnection[*gorm.DB] {
	return &dbConnectionImpl{db: db}
}

func (i *dbConnectionImpl) GetDBConnection() *gorm.DB {
	return i.db
}

var _ base.IDBConnection[*gorm.DB] = (*dbConnectionImpl)(nil)
