package models

import (
	"log"
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

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	log.Printf("Membuat User baru: %v", u)
	return
}

func (u *User) AfterCreate(tx *gorm.DB) (err error) {
	log.Printf("User berhasil dibuat: %v", u)
	return
}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	log.Printf("Memperbarui User: %v", u)
	return
}

func (u *User) AfterUpdate(tx *gorm.DB) (err error) {
	log.Printf("User berhasil diperbarui: %v", u)
	return
}
