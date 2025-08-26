# Banking API (pet-project)

Минималистичный REST API на Go про счета, транзакции и бонусы. Писал его для демонстрации AQA/Go навыков: структура кода, тесты, простые бизнес-правила.

## Стек
- Go 1.21, Gin
- Testify (assert/mock)
- In-memory DB (для простоты)

## Как запустить
```bash
go mod download
go run main.go
# http://localhost:8080/health
```
Docker:
```bash
docker build -t banking-api .
docker run -p 8080:8080 banking-api
```
Makefile (удобно):
```bash
make run          # build+run
make test         # все тесты
make test-coverage
```

## Основные эндпоинты
- Health: GET `/health`
- Accounts: GET `/api/v1/accounts/:id`, POST `/api/v1/accounts/`
- Transactions: POST `/api/v1/transactions/{transfer|deposit|withdrawal}`
- Bonuses: POST `/api/v1/bonuses/{welcome|use}`
- Users: GET/POST/PUT/DELETE `/api/v1/users/...`

Примеры запросов в `examples/api-examples.md`.

## Идея домена (очень кратко)
- Перевод: проверка валюты и достаточности средств, обновление балансов, статуса транзакции.
- Депозит/Списание: изменение баланса и фиксация транзакции.
- Бонусы: приветственный и за транзакции, проверка статуса/срока, списание в баланс.

## Тесты
- Unit-тесты сервисов с моками `testify/mock`.
- Тесты in-memory хранилища, включая конкурентный доступ.
```bash
go test ./...
```

## Структура
```
internal/
  api/        # handlers + server
  services/   # бизнес-логика
  database/   # in-memory реализация
  models/     # модели
  config/     # конфиг (env)
```

## План расширений (если будет время)
- PostgreSQL вместо in-memory
- Swagger, авторизация, метрики

Автор: личный пет-проект для портфолио.
