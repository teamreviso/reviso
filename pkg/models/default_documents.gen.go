// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package models

const TableNameDefaultDocument = "default_documents"

// DefaultDocument mapped from table <default_documents>
type DefaultDocument struct {
	Name  string `gorm:"column:name;primaryKey" json:"name"`
	DocID string `gorm:"column:doc_id;not null" json:"doc_id"`
}

// TableName DefaultDocument's table name
func (*DefaultDocument) TableName() string {
	return TableNameDefaultDocument
}
