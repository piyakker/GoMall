# 🛒 Go mall task board（MVP 版本）

## 🔨 Todo

- [ ] **T6**：設計商品資料表（PostgreSQL）
- [ ] **T7**：建立商品列表 API
- [ ] **T8**：加入 Redis 商品快取機制
- [ ] **T9**：加入商品查詢單元測試

- [ ] **T10**：設計訂單資料表（PostgreSQL）
- [ ] **T11**：建立下單 API
- [ ] **T12**：簡易訂單限流（Redis + Lua）
- [ ] **T13**：加入訂單單元測試

- [ ] **T14**：定義 MQ Topic / Exchange
- [ ] **T15**：建立 Go Worker 消費 MQ 訊息
- [ ] **T16**：加上錯誤重試與 Fail log 記錄
- [ ] **T17**：模擬寄信/推播處理（mock）

- [ ] **T18**：撰寫整合測試（下單流程）
- [ ] **T19**：撰寫壓力測試（可用 vegeta）
- [ ] **T20**：製作 Postman Collection

- [ ] **T21**：撰寫簡易 Dockerfile / docker-compose
- [ ] **T22**：撰寫專案 README 使用說明
- [ ] **T23**：製作 Demo 錄影或簡報

---

## 🏗 In Progress

---

## ✅ Done

- [x] **T1**：建立 Git repo + 架構 README
- [x] **T2**：建立 Go module + 基礎專案結構
- [x] **T3**：安裝並配置 RabbitMQ、Redis、PostgreSQL（docker-compose）
- [x] **T4**：設計 .env 檔與設定載入（dotenv）
- [x] **T5**：撰寫 HealthCheck API

---

## 💡 Bonus

- [ ] **B1**：商品快取加上自動失效策略（TTL + 回源）
- [ ] **B2**：訂單服務改為事件導向架構（Worker 解耦）
- [ ] **B3**：加入 Prometheus + Grafana 監控
- [ ] **B4**：health check 檢查連線狀態
