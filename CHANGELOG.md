# Changelog

## [Unreleased]

### Added
- doc.go files for all pkg subpackages (pkg.go.dev documentation support)
- Makefile with build, test, lint, tidy, vet, and release targets
- CHANGELOG.md for version tracking

### Fixed
- pkg/admin/init.go: corrected package name from "api" to "admin" to match directory structure
- pkg/upload/aliyunOss.go: fix memory leak — `make([]string, file.Size)` → `make([]string, 0)`
- pkg/middleware/jwtCheckByCasbin.go: remove reverse dependency on `pkg/admin/model`
- pkg/admin/model/sysJwt.go: move JWT logic to new `pkg/jwt` package
- pkg/admin/model/sysCasbinRule.go: move enforcer to new `pkg/casbin` package

### Added
- doc.go files for all pkg subpackages (pkg.go.dev documentation support)
- Makefile with build, test, lint, tidy, vet, and release targets
- CHANGELOG.md for version tracking
