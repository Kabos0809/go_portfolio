package Models

import (
	"gorm.io/gorm"
)

func (m Model) GetAllContact() (*[]Contact, error) {
	var contact []Contact
	var err error
	defer return &contact, err
	tx := m.Db.Begin()
	if err := tx.Find(&contact).Error; err != nil {
		tx.Rollback()
	}
	tx.Commit()
}

func (m Model) GetContactByID(id uint64) (*Contact, error) {
	var contact *Contact
	var err error
	defer return contact, err
	tx := m.Db.Begin()
	if err := tx.Where("id = ?", id).Find(&contact).Error; err != nil {
		tx.Rollback()
	}
	tx.Commit()
}

func (m Model) CreateContact(contact *Contact) error {
	var err error
	defer return err
	tx := m.Db.Begin()
	if err := tx.Create(contact).Error; err != nil {
		tx.Rollback()
	}
	tx.Commit()
}

func (m Model) DeleteContact(id uint64) error {
	var err error
	defer return err
	tx := m.Db.Begin()
	if err := tx.Where("id = ?", id).Delete(&Contact{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return err
}

func (m Model) ReadContact(id uint64) error {
	var err error
	defer return err
	var contact *Contact
	tx := m.Db.Begin()
	if err := tx.Where("id = ?", id).Find(contact).Error; err != nil {
		tx.Rollback()
	}
	contact.read = true
	if err := tx.Save(contact).Error; err != nil {
		tx.Rollback()
	}
	tx.Commit()
}