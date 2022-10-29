package Models

func (m Model) GetAllBlog() (*[]Blog, error) {
	var blogs []Blog
	//タグに対応できたら使う
	//tx := m.Db.Preload("Tag").Begin()
	tx := m.Db.Begin()
	if err := tx.Order("created_at").Find(&blogs).Error; err != nil {
		tx.Rollback()
		return &blogs, err
	}
	f := func(b []Blog) []Blog {
		l := len(b)
		for i := 0; i < l/2; i++ {
			b[i], b[l-i-1] = b[l-i-1], b[i]
		}
		return b
	}
	blogs = f(blogs)
	tx.Commit()
	return &blogs, nil
}

func (m Model) GetBlogByID(id uint64) (*Blog, error) {
	var blog *Blog
	//同上
	//tx := m.Db.Preload("Tag").Begin()
	tx := m.Db.Begin()
	if err := tx.Where("id = ?", id).Find(&blog).Error; err != nil {
		tx.Rollback()
		return blog, err
	}
	tx.Commit()
	return blog, nil
}

func (m Model) CreateBlog(blog *Blog) error {
	tx := m.Db.Begin()
	if err := tx.Create(blog).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (m Model) UpdateBlog(blog *Blog) error {
	tx := m.Db.Begin()
	if err := tx.Save(blog).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (m Model) DeleteBlog(id uint64) error {
	tx := m.Db.Begin()
	if err := tx.Where("id = ?", id).Delete(&Blog{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (m Model) IncrementSeeBlog(blog *Blog) error {
	tx := m.Db.Begin()
	blog.SeeCount = blog.SeeCount + 1
	err := tx.Save(blog).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (m Model) ChangeBlogIsActive(id uint64) error {
	tx := m.Db.Begin()
	blog, err := m.GetBlogByID(id)
	if err != nil {
		tx.Rollback()
		return err
	} else {
		if blog.IsActive == true {
			blog.IsActive = false
		} else {
			blog.IsActive = true
		}
		err = tx.Save(blog).Error
		if err != nil {
			tx.Rollback()
			return err
		}
		tx.Commit()
		return nil
	}
}

//func (m Model) GetAllBlogTag(tags *[]BlogTag) error {
//	if err := m.Db.Select("name").Group("name").Find(tags).Error; err != nil {
//		return err
//	}
//	return nil
//}
//
//func (m Model) GetBlogByTag(blogs *[]Blog, tag string) error {
//	tx := m.Db.Preload("Tag").Begin()
//	if err := tx.Joins("inner join blogtag on blog.id = blogtag.blog_id").Where("blogtag.name = ?", tag).Preload("Tag").Find(blogs).Error; err != nil {
//		tx.Rollback()
//		return err
//	}
//	tx.Commit()
//	return nil
//}