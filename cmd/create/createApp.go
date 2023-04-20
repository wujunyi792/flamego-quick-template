package create

import (
	"bytes"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/wujunyi792/flamego-quick-template/pkg/colorful"
	"github.com/wujunyi792/flamego-quick-template/pkg/fs"
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
				println(colorful.Red(err.Error()))
				os.Exit(1)
			}
			println(colorful.Green("App " + appName + " generate success under " + dir))
		},
	}
)

func init() {
	StartCmd.PersistentFlags().StringVarP(&appName, "name", "n", "", "create a new app with provided name")
	StartCmd.PersistentFlags().StringVarP(&dir, "path", "p", "internal/app", "new file will generate under provided path")
	StartCmd.PersistentFlags().BoolVarP(&force, "force", "f", false, "Force generate the app")
}

func load() error {
	if appName == "" {
		return errors.New("app name should not be empty, use -n")
	}

	root := path.Join(dir, appName)
	router := path.Join(dir, appName, "router")
	handler := path.Join(dir, appName, "handler")
	dto := path.Join(dir, appName, "dto")
	service := path.Join(dir, appName, "service")
	trigger := path.Join(dir, "appInitialize")

	_ = fs.IsNotExistMkDir(root)
	_ = fs.IsNotExistMkDir(router)
	_ = fs.IsNotExistMkDir(handler)
	_ = fs.IsNotExistMkDir(dto)
	_ = fs.IsNotExistMkDir(service)
	_ = fs.IsNotExistMkDir(trigger)

	m := map[string]string{}
	m["appNameExport"] = strings.ToUpper(appName[:1]) + appName[1:]
	m["appName"] = strings.ToLower(appName[:1]) + appName[1:]

	root += "/init.go"
	router += "/" + m["appName"] + ".go"
	handler += "/" + m["appName"] + ".go"
	dto += "/" + m["appName"] + ".go"
	trigger += "/" + m["appName"] + ".go"

	if !force && (fs.FileExist(router) || fs.FileExist(handler) || fs.FileExist(dto) || fs.FileExist(trigger) || fs.FileExist(root)) {
		return errors.New("target file already exist, use -f flag to cover")
	}

	if rt, err := template.ParseFiles("template/router.template"); err != nil {
		return err
	} else {
		var b bytes.Buffer
		_ = rt.Execute(&b, m)
		fs.FileCreate(b, router)
	}

	if rt, err := template.ParseFiles("template/handler.template"); err != nil {
		return err
	} else {
		var b bytes.Buffer
		_ = rt.Execute(&b, m)
		fs.FileCreate(b, handler)
	}

	if rt, err := template.ParseFiles("template/dto.template"); err != nil {
		return err
	} else {
		var b bytes.Buffer
		_ = rt.Execute(&b, m)
		fs.FileCreate(b, dto)
	}

	if rt, err := template.ParseFiles("template/trigger.template"); err != nil {
		return err
	} else {
		var b bytes.Buffer
		_ = rt.Execute(&b, m)
		fs.FileCreate(b, trigger)
	}

	if rt, err := template.ParseFiles("template/init.template"); err != nil {
		return err
	} else {
		var b bytes.Buffer
		_ = rt.Execute(&b, m)
		fs.FileCreate(b, root)
	}

	return nil
}
