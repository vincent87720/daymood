# Daymood Supply Chain Management System
Daymood進銷存管理系統提供進出貨資料紀錄功能，並提供多種歷史紀錄及報表查詢，讓營運狀況更一目瞭然。

- [功能](#功能)
- [安裝](#安裝)
    - [快速安裝](#快速安裝)
    - [手動安裝](#手動安裝)
- [設計](#設計)
- [開發](#開發)
- [建置](#建置)
    - [快速建置](#快速建置)
    - [手動建置](#手動建置)

# 功能
- 提供廠商、商品、採購、採購明細、出貨、出貨明細及折扣等資料的增刪查改功能
- 依照商品及進出貨之間的關聯列出商品的進出貨歷史紀錄
- 依據歷史紀錄自動計算成本、毛利率及總收支
- 提供方便快速的搜尋功能，使尋找資料更便利
- 可設定抽成百分比、國際運費、匯率、關稅及成本，讓成本計算更精準

# 預覽
![](/assets/img/0.0.png)
![](/assets/img/0.1.png)
![](/assets/img/0.4.png)
![](/assets/img/0.2.png)
![](/assets/img/0.3.png)

# 安裝

## 快速安裝
```sh
make deploy
```

## 手動安裝
### Database
```
docker run -itd \
    -p 5432:5432 \
    --name daymood-database \
    -v "$(PWD)/database/postgres:/var/lib/postgresql/data" \
    -e POSTGRES_USER="daymood"\
    -e POSTGRES_PASSWORD="daymood"\
    vincent87720/daymood-database
```
daymood-database與postgres相同，使用5432port作為預設連接埠
`-p 5432:5432`可以將容器內的5432port綁定到本機的5432port

`--name daymood-database`將容器的名稱指定為`daymood-database`，便於識別及設定

`-v`可將本地的資料夾掛載到容器中，用法為`-v "本地絕對路徑:/var/lib/postgresql/data"`，`/var/lib/postgresql/data`為daymood-database預設的資料存放目錄

`-e`可設定容器的環境變數，daymood-database所使用的環境變數如下：

- POSTGRES_USER: 指定使用者名稱，同時也作為資料庫role的名稱
- POSTGRES_PASSWORD: 指定使用者密碼

### APP
```
docker run -itd \
    -p 8000:8000 \
    --name daymood-app \
    -e APP_MODE="PROD"\
    -e APP_HOST=""\
    -e APP_PORT="8000" \
    -e DB_HOSTNAME="daymood-postgres" \
    -e DB_DATABASE="daymood" \
    -e DB_USERNAME="daymood" \
    -e DB_PASSWORD="daymood" \
    -e SESSION_SECRET="daymood" \
    vincent87720/daymood-app
```
daymood-app使用8000port作為預設連接埠
`-p 8000:8000`可以將容器內的8000port綁定到本機的8000port

`--name daymood-app`將容器的名稱指定為`daymood-app`，便於識別及設定

`-e`可設定容器的環境變數，daymood-app所使用的環境變數如下：

- APP_MODE: 指定程式模式，模式包含`DEV`與`PROD`兩種，模式的設定會影響到程式中`rootPath`的值，進而影響到**settings.yaml**與**systemConfigs.json**等檔案的檔案路徑設定
- APP_HOST: 可指定app的host，預設為空
- APP_PORT: 指定app的port，預設為8000
- DB_HOSTNAME: 指定資料庫的host，可填入資料庫的IP位址或域名，若使用docker-compose或docker network可使用資料庫容器的名稱作為hostname
- DB_DATABASE: 指定資料庫名稱
- DB_USERNAME: 指定資料庫使用者名稱
- DB_PASSWORD: 指定資料庫使用者密碼
- SESSION_SECRET: 指定login session的加密字串

### WebServer
```
docker run -itd \
    -p 80:80 \
    --name daymood-webserver \
    -e SERVER_HOST="0.0.0.0"\
    -e APP_HOST="daymood-app"\
    -e APP_PORT="8000"\
    vincent87720/daymood-webserver
```
daymood-app使用80port作為預設連接埠
`-p 80:80`可以將容器內的80port綁定到本機的80port

`--name daymood-webserver`將容器的名稱指定為`daymood-webserver`，便於識別及設定

`-e`可設定容器的環境變數，daymood-nginx所使用的環境變數如下：

- SERVER_HOST: 指定nginx的hostname，預設為`0.0.0.0`
- APP_HOST: 指定daymood-app的port，可填入daymood-app的IP位址或域名，若使用docker-compose或docker network可使用daymood-app容器的名稱作為hostname
- APP_PORT: 指定daymood-app的port，預設為8000

# 設計
此系統採用主從式架構進行設計，以網頁的形式提供服務
此系統使用以下框架及軟體進行開發和部署：
- ReverseProxy
    - Nginx
- Web
    - Vue, Vuetify
- API
    - Golang, Gin
- Database
    - PostgreSQL
- Dev, Build & Deploy
    - Docker, Buildx, DockerCompose, DockerHub

# 開發
預設使用以下port進行開發
- Database: 5432
- APP: 8000
- Web: 8001

## 開發文件
執行前請參閱各項服務的開發文件
- [Database](/database/README.md)
- [APP](/app/README.md)
- [Web](/web/README.md)
- [WebServer](/webserver/README.md)

## 使用DockerCompose開發
```
docker-compose up
```

# 建置

## 快速建置
```
make build
```

## 手動建置
### 使用Buildx建置並上傳多平台映像檔
```sh
# Database
docker buildx build --push --rm \
    --platform linux/amd64,linux/arm64 \
    -t vincent87720/daymood-database:latest \
    -f database/Dockerfile .

# APP
docker buildx build --push --rm \
    --platform linux/amd64,linux/arm64 \
    -t vincent87720/daymood-app:latest \
    -f app/Dockerfile .

# Webserver
docker buildx build --push --rm \
    --platform linux/amd64,linux/arm64 \
    -t vincent87720/daymood-webserver:latest \
    -f webserver/Dockerfile .
```