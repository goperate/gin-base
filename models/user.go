package models

import (
	"time"
)

type User struct {
	ID    uint   `gorm:"primaryKey;column:id;autoIncrement"`
	Name  string `gorm:"index;size:50"`
	Email string `gorm:"size:50"`
	// autoCreateTime, autoUpdateTime gorm内置的自动设置创建时间/更新时间
	CreatedTime time.Time `gorm:"->;type:datetime;not null;default:CURRENT_TIMESTAMP"`
	UpdatedTime time.Time `gorm:"->;type:datetime;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}
