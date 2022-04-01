package service

import (
	"strconv"

	"github.com/Picus-Security-Golang-Bootcamp/homework-4-week-5-eibrahimarisoy/model"
	"gorm.io/gorm"
)

// Search adds where to search keywords
func Search(search string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if search != "" {
			db = db.Where("name ILIKE ?", "%"+search+"%")
			db = db.Or("stock_code ILIKE ?", "%"+search+"%")
		}
		return db
	}
}

// Paginate adds limit and offset to query
func Paginate(args model.Args) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		offset, _ := strconv.Atoi(args.Offset)
		limit, _ := strconv.Atoi(args.Limit)

		switch {
		case limit > 100:
			limit = 100
		case limit <= 0:
			limit = 10
		}

		// offset := (offset - 1) * limit
		return db.Offset(offset).Limit(limit)
	}
}
