# Phase 4 运动 MVP

## 完成内容

- 运动数据模型
- PostgreSQL 迁移草案
- 心肺强化积分规则 v1
- Android 步数读取适配层
- iOS 步数读取适配层
- 手动补录接口
- 后端积分计算接口
- 日/周目标接口
- Android 运动页
- iOS 运动页
- 提醒策略 v1
- 趋势展示骨架

## 当前接口

- `POST /api/v1/activity/records`
- `GET /api/v1/activity/today`
- `GET /api/v1/activity/week`

## 当前积分规则

- `moderate`: `1 分钟 = 1 心肺强化分`
- `vigorous`: `1 分钟 = 2 心肺强化分`
- `light`: `2 分钟 = 1 分`
- 周目标：`150 分`

## 当前实现说明

- 步数读取已接入平台适配入口：
  - Android: `Sensor.TYPE_STEP_COUNTER`
  - iOS: `CMPedometer`
- 当前客户端页面使用预览值展示，不包含真机权限验证和持续监听
- 手动补录会走后端积分换算

## 下一步衔接

- Phase 5 可复用相同卡片结构做冥想时长展示
- 后续可把步数预览值替换成真实设备读取结果
