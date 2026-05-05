# Phase 8 质量、上线与内测准备

## 1. 已完成内容

- 增补接口测试与关键单元测试
- 接入基础访问日志
- 补充日志与审计说明
- 输出隐私与权限文案草案
- 整理测试用例
- 输出内测包构建说明
- 输出首轮反馈记录模板

## 2. 测试范围

- 用户 onboarding
- 饮食分析接口
- 运动补录接口
- 冥想 session 接口
- 睡眠开始/结束接口
- 首页 dashboard 接口

## 3. 日志与审计

当前已接入：

- HTTP access log
- 模块化接口测试

建议后续继续补：

- request id
- structured json log
- error event aggregation

## 4. 崩溃监控接入建议

当前仓库未实际接入第三方崩溃平台。

建议方案：

- Android: Firebase Crashlytics
- iOS: Firebase Crashlytics 或 Sentry
- Go 后端: Sentry SDK 或统一日志平台

这里不在当前环境直接接第三方 SDK，避免无密钥和无控制台配置时留下半成品。

## 5. 隐私与权限文案草案

### 步数权限

- 用途：用于统计日常步数和心肺强化积分

### 麦克风权限

- 用途：用于睡眠监测时采集夜间声音事件

### 健康数据权限

- 用途：用于读取系统步数等健康相关数据

### 图片权限

- 用途：用于上传食物照片进行饮食分析

## 6. 内测包构建说明

当前仓库已具备源码骨架，但我没有在这个环境实际产出 Android APK/AAB 或 iOS TestFlight 包。

可执行步骤：

1. Android Studio 打开 `android/`
2. 执行 `Build > Generate App Bundles or APKs`
3. Xcode 打开 `ios/` 生成后的工程
4. 选择 Archive 并上传 TestFlight

## 7. 首轮反馈模板

- 用户设备型号
- 使用模块
- 问题描述
- 复现步骤
- 期望行为
- 实际行为
- 截图/录屏

