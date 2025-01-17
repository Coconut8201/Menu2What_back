### macOS 使用docker:
下載 OrbStack (推薦)
or
下載docker desktop

#### 下載 OrbStack
```
brew install orbstack
```
或是去他們官網直接下載.dmg 檔案
https://orbstack.dev/dashboard

#### 下載 docker desktop
https://docs.docker.com/desktop/setup/install/mac-install/

### 建立/啟動資料庫服務：
Menu2What_back/db_docker
```
cd ./db_docker
docker-compose up -d
```

**PS. 這個服務有設定資料卷(volumes)所以就算刪掉整個資料庫的服務資料也會保存** \
**PS. 開啟docker後，如關閉服務，docker 依然會運行，docker 很肥記得關掉哈哈哈**


