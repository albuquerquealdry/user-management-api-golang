package controller

import (
	"net/http"
	"user-management/src/config"

	"github.com/gin-gonic/gin"
)

type HealthController struct{}

func NewHealthController() *HealthController {
	return &HealthController{}
}

func (h *HealthController) Liveness(ctx *gin.Context) {
	ctx.Status(http.StatusOK)
}

func (h *HealthController) Readiness(ctx *gin.Context) {
	if !config.IsDatabaseReady() {
		ctx.JSON(http.StatusServiceUnavailable, gin.H{
			"status": "database not ready",
		})
		return
	}
	ctx.Status(http.StatusOK)
}
