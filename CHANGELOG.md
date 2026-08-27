# Changelog

## 未发布

- 契约测试重写为 Lesson 体系（tests/）：Schema 与 Fixture 以服务端 API camelCase 标签为准，Go 包 contract_test 落地验证；python/dart 旧契约测试挂起待迁移
- 新增 CONTRIBUTING.md：Go 包双标签规范与原因说明

## [go/v0.1.0-alpha.1] - 2026-08-26

- 初始化项目结构
- 新增 Go SDK 包（packages/go）：封装 Course / Lesson / Scene / Criterion 共通领域模型
