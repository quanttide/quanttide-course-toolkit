# Changelog

## [go/v0.1.0-alpha.3] - 2026-08-26

### Added

- API 路由契约：Route* 模板常量（含通配段名，服务端注册用）+ 具体路径构造函数（消费方拼接用），同文件互锁
- 契约测试：contract_test 以根 tests/ 的 Schema/Fixture 验证四实体模型

## [go/v0.1.0-alpha.2] - 2026-08-26

### Fixed

- Go 子模块拉取修复：Go 要求子目录模块的标签含完整子路径（`packages/go/vX.Y.Z`），单一 `go/` 前缀标签无法被 `go get` 解析
- 落实双标签制：主标签保持仓库统一惯例 `<scope>/vX.Y.Z` 并挂 GitHub Release；同 commit 机械追加 `packages/go/X.Y.Z` 工具链别名标签（无 Release）。此后所有 Go 包发布遵循此机制
- 代码内容与 alpha.1 相同

## [go/v0.1.0-alpha.1] - 2026-08-26

### Added

- 初始化 Go 包：对齐 qtcloud-course 最新领域模型——Course / Lesson / Scene / Criterion 及 Step / Choice，JSON 标签与 API 一一对应
