package controller

import (
	"net/http"

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

func (controller *DataController) GetData(c *gin.Context) {
	date := c.Query("date")
	data, err := controller.dataService.GetDataByDate(c.Request.Context(), date)
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

func (controller *DataController) GetDataCity(c *gin.Context) {
	date := c.Query("date")
	city := c.Query("city")

	data, err := controller.dataService.GetDataCity(c.Request.Context(), date, city)
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

func (controller *DataController) GetFilterName(c *gin.Context) {

	data, err := controller.dataService.GetFilterName(c.Request.Context())
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
