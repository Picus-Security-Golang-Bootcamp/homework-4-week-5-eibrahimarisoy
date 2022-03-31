package service

import (
	"github.com/Picus-Security-Golang-Bootcamp/homework-4-week-5-eibrahimarisoy/model"
	"gorm.io/gorm"
)

type AuthorRepository struct {
	db *gorm.DB
}

// NewAuthorRepository returns a new AuthorRepository
func NewAuthorRepository(db *gorm.DB) *AuthorRepository {
	return &AuthorRepository{db: db}

}

// Migrations runs the database migrations
func Migrations(db *gorm.DB) {
	db.AutoMigrate(&model.Author{})
}

// InsertSampleData inserts sample data into the database
func InsertSampleData(db *gorm.DB, author *model.Author) model.Author {
	result := db.Unscoped().Where("name = ?", author.Name).FirstOrCreate(author)

	if result.Error != nil {
		panic(result.Error) // TODO: handle error
	}
	return *author
}

// GetByID returns an author by id
func GetByID(db *gorm.DB, id int) (model.Author, error) {
	var author model.Author

	result := db.Where("id = ?", id).First(&author)
	if result.Error != nil {
		return model.Author{}, result.Error
	}
	return author, nil
}

// FindByName returns an author by name
func FindByName(db *gorm.DB, name string) ([]model.Author, error) {
	var authors []model.Author

	result := db.Where("name ILIKE ?", "%"+name+"%").Find(&authors)
	if result.Error != nil {
		return nil, result.Error
	}
	return authors, nil
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

// **************EXTRA QUERIES************** //

// DeleteByID deletes author by id
func DeleteAuthorByID(db *gorm.DB, id int) error {
	result := db.Where("id = ?", id).Delete(&model.Author{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// UpdateAuthorName updates author name
func UpdateAuthorName(db *gorm.DB, author *model.Author) error {
	result := db.Model(&author).Where("id = ?", author.ID).Update("name", author.Name)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
