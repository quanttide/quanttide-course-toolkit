# 契约测试

跨语言数据模型与 API 一致性验证。**单一事实源是根 `tests/` 下的 Schema 与 Fixture**，各语言在自身包内编写契约测试并引用它们。

## 与 Go 的对齐现状

| 包 | 状态 | 说明 |
|---|------|------|
| `go/` | ✅ 已对齐 | 模型 = Lesson 体系；字段标签 camelCase，与服务端 API 一一对应 |
| `python/` | ⏳ 待迁移 | 仍为旧 Lecture 模型，契约测试挂起（skip） |
| `dart/` | ⏳ 待迁移 | 同上 |

## 设计

- **字段命名：camelCase**——即 qtcloud-course 服务端 API 的 JSON 标签。各语言包不再做 case 转换层：Dart / JS 直接可用，Python 的 snake_case 由其包内边界自行处理（迁移时落实）。
- **实体范围**：Course → Lesson → Scene 三级 + Criterion 验收标准（学习云 completion.criterion_id 直连 Criterion.id）。
- **API 路由**也是契约的一部分：Go 包 `pkg/api.go` 以常量持有路由模板，`api_test.go` 用独立标本防漂移。

```
tests/
  README.md
  schemas/    # JSON Schema（draft-07）：course / lesson / scene / criterion
  fixtures/   # 同源课程树标本：course / lesson / scene / criterion
packages/
  go/pkg/contract_test.go    # 当前唯一已落地的契约验证
```

## 测试内容

1. **必填字段存在性** — fixture 覆盖 Schema 的 `required`；
2. **Fixture 反序列化** — 共享标本能被正确解析且关键字段值符合预期；
3. **Round-trip** — 反序列化→再序列化→反序列化，值不变；
4. 完整 JSON Schema 校验待 Go 侧引入校验依赖后启用（当前轻量实现见 `contract_test.go` 头注）。

## 运行

```bash
# Go（唯一生效方）
cd packages/go && go test ./... -count=1
```

## 工作流

1. **新增字段**：先改对应 Schema + fixture，再同步各语言模型与 API 标签；
2. **其他语言完成 Lesson 迁移时**：移除各自契约测试的 skip 标记并接入本目录。
