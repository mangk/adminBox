# AdminBox

[![Go Version](https://img.shields.io/badge/Go-1.23+-00ADD8.svg)](https://go.dev/)
[![Go Reference](https://pkg.go.dev/badge/github.com/mangk/adminBox.svg)](https://pkg.go.dev/github.com/mangk/adminBox)
[![License](https://img.shields.io/badge/License-AGPLv3-green.svg)](https://www.gnu.org/licenses/agpl-3.0.html)
[![Gin Framework](https://img.shields.io/badge/Framework-Gin-blue.svg)](https://gin-gonic.com/)

AdminBox 是一个模块化、可组合的 Go 后台开发套件。每个模块可独立导入、独立使用，按需组装。

## 设计理念

**"后台优先，按需组装"** — AdminBox 不是一个大而全的框架，而是一组可独立使用的功能积木：

- **`httpServer`** 是底座 — Gin + Cobra + 服务管理，可以独立使用
- **`admin`** + **`front`** 是可选插件 — 通过 `init()` 注册路由，插上即用
- **`log` / `config` / `db` / `cache` / `jwt` / `casbin` / `middleware` / `upload` / `util`** 是独立工具包 — 可各自 import 到任何项目

## 快速上手

### 最小化服务（无需数据库）

```go
package main

import (
    "github.com/gin-gonic/gin"
    "github.com/mangk/adminBox/httpServer"
    "github.com/mangk/adminBox/response"
)

func main() {
    httpServer.SetRouter(func(root *gin.Engine) {
        root.GET("/ping", func(ctx *gin.Context) {
            response.OkWithMsg(ctx, "pong")
        })
    })
    httpServer.Execute("my-app", "My AdminBox App")
}
```

```yaml
# config.yaml
server:
  name: my-app
  port: 8910
log:
  output:
    - console
```

```bash
go run main.go run
```

### 完整后台管理（需 MySQL + Redis）

```go
package main

import (
    "github.com/gin-gonic/gin"
    _ "github.com/mangk/adminBox/admin"        // 自动注册后台API
    _ "github.com/mangk/adminBox/admin/front"   // 嵌入前端UI
    "github.com/mangk/adminBox/httpServer"
)

func main() {
    httpServer.SetRouter(func(root *gin.Engine) {
        root.GET("/hello", func(ctx *gin.Context) {
            ctx.String(200, "Hello AdminBox!")
        })
    })
    httpServer.Execute("my-admin", "My Admin Panel")
}
```

导入 `admin` 和 `admin/front` 后，系统会自动注册所有管理后台接口和前端页面。

### 独立使用工具包

```go
import "github.com/mangk/adminBox/log"
import "github.com/mangk/adminBox/config"
import "github.com/mangk/adminBox/jwt"
```

每个工具包都可单独引入，不强制依赖其他模块。

## 示例项目

| 示例 | 说明 | 运行 |
|---|---|---|
| [minimal](./examples/minimal) | 仅 httpServer + 自定义路由 | `go run ./examples/minimal/ run` |
| [with-admin](./examples/with-admin) | 完整后台 + 前端管理界面 | `go run ./examples/with-admin/ run` |

## 模块一览

| 包 | 说明 | 独立使用 |
|---|---|---|
| `pkg/httpServer` | HTTP 服务底座 (Gin + Cobra + daemon) | ✅ |
| `pkg/admin` | 后台管理业务模块 | 需 httpServer |
| `pkg/admin/front` | 嵌入式 Vue 前端界面 | 需 httpServer |
| `pkg/config` | 配置管理 (Viper, YAML, 热加载) | ✅ |
| `pkg/log` | 结构化日志 (Zap, 文件轮转) | ✅ |
| `pkg/db` | 数据库 ORM (GORM, 多驱动) | ✅ |
| `pkg/cache` | 缓存（本地 + Redis） | ✅ |
| `pkg/jwt` | JWT 令牌创建与解析 | ✅ |
| `pkg/casbin` | 权限引擎 (Casbin + GORM) | ✅ |
| `pkg/middleware` | Gin 中间件集合 | ✅ |
| `pkg/response` | 统一响应格式 | ✅ |
| `pkg/request` | 请求参数解析 | ✅ |
| `pkg/upload` | 文件上传 (local/OSS/COS/S3/Qiniu) | ✅ |
| `pkg/util` | 通用工具函数 | ✅ |

## 项目结构

```
adminBox/
├── examples/             # 示例代码
│   ├── minimal/          # 最小化服务示例
│   └── with-admin/       # 完整后台管理示例
├── pkg/                  # 核心功能模块（全部对外公开）
│   ├── admin/            # 后台管理
│   │   ├── handler/      #   HTTP 处理器
│   │   ├── model/        #   GORM 模型
│   │   └── crud/         #   通用 CRUD 引擎
│   ├── httpServer/       # HTTP 服务启动器
│   ├── config/           # 配置管理
│   ├── log/              # 结构化日志
│   ├── db/               # 数据库 ORM
│   ├── cache/            # 缓存
│   ├── jwt/              # JWT 令牌
│   ├── casbin/           # 权限引擎
│   ├── middleware/       # Gin 中间件
│   ├── request/          # 请求参数
│   ├── response/         # 响应格式
│   ├── upload/           # 文件上传
│   └── util/             # 工具函数
├── Makefile
├── CHANGELOG.md
├── go.mod
└── README.md
```

## 主要特性

- **高度模块化**: 每个 `pkg/*` 下的子包都可单独 import，零耦合
- **可插拔前端**: 嵌入式 Vue 前端，一行 import 即可激活管理界面
- **配置热加载**: 基于 Viper，支持 YAML，运行时修改即时生效
- **多数据库**: MySQL / PostgreSQL / SQLServer，多实例连接池
- **认证授权**: JWT + Casbin RBAC，中间件即插即用
- **文件上传**: 统一接口，支持本地 / 阿里云 OSS / 腾讯云 COS / AWS S3 / 七牛云
- **CRUD 引擎**: 根据 GORM 模型自动生成前后端 CRUD 接口
- **服务管理**: 内置 daemon 模式，支持 install/start/stop/uninstall

## 运行方式

```bash
# 开发运行
go run main.go run

# 守护进程模式（需 root）
go run main.go daemon install
go run main.go daemon start
go run main.go daemon stop
```

## 许可证

本项目基于 [AGPLv3](https://www.gnu.org/licenses/agpl-3.0.html) 开源。
