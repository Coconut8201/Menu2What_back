待辦事項:
- [X] 加入gemini api
- [X] 新增gemini api 上傳圖片和文字功能
- [ ] 修改gemini api 回傳內容（應該使用< apiResult >）
- [ ] 修改確認是否需要上傳pdf 的功能
- [ ] 修復logger 系統

下載依賴：
```
go mod tidy
```

**開發工具：**
```
air
```

啟動：
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