# CONTRIBUTING.md - quanttide-course-toolkit

仓库约定与常见任务的贡献指南。定位、结构与领域模型事实源见 [AGENTS.md](AGENTS.md)。

## 发布标签规范（Go 包）

本仓库所有发布标签遵循统一格式 `<scope>/vX.Y.Z`，如 `go/v0.1.0-alpha.2`、`dart/v0.1.1`。**主标签一律挂 GitHub Release**，作为人可读的版本事实源。

### Go 包的双标签制

Go 子目录模块（拥有独立 go.mod 的 `packages/go`）有社区硬性要求：标签必须包含完整子路径前缀，即 `packages/go/X.Y.Z`——否则 `go get` 报 `invalid version: unknown revision`，无法拉取。

为同时满足仓库统一惯例与 Go 工具链，Go 包发布时打**两条指向同一 commit 的标签**：

| 标签 | 作用 | 是否挂 Release |
|------|------|---------------|
| `go/vX.Y.Z` | 主标签：人类惯例、CHANGELOG 版本头、GitHub Release | 是 |
| `packages/go/X.Y.Z` | 工具链别名：仅供 `go get` 解析 | 否 |

示例（`packages/go` 发版）：

```bash
VERSION="v0.1.0-alpha.3"
git tag "go/$VERSION"                       # 主标签 + GH Release
git tag "packages/go/$VERSION"              # 工具链别名
git push origin "go/$VERSION" "packages/go/$VERSION"
gh release create "go/$VERSION" --prerelease ...
```

注意事项：

- 别名标签是发布机制的机器侧产物：不设 Release、不出现在 CHANGELOG 版本头，也不要手工删除或移动；
- 漏打别名标签时，包虽然"已发布"但实际不可被任何消费者获取（alpha.1 即此教训）；
- 其余语言包（python/dart）不受此规则影响，只打主标签。

## 提交规范

遵循 Conventional Commits（`feat:` / `fix:` / `docs:` / `chore:` 等）；破坏性变更标 `!` 并在 body 说明迁移方式。

## 子模块协作

本仓库作为子模块挂载于 quanttide-course 的 `packages/default`：

1. 在本仓库完成修改并提交推送；
2. 回父仓库 `git add packages/default && git commit` 更新引用指针并推送。
