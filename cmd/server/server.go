package server

import (
	"context"
	"errors"
	"fmt"
	"github.com/flamego/cors"
	"github.com/flamego/flamego"
	"github.com/spf13/cobra"
	"github.com/wujunyi792/flamego-quick-template/config"
	"github.com/wujunyi792/flamego-quick-template/internal/app/appInitialize"
	"github.com/wujunyi792/flamego-quick-template/internal/core/cache"
	"github.com/wujunyi792/flamego-quick-template/internal/core/database"
	"github.com/wujunyi792/flamego-quick-template/internal/core/kernel"
	"github.com/wujunyi792/flamego-quick-template/internal/core/logx"
	"github.com/wujunyi792/flamego-quick-template/internal/middleware/gw"
	"github.com/wujunyi792/flamego-quick-template/pkg/colorful"
	"github.com/wujunyi792/flamego-quick-template/pkg/ip"
	"go.uber.org/zap/zapcore"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	configYml string
	engine    *kernel.Engine
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
	log = logx.NameSpace("cmd.server")
)

func init() {
	StartCmd.PersistentFlags().StringVarP(&configYml, "config", "c", "config/config.yaml", "Start server with provided configuration file")
}

func setUp() {
	// 初始化全局 ctx
	ctx, cancel := context.WithCancel(context.Background())

	// 初始化资源管理器
	engine = &kernel.Engine{Ctx: ctx, Cancel: cancel}

	// 顺序不能变 logger依赖config logger后面的同时依赖logger和config 否则crash
	config.LoadConfig(configYml)
	if config.GetConfig().MODE == "" || config.GetConfig().MODE == "debug" {
		logx.Init(zapcore.DebugLevel)
	} else {
		logx.Init(zapcore.InfoLevel)
	}

	flamego.SetEnv(flamego.EnvType(config.GetConfig().MODE))
	engine.Fg = flamego.New()
	engine.Fg.Use(flamego.Recovery(), gw.RequestLog(), flamego.Renderer(), cors.CORS(cors.Options{
		AllowCredentials: true,
		Methods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodDelete,
			http.MethodOptions,
		},
	}))

	database.InitDB()
	cache.InitCache()
}

func load() {
	modules := appInitialize.GetApps()
	for _, module := range modules {
		_err := module.PreInit(engine)
		if _err != nil {
			log.Errorw("failed to pre init app", _err)
			os.Exit(1)
		}
	}
	for _, module := range modules {
		_err := module.Init(engine)
		if _err != nil {
			log.Errorw("failed to init app", _err)
			os.Exit(1)
		}
	}
	for _, module := range modules {
		_err := module.PostInit(engine)
		if _err != nil {
			log.Errorw("failed to post init app", _err)
			os.Exit(1)
		}
	}
	for _, module := range modules {
		_err := module.Load(engine)
		if _err != nil {
			log.Errorw("failed to load app", _err)
			os.Exit(1)
		}
	}
	for _, module := range modules {
		_err := module.Start(engine)
		if _err != nil {
			log.Errorw("failed to start app", _err)
			os.Exit(1)
		}
	}
}

func run() {
	srv := &http.Server{
		Addr:    config.GetConfig().Host + ":" + config.GetConfig().Port,
		Handler: engine.Fg,
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

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	println(colorful.Blue("Shutting down server..."))

	ctx, cancel := context.WithTimeout(engine.Ctx, 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		println(colorful.Yellow("Server forced to shutdown: " + err.Error()))
	}

	println(colorful.Green("Server exiting Correctly"))
}
