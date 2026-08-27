# 量潮课程研发工具箱

课程领域的共通代码工具箱：封装领域模型、API 客户端等公共代码，供各应用与工具复用，提高系统一致性。领域模型的单一事实源是 qtcloud-course 服务端（Course → Lesson → Scene + Criterion）。

## 项目结构

```
quanttide-course-toolkit/
├── AGENTS.md          # AI 开发指引
├── CHANGELOG.md       # 变更日志
├── ROADMAP.md         # 路线图
├── packages/
│   ├── go/            # Go SDK
│   ├── python/        # Python SDK（待迁移：Lecture → Lesson）
│   └── dart/          # Dart SDK（待迁移：Lecture → Lesson）
└── tests/             # 跨语言契约测试与共享标本
```
