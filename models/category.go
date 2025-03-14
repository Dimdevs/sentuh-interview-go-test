package models

import (
	"log"

	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name     string    `json:"name" xml:"name"`
	Products []Product `json:"-" xml:"-"`
}

func (c *Category) BeforeCreate(tx *gorm.DB) (err error) {
	log.Printf("Membuat Category baru: %v", c)
	return
}

func (c *Category) AfterCreate(tx *gorm.DB) (err error) {
	log.Printf("Category berhasil dibuat: %v", c)
	return
}

func (c *Category) BeforeUpdate(tx *gorm.DB) (err error) {
	log.Printf("Memperbarui Category: %v", c)
	return
}

func (c *Category) AfterUpdate(tx *gorm.DB) (err error) {
	log.Printf("Category berhasil diperbarui: %v", c)
	return
}
