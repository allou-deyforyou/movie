.PHONY: build clean deploy

build:
	bash
	# env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/film_latest handler/film_latest/main.go

clean:
	rm -rf ./bin ./vendor Gopkg.lock

deploy: clean build
	sls deploy --verbose
