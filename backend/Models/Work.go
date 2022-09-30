package Models

import (
	"gorm.io/gorm"
)

func (m Model) GetAllWork() (*[]Work, error) {
	var work []Work
	tx := m.Db.Preload("Tag").Begin()
	r, err := tx.Find(&work).Error
	if err != nil {
		tx.Rollback()
		return &work, err
	}
	tx.Commit()
	return &work, err
}

func (m Model) GetWorkByID(id uint64) (*Work, error) {
	var work *Work
	tx := m.Db.Preload("Tag").Begin()
	r, err := tx.Where("id = ?", id).Find(&work).Error
	if err != nil {
		tx.Rollback()
		return &work, err
	}
	tx.Commit()
	return &work, err
}

func (m Model) CreateWork(work *Work) error {
	tx := m.Db.Begin()
	err := tx.Create(work).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return err
}

func (m Model) UpdateWork(work *Work) error {
	tx := m.Db.Begin()
	err := m.Save(work).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return err
}

func (m Model) DeleteWork(id uint64) error {
	tx := m.Db.Begin()
	err := tx.Where("id = ?", id).Delete(&Work{}).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return err
}