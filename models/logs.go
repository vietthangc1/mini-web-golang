package models

import (
	"time"
)

type Log struct {
	ID         uint `gorm:"primaryKey"`
	UserEmail  string
	TableModel string
	EntityID   uint64
	OldValue   string
	NewValue   string
	Timestamp  time.Time `gorm:"autoUpdateTime"`
}
