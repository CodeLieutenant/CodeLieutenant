package models

import "time"

type Model struct {
	ID        uint64    `json:"id,omitempty" yaml:"id,omitempty" gorm:"primarykey"`
	CreatedAt time.Time `json:"created_at,omitempty" yaml:"created_at,omitempty"`
	UpdatedAt time.Time `json:"-" yaml:"-"`
}
