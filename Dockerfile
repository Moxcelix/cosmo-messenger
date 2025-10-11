FROM golang:1.23-alpine as BUILDER

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем файлы модулей и загружаем зависимости
COPY go.mod go.sum ./
RUN go mod download

# Копируем исходный код
COPY . .

# Собираем приложение
#RUN go build -o main .

# Открываем порт (если приложение использует сеть)
EXPOSE 8080

# Запускаем приложение
#CMD ["./main"]
CMD ["go", "run", "main.go"]