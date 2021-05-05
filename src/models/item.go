package models

import "time"

type Item struct {
	ID         string `gorm:"primary_key"`
	Name       string
	CategoryID string
	CreatedAt  time.Time
}
