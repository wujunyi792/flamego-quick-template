# Flamego-Template

全新的后端快速开发模板

## 有什么
- flameGo
- 定时任务
- Gorm + 多数据库支持
- memory cache, redis 多缓存支持
- 腾讯云短信
- 阿里云对象存储
- 邮件服务
- Viper 配置文件
- 日志服务
- 容器化部署方案 dockerfile + docker compose
- ip 检测
- jwt鉴权套件
- 随机数生成工具
- 彩色控制台输出
- websocket

## How to use
- [ ] Globally replace the package name with your own repository
- [ ] Edit config/vars GlobalConfig. **It is recommended to make changes on the existing basis. Try not to change the existing structure, if you change, you need to change part of the code synchronously**
- [ ] Exec ` go run .\main.go config` and run, `config.yaml` will generate under `config/`
- [ ] Complete the config
- [ ] If you deploy with docker engine, edit `docker-copmose.yml`, Especially port mappings and service names


## 新增服务
```shell
go run main.go create -n [newAppName]
```
然后会在`internal/app`目录下生成`newAppName`目录，里面包含了`handle`、`service`、`router`等目录，新增的路由会被自动注册

## 启动服务
```shell
go run main.go server
```
