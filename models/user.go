package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `json:"id" xml:"id"`
	Name      string         `json:"name" xml:"name"`
	Email     string         `gorm:"type:varchar(191);unique" json:"email" xml:"email"`
	Password  string         `json:"-" xml:"password,omitempty"`
	CreatedAt time.Time      `json:"created_at,omitempty" xml:"-"`
	UpdatedAt time.Time      `json:"updated_at,omitempty" xml:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-" xml:"-"`
}
