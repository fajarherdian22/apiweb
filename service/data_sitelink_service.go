package service

import (
	"context"
	"database/sql"
	"fmt"

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
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("No Data to retrieve")
		}
		return data, err
	}

	return data, nil
}

func (service *DataServiceImpl) ListCity(ctx context.Context) ([]string, error) {
	data, err := service.q.ListCity(ctx)
	helper.IsError(err)

	return data, nil
}

func (service *DataServiceImpl) GetTopoCity(ctx context.Context, city string) ([]repository.TopoLink, error) {

	data, err := service.q.GetTopoCity(ctx, city)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("No Data to retrieve")
		}
		return data, err
	}
	return data, nil
}

func (service *DataServiceImpl) ListCityTopo(ctx context.Context) ([]string, error) {
	data, err := service.q.ListTopoCity(ctx)
	helper.IsError(err)

	return data, nil
}
