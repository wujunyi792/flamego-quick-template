package server

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/wujunyi792/gin-template-new/config"
	_ "github.com/wujunyi792/gin-template-new/internal/controller/example"
	_ "github.com/wujunyi792/gin-template-new/internal/corn"
	"github.com/wujunyi792/gin-template-new/internal/logger"
	"github.com/wujunyi792/gin-template-new/internal/middleware"
	"github.com/wujunyi792/gin-template-new/internal/redis"
	v1 "github.com/wujunyi792/gin-template-new/internal/router/v1"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var E *gin.Engine

func init() {
	logger.Info.Println("start init gin")
	gin.SetMode(config.GetConfig().MODE)
	E = gin.New()
	E.Use(middleware.GinRequestLog, gin.Recovery(), middleware.Cors(E))
}

func Run() {
	redis.GetRedis()
	v1.MainRouter(E)
	srv := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: E,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Error.Println("Got Server Err: ", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Info.Println("Start shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Error.Fatalln("Server forced to shutdown:", err)
	}

	logger.Info.Println("Server exiting Correctly")
}
