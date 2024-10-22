package controller

import (
	"net/http"
	"sort"

	"github.com/fajarherdian22/apiweb/helper"
	"github.com/fajarherdian22/apiweb/model/web"
	"github.com/fajarherdian22/apiweb/service"
	"github.com/gin-gonic/gin"
)

type DataController struct {
	dataService *service.DataServiceImpl
}

func NewDataController(dataService *service.DataServiceImpl) *DataController {
	return &DataController{dataService: dataService}
}

func (controller *DataController) GetCityLink(c *gin.Context) {

	type CityLinkRequest struct {
		Date string `json:"date" binding:"required"`
		City string `json:"city" binding:"required"`
	}

	var req CityLinkRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, helper.ErrorResponse(err))
		return
	}

	data, err := controller.dataService.GetDataCity(c.Request.Context(), req.Date, req.City)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve data"})
		return
	}
	WebResponse := web.WebResponse{
		Code:   200,
		Data:   data,
		Status: true,
	}

	helper.HandleEncodeWriteJson(c, WebResponse)
}

func (controller *DataController) ListCity(c *gin.Context) {
	data, err := controller.dataService.ListCity(c)
	if err != nil {
		helper.ErrorResponse(err)
	}
	sort.Strings(data)
	WebResponse := web.WebResponse{
		Code:   200,
		Data:   data,
		Status: true,
	}

	helper.HandleEncodeWriteJson(c, WebResponse)
}
