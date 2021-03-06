package statistics

import (
	"github.com/puti-projects/puti/internal/admin/api"
	"github.com/puti-projects/puti/internal/admin/service"
	"github.com/puti-projects/puti/internal/pkg/errno"

	"github.com/gin-gonic/gin"
)

// Dashboard get some basic statistics
func Dashboard(c *gin.Context) {
	svc := service.New(c.Request.Context())
	data, err := svc.GetDashboardStatisticsData()
	if err != nil {
		api.SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	api.SendResponse(c, nil, data)
}
