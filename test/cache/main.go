package main

import (
	"fmt"
	"github.com/wujunyi792/gin-template-new/config"
	"github.com/wujunyi792/gin-template-new/internal/cache"
	"github.com/wujunyi792/gin-template-new/pkg/logx"
)

func main() {
	config.LoadConfig("config/config.yaml")
	logx.InitLogger()
	cache.InitCache()
	//for i := 'A'; i < 'z'; i++ {
	//	cache.GetCache("*").Set(string(i), string(i), 2*time.Second)
	//}
	for i := 'A'; i < 'z'; i++ {
		value, exist := cache.GetCache("3").GetString(string(i))
		fmt.Println(value, exist)
	}
}
