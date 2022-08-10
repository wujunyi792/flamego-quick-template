package server

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/wujunyi792/gin-template-new/config"
	"github.com/wujunyi792/gin-template-new/internal/middleware"
	v1 "github.com/wujunyi792/gin-template-new/internal/router/v1"
	"github.com/wujunyi792/gin-template-new/pkg/colorful"
	"github.com/wujunyi792/gin-template-new/pkg/ip"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	configYml string
	E         *gin.Engine
	StartCmd  = &cobra.Command{
		Use:     "server",
		Short:   "Set Application config info",
		Example: "main server -c config/settings.yml",
		PreRun: func(cmd *cobra.Command, args []string) {
			println("Loading config...")
			setUp()
			println("Loading config complete")
			println("Loading Api...")
			load()
			println("Loading Api complete")
		},
		Run: func(cmd *cobra.Command, args []string) {
			println("Starting Server...")
			run()
		},
	}
)

func init() {
	StartCmd.PersistentFlags().StringVarP(&configYml, "config", "c", "config/config.yaml", "Start server with provided configuration file")
}

func setUp() {
	config.LoadConfig(configYml)
}

func load() {
	gin.SetMode(config.GetConfig().MODE)
	E = gin.New()
	E.Use(middleware.GinRequestLog, gin.Recovery(), middleware.Cors(E))

	v1.MainRouter(E)

}

func run() {
	srv := &http.Server{
		Addr:    "0.0.0.0:" + config.GetConfig().Port,
		Handler: E,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			println(colorful.Red("Got Server Err: " + err.Error()))
		}
	}()

	println(colorful.Green("Server run at:"))
	println(fmt.Sprintf("-  Local:   http://localhost:%s", config.GetConfig().Port))
	for _, host := range ip.GetLocalHost() {
		println(fmt.Sprintf("-  Network: http://%s:%s", host, config.GetConfig().Port))
	}

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	println(colorful.Blue("Shutting down server..."))

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		println(colorful.Yellow("Server forced to shutdown: " + err.Error()))
	}

	println(colorful.Green("Server exiting Correctly"))
}
