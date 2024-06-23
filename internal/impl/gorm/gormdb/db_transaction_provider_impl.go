package gormdb

import (
	"github.com/yohanesmario/online-book-store/internal/domain/repository/base"
	"gorm.io/gorm"
)

type dbTransactionProviderImpl struct {
	db base.IDBConnection[*gorm.DB]
}

func NewDBTransactionProvider(db base.IDBConnection[*gorm.DB]) base.DBTransactionProvider[*gorm.DB] {
	return &dbTransactionProviderImpl{db: db}
}

func (i *dbTransactionProviderImpl) Transaction(txBody func(txConnection base.IDBConnection[*gorm.DB]) error) error {
	return i.db.GetDBConnection().Transaction(func(tx *gorm.DB) error {
		return txBody(&dbConnectionImpl{db: tx})
	})
}

var _ base.DBTransactionProvider[*gorm.DB] = (*dbTransactionProviderImpl)(nil)
