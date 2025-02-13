package http

import (
	"net/http"
	"time"

	"github.com/GeovanniGomes/blacklist/internal/application/interfaces"
	"github.com/GeovanniGomes/blacklist/internal/application/service"
	"github.com/GeovanniGomes/blacklist/internal/infrastructure/depedence_injector"
	"github.com/gin-gonic/gin"
)

type BlackListHanhler struct {
	gin_engine       *gin.Engine
	serviceBlacklist *service.BlacklistService
}

func NewBlackListHanhler(gin_engine *gin.Engine, container depedence_injector.ContainerInjection) *BlackListHanhler {
	serviceBlacklist, err := container.GetBlacklistService()
	if err != nil {
		panic(err)
	}
	return &BlackListHanhler{gin_engine: gin_engine, serviceBlacklist: serviceBlacklist}
}

func (h *BlackListHanhler) BlacklistRoutes() {
	blacklistGroup := h.gin_engine.Group("api/v1/blacklist")
	{
		blacklistGroup.POST("/", h.addToBlacklist)
		blacklistGroup.GET("/check", h.checkBlacklist)
		blacklistGroup.DELETE("/remove", h.removeBlacklist)
		blacklistGroup.POST("/report", h.generateReportBlacklist)
	}
}

func (h *BlackListHanhler) addToBlacklist(c *gin.Context) {
	var entry interfaces.BlacklistInput

	if err := c.ShouldBindJSON(&entry); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	if err := h.serviceBlacklist.AddBlacklist(entry); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Added to blacklist"})
}

func (h *BlackListHanhler) checkBlacklist(c *gin.Context) {
	var entry interfaces.BlacklistInputCheck
	if err := c.ShouldBindJSON(&entry); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	resultOutput, err := h.serviceBlacklist.CheckBlacklist(entry)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": resultOutput})
}

func (h *BlackListHanhler) removeBlacklist(c *gin.Context) {
	var entry interfaces.BlacklistInputRemove
	if err := c.ShouldBindJSON(&entry); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	if err := h.serviceBlacklist.RemoveBlacklist(entry); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, gin.H{"message": ""})
}

func (h *BlackListHanhler) generateReportBlacklist(c *gin.Context) {
	var entry interfaces.BlacklistInputReport
	if err := c.ShouldBindJSON(&entry); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	startDate := entry.StartDate.ToTime()
	endDate := entry.EndDate.ToTime()

	startDate = time.Date(startDate.Year(), startDate.Month(), startDate.Day(), 0, 0, 0, 0, startDate.Location())
	endDate = time.Date(endDate.Year(), endDate.Month(), endDate.Day(), 23, 59, 59, 0, startDate.Location())

	if startDate.After(endDate) {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": "Start date is longer than the end date"})
		return
	}
	if endDate.Before(startDate) {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": "End date is longer than the start date"})
		return
	}

	h.serviceBlacklist.SendGenerateReport(startDate, endDate)

	c.JSON(http.StatusAccepted, gin.H{"message": ""})
}
