package base

type BaseTransactionalRepository[TRepository any, TDBConnection any] interface {
	Use(txConnection IDBConnection[TDBConnection]) TRepository
}
