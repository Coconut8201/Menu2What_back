待辦事項:
- [X] 加入gemini api
- [X] 新增gemini api 上傳圖片和文字功能
- [ ] 修改gemini api 回傳內容（應該使用< apiResult >）
- [X] 修復logger 系統
- [x] user 創建使用者功能
- [x] user 登入登出功能

1. 下載：
```
clone https://github.com/Coconut8201/Menu2What_back.git
cd Menu2What_back
```
複製.env.example 到 .env：
```
cp .env.example .env
```
gemini api key (必填)


2. 下載依賴：
```
go mod tidy
```
3. 啟動：
```
go run .
```
路徑啟動在：
```
http://localhost:6382
```

服務狀態：
```
http://localhost:6382/ping
```

更新Swagger API Docs:
```
swag init
```

Swagger API Docs:
```
http://localhost:6382/swagger/index.html#/
```


**開發工具：**
```
air
```