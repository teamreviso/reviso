// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package models

import (
	"time"

	"gorm.io/gorm"
)

const TableNameDocument = "documents"

// Document mapped from table <documents>
type Document struct {
	ID            string         `gorm:"column:id;primaryKey;default:uuid_generate_v4()" json:"id"`
	Title         string         `gorm:"column:title;not null" json:"title"`
	CreatedAt     time.Time      `gorm:"column:created_at;not null;default:now()" json:"created_at"`
	UpdatedAt     time.Time      `gorm:"column:updated_at;not null;default:now()" json:"updated_at"`
	ParentID      *string        `gorm:"column:parent_id" json:"parent_id"`
	DeletedAt     gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	IsPublic      bool           `gorm:"column:is_public;not null;default:true" json:"is_public"`
	RogueVersion  string         `gorm:"column:rogue_version" json:"rogue_version"`
	RootParentID  string         `gorm:"column:root_parent_id;not null" json:"root_parent_id"`
	ParentAddress *string        `gorm:"column:parent_address" json:"parent_address"`
	IsFolder      bool           `gorm:"column:is_folder" json:"is_folder"`
	FolderID      *string        `gorm:"column:folder_id" json:"folder_id"`
}

// TableName Document's table name
func (*Document) TableName() string {
	return TableNameDocument
}
