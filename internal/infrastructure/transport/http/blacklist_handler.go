package http

import (
	"net/http"

	"github.com/GeovanniGomes/blacklist/internal/application/interfaces"
	"github.com/GeovanniGomes/blacklist/internal/application/service"
	"github.com/GeovanniGomes/blacklist/internal/infrastructure/depedence_injector"
	"github.com/gin-gonic/gin"
)

type BlackListHanhler struct {
	gin_engine       *gin.Engine
	serviceBlacklist *service.BlacklistService
}

func NewBlackListHanhler(gin_engine *gin.Engine, container_injection *depedence_injector.ContainerInjection) *BlackListHanhler {
	serviceBlacklist, err := container_injection.GetBlacklistService()
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
