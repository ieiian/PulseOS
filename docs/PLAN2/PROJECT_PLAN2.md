# PulseOS 实施计划 2 — 端到端功能落地

## 背景

Phase 0–7 的骨架搭建已全部完成（PROJECT_PLAN.md 中标记为 `[x]`），但三端代码目前处于"UI 外壳 + 硬编码数据"状态：

- **后端**：所有 repository 是内存 stub，尽管包名是 `postgres`；SQL migrations 存在但 Go 代码从未连接数据库
- **iOS**：22 个 Swift 文件，8 个空目录（Network、DTO、Repository 等）；零 URLSession 调用
- **Android**：26 个 Kotlin 文件，无 HTTP 客户端（无 Retrofit/OkHttp）；无 core/network 目录

本计划的目标：**将所有模块从骨架变为可运行的端到端应用**，后端接真实 PostgreSQL，客户端调真实 API。AI Provider 保持 Mock。

---

## Phase 2A：后端数据库集成

### 步骤 2A-1：添加 Go 依赖

- [x] `go get github.com/jackc/pgx/v5`
- [x] `go get gopkg.in/yaml.v3`

### 步骤 2A-2：创建 users 表 migration

- [x] 创建 users 表
- [x] 创建 user_settings 表
- [x] 创建 user_stats 表

### 步骤 2A-3：给 domain model 添加 UserID 和时间戳字段

- [x] `internal/domain/user/user.go`
- [x] `internal/domain/diet/diet.go`
- [x] `internal/domain/activity/activity.go`
- [x] `internal/domain/meditation/meditation.go`
- [x] `internal/domain/sleep/sleep.go`
- [x] `internal/domain/scoring/scoring.go`

所有领域类型添加 `UserID int64` 和 `CreatedAt time.Time`：

- `internal/domain/user/user.go` — Profile 加 `ID int64`, `CreatedAt`；Settings/Stats 加 `UserID`
- `internal/domain/diet/diet.go` — Record 加 `UserID`, `CreatedAt`
- `internal/domain/activity/activity.go` — Record 加 `UserID`, `CreatedAt`
- `internal/domain/meditation/meditation.go` — Session 加 `UserID`, `CreatedAt`
- `internal/domain/sleep/sleep.go` — Session 加 `UserID`, `CreatedAt`
- `internal/domain/scoring/scoring.go` — DailyScore 加 `UserID`

### 步骤 2A-4：创建数据库连接包

- [x] 新建 `internal/database/db.go`

### 步骤 2A-5：更新 app.go 初始化数据库

- [x] 修改 `internal/app/app.go`

### 步骤 2A-6：迁移内存 stub 到 memory 包

- [x] 新建 `internal/repository/memory/` 6 个文件
- [x] 更新 `backend/tests/*.go` import

### 步骤 2A-7：用真实 SQL 重写 UserRepository

- [x] 重写 `internal/repository/postgres/user_repository.go`

### 步骤 2A-8：用真实 SQL 重写 DietRepository

- [x] 重写 `internal/repository/postgres/diet_repository.go`

### 步骤 2A-9：用真实 SQL 重写 ActivityRepository

- [x] 重写 `internal/repository/postgres/activity_repository.go`

### 步骤 2A-10：用真实 SQL 重写 MeditationRepository

- [x] 重写 `internal/repository/postgres/meditation_repository.go`

### 步骤 2A-11：用真实 SQL 重写 SleepRepository

- [x] 重写 `internal/repository/postgres/sleep_repository.go`

### 步骤 2A-12：用真实 SQL 重写 ScoringRepository

- [x] 重写 `internal/repository/postgres/scoring_repository.go`

### 步骤 2A-13：消除 service 层硬编码数据

- [x] 修改 `internal/service/scoring_service.go`
- [x] 修改 `internal/service/sleep_service.go`
- [x] 修改 `internal/service/activity_service.go`

### 步骤 2A-14：添加 CORS 中间件

- [x] 新建 `internal/middleware/cors.go`
- [x] 修改 `internal/app/app.go` 包裹 mux

### 步骤 2A-15：补充缺失的 GET 接口

- [x] `GET /api/v1/diet/records`
- [x] `GET /api/v1/activity/records`

### 步骤 2A-16：更新 bootstrap config 解析

- [x] 用 `gopkg.in/yaml.v3` 替换手写解析器

### 步骤 2A-17：后端端到端冒烟测试

- [x] 新建 `backend/scripts/e2e_test.sh`

---

## Phase 2B：iOS 网络层

### 步骤 2B-1：创建 APIClient

- [x] 新建 `ios/PulseOS/Core/Network/APIClient.swift`

### 步骤 2B-2：创建响应 DTO

- [x] `ios/PulseOS/Data/DTO/UserDTOs.swift`
- [x] `ios/PulseOS/Data/DTO/DietDTOs.swift`
- [x] `ios/PulseOS/Data/DTO/ActivityDTOs.swift`
- [x] `ios/PulseOS/Data/DTO/MeditationDTOs.swift`
- [x] `ios/PulseOS/Data/DTO/SleepDTOs.swift`
- [x] `ios/PulseOS/Data/DTO/HomeDTOs.swift`

### 步骤 2B-3：创建 API Service 类

- [x] `ios/PulseOS/Data/API/UserService.swift`
- [x] `ios/PulseOS/Data/API/DietService.swift`
- [x] `ios/PulseOS/Data/API/ActivityService.swift`
- [x] `ios/PulseOS/Data/API/MeditationService.swift`
- [x] `ios/PulseOS/Data/API/SleepService.swift`
- [x] `ios/PulseOS/Data/API/HomeService.swift`

### 步骤 2B-4：让 Domain Model 支持 Decodable

- [x] `ios/PulseOS/Domain/Models/HomeModels.swift`
- [x] `ios/PulseOS/Domain/Models/DietModels.swift`
- [x] `ios/PulseOS/Domain/Models/ActivityModels.swift`
- [x] `ios/PulseOS/Domain/Models/SleepModels.swift`
- [x] `ios/PulseOS/Domain/Models/MeditationModels.swift`

### 步骤 2B-5：更新 project.yml

- [x] 修改 `ios/project.yml`（XcodeGen 自动发现源文件，无需修改）

---

## Phase 2C：iOS 页面接入真实数据

### 步骤 2C-1：HomeDashboardView 接入 API

- [x] 修改 `ios/PulseOS/Features/Home/HomeDashboardView.swift`

### 步骤 2C-2：DietView 接入 API

- [x] 修改 `ios/PulseOS/Features/Diet/DietView.swift`

### 步骤 2C-3：ActivityView 接入 API

- [x] 修改 `ios/PulseOS/Features/Activity/ActivityView.swift`

### 步骤 2C-4：MeditationView 接入 API

- [x] 修改 `ios/PulseOS/Features/Meditation/MeditationView.swift`

### 步骤 2C-5：SleepView 接入 API

- [x] 修改 `ios/PulseOS/Features/Sleep/SleepView.swift`

### 步骤 2C-6：UserView 接入 API

- [x] 修改 `ios/PulseOS/Features/User/UserView.swift`

### 步骤 2C-7：Onboarding 同步到后端

- [x] 修改 `ios/PulseOS/Features/Onboarding/OnboardingView.swift`

---

## Phase 2D：Android 网络层

### 步骤 2D-1：添加依赖

- [x] 修改 `android/app/build.gradle.kts`

### 步骤 2D-2：添加 INTERNET 权限

- [x] 修改 `android/app/src/main/AndroidManifest.xml`

### 步骤 2D-3：创建 APIClient

- [x] 新建 `core/network/ApiClient.kt`

### 步骤 2D-4：创建响应 DTO

- [x] `data/dto/UserDTOs.kt`
- [x] `data/dto/DietDTOs.kt`
- [x] `data/dto/ActivityDTOs.kt`
- [x] `data/dto/MeditationDTOs.kt`
- [x] `data/dto/SleepDTOs.kt`
- [x] `data/dto/HomeDTOs.kt`

### 步骤 2D-5：创建 API Service 类

- [x] `data/api/UserService.kt`
- [x] `data/api/DietService.kt`
- [x] `data/api/ActivityService.kt`
- [x] `data/api/MeditationService.kt`
- [x] `data/api/SleepService.kt`
- [x] `data/api/HomeService.kt`

### 步骤 2D-6：更新 AppContainer

- [x] 修改 `app/AppContainer.kt`

---

## Phase 2E：Android 页面接入真实数据

### 步骤 2E-1：更新 PulseNavHost 传递 AppContainer

- [x] 修改 `navigation/PulseNavHost.kt`

### 步骤 2E-2：HomeScreen 接入 API

- [x] 修改 `feature/home/HomeScreen.kt`

### 步骤 2E-3：DietScreen 接入 API

- [x] 修改 `feature/diet/DietScreen.kt`

### 步骤 2E-4：ActivityScreen 接入 API

- [x] 修改 `feature/activity/ActivityScreen.kt`

### 步骤 2E-5：MeditationScreen 接入 API

- [x] 修改 `feature/meditation/MeditationScreen.kt`

### 步骤 2E-6：SleepScreen 接入 API

- [x] 修改 `feature/sleep/SleepScreen.kt`

### 步骤 2E-7：UserScreen 接入 API

- [x] 修改 `feature/user/UserScreen.kt`

### 步骤 2E-8：Onboarding 同步到后端

- [x] 修改 `feature/onboarding/OnboardingScreen.kt`

---

## Phase 2F：端到端验证

### 步骤 2F-1：后端集成测试

- [x] 修复 e2e_test.sh 路径和状态码（7 处错误）
- [ ] 运行 `docker compose up -d && bash scripts/e2e_test.sh`（需本地 Docker）

### 步骤 2F-2：iOS 模拟器验证

- [ ] 逐屏手动测试

### 步骤 2F-3：Android 模拟器验证

- [ ] 逐屏手动测试

### 步骤 2F-4：更新 CI

- [x] 修改 `.github/workflows/ci.yml`

---

## 依赖总结

| 平台 | 新增依赖 | 用途 |
|------|---------|------|
| Go | `pgx/v5`, `yaml.v3` | PostgreSQL 驱动 + YAML 解析 |
| iOS | 无（URLSession + Codable 是标准库） | — |
| Android | OkHttp 4.12, Gson 2.11, coroutines 1.8 | HTTP + JSON + 异步 |

## 执行顺序

严格按 Phase 2A → 2B → 2C → 2D → 2E → 2F 顺序执行，每个 Phase 内的步骤按编号顺序。

- 2A 是基础（后端数据库），2B/2D 可并行（iOS/Android 网络层），2C/2E 依赖各自网络层完成
- 预计涉及 ~42 个文件修改，~32 个新文件
