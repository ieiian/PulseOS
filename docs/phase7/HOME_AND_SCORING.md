# Phase 7 统一评分与首页聚合

## 完成内容

- `daily_scores` 结构
- 饮食 / 运动 / 睡眠评分汇总
- 首页 dashboard 接口
- Android 首页总览升级
- iOS 首页总览升级
- 基础趋势数据展示

## 当前接口

- `GET /api/v1/home/dashboard`

## 当前评分说明

- `diet_score`: 当前为 MVP 占位计算
- `activity_score`: 取运动积分并封顶 100
- `sleep_score`: 使用睡眠评分
- `total_score`: 三项平均

## 当前首页内容

- 今日总分
- 当前待完成事项
- 饮食摘要
- 运动摘要
- 睡眠摘要
- 冥想摘要
- 趋势数据

