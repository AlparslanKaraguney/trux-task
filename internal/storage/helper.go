package storage

import (
	"github.com/AlparslanKaraguney/trux-task/internal/entities"
	"gorm.io/gorm"
)

func paginate(limit, offset int, query *gorm.DB) (*gorm.DB, *entities.Pagination) {

	totalRows := new(int64)
	query.Count(totalRows)
	pagination := &entities.Pagination{
		Limit:     limit,
		TotalRows: *totalRows,
		Offset:    offset,
	}
	query = query.Offset(offset).Limit(limit)

	return query, pagination

}
