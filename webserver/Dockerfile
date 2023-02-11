#在根目錄下將web資後夾複製到node環境的/web目錄中，執行"npm run build"指令打包前端網頁
FROM node:19.2.0 AS build
WORKDIR /web
COPY web /web
RUN npm run build

#將打包過後的前端網頁資料夾連同nginx的設定檔一同放入nginx的環境
#使用SERVER_HOST, APP_HOST及APP_PORT環境變數設定nginx設定檔內的數值
FROM nginx:1.23 AS release
COPY webserver/nginx.conf.template /etc/nginx/templates/default.conf.template
COPY --from=build /web/daymoodui /web
ENV SERVER_HOST=0.0.0.0 \
    APP_HOST=daymood-app \
    APP_PORT=8000