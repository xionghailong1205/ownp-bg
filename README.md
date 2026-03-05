# ownp-bg
ownp 在线小说平台后端

## Gin + OpenAPI + PostgreSQL 初始化

### 1. 环境变量

复制示例配置：

```bash
cp .env.example .env
```

关键配置项：

- `APP_ENV`：`development` 或 `production`
- `DATABASE_URL_DEV`：开发环境 PostgreSQL 连接串
- `DATABASE_URL_PROD`：生产环境 PostgreSQL 连接串
- `SERVER_ADDR`：服务监听地址（默认 `:8080`）

### 2. 启动服务

```bash
go run ./cmd/server
```

### 3. 可用接口

- `GET /api/v1/health`：返回应用健康状态和数据库时间
- `GET /openapi.json`：返回 OpenAPI 3.0 文档
