package models

import "time"

type Item struct {
	ID         string
	Name       string
	CategoryID string
	CreatedAt  time.Time
}
