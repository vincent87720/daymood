run:
	docker-compose up

build: builddb buildapp buildserver

builddb:
	docker-compose up --detach database
	sleep 5
	rm -f database/initdb.d/01-init.sql
	docker exec -i daymood-database-1 pg_dump -U daymood --schema-only daymood >> database/initdb.d/01-init.sql
	docker-compose stop database
	docker buildx build --push --rm --platform linux/amd64,linux/arm64 -t vincent87720/daymood-database:latest -f database/Dockerfile .

buildapp:
	docker buildx build --push --rm --platform linux/amd64,linux/arm64 -t vincent87720/daymood-app:latest -f app/Dockerfile .

buildserver:
	docker buildx build --push --rm --platform linux/amd64,linux/arm64 -t vincent87720/daymood-webserver:latest -f webserver/Dockerfile .

deploy:
	docker run -itd \
		--name daymood-database \
		-p 5432:5432 \
		--network bridge-net \
		-v "$(PWD)/database/postgres:/var/lib/postgresql/data" \
		-e POSTGRES_USER="daymood"\
		-e POSTGRES_PASSWORD="daymood"\
		vincent87720/daymood-database
	docker run -itd \
		--name daymood-app \
		-p 8000:8000 \
		--network bridge-net \
		-e APP_MODE="PROD"\
		-e APP_HOST=""\
		-e APP_PORT="8000" \
		-e DB_HOSTNAME="daymood-database" \
		-e DB_DATABASE="daymood" \
		-e DB_USERNAME="daymood" \
		-e DB_PASSWORD="daymood" \
		-e SESSION_SECRET="daymood" \
		vincent87720/daymood-app
	docker run -itd \
		--name daymood-webserver \
		-p 80:80 \
		--network bridge-net \
		-e SERVER_HOST="0.0.0.0"\
		-e APP_HOST="daymood-app"\
		-e APP_PORT="8000"\
		vincent87720/daymood-webserver