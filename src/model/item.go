package model

import "time"

type Item struct {
	ID         string `gorm:"primary_key"`
	Name       string
	CategoryID string
	Color      Color
	CreatedAt  time.Time
}
