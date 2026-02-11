# AdminBox

[![Go](https://img.shields.io/badge/Go-1.18+-blue.svg)](https://golang.org/)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](https://opensource.org/licenses/MIT)

一个用于快速构建后台管理系统的、模块化的 Go 语言开发套件。

## 核心理念

AdminBox 的设计哲学是“**后台优先，按需组装**”。它不是一个大而全的笨重框架，而是一个由多个独立的“功能积木”组成的工具箱。你可以根据你的项目需求，只选择你需要的部分，构建一个轻量、高效、定制化的后台服务。

同时，它也提供了一个开箱即用的前端界面，只需一行 `import` 即可自动集成，实现真正的“开箱即用”。

## 主要特性

- **极致模块化**: 项目的 `pkg` 目录下的每个包（如 `log`, `database`, `config`）都是一个独立的功能模块，可以被任何项目单独引用。
- **可插拔的UI**: 只需在你的 `main.go` 中匿名导入 `front` 包，一个功能完善的后台管理界面就会被自动嵌入并提供服务。
- **清晰的架构**: 采用 Go 社区推崇的 `pkg`/`internal` 结构，清晰地分离了公共库与内部业务逻辑。
- **功能完备**: 内置了配置加载、结构化日志、数据库(GORM)、缓存、中间件、文件上传、CRUD引擎等后台开发的常用组件。
- **易于扩展**: 清晰的职责划分和接口设计，方便你进行二次开发或替换任意模块。

## 快速上手

以下是一个构建最小化后台服务的示例。

1.  **创建你的项目目录和 `go.mod` 文件:**

    ```bash
    mkdir my-admin
    cd my-admin
    go mod init my-admin
    go get github.com/mangk/adminBox
    ```

2.  **创建 `config.yaml` 配置文件:**

    ```yaml
    server:
      port: ":8080"
    log:
      prefix: my-admin
      output:
        - console
    ```

3.  **创建 `main.go` 文件:**

    ```go
    package main

    import (
    	"fmt"
    	"github.com/mangk/adminBox/internal/bootstrap"
    	"github.com/mangk/adminBox/pkg/config"
    	"github.com/mangk/adminBox/pkg/log"

    	// 匿名导入，即可拥有开箱即用的前端UI
    	_ "github.com/mangk/adminBox/front"
    )

    func main() {
    	// 1. 加载配置
    	if err := config.Load("config.yaml", nil); err != nil {
    		panic(fmt.Sprintf("Failed to load config: %v", err))
    	}

    	// 2. 初始化日志
    	log.Init()
    	log.Info("Server starting...")

    	// 3. 初始化路由引擎
    	router := bootstrap.InitRouter()

    	// 4. 在这里注册你自己的业务路由...
    	// router.GET("/ping", func(c *gin.Context) {
    	// 	c.JSON(200, gin.H{"message": "pong"})
    	// })

    	// 5. 启动服务
    	port := config.ServerCfg().Port
    	if port == "" {
    		port = ":8080" // 默认端口
    	}
    	log.Infof("Server listening on port %s", port)
    	if err := router.Run(port); err != nil {
    		log.Errorf("Server failed to start: %v", err)
    	}
    }
    ```

4.  **运行你的服务:**

    ```bash
    go run main.go
    ```
    现在，访问 `http://localhost:8080`，你将看到 AdminBox 的登录界面。

## 项目结构解析

-   `pkg/` - **功能积木 (公共库)**
    这是 AdminBox 的核心价值所在。此目录下的每个子包都是一个可被外部项目独立使用的功能模块。
    -   `config`: 基于 `viper` 的配置加载模块。
    -   `log`: 基于 `zap` 的结构化日志模块。
    -   `database`: 基于 `gorm` 的数据库连接模块。
    -   `cache`: 缓存模块（支持内存和Redis）。
    -   `middleware`: Gin 中间件集合。
    -   `request` / `response`: 标准化的请求获取和响应输出工具。
    -   `upload`: 支持本地、腾讯云COS、阿里云OSS、七牛云的文件上传模块。
    -   `crud`: 一个通用的CRUD后端引擎。
-   `internal/` - **内部逻辑**
    这部分是 AdminBox 示例应用的私有业务逻辑，**不应该**被你的项目直接导入。
-   `front/` - **前端界面**
    包含了编译好的前端静态资源。通过 `embed` 和 `init()` 实现自动注册。
-   `examples/` - **示例代码**
    提供了一个功能更完整的 `simple-server` 示例，展示了如何组装所有模块来构建一个全功能的后台。

## 运行完整示例

要运行功能更完整的官方示例：

```bash
# 1. 进入示例目录
cd examples/simple-server

# 2. 整理依赖
go mod tidy

# 3. 运行
go run main.go
```

## 贡献

欢迎任何形式的贡献！如果你有好的想法或发现了Bug，请随时提交 Pull Request 或 Issue。

## 许可证

本项目基于 [MIT License](https://opensource.org/licenses/MIT) 开源。
