# Changelog

本文件只记录主仓库级事件（结构、约定、跨包变更）；各语言包的版本变更见各包自己的 CHANGELOG。

## 未发布

- 契约测试重写为 Lesson 体系（tests/）：Schema 与 Fixture 以服务端 API camelCase 标签为准，Go 包 contract_test 落地验证；python/dart 旧契约测试挂起待迁移
- 新增 CONTRIBUTING.md：版本号以契约测试为准、Go 包双标签规范与原因说明

## 历史

- 2026-08-26：初始化项目结构；新增 Go SDK 包（packages/go）
