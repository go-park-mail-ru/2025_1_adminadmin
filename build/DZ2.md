# ДЗ 2 Администрирование СУБД

Все данные по ДЗ2 дисциплины СУБД находятся в директории `build/sql`

## 1.	Безопасность сервера СУБД
Работа с БД через сервисную учетную запись
Создан скрипт создания сервисного пользователя app_user, который имеет минимально необходимые права для работы приложения:

Права только на SELECT, INSERT, UPDATE и DELETE в рамках конкретных таблиц.
Нет прав на создание объектов БД или управление пользователем.
Все действия выполняются в схеме public.
Скрипт находится по пути:

```
build/sql/scripts/app_user.sql
```

## 2.	Защита от SQL Injections
Валидация входных данных
Проверка корректности входящих данных производится на уровне delivery. Пример проверки номера телефона:

```go
func isValidPhone(phone string) bool {
	if len(phone) < minPhoneLength || len(phone) > maxPhoneLength {
		return false
	}
	for _, r := range phone {
		if r < '0' || r > '9' {
			return false
		}
	}
	return true
}
```

Экранирование спецсимволов
Используется функция ValidateInputText() из утилитарного пакета escapingutil:

```go
func ValidateInputText(text string) (string, error) {
    trimmed := strings.TrimSpace(html.EscapeString(text))
    if len(trimmed) == 0 {
        return "", errors.New("input text is empty or invalid")
    }
    return trimmed, nil
}
```

Пример использования:

```go
reviewText, err := escapingutil.ValidateInputText(input.ReviewText)
if err != nil {
    return nil, err
}
```

### Использование Prepared Statements
Работа с БД осуществляется через подготовленные запросы (Prepared Statements) в слое repository. Пример:

```go
func (r *RestaurantRepository) GetRecommendedProducts(ctx context.Context, productIDs []string, restaurantID string) ([]models.CartItem, error) {
    logger := log.GetLoggerFromContext(ctx).With(slog.String("func", log.GetFuncName()))

    var ids []uuid.UUID
    for _, pid := range productIDs {
        id, err := uuid.FromString(pid)
        if err != nil {
            logger.Error("invalid product ID", slog.String("error", err.Error()))
            return nil, fmt.Errorf("invalid product ID")
        }
        ids = append(ids, id)
    }

    restID, err := uuid.FromString(restaurantID)
    if err != nil {
        logger.Error("invalid restaurant ID", slog.String("error", err.Error()))
        return nil, fmt.Errorf("invalid restaurant ID")
    }

    rows, err := r.db.Query(ctx, getRecommendsByPtoducts, pq.Array(ids), restID)
    if err != nil {
        logger.Error("failed to execute query", slog.String("error", err.Error()))
        return nil, fmt.Errorf("database query error: %w", err)
    }
    defer rows.Close()

    var result []models.CartItem
    for rows.Next() {
        var item models.CartItem
        err := rows.Scan(
            &item.Id,
            &item.Name,
            &item.Price,
            &item.ImageURL,
            &item.Weight,
        )
        if err != nil {
            logger.Error("failed to scan row", slog.String("error", err.Error()))
            return nil, fmt.Errorf("row scan error: %w", err)
        }
        item.Amount = 1
        result = append(result, item)
    }

    return result, nil
}
```

## 3.   Настройка параметров сервера и клиента
Файл конфигурации PostgreSQL расположен по пути:

```
build/sql/postgresql.conf
```

### listen_addresses

```
listen_addresses = 'localhost'
```

### max_connections

```
max_connections = 100
```

### pool_size в приложении
На стороне приложения реализован пул соединений:

```go
connStr := os.Getenv("POSTGRES_CONN")
config, err := pgxpool.ParseConfig(connStr)
if err != nil {
    return nil, fmt.Errorf("failed to parse connection string: %w", err)
}
config.MaxConns = 20                             
config.MinConns = 10                             
config.MaxConnLifetime = time.Minute * 5         
config.HealthCheckPeriod = time.Minute           
config.MaxConnIdleTime = time.Minute * 2
```

## 4.   Таймауты

```
statement_timeout = 30s
lock_timeout = 5s
```

## 5.   pg_stat_statements

```
pg_stat_statements.max = 10000
pg_stat_statements.track = all
pg_stat_statements.save = on
```

## 6.   pg_stat_statements

```
auto_explain.log_min_duration = 150
auto_explain.log_analyze = true
auto_explain.log_buffers = true
auto_explain.log_verbose = true
auto_explain.log_timing = true
```

## 7.   Логгирование медленных запросов в формате PGBadger
Логгирование включено:

```
logging_collector = on
log_directory = 'log'
log_filename = 'postgresql-%Y-%m-%d_%H%M%S.log'
log_rotation_age = 1d
log_min_duration_statement = 150
log_format = 'csvlog'
log_line_prefix = '%t [%p]: [%l-1] user=%u,db=%d '
```

### Для запуска PGBadger

```
docker exec -it restaurant_db pgbadger /var/lib/postgresql/data/log/*.log -o report.html
docker cp restaurant_db:/report.html ./report.html
```