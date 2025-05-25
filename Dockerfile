# Используем базовый образ для сборки
FROM golang:1.23 AS builder

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем go.mod и go.sum
COPY . .
COPY go.mod go.sum ./

# Загружаем зависимости
RUN apk add --no-cache git
RUN go mod download

# Копируем весь проект
COPY . .

# Сборка проекта и установка прав на исполнимость
RUN go build -o /app/adtime-bot ./cmd/adtime/main.go && chmod +x /app/adtime-bot

# Минимальный образ для запуска
FROM gcr.io/distroless/static-debian11

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем скомпилированный файл
COPY --from=builder /app/pompon-bot /app/
COPY --from=builder /app/migrations /app/migrations
# Открываем порт для бота (опционально)
EXPOSE 8080

# Запускаем приложение
CMD ["/app/adtime-bot"]
