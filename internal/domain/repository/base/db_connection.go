package base

type IDBConnection[TDBConnection any] interface {
	GetDBConnection() TDBConnection
}
