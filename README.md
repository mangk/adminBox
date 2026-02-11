# AdminBox

[![Go Version](https://img.shields.io/badge/Go-1.23+-00ADD8.svg)](https://go.dev/)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](https://opensource.org/licenses/MIT)
[![Gin Framework](https://img.shields.io/badge/Framework-Gin-blue.svg)](https://gin-gonic.com/)
[![GORM](https://img.shields.io/badge/ORM-GORM-lightgrey.svg)](https://gorm.io/)

**AdminBox** 是一个模块化、可插拔的 Go 后台开发套件，旨在帮助开发者快速构建功能完备、易于扩展的后台管理系统。

## 设计理念

AdminBox 的核心设计哲学是 **“后台优先，按需组装”**。它并非一个大而全的笨重框架，而是一个由多个独立“功能积木”组成的工具箱。您可以根据项目需求，仅选择需要的部分，构建一个轻量、高效、定制化的后台服务。

同时，项目内置了一个基于 Vue 的现代化前端界面，通过 `embed` 特性实现编译时嵌入，只需一行 `import` 即可自动集成，实现真正的“开箱即用”。

## 主要特性

-   **高度模块化**: `pkg` 目录下的每个包（如 `log`, `db`, `config`）都是独立的功能模块，可以被任何项目单独引用。
-   **可插拔的前端**: 只需在 `main.go` 中分别匿名导入 `pkg/admin`（后端API）和 `pkg/admin/front`（前端UI），一个功能完善的后台管理界面就会被自动嵌入并提供服务。
-   **清晰的架构**: 遵循 Go 社区推崇的项目结构，清晰地分离了公共库与内部业务逻辑。
-   **丰富的功能集**:
    -   **配置管理**: 基于 `Viper`，支持 YAML 文件格式，可热加载。
    -   **结构化日志**: 基于 `Zap`，支持多级别、多输出（控制台、文件）。
    -   **数据库 ORM**: 基于 `GORM`，支持 `MySQL`, `PostgreSQL`, `SQLServer` 等多种数据库。
    -   **缓存**: 内置 `Redis` 客户端支持。
    -   **Web 框架**: 使用高性能的 `Gin` 作为路由和中间件核心。
    -   **认证与授权**: 内置 `JWT` 和 `Casbin` 的集成方案。
    -   **文件上传**: 支持本地存储、阿里云 OSS、腾讯云 COS、AWS S3 和七牛云。
    -   **CRUD 引擎**: 提供通用的增删改查后端引擎，简化业务开发。
-   **易于扩展**: 清晰的职责划分和接口设计，方便进行二次开发或替换任意模块。

## 快速上手

以下步骤将引导您快速启动一个最小化的 AdminBox 服务。

### 1. 环境准备

-   Go 1.23+
-   MySQL 5.7+ 或其他 GORM 支持的数据库
-   Redis

### 2. 初始化项目

```bash
# 创建项目目录
mkdir my-admin && cd my-admin

# 初始化 Go Module
go mod init my-admin

# 获取 AdminBox
go get github.com/mangk/adminBox
```

### 3. 配置文件 `config.yaml`

在项目根目录创建 `config.yaml` 文件。这是一个最小化的配置示例：

```yaml
server:
  name: my-admin
  env: debug
  port: 8910
  jwt:
    signingKey: "a-secret-key-change-it" # 请务必修改为你的密钥
    expiresTime: 7d
    bufferTime: 1d
    issuer: my-admin

db:
  default:
    driver: mysql
    path: 127.0.0.1
    port: 3306
    dbname: adminbox # 数据库名
    username: root # 数据库用户名
    password: "your-password" # 数据库密码
    config: "charset=utf8mb4&parseTime=True&loc=Local"
    logMode: 3 # 日志级别

cache:
  default:
    addr: "127.0.0.1:6379"
    password: ""
    db: 0

log:
  level: debug
  format: console
  output:
    - console
    - logs
```

### 4. 创建数据库

根据 `config.yaml` 中的配置，创建名为 `adminbox` 的数据库。

### 5. 编写 `main.go`

```go
package main

import (
	"github.com/gin-gonic/gin"
	// 匿名导入 AdminBox 后端模块，自动注册后台API
	_ "github.com/mangk/adminBox/pkg/admin"
	// 匿名导入 AdminBox 前端模块，自动注册UI界面
	_ "github.com/mangk/adminBox/pkg/admin/front"
	"github.com/mangk/adminBox/pkg/httpServer"
)

func main() {
	// (可选) 在这里注册你自己的业务路由
	httpServer.SetRouter(func(root *gin.Engine) {
		// example: root.GET("/ping", func(c *gin.Context) { c.String(200, "pong") })
	})

	// 启动服务
	httpServer.Execute("my-admin", "My First AdminBox Server")
}
```

### 6. 运行

```bash
go run main.go
```

服务启动后，访问 `http://127.0.0.1:8910` 即可看到 AdminBox 的登录界面。系统会自动初始化数据库表结构。

## 项目结构

```
/
├── pkg/            # 核心功能模块 (可被外部引用的公共库)
│   ├── admin/      # 后台管理核心逻辑 (包含前端资源)
│   ├── cache/      # 缓存 (Redis)
│   ├── config/     # 配置加载 (Viper)
│   ├── db/         # 数据库 (GORM)
│   ├── httpServer/ # HTTP 服务启动器 (Cobra + Gin)
│   ├── log/        # 日志 (Zap)
│   ├── middleware/ # Gin 中间件
│   ├── upload/     # 文件上传模块
│   └── util/       # 通用工具函数
├── example/        # 示例代码
├── go.mod          # Go 模块文件
└── README.md       # 项目说明
```

## 贡献

我们欢迎任何形式的贡献！如果您有好的想法或发现了 Bug，请随时提交 Pull Request 或 Issue。

## 许可证

本项目基于 [AGPLv3](https://www.gnu.org/licenses/agpl-3.0.en.html) 开源。
