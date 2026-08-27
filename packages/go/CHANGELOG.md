# Changelog

## [packages/go/v0.1.0-alpha.2] - 2026-08-26

### Fixed

- 发布标签改为 `packages/go/` 前缀：Go 子目录模块要求标签含完整子路径，`go/v0.1.0-alpha.1` 无法被 `go get` 解析（alpha.1 为无效发布）
- 代码内容与 alpha.1 相同

## [go/v0.1.0-alpha.1] - 2026-08-26

### Added

- 初始化 Go 包：对齐 qtcloud-course 最新领域模型——Course / Lesson / Scene / Criterion 及 Step / Choice，JSON 标签与 API 一一对应
