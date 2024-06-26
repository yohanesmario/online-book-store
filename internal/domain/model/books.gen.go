// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameBook = "books"

// Book mapped from table <books>
type Book struct {
	ID        int32          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Title     string         `gorm:"column:title;not null" json:"title"`
	Author    string         `gorm:"column:author;not null" json:"author"`
	Isbn      string         `gorm:"column:isbn;not null" json:"isbn"`
	Price     float64        `gorm:"column:price;not null" json:"price"`
	CreatedAt *time.Time     `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
	CreatedBy *string        `gorm:"column:created_by" json:"created_by"`
	UpdatedAt *time.Time     `gorm:"column:updated_at;default:CURRENT_TIMESTAMP" json:"updated_at"`
	UpdatedBy *string        `gorm:"column:updated_by" json:"updated_by"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	DeletedBy *string        `gorm:"column:deleted_by" json:"deleted_by"`
}

// TableName Book's table name
func (*Book) TableName() string {
	return TableNameBook
}
