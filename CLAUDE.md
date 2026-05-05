# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

PulseOS is a multi-platform health tracking app (Diet, Activity, Meditation, Sleep, Home dashboard). Documentation, UI strings, and code comments are primarily in Chinese. The repo contains three independent build systems: a Go backend, an Android app (Kotlin/Compose), and an iOS app (Swift/SwiftUI).

## Common Commands

### Backend (Go)
```bash
cd backend
docker compose -f deploy/docker-compose.yml up -d   # start PostgreSQL, Redis, MinIO
bash scripts/dev.sh                                   # run server (go run ./cmd/server)
go test ./...                                         # run all tests
go test ./tests/diet_test.go ./tests/...              # run a single test file
gofmt -l .                                            # check formatting (enforced in CI)
```

### Android
Open `android/` in Android Studio. AGP 8.5.2, Kotlin 1.9.24, Compose BOM 2024.09.00, compileSdk 35, minSdk 26, Java 17.

### iOS
```bash
cd ios
xcodegen generate   # generate Xcode project from project.yml
```
Deployment target iOS 17.0, Swift 5.10, no SPM/CocoaPods dependencies.

## Architecture

### Backend (Clean / Layered)
```
cmd/server/main.go → internal/app/app.go (wiring) → handler → service → repository
```
- **`internal/domain/`** — Pure domain models, one subpackage per module (diet, activity, meditation, sleep, scoring)
- **`internal/handler/`** — HTTP handlers registering routes on `http.ServeMux`. Routes: `/api/v1/{module}/{action}`
- **`internal/service/`** — Business logic per domain. Wired together via constructor injection in `app.go`
- **`internal/repository/postgres/`** — Currently **in-memory stubs** (sync.RWMutex + slices), not connected to PostgreSQL despite the package name
- **`internal/ruleengine/`** — Health rule evaluation (diet flags, calorie thresholds, etc.)
- **`internal/ai/`** — AI provider abstraction, currently using a MockProvider
- **`internal/middleware/`** — `RequireAuth` (passthrough) and `WithAccessLog`
- **`internal/bootstrap/`** — Config loading from YAML with a hand-rolled parser

**No external Go dependencies** — the backend uses only the standard library (no YAML lib, no HTTP router framework).

### Android
Package root: `android/app/src/main/java/com/pulseos/`
- **`app/`** — Application class + `AppContainer` (manual DI)
- **`core/`** — Config, design tokens (`PulseDesignTokens`), local storage
- **`domain/model/`** — Data classes per domain
- **`feature/`** — Feature modules: onboarding, home, diet, activity, meditation, sleep, user
- **`navigation/`** — `PulseNavHost` with bottom tab bar
- API base URL: `http://10.0.2.2:8080` (emulator localhost)

### iOS
Root: `ios/PulseOS/`
- **`App/`** — SwiftUI app entry point
- **`Core/`** — Config, `PulseTheme`, `ProfileLocalStore`
- **`Domain/Models/`** — Swift structs mirroring backend models
- **`Features/`** — Feature modules matching Android
- API base URL: `http://localhost:8080`

### Infrastructure
- PostgreSQL 15, Redis 7, MinIO (S3-compatible) via `backend/deploy/docker-compose.yml`
- SQL migrations in `backend/migrations/` (001–006)
- CI (`.github/workflows/ci.yml`) runs `gofmt`, `go build`, `go test` for backend only

### Project Planning
`docs/` contains phased development plans (phase 0–8). Mark completed tasks with `- [x]` in the plan docs.
