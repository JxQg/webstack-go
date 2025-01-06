// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameStCategory = "st_category"

// StCategory mapped from table <st_category>
type StCategory struct {
	ID        int        `gorm:"column:id;type:INTEGER;primaryKey" json:"id"`
	ParentID  int        `gorm:"column:parent_id;type:int(11);not null" json:"parent_id"`
	Sort      int        `gorm:"column:sort;type:int(11);not null" json:"sort"`
	Title     string     `gorm:"column:title;type:varchar(50);not null" json:"title"`
	Icon      string     `gorm:"column:icon;type:varchar(20);not null" json:"icon"`
	Level     int32      `gorm:"column:level;type:integer;not null" json:"level"`
	IsUsed    bool       `gorm:"column:is_used;type:bool;default:false" json:"is_used"`
	CreatedAt *time.Time `gorm:"column:created_at;type:datetime;not null;default:CURRENT_TIMESTAMP not null" json:"created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at;type:datetime;not null;default:CURRENT_TIMESTAMP not null" json:"updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at;type:datetime" json:"deleted_at"`
}

// TableName StCategory's table name
func (*StCategory) TableName() string {
	return TableNameStCategory
}