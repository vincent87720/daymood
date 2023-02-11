run:
	npm run serve

dockerrun:
	docker run -it \
		--name daymood-web\
		-p 8001:8001\
		--network bridge-net \
		-w /web\
		-v "$(PWD):/web"\
		node:19.2.0 npm run serve -- --port "8001"
build:
	npm run build