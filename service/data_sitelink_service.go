package service

import (
	"context"
	"database/sql"

	"github.com/fajarherdian22/apiweb/helper"
	"github.com/fajarherdian22/apiweb/repository"
)

type DataServiceImpl struct {
	q  *repository.Queries
	DB *sql.DB
}

func NewService(q *repository.Queries, db *sql.DB) *DataServiceImpl {
	return &DataServiceImpl{
		q:  q,
		DB: db,
	}
}

func (service *DataServiceImpl) GetDataCity(ctx context.Context, Date, city string) ([]repository.DemarcationSiteLink, error) {
	payload := repository.GetCityDataParams{
		Date: Date,
		City: city,
	}
	data, err := service.q.GetCityData(ctx, payload)
	helper.IsError(err)

	return data, nil
}
