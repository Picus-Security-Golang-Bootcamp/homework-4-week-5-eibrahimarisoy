package service

import (
	"github.com/Picus-Security-Golang-Bootcamp/homework-4-week-5-eibrahimarisoy/model"
	"gorm.io/gorm"
)

// CreateAuthor creates an author
func CreateAuthor(db *gorm.DB, author *model.Author) (*model.Author, error) {
	result := db.Where("name = ?", author.Name).FirstOrCreate(author)

	if result.Error != nil {
		return nil, result.Error
	}

	return author, nil
}

// GetAuthorByID returns an author by id
func GetAuthorByID(db *gorm.DB, id int) (model.Author, error) {
	var author model.Author

	result := db.Where("id = ?", id).First(&author)
	if result.Error != nil {
		return model.Author{}, result.Error
	}
	return author, nil
}

// UpdateAuthor updates author
func UpdateAuthor(db *gorm.DB, author *model.Author) (*model.Author, error) {
	if err := db.Save(&author).Error; err != nil {
		return author, err
	}
	return author, nil
}

// GetByIDWithBooks returns author with books
func GetByIDWithBooks(db *gorm.DB, id int) (model.Author, error) {
	var author model.Author

	result := db.Preload("Books").Where("id = ?", id).First(&author)
	if result.Error != nil {
		return model.Author{}, result.Error
	}
	return author, nil
}

// GetAuthorsWithBooks returns authors with books
func GetAuthorsWithBooks(db *gorm.DB) ([]model.Author, error) {
	var authors []model.Author
	result := db.Preload("Books").Find(&authors)
	if result.Error != nil {
		return []model.Author{}, result.Error
	}
	return authors, nil
}

// DeleteByID deletes author by id
func DeleteAuthorByID(db *gorm.DB, id int) error {
	result := db.Where("id = ?", id).Delete(&model.Author{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
