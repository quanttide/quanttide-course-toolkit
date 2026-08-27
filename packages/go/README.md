# quanttide-course Go SDK

量潮课程研发 Go SDK：课程领域模型与数据校验，字段语义对齐 Dart / Python SDK。

```go
import course "github.com/quanttide/quanttide-course-toolkit/packages/go/pkg"

var lec course.Lecture
json.Unmarshal(data, &lec)
err := lec.Validate()
```

## 许可

本项目采用 [Apache License 2.0](./LICENSE)。
