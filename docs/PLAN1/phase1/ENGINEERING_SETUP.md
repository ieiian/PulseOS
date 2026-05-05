# Phase 1 工程搭建

## 完成内容

- Android `Kotlin + Compose` 工程骨架
- iOS `Swift + SwiftUI` 源码骨架
- `XcodeGen` 项目描述文件
- Go 后端单体骨架
- PostgreSQL / Redis / MinIO 本地开发环境
- 配置文件与 env 模板
- Design System 占位结构
- GitHub Actions 基础 CI

## 目录说明

- `android/`: Android Studio 工程
- `ios/`: iOS SwiftUI 源码与 `project.yml`
- `backend/`: Go 服务代码、配置和脚本
- `docs/config/`: 共享配置说明

## 本地开发

### 后端

```bash
cd backend
go run ./cmd/server
```

### Android

用 Android Studio 打开 `android/`。

### iOS

如果本机安装了 `xcodegen`，在 `ios/` 下执行：

```bash
xcodegen generate
```

然后用 Xcode 打开生成的工程。

## 环境依赖

- Go 1.22
- JDK 17
- Android SDK 35
- Xcode 15+
- Docker / Docker Compose

## 本地基础服务

```bash
docker compose -f backend/deploy/docker-compose.yml up -d
```

启动后本地可用地址：

- PostgreSQL: `localhost:5432`
- Redis: `localhost:6379`
- MinIO API: `http://localhost:9000`
- MinIO Console: `http://localhost:9001`

默认开发凭据：

- PostgreSQL: `pulseos / pulseos_password / pulseos`
- MinIO: `pulseos / pulseos_secret`
- 默认 bucket: `pulseos-dev`

如果你使用 `docker + colima`：

```bash
colima start
docker compose -f backend/deploy/docker-compose.yml up -d
```

`minio-init` 会在 MinIO 健康后自动创建默认 bucket，无需手动执行。
