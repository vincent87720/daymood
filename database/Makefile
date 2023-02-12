dockerrun:
	docker run -d \
		--name daymood-database \
		-p 5432:5432 \
		--network bridge-net \
		-v "$(PWD)/postgres:/var/lib/postgresql/data" \
		-e POSTGRES_PASSWORD="admin"\
		postgres:15.1

dump:
	rm init.sql |\
	docker run -d \
		--name daymood-build-database \
		--network bridge-net \
		-v "$(PWD)/postgres:/var/lib/postgresql/data" \
		-e POSTGRES_PASSWORD="admin"\
		postgres:15.1

# Waiting for the container to be ready
	sleep 2

# dump schema
	docker exec -i daymood-build-database pg_dump -U daymood --schema-only daymood >> init.sql

	docker stop daymood-build-database
	
	docker rm daymood-build-database