# quanttide-course Go SDK

量潮课程研发 Go SDK：封装课程领域共通模型与公共代码，供各应用与工具复用，保证系统一致性。

## 领域模型

以 [qtcloud-course](https://github.com/quanttide/qtcloud-course) 服务端最新实现为准，JSON 标签与 API 字段一一对应：

| 模型 | 说明 |
|------|------|
| `Course` | 课程（目录展示单元） |
| `Lesson` | 课时（教学内容最小组织单元） |
| `Scene` | 场景 / 视频片段（互动课时基本单元） |
| `Criterion` | 验收标准（跨领域对接的原子单元，学习云直连其 ID） |

```go
import course "github.com/quanttide/quanttide-course-toolkit/packages/go/pkg"

lessons, err := course.ParseLessons(data)
```

## 许可

本项目采用 [Apache License 2.0](./LICENSE)。
