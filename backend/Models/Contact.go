package Models

func (m Model) GetAllContact() (*[]Contact, error) {
	var contact []Contact
	tx := m.Db.Begin()
	if err := tx.Find(&contact).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return &contact, nil
}

func (m Model) GetContactByID(id uint64) (*Contact, error) {
	var contact *Contact
	tx := m.Db.Begin()
	if err := tx.Where("id = ?", id).Find(&contact).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return contact, nil
}

func (m Model) CreateContact(contact *Contact) error {
	tx := m.Db.Begin()
	if err := tx.Create(contact).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (m Model) DeleteContact(id uint64) error {
	tx := m.Db.Begin()
	if err := tx.Where("id = ?", id).Delete(&Contact{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (m Model) ReadContact(id uint64) error {
	var contact *Contact
	tx := m.Db.Begin()
	if err := tx.Where("id = ?", id).Find(contact).Error; err != nil {
		tx.Rollback()
		return err
	}
	contact.Read = true
	if err := tx.Save(contact).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}