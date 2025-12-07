# Переменные для удобства
BINARY_NAME=main
LOCAL_BIN:=$(CURDIR)/bin

# .PHONY указывает, что это команды, а не файлы
.PHONY: all build test clean run lint deps

all: build

# Сборка приложения
build:
	go build -o $(BINARY_NAME) cmd/app/main.go

# Запуск (для локальной разработки)
run:
	CONFIG_PATH=./config/local.yaml go run cmd/app/main.go

# Запуск тестов (с проверкой на гонку потоков - race condition)
test:
	go test -v -race -timeout 30s ./...

# Запуск линтера
lint:
	golangci-lint run

# Установка зависимостей и чистка go.mod
deps:
	go mod download
	go mod tidy

# Очистка артефактов
clean:
	go clean
	rm -f $(BINARY_NAME)

# Запуск в докере (на будущее)
docker-build:
	docker-compose build
docker-up:
	docker-compose up -d
docker-down:
	docker-compose down