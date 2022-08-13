package create

import (
	"bytes"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/wujunyi792/gin-template-new/pkg/fs"
	"os"
	"path"
	"strings"
	"text/template"
)

var (
	appName  string
	dir      string
	force    bool
	StartCmd = &cobra.Command{
		Use:     "create",
		Short:   "create a new app",
		Example: "app create -n users",
		Run: func(cmd *cobra.Command, args []string) {
			err := load()
			if err != nil {
				println(err.Error())
				os.Exit(1)
			}
		},
	}
)

func init() {
	StartCmd.PersistentFlags().StringVarP(&appName, "name", "n", "example", "create a new app with provided name")
	StartCmd.PersistentFlags().StringVarP(&dir, "path", "p", "internal/app", "new file will generate under provided path")
	StartCmd.PersistentFlags().BoolVarP(&force, "force", "f", false, "Force generate the app")
}

func load() error {
	if appName == "" {
		return errors.New("app name should not be empty")
	}

	router := path.Join(dir, appName, "router")
	handle := path.Join(dir, appName, "handle")
	dto := path.Join(dir, appName, "dto")
	service := path.Join(dir, appName, "service")

	_ = fs.IsNotExistMkDir(router)
	_ = fs.IsNotExistMkDir(handle)
	_ = fs.IsNotExistMkDir(dto)
	_ = fs.IsNotExistMkDir(service)

	router += "/init.go"
	handle += "/example.go"
	dto += "/example.go"

	if !force && !(fs.FileExist(router) && fs.FileExist(handle) && fs.FileExist(dto)) {
		return errors.New("target file already exist, use -f flag to cover")
	}

	m := map[string]string{}
	m["appNameExport"] = strings.ToUpper(appName[:1]) + appName[1:]
	m["appName"] = strings.ToLower(appName[:1]) + appName[1:]

	if rt, err := template.ParseFiles("template/router.template"); err != nil {
		return err
	} else {
		var b bytes.Buffer
		err = rt.Execute(&b, m)
		fs.FileCreate(b, router)
	}

	if rt, err := template.ParseFiles("template/handle.template"); err != nil {
		return err
	} else {
		var b bytes.Buffer
		err = rt.Execute(&b, m)
		fs.FileCreate(b, handle)
	}

	if rt, err := template.ParseFiles("template/dto.template"); err != nil {
		return err
	} else {
		var b bytes.Buffer
		err = rt.Execute(&b, m)
		fs.FileCreate(b, dto)
	}

	return nil
}
