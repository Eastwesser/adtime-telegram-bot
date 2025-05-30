# adTime tgbot on Golang
Это MVP проект для телеграм-бота команды adTime!

adTime занимается подарками и упаковками, оформлением различных индивидуальных и корпоративных подарочных упаковок,
а также проводит мастер-классы! 🎨

## Структура проекта:

```
pompon-bot/
├── cmd/                    // Точка входа в приложение
│   └── adtime/             // Основной бинарник бота
│       └── main.go         // Точка запуска
├── internal/               // Вся внутренняя логика приложения
│   ├── config/             // Конфигурация и база данных
│   │   ├── config.go       // Настройки приложения
│   │   └── database.go     // Подключение к PostgreSQL
│   ├── handlers/           // Обработчики Telegram команд
│   │   ├── about.go        // Команда /about
│   │   ├── catalog.go      // Команда /catalog
│   │   ├── order.go        // Команда /order
│   │   └── subscribe.go    // Команда /subscribe
│   ├── keyboards/          // Клавиатуры для бота
│   │   └── keyboards.go    // Генерация клавиатур
│   ├── models/             // Модели данных и работа с PostgreSQL
│   │   ├── order.go        // Модель заказов
│   │   ├── product.go      // Модель товаров
│   │   └── subscriber.go   // Модель подписчиков
│   ├── services/           // Логика работы приложения
│   │   ├── catalog_service.go   // Логика каталога
│   │   ├── notify_service.go    // Логика уведомлений
│   │   ├── order_service.go     // Логика заказов
│   │   └── subscribe_service.go // Логика подписок
│   └── utils/              // Утилиты и вспомогательные функции
│       └── helpers.go      // Разные утилиты
├── sql/                    // SQL-скрипты для базы данных
│   ├── schema.sql          // Схема базы данных
│   ├── seed.sql            // Тестовые данные
│   └── README.md           // Описание структуры базы
├── .env-template           // Шаблон для переменных окружения
├── .gitignore              // Игнорируемые файлы
├── docker-compose.yml      // Docker Compose для запуска PostgreSQL и бота
├── Dockerfile              // Dockerfile для сборки приложения
├── go.mod                  // Go модули
├── README.md               // Документация проекта
└── .env                    // Настройки окружения (сейчас в .gitignore, используйте шаблон .env-template)
```

## Команды бота:

- `/start` — Начало работы
- `/about` — Информация о магазине
- `/catalog` — Каталог товаров
- `/order` — Оформить заказ
- `/subscribe` — Подписаться на новости

## Запуск проекта

### 1. Установите Docker и Docker Compose
Убедитесь, что у вас установлены Docker и Docker Compose. 
[Инструкция по установке Docker](https://docs.docker.com/get-docker/).
### 2. Создайте файл `.env`
В корневой директории создайте файл `.env` со следующими данными:
```
BOT_TOKEN=ваш_токен_бота 
DATABASE_URL=postgres://postgres:password@db:5432/adtime
```
### 3. Запустите контейнеры
Выполните команду для запуска проекта:
```bash
docker-compose up --build -d
```
Бот будет запущен и подключится к базе данных PostgreSQL.

## Как запустить (без Docker):

1. Скопируйте `.env-template` в `.env` и заполните данные. Если нет `.env`, то создайте его.
2. Установите зависимости: `go mod tidy`.
3. Запустите PostgreSQL и примените миграции.
4. Запустите бота: `go run cmd/adtime/main.go`.