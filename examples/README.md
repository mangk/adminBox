# Examples

AdminBox 的示例项目，展示不同粒度的模块组合方式。

## 目录

| 示例 | 说明 | 使用模块 |
|---|---|---|
| [minimal](./minimal) | 最小化服务，仅 httpServer + 自定义路由 | httpServer, response |
| [with-admin](./with-admin) | 完整后台管理，含前端界面 | httpServer, admin, admin/front, response |

## 运行方式

```bash
# 在项目根目录下运行

# 最小化服务
go run ./examples/minimal/ run

# 完整后台管理（需 MySQL + Redis）
go run ./examples/with-admin/ run
```
