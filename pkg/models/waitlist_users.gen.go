// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package models

import (
	"time"
)

const TableNameWaitlistUser = "waitlist_users"

// WaitlistUser mapped from table <waitlist_users>
type WaitlistUser struct {
	Email       string    `gorm:"column:email;not null" json:"email"`
	CreatedAt   time.Time `gorm:"column:created_at;not null;default:now()" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at;not null;default:now()" json:"updated_at"`
	AllowAccess bool      `gorm:"column:allow_access;not null" json:"allow_access"`
}

// TableName WaitlistUser's table name
func (*WaitlistUser) TableName() string {
	return TableNameWaitlistUser
}
