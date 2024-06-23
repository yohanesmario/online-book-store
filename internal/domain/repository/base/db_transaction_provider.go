package base

type DBTransactionProvider[TDBConnection any] interface {
	Transaction(txBody func(txConnection IDBConnection[TDBConnection]) error) error
}
