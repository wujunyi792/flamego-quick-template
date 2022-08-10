package main

import (
	"github.com/spf13/cobra"
	"github.com/wujunyi792/gin-template-new/cmd/config"
	"github.com/wujunyi792/gin-template-new/cmd/server"
)

func main() {
	var rootCmd = &cobra.Command{}
	rootCmd.AddCommand(server.StartCmd)
	rootCmd.AddCommand(config.StartCmd)
	rootCmd.Execute()
	//sign := make(chan os.Signal, 1)
	//signal.Notify(sign, syscall.SIGINT, syscall.SIGTERM)
	//<-sign
}
