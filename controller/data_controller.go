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

	var req struct {
		Date string `json:"date" binding:"required"`
		City string `json:"city" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, helper.ErrorResponse(err))
		return
	}

	data, err := controller.dataService.GetDataCity(c.Request.Context(), req.Date, req.City)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve data"})
		return
	}

	responseCode := 200
	responseStatus := true

	if helper.HandlingNullData(data) {
		responseCode = 400
		responseStatus = false
	}

	WebResponse := web.WebResponse{
		Code:   responseCode,
		Data:   data,
		Status: responseStatus,
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

func (controller *DataController) GetCityTopo(c *gin.Context) {

	var req struct {
		City string `json:"city" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, helper.ErrorResponse(err))
		return
	}

	data, err := controller.dataService.GetTopoCity(c.Request.Context(), req.City)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve data"})
		return
	}

	responseCode := 200
	responseStatus := true

	if len(data) == 0 {
		responseCode = 400
		responseStatus = false
	}

	WebResponse := web.WebResponse{
		Code:   responseCode,
		Data:   data,
		Status: responseStatus,
	}

	helper.HandleEncodeWriteJson(c, WebResponse)
}

func (controller *DataController) ListCityTopo(c *gin.Context) {
	data, err := controller.dataService.ListCityTopo(c)
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
