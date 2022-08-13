package cmd

import (
	"github.com/spf13/cobra"
	"github.com/wujunyi792/gin-template-new/cmd/config"
	"github.com/wujunyi792/gin-template-new/cmd/create"
	"github.com/wujunyi792/gin-template-new/cmd/server"
	"os"
)

var rootCmd = &cobra.Command{
	Use:          "app",
	Short:        "app",
	SilenceUsage: true,
	Long:         `app`,
}

func init() {
	rootCmd.AddCommand(server.StartCmd)
	rootCmd.AddCommand(config.StartCmd)
	rootCmd.AddCommand(create.StartCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}
