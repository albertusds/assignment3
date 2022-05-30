package controllers

import (
	"assignment3/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type WebController struct {
	webService *services.WebService
}

func NewWebController(service *services.WebService) *WebController {
	return &WebController{
		webService: service,
	}
}

func (w *WebController) GetStatus(c *gin.Context) {

	response, err := w.webService.GetStatus()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "/index.html", err)
	}

	c.HTML(http.StatusOK, "index.html", response)
}
