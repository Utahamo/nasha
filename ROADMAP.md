# nasha — 项目发展路线图

## Context

nasha 是一个自托管的多协议存储聚合网关。目前 Demo 版本已实现：
- LocalDriver 全部 7 个方法
- VFS 挂载点路由
- JWT 登录 (admin/admin123)
- 文件浏览 (列表 + 读取)
- 前端登录页 + 文件浏览器
- Docker 构建

但仍有大量功能缺失：上传/删除/重命名/建目录的 API 和前端操作、数据库模型、其余 4 个 driver (S3/SFTP/SMB/WebDAV) 全是 panic 桩、无测试、无搜索/预览/分享等。

---

## 功能总览表 (Feature Matrix)

| 功能 | 阶段 | 优先级 | 状态 | 说明 |
|---|---|---|---|---|
| **Config 加载** | — | P0 | ✅ Done | `internal/config/config.go` |
| **LocalDriver** | — | P0 | ✅ Done | 全部 7 方法 |
| **VFS 挂载路由** | — | P0 | ✅ Done | 最长前缀匹配 |
| **JWT 中间件** | — | P0 | ✅ Done | `internal/auth/auth.go` |
| **GET /api/v1/fs/*** | — | P0 | ✅ Done | 列表/读取文件 |
| **登录端点** | — | P0 | ✅ Done | 硬编码 admin/admin123 |
| **SPA 静态服务** | — | P0 | ✅ Done | `web/dist/` + 回退 |
| **CORS** | — | P0 | ✅ Done | 全放通 |
| **优雅关闭** | — | P0 | ✅ Done | SIGINT/SIGTERM |
| **Docker 构建** | — | P0 | ✅ Done | 多阶段构建 |
| **前端登录页** | — | P0 | ✅ Done | 已对接 API |
| **前端文件浏览器** | — | P0 | ✅ Done | 只读浏览 |
| | | | | |
| **POST 上传 API** | 1 | P0 | 📋 Planned | `fsGroup.Post("/*")` multipart |
| **POST 建目录 API** | 1 | P0 | 📋 Planned | `fsGroup.Post("/mkdir")` |
| **DELETE 删除 API** | 1 | P0 | 📋 Planned | `fsGroup.Delete("/*")` |
| **PATCH 重命名 API** | 1 | P0 | 📋 Planned | `fsGroup.Patch("/*")` |
| **DB 模型 (User/Mount/ShareLink)** | 1 | P0 | 📋 Planned | GORM structs + AutoMigrate |
| **密码哈希 (bcrypt)** | 1 | P0 | 📋 Planned | 替代硬编码 admin/admin123 |
| **可配置 JWT Secret** | 1 | P0 | 📋 Planned | 从 config.yaml 读取 |
| **前端上传按钮** | 1 | P0 | 📋 Planned | FormData POST |
| **前端删除/重命名/新建文件夹** | 1 | P0 | 📋 Planned | 对话框 + API 调用 |
| **API client 工具函数** | 1 | P1 | 📋 Planned | 集中 fetch + auth 头 |
| **AuthGuard 路由守卫** | 1 | P1 | 📋 Planned | 集中式路由保护 |
| **Emoji → lucide-react 图标** | 1 | P1 | 📋 Planned | lucide-react 已装 |
| **404 路由** | 1 | P1 | 📋 Planned | 前端 NotFound 页 |
| **页面标题修正** | 1 | P1 | 📋 Planned | "web" → "nasha" |
| | | | | |
| **S3Driver 实现** | 2 | P0 | 📋 Planned | aws-sdk-go-v2 |
| **SFTPDriver 实现** | 2 | P0 | 📋 Planned | pkg/sftp + crypto/ssh |
| **SMBDriver 实现** | 2 | P0 | 📋 Planned | go-smb2 |
| **WebDAVDriver 实现** | 2 | P0 | 📋 Planned | net/http + PROPFIND |
| **main.go 多类型挂载** | 2 | P0 | 📋 Planned | switch 分发 |
| **驱动连接生命周期** | 2 | P1 | 📋 Planned | 连接/重连/关闭 |
| **驱动配置验证** | 2 | P1 | 📋 Planned | Validate() 方法 |
| | | | | |
| **文件预览 (文本/图片/视频/PDF)** | 3 | P1 | 📋 Planned | 模态框内预览 |
| **搜索 API + UI** | 3 | P1 | 📋 Planned | 递归搜索 + 防抖输入 |
| **分享链接** | 3 | P1 | 📋 Planned | 生成/访问/撤销 |
| **目录缓存 (TTL LRU)** | 3 | P1 | 📋 Planned | `internal/cache/` |
| **设置页 - 挂载管理** | 3 | P1 | 📋 Planned | CRUD 挂载点 |
| **设置页 - 用户管理** | 3 | P1 | 📋 Planned | CRUD 用户 |
| **排序/骨架屏/Toast** | 3 | P2 | 📋 Planned | 前端体验优化 |
| | | | | |
| **单元测试 (driver)** | 4 | P0 | 📋 Planned | 表驱动测试 |
| **集成测试 (API)** | 4 | P0 | 📋 Planned | httptest |
| **速率限制** | 4 | P1 | 📋 Planned | Fiber limiter |
| **上传大小限制** | 4 | P1 | 📋 Planned | 可配置 |
| **RBAC (admin/editor/viewer)** | 4 | P1 | 📋 Planned | 角色中间件 |
| **审计日志** | 4 | P2 | 📋 Planned | 结构化请求日志 |
| **CI/CD (GitHub Actions)** | 4 | P1 | 📋 Planned | lint + test + build |
| **后台任务队列** | 4 | P2 | 📋 Planned | 异步复制/压缩 |
| **缩略图生成** | 4 | P2 | 📋 Planned | 图片 + ffmpeg 视频 |
| | | | | |
| **WebDAV 服务端** | 5 | P1 | 📋 Planned | 将 nasha 暴露为 WebDAV |
| **分块上传 (TUS)** | 5 | P2 | 📋 Planned | 断点续传 |
| **i18n 国际化** | 5 | P2 | 📋 Planned | react-i18next |
| **移动端适配** | 5 | P2 | 📋 Planned | 响应式布局 |
| **2FA 双因素认证** | 5 | P2 | 📋 Planned | TOTP |
| **多数据库支持 (PG/MySQL)** | 5 | P2 | 📋 Planned | GORM 切换 |

---

## 详细阶段规划

### Phase 1 — 核心操作 (完成 MVP)

**目标**: 从只读 Demo 变为可增删改查的文件管理器

**涉及文件**:
- [internal/api/api.go](internal/api/api.go) — 新增 POST/DELETE/PATCH 路由
- [internal/api/api.go](internal/api/api.go) — 登录改为查询 DB + bcrypt 验证
- [internal/db/models.go](internal/db/models.go) (新建) — User / Mount / ShareLink GORM 模型
- [internal/db/db.go](internal/db/db.go) — 添加 AutoMigrate + 默认 admin 种子
- [internal/auth/auth.go](internal/auth/auth.go) — JWT secret 从 config 加载
- [internal/auth/auth.go](internal/auth/auth.go) — TokenTTL 从 config 解析
- [internal/auth/hash.go](internal/auth/hash.go) (新建) — bcrypt 哈希/验证
- [cmd/server/main.go](cmd/server/main.go) — 创建并传递 DB 实例
- [web/src/lib/api.ts](web/src/lib/api.ts) (新建) — API 客户端封装
- [web/src/components/AuthGuard.tsx](web/src/components/AuthGuard.tsx) (新建) — 路由守卫
- [web/src/pages/FileBrowser.tsx](web/src/pages/FileBrowser.tsx) — 添加上传/删除/重命名/新建文件夹
- [web/src/pages/NotFound.tsx](web/src/pages/NotFound.tsx) (新建) — 404 页面
- [web/src/App.tsx](web/src/App.tsx) — 添加 AuthGuard 和 404 路由
- [web/index.html](web/index.html) — 标题改为 nasha

**关键设计决策**: API 构造函数签名改为 `New(v *vfs.VFS, db *db.DB, cfg *config.Config)`，通过 closure 注入依赖。

**估计工作量**: 5-7 天

---

### Phase 2 — 多协议支持

**目标**: 所有 4 个 driver 从 panic stub 变为可用实现

**涉及文件**:
- [internal/driver/s3.go](internal/driver/s3.go) — S3 全部 7 方法
- [internal/driver/sftp.go](internal/driver/sftp.go) — SFTP 全部 7 方法
- [internal/driver/smb.go](internal/driver/smb.go) — SMB 全部 7 方法
- [internal/driver/webdav.go](internal/driver/webdav.go) — WebDAV 全部 7 方法
- [cmd/server/main.go](cmd/server/main.go) — switch 分发 mount 类型

**各 Driver 关键点**:
| Driver | 难度 | 特殊注意 |
|---|---|---|
| WebDAV | M | 手动解析 PROPFIND XML，MOVE 需 Destination header |
| SFTP | M | SSH 连接生命周期 + 密码/密钥认证 |
| SMB | M | TCP:445 + SMB2 协商 + session/mount |
| S3 | L | 无目录概念，Rename=copy+delete，List 需分页 |

**实现顺序建议**: WebDAV → SFTP → SMB → S3

**估计工作量**: 8-12 天

---

### Phase 3 — 体验与功能完善

**目标**: 预览、搜索、分享、设置、缓存

**涉及文件**:
- [internal/api/api.go](internal/api/api.go) — 预览/搜索/分享端点
- [internal/cache/cache.go](internal/cache/cache.go) — LRU + TTL 缓存实现
- [internal/vfs/vfs.go](internal/vfs/vfs.go) — 缓存集成 + hot-reload
- [web/src/components/FilePreview.tsx](web/src/components/FilePreview.tsx) (新建) — 文件预览模态框
- [web/src/components/SearchBar.tsx](web/src/components/SearchBar.tsx) (新建) — 搜索组件
- [web/src/pages/Settings.tsx](web/src/pages/Settings.tsx) — 完整设置页替换占位符

**估计工作量**: 10-15 天

---

### Phase 4 — 生产就绪

**目标**: 测试、安全、CI/CD、运营基础设施

**涉及文件**:
- 各 `*_test.go` — 单元测试 + 集成测试
- [internal/auth/auth.go](internal/auth/auth.go) — RBAC 中间件
- [internal/auth/audit.go](internal/auth/audit.go) (新建) — 审计日志
- `.github/workflows/ci.yml` (新建) — CI 流水线
- [internal/task/queue.go](internal/task/queue.go) (新建) — 后台任务队列

**测试优先级**: LocalDriver → Auth → VFS → API → 各远程 Driver

**估计工作量**: 12-18 天

---

### Phase 5 — 高级功能

**目标**: 企业级功能、国际化、移动端

- WebDAV 服务端端点 (`golang.org/x/net/webdav`)
- TUS 分块上传协议
- i18n (react-i18next)
- 移动端响应式布局
- 2FA (TOTP)

**估计工作量**: 15-20 天

---

## 架构关注点

| 关注点 | P0 | P1 | P2 |
|---|---|---|---|
| **安全** | JWT secret 从配置读 | 密码加密 bcrypt | CORS 收紧 | 2FA |
| **安全** | 路径穿越防护 (已有) | 文件大小限制 | Docker 非 root 运行 | 防病毒扫描 |
| **测试** | — | LocalDriver 测试 | API 集成测试 | 远程 Driver 测试 |
| **基础设施** | — | CI/CD | Docker Compose dev | 预提交钩子 |
| **可观测性** | — | — | 审计日志 | Prometheus 指标 |
| **性能** | — | 目录缓存 | S3 多部分上传 | CDN/302 重定向 |

---

## 验证方式

每个阶段完成后验证：
1. `go build ./...` + `cd web && npm run build` 编译通过
2. 启动服务 `go run ./cmd/server`
3. 测试 API: 登录 → 文件操作 → 所有端点返回正确状态码
4. 前端: 登录 → 浏览 → 操作 → 退出 完整链路
