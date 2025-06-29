# 🛒 GoMall Task Board（MVP + 可擴展）

## 🔨 Todo

### 🛍 商品模組

- [ ] **T7**：建立商品列表 API（支援分頁查詢）
- [ ] **T8**：加入 Redis 商品快取機制（read-through cache）
- [ ] **T9**：加入商品查詢單元測試

### 🧾 訂單模組

- [ ] **T10**：設計訂單資料表（含狀態欄位）
- [ ] **T11**：建立下單 API（預留 user_id）
- [ ] **T12**：簡易訂單限流（Redis + Lua）
- [ ] **T13**：加入訂單單元測試

### 👤 使用者與驗證

- [ ] **T24**：設計使用者資料表（PostgreSQL）
- [ ] **T25**：建立註冊 / 登入 API（JWT）
- [ ] **T26**：加入 JWT 驗證中介層（middleware）

### 💳 金流模擬

- [ ] **T27**：模擬付款 API（mock payment）
- [ ] **T28**：付款後更新訂單狀態
- [ ] **T29**：付款失敗流程模擬（取消訂單 / 回補庫存）

### 📩 非同步處理與 MQ

- [ ] **T14**：定義 MQ Topic / Exchange
- [ ] **T15**：建立 Go Worker 消費 MQ 訊息
- [ ] **T16**：加上錯誤重試與 Fail log 記錄
- [ ] **T17**：模擬寄信 / 推播處理（mock notify service）

### 🧪 測試與工具

- [ ] **T18**：撰寫整合測試（下單流程）
- [ ] **T19**：撰寫壓力測試（可用 vegeta）
- [ ] **T20**：製作 Postman Collection
- [ ] **T30**：加入 log middleware（請求/錯誤日誌）
- [ ] **T31**：加入 metrics（Prometheus 支援點）

### 📦 部署與交付

- [ ] **T21**：撰寫 Dockerfile / docker-compose
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
- [x] **T5**：撰寫 HealthCheck API（含 Redis / DB / MQ 檢查）
- [x] **T6**：設計商品資料表（PostgreSQL）

---

## 💡 Bonus（延展性強化）

- [ ] **B1**：商品快取加上自動失效策略（TTL + 回源）
- [ ] **B2**：訂單服務改為事件導向架構（Worker 解耦）
- [ ] **B3**：加入 Prometheus + Grafana 監控
- [ ] **B4**：health check 檢查連線狀態（Prometheus 探針）
- [ ] **B5**：事件儲存機制（event log for replay）
- [ ] **B6**：開放商品 API 給未登入者（模擬前台）
