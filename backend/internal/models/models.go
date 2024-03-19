package models

import "gorm.io/gorm"

type Link struct {
	gorm.Model
	Redirect string `json:"redirect" gorm:"not null"`
	Link     string `json:"link" gorm:"unique;not null"`
	Clicked  uint8  `json:"clicked"`
	Random   bool   `json:"random"`
}

type LinkModel struct{}

func (l *LinkModel) GetAllLinks() ([]Link, error) {
	var links []Link

	return links, nil
}
