# nasha

> A self-hosted, open-source multi-protocol storage aggregation gateway.
> Aggregate local disks, WebDAV shares, SMB/CIFS (Samba), S3-compatible object
> stores, and SFTP servers behind a single, beautiful Web UI.

---

## Architecture overview

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
                 │   │ local webdav    │  │
                 │   │ smb   s3  sftp  │  │
                 │   └─────────────────┘  │
                 └─────────────────────────┘
```

## Directory structure

```
nasha/
├── cmd/
│   └── server/          # Binary entry point (main.go)
├── internal/
│   ├── driver/          # StorageDriver interface + per-protocol stubs
│   │   ├── driver.go    #   interface & FileInfo type
│   │   ├── local.go     #   local filesystem
│   │   ├── webdav.go    #   WebDAV client
│   │   ├── smb.go       #   SMB/CIFS (go-smb2)
│   │   ├── s3.go        #   S3-compatible (aws-sdk-go-v2)
│   │   └── sftp.go      #   SFTP (pkg/sftp)
│   ├── vfs/             # Virtual filesystem – mount-point routing
│   ├── api/             # Fiber routes & handlers
│   ├── auth/            # JWT issuance & Fiber middleware
│   └── cache/           # Thumbnail & directory-listing cache
├── web/                 # React + Vite + TailwindCSS frontend
├── config.yaml          # Example configuration
├── Dockerfile           # Multi-stage Go + static-asset image
└── docker-compose.yml   # One-command local deployment
```

## Key dependencies

| Layer | Package | Purpose |
|---|---|---|
| HTTP server | `github.com/gofiber/fiber/v2` | High-performance HTTP |
| Auth | `github.com/golang-jwt/jwt/v5` | JWT tokens |
| ORM | `gorm.io/gorm` + `gorm.io/driver/sqlite` | Metadata storage |
| SMB | `github.com/hirochachacha/go-smb2` | SMB/CIFS client |
| SFTP | `github.com/pkg/sftp` + `golang.org/x/crypto/ssh` | SFTP client |
| S3 | `github.com/aws/aws-sdk-go-v2/service/s3` | S3-compatible stores |
| WebDAV | `golang.org/x/net/webdav` | WebDAV |
| UI | React + Vite + TailwindCSS | Frontend |
| Router | `react-router-dom` | Client-side routing |

## Development

### Prerequisites

- Go 1.25+
- Node.js 20+
- (Optional) Docker & Docker Compose

### Backend

```bash
# Run the API server (hot-reload with Air recommended)
go run ./cmd/server
```

### Frontend

```bash
cd web
npm install
npm run dev        # starts Vite dev server on :5173 with API proxy to :8080
```

### Build

```bash
# Build the React bundle into web/dist/
cd web && npm run build && cd ..

# Build the Go binary (embeds the web/dist/ directory)
go build -o nasha ./cmd/server
```

## Docker

```bash
# Build and start everything with Docker Compose
docker compose up --build
```

The service will be available at `http://localhost:8080`.

## Configuration

See [`config.yaml`](config.yaml) for a fully-commented example.

## Development roadmap

| Phase | Description |
|---|---|
| **1 – MVP** | Local filesystem driver + REST API + basic WebUI (browse, upload, download) |
| **2** | WebDAV / SFTP drivers + JWT auth + Docker packaging |
| **3** | S3 / SMB drivers + thumbnail cache + share links + inline preview |
| **4** | Multi-user RBAC + background task queue (copy/unzip) + mobile layout |

## License

MIT
