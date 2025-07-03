package models

import "gorm.io/gorm"

type Slip struct {
	gorm.Model
	Stake      string
	Status     string
	Selections []Selection
}

// has many selections

type Selection struct {
	gorm.Model
	IncomeId uint
	Status   string
	SlipID   uint
}
