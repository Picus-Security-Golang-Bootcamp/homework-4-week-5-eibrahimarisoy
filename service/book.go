package service

import (
	"github.com/Picus-Security-Golang-Bootcamp/homework-4-week-5-eibrahimarisoy/model"
	"gorm.io/gorm"
)

// InsertSampleData inserts sample data into the database
func SaveBook(db *gorm.DB, book *model.Book) (*model.Book, error) {
	result := db.Omit("Author").Where(model.Book{Name: book.Name, StockCode: book.StockCode}).
		FirstOrCreate(&book)
	if result.Error != nil {
		return nil, result.Error
	}

	return book, nil
}

// GetAuthorWithoutAuthorInformation returns only books
// func GetAllBooksWithoutAuthorInformation() ([]model.Book, error) {
// 	var books []model.Book
// 	result := db.Find(&books)

// 	if result.Error != nil {
// 		return nil, result.Error
// 	}
// 	return books, nil
// }

// GetBooksWithAuthor returns books with author
func GetBooksWithAuthor(db *gorm.DB) ([]model.Book, error) {
	var books []model.Book

	result := db.Preload("Author").Order("id").Find(&books)
	if result.Error != nil {
		return nil, result.Error
	}
	return books, nil
}

// // FindByName returns books by name
// func FindByName(keyword string) ([]model.Book, error) {
// 	var books []model.Book

// 	result := db.Preload("Author").Where("name ILIKE ?", "%"+keyword+"%").Find(&books)
// 	if result.Error != nil {
// 		return nil, result.Error
// 	}
// 	return books, nil
// }

// GetByIDWithAuthor returns books by ID with author
func GetByIDWithAuthor(db *gorm.DB, id int) (model.Book, error) {
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

// // UpdateBookStockCountByID updates book stock count by ID
// func UpdateBookStockCountByID(id, newStockCount int) (model.Book, error) {
// 	instance, _ := GetByIDWithAuthor(id)
// 	instance.StockCount = uint(newStockCount)
// 	db.Model(&instance).Update("stock_count", newStockCount)

// 	return instance, nil
// }

// // **************EXTRA QUERIES************** //

// // UpdateBookName updates book name
// func UpdateBookName(book model.Book, newName string) (model.Book, error) {
// 	book.Name = newName
// 	db.Model(&book).Update("name", newName)

// 	return book, nil
// }

// // UpdateBookPrice updates book price
// func UpdateBookPrice(book model.Book, newPrice float64) (model.Book, error) {
// 	book.Price = newPrice
// 	db.Model(&book).Update("price", newPrice)

// 	return book, nil
// }

// // FilterBookByPriceRange filters book by price range
// func FilterBookByPriceRange(min, max float64) ([]model.Book, error) {
// 	var books []model.Book

// 	result := db.Unscoped().Where("price BETWEEN ? AND ?", min, max).Find(&books)
// 	if result.Error != nil {
// 		return nil, result.Error
// 	}
// 	return books, nil
// }

// // GetBooksWithIDs returns books by IDs
// func GetBooksWithIDs(ids []int) ([]model.Book, error) {
// 	var books []model.Book

// 	result := db.Where("id IN ?", ids).Find(&books)
// 	if result.Error != nil {
// 		return nil, result.Error
// 	}
// 	return books, nil
// }

// // FilterBookByCreatedAtRange filters book by created at range
// func FilterBookByCreatedAtRange(min, max string) ([]model.Book, error) {
// 	var books []model.Book

// 	result := db.Unscoped().Where("created_at BETWEEN ? AND ?", min, max).Find(&books)
// 	if result.Error != nil {
// 		return nil, result.Error
// 	}
// 	return books, nil
// }

// // SearchBookByNameAndStockCode searches book by name and stock code
// func SearchBookByNameOrStockCode(keyword string) ([]model.Book, error) {
// 	var books []model.Book

// 	result := db.Unscoped().Where("name ILIKE ? ", "%"+keyword+"%").Or("stock_code ILIKE ?", "%"+keyword+"%").Find(&books)
// 	if result.Error != nil {
// 		return nil, result.Error
// 	}
// 	return books, nil
// }

// // GetAllBooksOrderByPriceAsc returns all books order by price asc
// func GetAllBooksOrderByPriceAsc() ([]model.Book, error) {
// 	var books []model.Book

// 	result := db.Unscoped().Order("price asc").Find(&books)
// 	if result.Error != nil {
// 		return nil, result.Error
// 	}
// 	return books, nil
// }

// // GetFirstTenBooks returns first ten books
// func GetFirstTenBooks() ([]model.Book, error) {
// 	var books []model.Book

// 	result := db.Unscoped().Limit(10).Find(&books)
// 	if result.Error != nil {
// 		return nil, result.Error
// 	}
// 	return books, nil
// }

// // GetBooksCount returns books count
// func GetCount() (int64, error) {
// 	var count int64

// 	result := db.Model(&model.Book{}).Count(&count)
// 	if result.Error != nil {
// 		return 0, result.Error
// 	}
// 	return count, nil
// }

// // GetTotalStockValue returns total stock value
// func GetTotalStockValue() (int64, error) {
// 	var count int64

// 	err := db.Model(&model.Book{}).Select("sum(stock_count)").Row().Scan(&count)
// 	if err != nil {
// 		return 0, err
// 	}
// 	return count, nil
// }

// // GetAvgPrice returns average price
// func GetAvgPrice() (float64, error) {
// 	var avgPrice float64

// 	err := db.Model(&model.Book{}).Select("avg(price)").Row().Scan(&avgPrice)
// 	if err != nil {
// 		return 0, err
// 	}
// 	return avgPrice, nil
// }
