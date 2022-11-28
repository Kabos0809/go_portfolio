package Models

func (m Model) GetAllWork() (*[]Work, error) {
	var work []Work
	//タグに対応できたら使う
	//tx := m.Db.Preload("Tag").Begin()
	tx := m.Db.Begin()
	err := tx.Order("created_at").Find(&work).Error
	if err != nil {
		tx.Rollback()
		return &work, err
	}
	f := func(b []Work) []Work {
		l := len(b)
		for i := 0; i < l/2; i++ {
			b[i], b[l-i-1] = b[l-i-1], b[i]
		}
		return b
	}
	work = f(work)
	tx.Commit()
	return &work, err
}

func (m Model) GetWorkByID(id uint64) (*Work, error) {
	var work *Work
	//同上
	//tx := m.Db.Preload("Tag").Begin()
	tx := m.Db.Begin()
	err := tx.Where("id = ?", id).Find(&work).Error
	if err != nil {
		tx.Rollback()
		return work, err
	}
	tx.Commit()
	return work, err
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
	err := tx.Save(work).Error
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

func (m Model) IncrementSeeWork(work *Work) error {
	tx := m.Db.Begin()
	work.SeeCount = work.SeeCount + 1
	if err := tx.Save(work).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (m Model) ChangeWorkIsActive(id uint64) error {
	tx := m.Db.Begin()
	work, err := m.GetWorkByID(id)
	if err != nil {
		tx.Rollback()
		return err
	} else {
		if work.IsActive == true {
			work.IsActive = false
		} else {
			work.IsActive = true
		}
		if err := tx.Save(work).Error; err != nil {
			tx.Rollback()
			return err
		}
		tx.Commit()
		return nil
	}
}

//func (m Model) GetAllWorkTag(tags *[]WorkTag) error {
//	if err := m.Db.Select("name").Group("name").Find(tags).Error; err != nil {
//		return err
//	}
//	return nil
//}
//
//func (m Model) GetWorkByTag(works *[]Work, tag string) error {
//	tx := m.Db.Preload("Tag").Begin()
//	if err := tx.Joins("inner join worktag on work.id = worktag.work_id").Where("worktag.name = ?", tag).Preload("Tag").Find(works).Error; err != nil {
//		tx.Rollback()
//		return err
//	}
//	tx.Commit()
//	return nil
//}