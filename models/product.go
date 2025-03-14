package models

import (
	"errors"
	"log"

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
	if err = p.validate(); err != nil {
		return err
	}
	log.Printf("[HOOK] Creating product: %s with price: %.2f", p.Name, p.Price)
	return
}

func (p *Product) AfterCreate(tx *gorm.DB) (err error) {
	log.Printf("[HOOK] Product created: %s with price: %.2f", p.Name, p.Price)
	return
}

func (p *Product) BeforeUpdate(tx *gorm.DB) (err error) {
	if err = p.validate(); err != nil {
		return err
	}
	log.Printf("[HOOK] Updating product: %s with new price: %.2f", p.Name, p.Price)
	return
}

func (p *Product) AfterUpdate(tx *gorm.DB) (err error) {
	log.Printf("[HOOK] Product updated: %s with new price: %.2f", p.Name, p.Price)
	return
}

func (p *Product) validate() error {
	if p.Name == "" {
		return errors.New("product name cannot be empty")
	}
	if p.Price < 0 {
		return errors.New("price cannot be negative")
	}
	return nil
}
