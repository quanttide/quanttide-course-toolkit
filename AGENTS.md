# AGENTS.md - quanttide-course-toolkit

## 定位

课程领域的**共通代码工具箱**：封装领域模型、API 客户端等公共代码，供各应用与工具复用，目标是提高系统一致性——同一份领域事实，所有语言一个口径。

- **领域模型的单一事实源是 [qtcloud-course](https://github.com/quanttide/qtcloud-course) 服务端**（`src/provider/internal/domain`）。各语言包的模型与 JSON 标签必须与其一一对应，不许自创字段或语义。
- 领域模型当前形态：课程树三级 `Course → Lesson → Scene`，加验收标准 `Criterion`（学习云的 completion.criterion_id 直连其 ID，无本地副本/映射表）。旧 Lecture 模型已废弃。
- 跨语言契约测试放 `tests/`，共享标本用 `tests/fixtures/`；新增模型时同步补标本与各语言用例。

## 项目结构

```
quanttide-course-toolkit/
├── AGENTS.md
├── CHANGELOG.md
├── LICENSE
├── README.md
├── ROADMAP.md
├── packages/
│   ├── go/       # Go SDK（已对齐 Lesson 体系）
│   ├── python/   # Python SDK（待迁移：Lecture → Lesson）
│   └── dart/     # Dart SDK（待迁移：Lecture → Lesson）
└── tests/        # 跨语言契约测试与共享标本
```

每个语言包自带 README / CHANGELOG / AGENTS / LICENSE；Go 包遵循 delib-toolkit 惯例（module = `<repo>/packages/go`、pkg/ 布局、零依赖、就地测试）。

## 提交约定

参见 [数据结构工具箱 AGENTS.md](https://github.com/quanttide/quanttide-data-toolkit/blob/main/AGENTS.md) 的提交规范。本仓库作为子模块挂载于 quanttide-course 的 `packages/default`：先在本仓库提交推送，再回父仓库更新引用指针。

## 发布流程

参见 [数据结构工具箱 AGENTS.md](https://github.com/quanttide/quanttide-data-toolkit/blob/main/AGENTS.md) 的发布流程。版本变更记录进各语言包自己的 CHANGELOG，仓库级 CHANGELOG 只记结构性与新增包级别的事件。
