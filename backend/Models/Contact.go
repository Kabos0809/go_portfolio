package Models

import (
	"gorm.io/gorm"
)

func (m Model) GetAllContact() (*[]Contact, error) {
	var contact []Contact
	tx := m.Db.Begin()
	r, err := tx.Find(&contact).Error
	if err != nil {
		tx.Rollback()
		return r, err
	}
	tx.Commit()
	return r, err
}

func (m Model) GetContactByID(id uint64) (*Contact, error) {
	var contact *Contact
	tx := m.Db.Begin()
	r, err := tx.Where("id = ?", id).Find(&contact).Error
	if err != nil {
		tx.Rollback()
		return r, err
	}
	tx.Commit()
	return r, err
}

func (m Model) CreateContact(contact *Contact) error {
	tx := m.Db.Begin()
	err := tx.Create(contact).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return err
}

func (m Model) DeleteContact(id uint64) error {
	tx := m.Db.Begin()
	err := tx.Where("id = ?", id).Delete(&Contact{}).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return err
}