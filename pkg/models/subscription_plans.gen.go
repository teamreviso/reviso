// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package models

import (
	"time"
)

const TableNameSubscriptionPlan = "subscription_plans"

// SubscriptionPlan mapped from table <subscription_plans>
type SubscriptionPlan struct {
	ID            string    `gorm:"column:id;primaryKey;default:uuid_generate_v4()" json:"id"`
	Name          string    `gorm:"column:name;not null" json:"name"`
	PriceCents    int32     `gorm:"column:price_cents;not null" json:"price_cents"`
	Currency      string    `gorm:"column:currency;not null;default:USD" json:"currency"`
	Interval      string    `gorm:"column:interval;not null" json:"interval"`
	Status        string    `gorm:"column:status;default:active" json:"status"`
	StripePriceID string    `gorm:"column:stripe_price_id;not null" json:"stripe_price_id"`
	CreatedAt     time.Time `gorm:"column:created_at;not null;default:now()" json:"created_at"`
	UpdatedAt     time.Time `gorm:"column:updated_at;not null;default:now()" json:"updated_at"`
}

// TableName SubscriptionPlan's table name
func (*SubscriptionPlan) TableName() string {
	return TableNameSubscriptionPlan
}
