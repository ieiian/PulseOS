# Phase 2 账户与 Onboarding

## 完成内容

- 用户基础资料结构
- Android onboarding 页面
- iOS onboarding 页面
- Android 本地资料持久化
- iOS 本地资料持久化
- 后端用户资料 API
- 首页基础骨架
- 用户页信息架构
- 设置/权限/通知结构占位

## 用户资料结构

- `name`
- `age`
- `gender`
- `height_cm`
- `weight_kg`
- `primary_goal`
- `secondary_goals`
- `health_flags`

## 后端接口

- `POST /api/v1/users/onboarding`
- `GET /api/v1/users/profile`
- `PUT /api/v1/users/profile`
- `GET /api/v1/users/settings`
- `PUT /api/v1/users/settings`
- `GET /api/v1/users/stats`

## 当前实现说明

- 当前后端用户仓库采用内存实现，放在 `repository/postgres` 目录中作为 Phase 2 过渡版本
- Phase 3/Phase 4 再替换为真实 PostgreSQL 持久化
- 双端已经有可运行的 onboarding -> 首页/用户页基础流

## 下一步衔接

- 饮食模块直接消费 `primary_goal`、`height_cm`、`weight_kg`
- 运动模块直接消费用户目标和权限状态
- 睡眠模块直接消费通知和麦克风权限状态
