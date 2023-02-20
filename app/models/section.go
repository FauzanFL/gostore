package models

import "time"

type Section struct {
	ID        string `grom:"size:36;not null;uniqueIndex;primary_key"`
	Name      string `grom:"size:100"`
	Slug      string `grom:"size:100"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Categories []Category
}