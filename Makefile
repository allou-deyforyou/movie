.PHONY: build clean deploy

build:
	# dep ensure -v (deyforyou)
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/film_article handler/film_article/main.go
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/film_latest handler/film_latest/main.go
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/film_search handler/film_search/main.go

	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/serie_article handler/serie_article/main.go
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/serie_latest handler/serie_latest/main.go
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/serie_search handler/serie_search/main.go

	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/manga_article handler/manga_article/main.go
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/manga_latest handler/manga_latest/main.go
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/manga_search handler/manga_search/main.go


clean:
	rm -rf ./bin ./vendor Gopkg.lock

deploy: clean build
	sls deploy --verbose
