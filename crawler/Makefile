run:
	go run cmd/crawler/main.go

build:
	go mod vendor && go build -o crawler cmd/crawler/main.go

start:
	./crawler

tidy:
	go mod tidy

clear_redis:
	go run scripts/reset_redis.go
