package service

import (
	"fmt"

	"github.com/Picus-Security-Golang-Bootcamp/homework-4-week-5-eibrahimarisoy/model"
	"gorm.io/gorm"
)

// SaveBook inserts sample data into the database
func SaveBook(db *gorm.DB, book *model.Book) (*model.Book, error) {
	result := db.Omit("Author").Where(model.Book{Name: book.Name, StockCode: book.StockCode}).
		FirstOrCreate(&book)
	if result.Error != nil {
		return nil, result.Error
	}

	return book, nil
}

// GetBooksWithAuthor returns books with author
func GetBooksWithAuthor(db *gorm.DB, args model.Args) (*model.Data, error) {
	books := []model.Book{}
	var filteredData int64

	table := "books"
	query := db.Select(table + ".*")
	query = query.Scopes(Search(args.Search))
	query = query.Scopes(Paginate(args))

	result := query.Preload("Author").Order("id").Find(&books)
	if result.Error != nil {
		return nil, result.Error
	}

	query.Offset(0)
	query.Table(table).Count(&filteredData)

	data := model.Data{
		FilteredData: filteredData,
		Data:         books,
	}

	return &data, nil

}

// GetByIDWithAuthor returns books by ID with author
func GetBookByIDWithAuthor(db *gorm.DB, id uint) (model.Book, error) {
	var book model.Book

	result := db.Preload("Author").Where("id = ?", id).First(&book)
	if result.Error != nil {
		return model.Book{}, result.Error
	}
	return book, nil
}

// GetBookByID returns books by ID
func GetBookByID(db *gorm.DB, id int) (model.Book, error) {
	var book model.Book

	result := db.Where("id = ?", id).First(&book)
	if result.Error != nil {
		return model.Book{}, result.Error
	}
	return book, nil
}

// DeleteBookByID deletes book by ID
func DeleteBookByID(db *gorm.DB, id int) error {
	result := db.Where("id = ?", id).Delete(&model.Book{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// UpdateBookByID updates book by ID
func UpdateBookByID(db *gorm.DB, book *model.Book) (*model.Book, error) {
	if err := db.Save(&book).Error; err != nil {
		return book, err
	}
	return book, nil
}

// UpdateBookStockCount updates book stock count
func UpdateBookStockCount(db *gorm.DB, book *model.Book, newStockCount uint) error {
	fmt.Println("newStockCount", newStockCount)
	book.StockCount = newStockCount
	result := db.Save(&book)
	fmt.Println(result)
	if result.Error != nil {
		fmt.Println("result")

		return result.Error
	}
	fmt.Println(book)
	return nil
}

// GetBooksByAuthorID returns books by author ID
func GetBooksByAuthorID(db *gorm.DB, authorID uint64) ([]model.Book, error) {
	var books []model.Book

	result := db.Preload("Author").Where("author_id = ?", authorID).Find(&books)
	if result.Error != nil {
		return nil, result.Error
	}
	return books, nil
}
