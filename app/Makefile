DARWIN := darwin
LINUX := linux
PACKAGE := github.com/vincent87720/daymood/app/cmd/daymood

build: rmbin release cpyml cpbin

run:
	go run ./cmd/daymood/main.go

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