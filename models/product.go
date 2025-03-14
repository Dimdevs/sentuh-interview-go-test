package models

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name       string    `json:"name" xml:"name" gorm:"not null"`
	Price      float64   `json:"price" xml:"price" gorm:"not null;default:0"`
	CategoryID *uint     `json:"category_id,omitempty" xml:"category_id,omitempty"`
	Category   *Category `json:"category,omitempty" xml:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
	if p.Name == "" {
		return errors.New("product name cannot be empty")
	}
	if p.Price < 0 {
		return errors.New("price cannot be negative")
	}
	fmt.Println("[HOOK] Creating product:", p.Name, "with price:", p.Price)
	return
}

func (p *Product) BeforeUpdate(tx *gorm.DB) (err error) {
	if p.Name == "" {
		return errors.New("product name cannot be empty")
	}
	if p.Price < 0 {
		return errors.New("price cannot be negative")
	}
	fmt.Println("[HOOK] Updating product:", p.Name, "with new price:", p.Price)
	return
}

func (p *Product) AfterCreate(tx *gorm.DB) (err error) {
	fmt.Println("[HOOK] Product created:", p.Name, "with price:", p.Price)
	return
}
