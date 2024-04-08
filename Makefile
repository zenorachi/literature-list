IMAGE=literature-list

.PHONY: build
build:
	docker build -t $(IMAGE) .

.PHONY: run
run: build
	docker run --publish 8080:8080 $(IMAGE)

.PHONY: stop
stop:
	docker stop $$(docker ps --filter ancestor=$(IMAGE) -q)

.PHONY: stop
lint:
	golangci-lint run

.PHONY: swagger
swagger:
	swag init -g ./internal/app/app.go
