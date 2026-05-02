# Phase 3 饮食 MVP

## 完成内容

- 饮食数据模型
- PostgreSQL 迁移草案
- 饮食推荐规则 v1
- 断食计划模板
- 常用餐与快速记录方案
- 拍照上传接口
- 食物识别 mock provider
- 进食分析接口
- AI 建议生成接口
- Android 饮食页
- iOS 饮食页
- 四档推荐状态展示
- 后端联调测试

## 当前接口

- `GET /api/v1/diet/plan/today`
- `POST /api/v1/diet/photo-upload`
- `POST /api/v1/diet/analyze`
- `POST /api/v1/diet/records`

## 当前实现说明

- 食物识别采用 mock 识别器，按关键词返回结构化食物数据
- AI 建议采用 mock provider，负责生成解释文案
- 规则优先级：
  - 糖尿病标记 + 高糖食物 -> `forbidden`
  - 热量超标 -> `not_recommended`
  - 油炸/高糖 -> `caution`
  - 其他 -> `recommended`

## 快速记录方案

- 常用餐模板
- 手动输入食物数组
- 图片上传后分析

## 下一步衔接

- Phase 4 运动模块可复用首页卡片和推荐展示风格
- 后续可把 mock 食物识别替换成真实视觉 API
