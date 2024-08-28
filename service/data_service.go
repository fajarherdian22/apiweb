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

func (service *DataServiceImpl) GetDataCity(ctx context.Context, Date, city string) ([]repository.PlmnTraffic, error) {
	arg := repository.GetPlmnDataCityParams{
		Date:        Date,
		RanSiteCity: city,
	}

	data, err := service.q.GetPlmnDataCity(ctx, arg)
	helper.IsError(err)

	return data, nil
}

func (service *DataServiceImpl) GetDataByDate(ctx context.Context, Date string) ([]repository.PlmnTraffic, error) {
	data, err := service.q.GetPlmnData(ctx, Date)
	helper.IsError(err)
	return data, nil
}

func (service *DataServiceImpl) GetFilterName(ctx context.Context) ([]repository.GetFilterNameRow, error) {
	data, err := service.q.GetFilterName(ctx)
	helper.IsError(err)
	return data, nil
}

func (service *DataServiceImpl) GetAllPlmn(ctx context.Context) ([]repository.PlmnTraffic, error) {
	data, err := service.q.GetAllPlmn(ctx)
	helper.IsError(err)
	return data, nil
}
