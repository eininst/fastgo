```text
project
├── api
│   ├── helloword
│   |   └── v1
│   │       └── helloword_api.go
├── build
|   ├── start.sh
├── cmd
|   ├── helloword
│   |   └── main.go (启动程序)
├── configs (配置文件)
|   ├── config.go
|   ├── helloword.yml
│   │
├── configs (配置文件)
|   ├── config.go
|   ├── helloword.yml
├── deployments (部署配置)
|   ├── api.yml
|   ├── nginx.yml
├── docs (项目文档)
|   ├── helloword_swagger.json
├── internal
│   ├── conf
│   |   └── conf.go (内部资源初始化)
│   ├── data
│   |   └── db.go (数据库配置)
│   |   └── rdb.go (Redis配置)
│   ├── model
│   |   └── model.go (POJO对象)
│   ├── service
│   |   └── user
│   │       └── user_service.go
│   │       └── type.go
│   ├── types
│   |   └── types.go (公用对象)
├── nginx
│   ├── conf
│   |   └── helloword.conf (应用nginx配置文件)
│   ├── ssl
│   |   └── xxxxx.pem
│   |   └── xxxxx.key
|   ├── Dockerfile
|   ├── nginx.conf (全局nginx配置)
├── pkg (公共模块)
│   ├── burst
│   ├── di
|   ├── grace
|   ├── redoc
|   ├── util
├── scripts (Makefile执行的sh文件)
│   ├── swarm.sh
│   ├── swagger.sh
├── test (测试)
│   ├── xxxx.go
├── website (网站页面)
│   ├── xxxx.html
│   ├── xxxx.html
├── Dockerfile
├── go.mod
├── Makefile
├── README.md
```