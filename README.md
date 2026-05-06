# nasha

> 自托管、开源的多协议存储聚合网关。
> 将本地磁盘、WebDAV、SMB/CIFS（Samba）、S3 兼容对象存储和 SFTP 服务器统一在一个 Web 界面下。

---

## 项目状态

nasha 目前处于 **Demo 阶段** — 核心架构已搭建完成，LocalDriver 可用，可浏览本地文件。更多功能逐步开发中，详见[路线图](ROADMAP.md)。

目前已完成：
- LocalDriver 全部 7 个方法 ✅
- VFS 挂载点路由（最长前缀匹配）✅
- JWT 登录 (admin/admin123) ✅
- 文件浏览（列表 + 内容读取）✅
- Docker 多阶段构建 ✅

## 架构

```
                 ┌─────────────────────────┐
 Browser / App ──►   React + Vite + TW     │  web/
                 └────────────┬────────────┘
                              │  REST API (JSON)
                 ┌────────────▼────────────┐
                 │   Fiber HTTP server     │  cmd/server/
                 │   ┌─────────────────┐  │
                 │   │  auth (JWT/RBAC)│  │  internal/auth/
                 │   └────────┬────────┘  │
                 │   ┌────────▼────────┐  │
                 │   │  Virtual FS     │  │  internal/vfs/
                 │   └────────┬────────┘  │
                 │   ┌────────▼────────┐  │
                 │   │ StorageDrivers  │  │  internal/driver/
                 │   │ Local  WebDAV   │  │
                 │   │ SMB    S3  SFTP │  │
                 │   └─────────────────┘  │
                 └─────────────────────────┘
```

## 目录结构

```
nasha/
├── cmd/server/          # 入口 (main.go)
├── internal/
│   ├── driver/          # StorageDriver 接口 + 5 种协议实现
│   ├── vfs/             # 虚拟文件系统 — 挂载点路由
│   ├── api/             # Fiber 路由和请求处理器
│   ├── auth/            # JWT 签发和中间件
│   ├── cache/           # 缓存（开发中）
│   └── db/              # GORM 数据库层（开发中）
├── web/                 # React + Vite + TailwindCSS 前端
├── config.yaml          # 配置文件
├── Dockerfile           # 多阶段构建
└── docker-compose.yml   # 一键部署
```

## 快速开始

### 前置要求

- Go 1.25+（CGO 必须启用，SQLite 需要）
- Node.js 20+

### 后端

```bash
CGO_ENABLED=1 go run ./cmd/server
```

默认监听 `:8080`，登录凭据 `admin / admin123`。

### 前端（开发模式，带热更新）

```bash
cd web
npm install
npm run dev
```

Vite 开发服务器运行在 `:5173`，自动代理 API 请求到 `:8080`。

### 构建

```bash
cd web && npm run build && cd ..
CGO_ENABLED=1 go build -o nasha ./cmd/server
```

### Docker

```bash
docker compose up --build
```

## 配置

参考 [`config.yaml`](config.yaml) 中的示例配置。

默认挂载 `./demo_data` 到根路径 `/`，包含示例文件可直接浏览。

## 技术栈

| 层 | 组件 |
|---|---|
| HTTP 服务 | `gofiber/fiber/v2` |
| 鉴权 | `golang-jwt/jwt/v5` |
| ORM | `gorm.io/gorm` + `gorm.io/driver/sqlite` |
| S3 | `aws/aws-sdk-go-v2/service/s3` |
| SFTP | `pkg/sftp` + `golang.org/x/crypto/ssh` |
| SMB | `hirochachacha/go-smb2` |
| 前端 | React 19 + Vite + TailwindCSS 4 |
| 路由 | `react-router-dom` v7 |

## 路线图

| 阶段 | 目标 |
|---|---|
| **Phase 1** | 增删改查 API + 前端操作 + 用户系统 + bcrypt 密码 |
| **Phase 2** | S3/SFTP/SMB/WebDAV 驱动实现 |
| **Phase 3** | 文件预览、搜索、分享链接、设置页 |
| **Phase 4** | 测试、RBAC、速率限制、CI/CD |
| **Phase 5** | WebDAV 服务端、TUS 分块上传、i18n、2FA、移动端 |

详见 [ROADMAP.md](ROADMAP.md)。

## License

MIT
