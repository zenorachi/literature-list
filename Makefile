build:
	go mod download && CGO_ENABLED=0 GOOS=linux go build -o ./.bin/app ./cmd/app/main.go

run: build
	docker-compose up --remove-orphans

rebuild: build
	docker-compose up --remove-orphans --build

stop:
	docker-compose down

lint:
	golangci-lint run

swag:
	swag init -g ./internal/app/app.go
