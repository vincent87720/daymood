DARWIN := darwin
PACKAGE := github.com/vincent87720/daymood.backend/cmd/daymood.backend

build: rmbin release cpui cpyml cpmacinfo cpbin

run:
	go run ./cmd/daymood/main.go

rmbin:
	rm -rf ./bin/* ./test/*

cpui:
	cp -r ./daymoodui ./bin/$(DARWIN)/daymood.app/Contents/MacOS

cpyml:
	cp ./settings.production.yaml ./bin/$(DARWIN)/daymood.app/Contents/MacOS

cpmacinfo:
	cp ./box.pdf ./bin/$(DARWIN)/daymood.app/Contents/MacOS
	cp ./box.icns ./bin/$(DARWIN)/daymood.app/Contents/Resources
	cp ./info.plist ./bin/$(DARWIN)/daymood.app/Contents


cpbin:
	cp -r ./bin/* ./test

swag:
	cd ./cmd/daymood; ../../tools/swag init

##########BUILD##########
.PHONY: builddarwin
builddarwin:
	mkdir -p bin/$(DARWIN)/daymood.app/Contents/MacOS
	mkdir -p bin/$(DARWIN)/daymood.app/Contents/Resources
	GOOS=$(DARWIN) GOARCH=amd64 go build -o bin/$(DARWIN)/daymood.app/Contents/MacOS/daymood $(PACKAGE)

.PHONY: release
release: builddarwin