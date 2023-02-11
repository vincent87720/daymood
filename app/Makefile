DARWIN := darwin
LINUX := linux
PACKAGE := github.com/vincent87720/daymood/app/cmd/daymood

run:
	go run ./cmd/daymood/main.go

dockerrun:
	docker run -it \
		--name daymood-app\
		-p 8000:8000\
		--network bridge-net \
		-w /app\
		-v "$(PWD):/app"\
		golang:1.19 go run $(PACKAGE)

build: rmbin release cpyml cpbin

rmbin:
	rm -rf ./bin/* ./test/*

cpyml:
	cp ./settings.yaml ./bin/$(DARWIN)

cpbin:
	cp -r ./bin/* ./test

#BUILD
.PHONY: builddarwin
builddarwin:
	mkdir -p bin/$(DARWIN)
	mkdir -p bin/$(LINUX)
	GOOS=$(DARWIN) GOARCH=amd64 go build -o bin/$(DARWIN) $(PACKAGE)
	GOOS=$(LINUX) GOARCH=amd64 go build -o bin/$(LINUX) $(PACKAGE)

.PHONY: release
release: builddarwin