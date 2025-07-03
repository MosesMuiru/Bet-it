package store

import (
	"fmt"
	db "kwik/storage"
	"kwik/storage/models"

	"gorm.io/gorm"
)

type Db struct {
	db *gorm.DB
}

func NewDb() *Db {
	return &Db{
		db.DB(),
	}
}

func (store *Db) CreateSlip(slip *models.Slip) (*models.Slip, error) {

	results := store.db.Create(&slip)

	if results.Error != nil {
		fmt.Println("Error creating the slip", results.Error)
		return nil, results.Error
	}

	return slip, nil
}

func (store *Db) GetAllSlips() (*[]models.Slip, error) {
	slips := &[]models.Slip{}
	results := store.db.Find(slips)

	if results.Error != nil {
		fmt.Println("No slip found", results.Error)
		return nil, results.Error
	}

	return slips, nil
}

func (store *Db) GetSlipById(id uint) (*models.Slip, error) {
	slip := &models.Slip{}

	results := store.db.Where("id = ?", id).Find(slip)
	if results.Error != nil {
		fmt.Println("this is the error", results.Error)
		return nil, results.Error
	}

	return slip, nil
}

func (store *Db) UpdateSlipStatusById(id uint, newStatus string) error {
	result := store.db.Model(&models.Slip{}).Where("id = ?", id).Update("status", newStatus)
	if result.Error != nil {
		fmt.Println("err", result.Error)
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("no slip found with ID %d", id)
	}
	return nil
}
