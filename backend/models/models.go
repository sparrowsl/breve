package models

import "gorm.io/gorm"

type Breve struct {
	gorm.Model
	Redirect string
	Link     string
	Clicked  uint8
	Random   bool
}
