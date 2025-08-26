# Примеры API запросов

Этот файл содержит примеры использования всех API endpoints.

## Предварительные требования

1. Запустите приложение: `go run main.go`
2. Сервер будет доступен по адресу: `http://localhost:8080`

## 1. Создание пользователя

```bash
curl -X POST http://localhost:8080/api/v1/users/ \
  -H "Content-Type: application/json" \
  -d '{
    "id": "user-2",
    "email": "john.doe@example.com",
    "name": "John Doe"
  }'
```

**Ожидаемый ответ:**
```json
{
  "id": "user-2",
  "email": "john.doe@example.com",
  "name": "John Doe",
  "created_at": "2024-01-15T10:30:00Z"
}
```

## 2. Создание банковского счета

```bash
curl -X POST http://localhost:8080/api/v1/accounts/ \
  -H "Content-Type: application/json" \
  -d '{
    "user_id": "user-2",
    "currency": "USD"
  }'
```

**Ожидаемый ответ:**
```json
{
  "id": "generated-uuid",
  "user_id": "user-2",
  "balance": 0,
  "currency": "USD",
  "created_at": "2024-01-15T10:30:00Z",
  "updated_at": "2024-01-15T10:30:00Z"
}
```

## 3. Создание второго счета в другой валюте

```bash
curl -X POST http://localhost:8080/api/v1/accounts/ \
  -H "Content-Type: application/json" \
  -d '{
    "user_id": "user-2",
    "currency": "EUR"
  }'
```

## 4. Пополнение счета

```bash
curl -X POST http://localhost:8080/api/v1/transactions/deposit \
  -H "Content-Type: application/json" \
  -d '{
    "account_id": "account-id-from-step-2",
    "amount": 1000.0,
    "description": "Initial deposit"
  }'
```

**Ожидаемый ответ:**
```json
{
  "id": "generated-uuid",
  "from_account": "",
  "to_account": "account-id-from-step-2",
  "amount": 1000,
  "type": "deposit",
  "status": "completed",
  "description": "Initial deposit",
  "created_at": "2024-01-15T10:30:00Z",
  "updated_at": "2024-01-15T10:30:00Z"
}
```

## 5. Создание приветственного бонуса

```bash
curl -X POST http://localhost:8080/api/v1/bonuses/welcome \
  -H "Content-Type: application/json" \
  -d '{
    "user_id": "user-2",
    "amount": 50.0
  }'
```

**Ожидаемый ответ:**
```json
{
  "id": "generated-uuid",
  "user_id": "user-2",
  "type": "welcome",
  "amount": 50,
  "status": "active",
  "expires_at": "2024-02-14T10:30:00Z",
  "created_at": "2024-01-15T10:30:00Z"
}
```

## 6. Использование бонуса

```bash
curl -X POST http://localhost:8080/api/v1/bonuses/use \
  -H "Content-Type: application/json" \
  -d '{
    "bonus_id": "bonus-id-from-step-5",
    "account_id": "account-id-from-step-2"
  }'
```

**Ожидаемый ответ:**
```json
{
  "message": "Bonus used successfully"
}
```

## 7. Создание перевода между счетами

```bash
curl -X POST http://localhost:8080/api/v1/transactions/transfer \
  -H "Content-Type: application/json" \
  -d '{
    "from_account": "account-id-from-step-2",
    "to_account": "account-id-from-step-3",
    "amount": 100.0,
    "description": "Transfer to EUR account"
  }'
```

**Ожидаемый ответ:**
```json
{
  "id": "generated-uuid",
  "from_account": "account-id-from-step-2",
  "to_account": "account-id-from-step-3",
  "amount": 100,
  "type": "transfer",
  "status": "completed",
  "description": "Transfer to EUR account",
  "created_at": "2024-01-15T10:30:00Z",
  "updated_at": "2024-01-15T10:30:00Z"
}
```

## 8. Просмотр информации о счете

```bash
curl http://localhost:8080/api/v1/accounts/account-id-from-step-2
```

## 9. Просмотр сводки по счету

```bash
curl http://localhost:8080/api/v1/accounts/account-id-from-step-2/summary
```

**Ожидаемый ответ:**
```json
{
  "account_id": "account-id-from-step-2",
  "user_id": "user-2",
  "balance": 950,
  "currency": "USD",
  "created_at": "2024-01-15T10:30:00Z",
  "updated_at": "2024-01-15T10:30:00Z",
  "total_transactions": 2
}
```

## 10. Просмотр истории транзакций

```bash
curl http://localhost:8080/api/v1/transactions/account/account-id-from-step-2
```

## 11. Просмотр активных бонусов пользователя

```bash
curl http://localhost:8080/api/v1/bonuses/user/user-2
```

## 12. Создание списания со счета

```bash
curl -X POST http://localhost:8080/api/v1/transactions/withdrawal \
  -H "Content-Type: application/json" \
  -d '{
    "account_id": "account-id-from-step-2",
    "amount": 50.0,
    "description": "ATM withdrawal"
  }'
```

## 13. Обновление информации о пользователе

```bash
curl -X PUT http://localhost:8080/api/v1/users/user-2 \
  -H "Content-Type: application/json" \
  -d '{
    "id": "user-2",
    "email": "john.doe.updated@example.com",
    "name": "John Doe Updated"
  }'
```

## 14. Удаление пустого счета

```bash
curl -X DELETE http://localhost:8080/api/v1/accounts/account-id-if-empty
```

## 15. Проверка здоровья сервиса

```bash
curl http://localhost:8080/health
```

**Ожидаемый ответ:**
```json
{
  "status": "ok",
  "service": "banking-api",
  "version": "1.0.0"
}
```

## Полный сценарий работы

1. **Создайте пользователя** (шаг 1)
2. **Создайте два счета** в разных валютах (шаги 2-3)
3. **Пополните USD счет** (шаг 4)
4. **Создайте приветственный бонус** (шаг 5)
5. **Используйте бонус** для пополнения USD счета (шаг 6)
6. **Переведите средства** с USD на EUR счет (шаг 7)
7. **Просмотрите результаты** (шаги 8-11)

## Обработка ошибок

### Недостаточно средств
```bash
curl -X POST http://localhost:8080/api/v1/transactions/transfer \
  -H "Content-Type: application/json" \
  -d '{
    "from_account": "account-id",
    "to_account": "another-account-id",
    "amount": 999999.0,
    "description": "Large transfer"
  }'
```

**Ожидаемый ответ:**
```json
{
  "error": "insufficient funds"
}
```

### Неподдерживаемая валюта
```bash
curl -X POST http://localhost:8080/api/v1/accounts/ \
  -H "Content-Type: application/json" \
  -d '{
    "user_id": "user-2",
    "currency": "GBP"
  }'
```

**Ожидаемый ответ:**
```json
{
  "error": "unsupported currency"
}
```

### Несуществующий счет
```bash
curl http://localhost:8080/api/v1/accounts/non-existent-id
```

**Ожидаемый ответ:**
```json
{
  "error": "account not found"
}
```

## Тестирование с помощью Postman

1. Импортируйте коллекцию в Postman
2. Создайте переменную окружения `base_url` со значением `http://localhost:8080`
3. Используйте `{{base_url}}` в URL запросов
4. Запустите запросы по порядку

## Тестирование с помощью curl в цикле

```bash
# Создание нескольких пользователей
for i in {1..5}; do
  curl -X POST http://localhost:8080/api/v1/users/ \
    -H "Content-Type: application/json" \
    -d "{
      \"id\": \"user-$i\",
      \"email\": \"user$i@example.com\",
      \"name\": \"User $i\"
    }"
  echo ""
done
```

## Мониторинг и логи

Приложение логирует все HTTP запросы. Для просмотра логов:

```bash
# Если запущено через Docker
docker logs banking-api

# Если запущено локально
# Логи выводятся в stdout
```
