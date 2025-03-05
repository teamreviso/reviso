// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package models

import (
	"time"
)

const TableNameDocumentTimelineCheckpoint = "document_timeline_checkpoints"

// DocumentTimelineCheckpoint mapped from table <document_timeline_checkpoints>
type DocumentTimelineCheckpoint struct {
	ID             string    `gorm:"column:id;primaryKey;default:uuid_generate_v4()" json:"id"`
	DocumentID     string    `gorm:"column:document_id;not null" json:"document_id"`
	UserID         string    `gorm:"column:user_id;not null" json:"user_id"`
	LastEventID    string    `gorm:"column:last_event_id;not null" json:"last_event_id"`
	ContentAddress string    `gorm:"column:content_address;not null" json:"content_address"`
	CreatedAt      time.Time `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt      time.Time `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

// TableName DocumentTimelineCheckpoint's table name
func (*DocumentTimelineCheckpoint) TableName() string {
	return TableNameDocumentTimelineCheckpoint
}
