package api

import (
	"github.com/Nolions/jobQueuePoc/config"
	"github.com/Nolions/jobQueuePoc/internal/jobQueue"
	"github.com/gin-gonic/gin"
	"github.com/redpkg/formula/db"
	"github.com/redpkg/formula/log"
	"github.com/redpkg/formula/redis"
	"net/http"
)

type Handler struct {
	CachePrefix string
}

func newHandler(cachePrefix string, appConf config.App, configApi config.Api, redisConf redis.Config, dbConf db.Config) Handler {
	d, err := db.New(dbConf)
	if appConf.Mode == gin.DebugMode || appConf.Mode == gin.TestMode {
		d.ShowSQL()
	}

	if err != nil {
		log.Fatal().Msgf("Failed to new db: [%v]", err)
	}

	return Handler{}
}

func (handler Handler) router(e *gin.Engine) {
	e.GET("/healthz", handler.healthz)
	e.POST("/task", handler.addTask)
}

func (handler Handler) healthz(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

func (handler Handler) addTask(ctx *gin.Context) {
	go jobQueue.AddJobTask()

	ctx.JSON(http.StatusNoContent, nil)
}
