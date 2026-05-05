# PulseOS 项目计划

## 1. 项目定位

PulseOS 是一款移动健康类 APP，包含五个核心模块：

- 用户：基础资料、健康档案、统计数据、设置与权限、设备与账号管理。
- 饮食：断食计划、饮食推荐、拍照识别、AI 建议、进食风险提示。
- 运动：步数统计、心肺强化积分、周目标、提醒。
- 冥想：呼吸训练、基础冥想引导、背景音频。
- 睡眠：睡眠监测、呼噜/梦话检测、睡眠质量分析、改进建议。

本计划以“先可用、再精细、再智能”为路线，不追求第一版做全，而追求：

- 功能路径清晰
- UI 观感成熟
- 工程结构可扩展
- 规则可配置
- AI 能随时接入和替换

## 2. 市场借鉴与产品策略

参考方向：

- 饮食记录与复用：借鉴 MyFitnessPal 的“快速记录、保存常用餐、按时间归档”
- 健康总览与趋势：借鉴 Apple Health 的“首页摘要、趋势、卡片化信息组织”
- 运动积分与目标：借鉴 Google Fit 一类产品的“积分目标、周视图、达成感反馈”
- 冥想与睡眠沉浸感：借鉴 Calm / Headspace 的“轻内容导航、情境化主题、睡前场景”
- 睡眠录音与事件回看：借鉴 SnoreLab 的“时间轴、重点事件回放、趋势对比”
- 用户中心与设置入口：借鉴成熟健康类 APP 的“资料、设备、权限、目标、历史统计集中管理”

落地原则：

- 借鉴功能结构，不直接拼贴 UI
- 首页只展示最重要的下一步动作和今日状态
- 每个模块都支持“快速进入 + 逐步深入”
- 复杂能力拆成渐进层级，先做高可行版本

## 3. 产品信息架构

推荐底部主导航：

1. 首页
2. 饮食
3. 运动
4. 睡眠
5. 用户

冥想建议不单独占底部 Tab，放在首页快捷入口和“用户”模块下的内容入口或单独二级页内。原因：

- 冥想使用频率通常低于饮食/运动/睡眠
- 底部导航控制在 5 个以内更稳定
- 后续若内容变多，再升级为独立一级模块

首页信息结构建议：

- 顶部问候与今日状态
- 今日总分卡
- 今日待完成事项
- 四大健康模块摘要卡
- 智能建议卡
- 趋势与连续打卡

用户模块信息结构建议：

- 用户头像与昵称
- 健康目标摘要
- 关键身体指标
- 周/月统计入口
- 设置与权限管理
- 设备连接与账号安全

## 4. UI 主题与视觉方向

目标不是“工具感很重”，而是“专业 + 温和 + 有节律感”。

### 视觉主题

- 总体风格：现代健康科技感，降低医疗系统的冰冷感
- 背景：浅暖灰或低饱和渐变底，减少纯白刺眼感
- 组件：圆角卡片、轻阴影、清晰留白
- 图表：简洁，重点突出趋势，不做复杂金融图式样

### 模块主题色

- 饮食：草木绿 / 柔和橙
- 运动：活力橙 / 珊瑚红
- 冥想：雾蓝 / 青绿
- 睡眠：深蓝 / 静谧靛蓝

### 设计原则

- 颜色用于模块识别，不用于大面积堆叠
- 首页统一风格，模块页允许局部主题变化
- 字体层级简单明确，优先提高可读性
- 动效只服务于状态反馈、呼吸节奏、睡眠场景切换

详细策略见 [docs/PRODUCT_UI_STRATEGY.md](/Users/tse/github/PulseOS/docs/PRODUCT_UI_STRATEGY.md)。

## 5. 技术架构决策

### 客户端

- Android：`Kotlin + Jetpack Compose`
- iOS：`Swift + SwiftUI`

### 后端

- API/BFF：`Go`
- 数据库：`PostgreSQL`
- 缓存：`Redis`
- 对象存储：`S3 兼容对象存储`
- 定时任务：先放在 `Go` 单体内，后续再拆

### AI 与算法

- LLM：通过 API 接入，后端统一封装 provider
- 食物识别：优先接成熟视觉 API，不自建模型
- 睡眠分析：先做规则与信号特征版，再逐步增强
- 规则引擎：关键决策先走后端强规则，AI 只做解释和补充建议

## 6. 总体架构

```text
[Android App]      [iOS App]
      \               /
       \             /
        ---- Go API / BFF ----
                  |
        -----------------------
        | user                |
        | diet                |
        | activity            |
        | meditation          |
        | sleep               |
        | scoring             |
        | ai                  |
        | rule_engine         |
        -----------------------
                  |
      PostgreSQL / Redis / Object Storage
```

## 7. 关键原则

- 业务逻辑尽量放后端，避免 Android/iOS 双端分叉。
- 前端只保留少量离线逻辑，如步数采集、本地缓存、简单分数展示。
- 所有名称、阈值、开关、模型选择、积分参数统一放配置层管理。
- 初期用 Go 单体模块化，不一开始拆微服务。
- 先做可行版本，再逐步增强 AI 与睡眠精度。
- 统一设计 token 与模块规范，避免双端视觉长期漂移。
- 每个模块都要先定义“最短完成路径”，减少首页与页面层级混乱。

## 8. 配置管理方案

### 后端统一配置

- APP 名称
- AI provider/base URL/model
- 饮食规则阈值
- 心肺强化积分换算参数
- 睡眠分析阈值
- 用户目标与推荐参数
- 功能开关
- 对象存储配置
- 第三方 API key

建议：

- 服务端使用 `configs/config.yaml + configs/env/*.env`
- 敏感配置只放环境变量，不写死进仓库
- 规则参数做成可配置结构，避免以后每次改参数都改代码

### 客户端配置

- 客户端只保留非敏感配置
- API Base URL
- 埋点开关
- 本地功能 flag
- UI 文案常量

建议再补一层共享产品配置说明：

- `docs/config/feature-flags.md`
- `docs/config/scoring-rules.md`
- `docs/config/design-tokens.md`

## 9. 推荐目录结构

```text
PulseOS/
├── README.md
├── docs/
│   └── PROJECT_PLAN.md
│   ├── PRODUCT_UI_STRATEGY.md
│   └── config/
│       ├── design-tokens.md
│       ├── feature-flags.md
│       └── scoring-rules.md
├── android/
│   ├── app/
│   │   ├── build.gradle.kts
│   │   └── src/main/java/com/pulseos/
│   │       ├── MainActivity.kt
│   │       ├── app/
│   │       │   ├── PulseApplication.kt
│   │       │   └── AppContainer.kt
│   │       ├── core/
│   │       │   ├── config/
│   │       │   ├── designsystem/
│   │       │   ├── network/
│   │       │   ├── storage/
│   │       │   ├── permissions/
│   │       │   └── utils/
│   │       ├── data/
│   │       │   ├── api/
│   │       │   ├── dto/
│   │       │   ├── repository/
│   │       │   └── mapper/
│   │       ├── domain/
│   │       │   ├── model/
│   │       │   └── usecase/
│   │       ├── feature/
│   │       │   ├── onboarding/
│   │       │   ├── home/
│   │       │   ├── user/
│   │       │   ├── diet/
│   │       │   ├── activity/
│   │       │   ├── meditation/
│   │       │   └── sleep/
│   │       ├── navigation/
│   │       └── ui/theme/
│   └── gradle/
├── ios/
│   └── PulseOS/
│       ├── App/
│       │   └── PulseOSApp.swift
│       ├── Core/
│       │   ├── Config/
│       │   ├── DesignSystem/
│       │   ├── Network/
│       │   ├── Storage/
│       │   ├── Permissions/
│       │   └── Utils/
│       ├── Data/
│       │   ├── API/
│       │   ├── DTO/
│       │   ├── Repository/
│       │   └── Mapper/
│       ├── Domain/
│       │   ├── Models/
│       │   └── UseCases/
│       ├── Features/
│       │   ├── Onboarding/
│       │   ├── Home/
│       │   ├── User/
│       │   ├── Diet/
│       │   ├── Activity/
│       │   ├── Meditation/
│       │   └── Sleep/
│       ├── Navigation/
│       └── UI/
├── backend/
│   ├── cmd/
│   │   └── server/
│   │       └── main.go
│   ├── configs/
│   │   ├── config.yaml
│   │   └── env/
│   │       ├── dev.env.example
│   │       ├── test.env.example
│   │       └── prod.env.example
│   ├── deploy/
│   │   └── docker-compose.yml
│   ├── internal/
│   │   ├── app/
│   │   ├── bootstrap/
│   │   ├── domain/
│   │   │   ├── user/
│   │   │   ├── diet/
│   │   │   ├── activity/
│   │   │   ├── meditation/
│   │   │   ├── sleep/
│   │   │   └── scoring/
│   │   ├── service/
│   │   ├── repository/
│   │   │   ├── postgres/
│   │   │   ├── redis/
│   │   │   └── objectstorage/
│   │   ├── handler/
│   │   ├── middleware/
│   │   ├── ai/
│   │   │   ├── provider/
│   │   │   ├── prompt/
│   │   │   └── schema/
│   │   ├── ruleengine/
│   │   ├── scheduler/
│   │   ├── analytics/
│   │   └── pkg/
│   ├── migrations/
│   ├── scripts/
│   ├── tests/
│   └── go.mod
└── assets/
    ├── audio/
    └── mock/
```

## 10. 模块边界

### 10.1 用户

- 基础资料录入
- 健康目标管理
- 身体指标管理
- 统计数据总览
- 账号与登录态
- 权限管理
- 通知与提醒设置
- 设备连接管理
- 隐私与数据导出预留

### 10.2 饮食

- 用户基础资料输入
- 目标热量与营养建议计算
- 断食计划模板
- 饮食推荐候选项
- 常用餐复用
- 用餐时间记录
- 拍照识别食物
- 单次进食分析
- 当日饮食总结
- AI 解释层

说明：

- 用户基础资料输入的采集入口在“用户模块”
- 饮食模块只消费用户资料，不负责资料管理

### 10.3 运动

- 系统步数读取
- 运动类型补录
- 心肺强化积分换算
- 周目标跟踪
- 提醒与鼓励机制
- 趋势与连续达标展示

### 10.4 冥想

- 呼吸训练模式
- 基础冥想引导
- 背景声音播放
- 时长统计
- 睡前快速入口

### 10.5 睡眠

- 睡眠会话开始/结束
- 音频采样
- 呼噜/梦话事件检测
- 睡眠时长统计
- 质量评分
- 建议与提示
- 重点事件时间轴回看

## 11. 产品能力分层

### L1：必做基础能力

- 用户 onboarding
- 用户资料页与设置页
- 今日首页
- 饮食记录与推荐
- 步数与运动积分
- 呼吸训练
- 睡眠开始/结束与基础分析

### L2：增强体验能力

- 用户统计页
- 用户目标调整
- 权限与设备管理
- AI 图片分析解释
- 常用餐与推荐模板
- 周目标趋势
- 冥想音频内容组织
- 睡眠事件时间轴

### L3：后续扩展能力

- 个性化计划自动生成
- 多日趋势与行为洞察
- 家庭/医生共享
- 穿戴设备扩展
- 更精细的睡眠分期和异常检测

## 12. 可行性与优先级

### 高可行

- onboarding
- 饮食规则推荐
- 拍照上传 + 第三方识别
- 步数接入
- 积分系统
- 基础呼吸训练
- 音频播放
- 睡眠基础录音与事件标记

### 中可行

- 用餐图像多食物识别
- 更精细营养评估
- 睡眠质量评分
- 打鼾强度趋势
- 个性化提醒策略

### 低可行或应后置

- 自建识别模型
- 医疗级睡眠判断
- 高置信度梦话识别
- 复杂多疾病联合风险分析

## 13. 数据与规则核心

### 统一评分

- `diet_score`
- `activity_score`
- `sleep_score`
- `total_score`

### 统一事件记录

- 所有关键行为按“记录”建模
- 饮食记录、运动记录、睡眠记录、冥想记录分开存
- 首页按日期聚合

### 规则优先级

1. 强规则：禁忌、风险、阈值判断
2. 评分规则：积分、目标达成度
3. AI 输出：解释、个性化描述、补充建议

## 14. 页面结构建议

### 首页

- 今日健康总览
- 今日三大关键指标
- 快速动作入口
- 今日建议
- 趋势摘要

### 用户页

- 个人信息卡
- 健康目标卡
- 身体指标摘要
- 统计数据入口
- 设置列表
- 权限与设备状态

### 饮食页

- 今日目标卡
- 当前进食状态
- 快速记录入口
- 推荐餐单
- 历史记录

### 运动页

- 今日步数与积分环
- 周目标进度
- 最近记录
- 补录入口

### 冥想页

- 快速开始
- 呼吸模式
- 推荐音频
- 最近使用

### 睡眠页

- 昨夜概览
- 开始睡眠按钮
- 事件时间轴
- 趋势图
- 睡前建议

## 15. 分阶段执行计划表

使用方式：

- 未完成：`- [ ]`
- 已完成：`- [x]`

### Phase 0：产品定义与技术预研

- [x] 明确 MVP 范围，只保留饮食/运动/冥想/睡眠四大主线中的基础能力
- [x] 定义用户画像、目标用户、健康目标分类
- [x] 明确第一版不做的内容，避免范围失控
- [x] 确定 AI provider、食物识别 API 预选方案
- [x] 确定睡眠音频采样与权限可行性
- [x] 输出 PRD v1、用户流程图、核心页面清单
- [x] 输出 UI 风格板、主题色、首页信息层级草案

### Phase 1：基础工程搭建

- [x] 初始化 Android 工程：Kotlin + Compose
- [x] 初始化 iOS 工程：Swift + SwiftUI
- [x] 初始化 Go 后端工程
- [x] 建立统一目录结构
- [x] 建立配置体系：`config.yaml + env`
- [x] 建立设计系统目录：颜色、字号、间距、卡片规范
- [x] 接入日志、错误码、基础鉴权
- [x] 建立本地开发环境：PostgreSQL/Redis/Object Storage
- [x] 配置 CI 基础检查

### Phase 2：账户与 Onboarding

- [x] 设计用户基础资料结构
- [x] 完成首次启动引导页
- [x] 完成健康目标录入页
- [x] 完成后端用户资料 API
- [x] 完成客户端用户资料持久化
- [x] 完成用户主页信息架构
- [x] 完成设置、权限、通知页结构
- [x] 完成基础首页骨架
- [x] 完成首页卡片组件和模块主题规范

### Phase 3：饮食 MVP

- [x] 设计饮食数据模型与数据库表
- [x] 完成饮食推荐规则 v1
- [x] 完成断食计划数据结构与模板
- [x] 完成常用餐与快速记录方案
- [x] 完成拍照上传接口
- [x] 接入食物识别 API
- [x] 完成进食分析接口
- [x] 完成 AI 建议生成接口
- [x] 完成 Android 饮食页面
- [x] 完成 iOS 饮食页面
- [x] 完成“建议食用/不建议食用/注意食用/禁止食用”展示逻辑
- [x] 完成饮食模块联调测试

### Phase 4：运动 MVP

- [x] 设计运动数据模型与数据库表
- [x] 定义心肺强化积分规则 v1
- [x] Android 接入步数读取
- [x] iOS 接入步数读取
- [x] 完成手动补录运动
- [x] 完成后端积分计算接口
- [x] 完成日/周目标接口
- [x] 完成 Android 运动页面
- [x] 完成 iOS 运动页面
- [x] 完成提醒策略 v1
- [x] 完成运动趋势图与达成反馈

### Phase 5：冥想 MVP

- [x] 设计冥想数据模型
- [x] 整理基础音频资源
- [x] 实现呼吸训练节奏配置
- [x] 完成 Android 冥想页面与播放控制
- [x] 完成 iOS 冥想页面与播放控制
- [x] 完成冥想记录上报接口
- [x] 完成时长统计与展示

### Phase 6：睡眠 MVP

- [x] 设计睡眠记录模型与数据库表
- [x] 设计睡眠会话开始/结束流程
- [x] 完成音频采样方案 v1
- [x] 完成基础呼噜检测规则
- [x] 完成睡眠时长统计
- [x] 完成睡眠评分规则 v1
- [x] 完成睡眠建议输出接口
- [x] 完成 Android 睡眠页面
- [x] 完成 iOS 睡眠页面
- [x] 完成事件时间轴与重点片段回看
- [x] 完成睡眠模块联调测试

### Phase 7：统一评分与首页聚合

- [x] 建立 `daily_scores` 结构
- [x] 汇总饮食/运动/睡眠评分
- [x] 完成首页总览接口
- [x] 完成 Android 首页
- [x] 完成 iOS 首页
- [x] 完成基础趋势图表

### Phase 8：质量与上线准备

- [x] 补充接口测试与关键单元测试
- [ ] 完成埋点与日志审计
- [ ] 完成崩溃监控接入
- [x] 评估隐私合规与权限文案
- [x] 整理测试用例
- [ ] 完成内测包构建
- [ ] 完成第一轮用户试用反馈修正

## 16. 第一阶段建议开发顺序

如果你现在就开始做，建议顺序固定为：

1. 先搭工程和配置体系
2. 先做用户 onboarding
3. 先打通饮食 MVP
4. 再做运动积分
5. 再补冥想
6. 最后做睡眠基础版

原因：

- 饮食最能体现 AI 价值，也最适合作为首个核心卖点
- 运动依赖系统 API，但实现稳定，适合作为第二主线
- 冥想实现简单，可快速补足产品完整度
- 睡眠难度最高，必须后置，否则会拖慢整体交付

## 17. 第一版 API 范围建议

### 用户

- `POST /api/v1/users/onboarding`
- `GET /api/v1/users/profile`
- `PUT /api/v1/users/profile`
- `GET /api/v1/users/stats`
- `GET /api/v1/users/settings`
- `PUT /api/v1/users/settings`
- `GET /api/v1/users/devices`

### 饮食

- `GET /api/v1/diet/plan/today`
- `POST /api/v1/diet/records`
- `POST /api/v1/diet/analyze`
- `GET /api/v1/diet/summary/today`

### 运动

- `POST /api/v1/activity/records`
- `GET /api/v1/activity/today`
- `GET /api/v1/activity/week`

### 冥想

- `POST /api/v1/meditation/sessions`
- `GET /api/v1/meditation/today`

### 睡眠

- `POST /api/v1/sleep/sessions/start`
- `POST /api/v1/sleep/sessions/end`
- `GET /api/v1/sleep/today`

### 首页

- `GET /api/v1/home/dashboard`

## 18. 数据表建议

第一版建议至少包含：

- `users`
- `user_health_profiles`
- `user_settings`
- `user_devices`
- `diet_plans`
- `food_records`
- `activity_records`
- `meditation_sessions`
- `sleep_records`
- `sleep_events`
- `daily_scores`

## 19. 里程碑定义

### M1：工程可运行

- Android/iOS/Backend 三端目录完成
- 配置与环境完成
- 本地服务能跑

### M2：用户可进入 APP

- Onboarding 完成
- 用户资料可提交和读取
- 用户页基础结构可用
- 首页骨架可展示

### M3：饮食 MVP 跑通

- 可拍照上传
- 可返回识别结果
- 可返回规则判断和 AI 解释

### M4：运动 MVP 跑通

- 可读取步数
- 可计算心肺积分
- 可展示周目标

### M5：睡眠基础版跑通

- 可开始/结束睡眠监测
- 可生成基础分析

### M6：首个内测版本

- 四大模块都有基础功能
- 首页评分可用
- 可进入测试与迭代

## 20. 下一步建议

按执行效率，最合理的下一步是：

1. 先创建三端脚手架和统一目录
2. 先补 `backend/configs` 与 env 模板
3. 先定义数据库表与 API 草案
4. 先做 onboarding 和饮食 MVP

如果你愿意，我下一步可以直接继续帮你做以下其中一项：

1. 生成三端工程目录骨架
2. 生成 Go 后端初始化代码
3. 生成数据库 SQL
4. 生成 API 文档草案
