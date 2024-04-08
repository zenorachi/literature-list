IMAGE=literature-list

build:
	docker build -t $(IMAGE) .

run: build
	docker run --publish 8080:8080 $(IMAGE)

lint:
	golangci-lint run

swag:
	swag init -g ./internal/app/app.go
