package models

import (
	"fmt"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name       string   `json:"name" xml:"name"`
	CategoryID uint     `json:"category_id" xml:"category_id"`
	Category   Category `json:"category" xml:"-"`
}

func (p *Product) BeforeUpdate(tx *gorm.DB) (err error) {
	fmt.Println("[HOOK] Updating product:", p.Name)
	return
}

func (p *Product) AfterCreate(tx *gorm.DB) (err error) {
	fmt.Println("[HOOK] Product created:", p.Name)
	return
}
