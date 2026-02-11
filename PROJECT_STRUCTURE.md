# AdminBox Project Structure

This document outlines the final, refactored project structure, designed for modularity and reusability as a Go framework/library.

```
/Users/mangk/Data/Code/adminBox/
├───.gitignore
├───go.mod
├───go.sum
├───README.md
├───权限页面配置.drawio
│
├───examples/                  # (新) 存放示例应用，取代原 demo 和 cmd
│   └───simple-server/
│       ├───main.go            # (新) 应用主入口，取代原根目录的 httpServer.go 等
│       └───config.yaml        # (移入) 从根目录移入，作为示例应用的配置
│
├───front/                     # 前端目录保持不变，但增加一个Go文件
│   ├───.eslintrc.cjs
│   ├───package.json
│   ├───vite.config.js
│   ├───... (其他前端源文件)
│   ├───dist/                  # (不变) 前端打包后的静态资源
│   │   └───...
│   └───register.go            # (新) 自动嵌入并注册前端UI的魔法文件
│
├───internal/                  # (新) 项目私有逻辑，不能被外部引用
│   ├───bootstrap/             # (新) 应用引导程序
│   │   └───router.go          # (新) 用于创建和获取全局Gin路由引擎
│   └───module/                # (移入) 核心业务模块，从原 admin 目录移入
│       ├───handler/
│       │   ├───api.go
│       │   ├───auth.go
│       │   └───... (department, menu, role, user)
│       └───model/
│           ├───model.go
│           ├───sysApi.go
│           └───... (sysAuth, sysCasbinRule, sysDepartment, etc.)
│
└───pkg/                       # (新) 可被外部引用的、通用的“功能积木”
    ├───cache/                 # (移入) 原 cache 目录
    │   ├───local.go
    │   └───redis.go
    ├───config/                # (移入) 原 config 目录 (Go结构体)
    │   ├───cache.go
    │   ├───config.go
    │   └───... (captcha, cors, db, etc.)
    ├───crud/                  # (移入) 从原 admin/crud 目录移入
    │   └───engine.go
    ├───database/              # (移入并重命名) 原 db 目录
    │   └───db.go
    ├───log/                   # (移入) 原 log 目录
    │   └───log.go
    ├───middleware/            # (移入) 原 http/middleware 目录
    │   ├───cors.go
    │   ├───jwtCheckByCasbin.go
    │   └───... (operation, request, response, traceLogger)
    ├───request/               # (新) 独立为 request 包，存放请求处理工具
    │   ├───crud.go
    │   └───request.go
    ├───response/              # (新) 独立为 response 包，存放响应格式化工具
    │   ├───crud.go
    │   ├───page.go            # (原 http/page.go)
    │   └───response.go
    ├───upload/                # (移入) 原 http/upload 目录
    │   ├───aliyunOss.go
    │   ├───local.go
    │   └───... (awsS3, qiniu, tencentCos, upload)
    └───util/                  # (移入) 原 util 目录
        ├───duration.go
        ├───hash.go
        └───... (md5, sha256, slices, string, typePtr)
```
