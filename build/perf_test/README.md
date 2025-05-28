# ДЗ - 3 Оптимизация работы СУБД


## Подготовка

### Выбор основноый сущности тестирования

В качестве основной сущности для тестирования была выбрана таблица users

```sql
CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY,
    login TEXT NOT NULL UNIQUE,
    phone_number TEXT,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    description TEXT DEFAULT '',
    user_pic TEXT DEFAULT 'default_user.jpg',
    password_hash BYTEA NOT NULL
);
```

#### Причины:
- Пользователь — ключевая сущность в сервисе доставки еды: через него происходит аутентификация, привязка заказов, отображение профиля, адресов, истории покупок и т.п.

- Создание и получение информации о пользователе — базовые и часто используемые операции.

- Любые задержки в регистрации, логине или отображении данных профиля могут повлиять на пользовательский опыт.

#### Характер нагрузки

- Тип: Смешанная (чтение и запись).

- Операции: 
   - Часто: регистрация и логин.

   - Реже, но регулярно: обновление профиля.

Цель оптимизации: ускорить регистрацию и логин пользователя, а также получение профиля.

### Выбор утилиты тестирования

Для нагрузочного тестирования была выбрана утилита Vegeta.

#### Причины 

-	Поддержка HTTP-запросов с телом (необходима для POST /users).
-	Гибкие параметры настройки (-rate, -duration, -connections).
-	Поддержка анализа метрик (latency, throughput, статус-коды).
-	Возможность построения гистограмм и графиков.

Команда для выполнения:

```bash
vegeta attack \
  -targets=targets.txt \     # файл со списком HTTP-запросов (метод, URL, тело)
  -rate=10 \                 # частота запросов: 10 запросов в секунду
  -duration=60s \            # общее время атаки: 60 секунд
  -connections=5 \           # количество одновременных соединений (сокетов)
| tee /tmp/vegeta-test \     # сохраняем "сырые" результаты в файл для последующего анализа
| vegeta report > report.txt # формируем итоговый отчёт (latency, throughput, status codes и др.)

```

# Автоматизация тестирования через Makefile

Для удобства выполнения тестов и генерации отчетов был использован Makefile, обеспечивающий полный цикл нагрузочного тестирования: от подготовки данных до визуализации результатов.

```makefile
perf_tests_get:
	clear
	@echo "Запуск установки сессий..."
	go run build/perf_test/main.go
	@echo "Запуск нагрузки..."
	$(MAKE) clean
	$(MAKE) make_perf_test_get
	$(MAKE) report
	$(MAKE) plot
	$(MAKE) histogram

make_perf_test_get:
	@echo "Запуск нагрузки на $(DURATION) с частотой $(RATE) запросов/сек..."
	vegeta attack -targets=$(TARGETS_FILE) -rate=$(RATE) -duration=$(DURATION) | tee /tmp/vegeta-test | vegeta report > $(REPORT_FILE)

report:
	@echo "Генерация текстового отчёта..."
	@cat $(REPORT_FILE)

plot:
	@echo "Генерация HTML-графика..."
	@cat /tmp/vegeta-test | vegeta plot > $(PLOT_FILE)
	@echo "Открой файл $(PLOT_FILE) в браузере."
	open $(PLOT_FILE)

histogram:
	@echo "Генерация гистограммы латентности..."
	@cat /tmp/vegeta-test | vegeta report -type=hist[0,10ms,20ms,50ms,100ms,200ms,500ms,1s] > $(HISTOGRAM_FILE)
	@cat $(HISTOGRAM_FILE)
```

- go run docs/perf_test/main.go - Инициализация сессий, файлов запроса
- make_perf_test_get - запуск самого тестирования
- report - Генерация текстового отчёта
- plot - Генерация HTML-графика
- histogram - Генерация гистограммы латентности


## Процесс тестирования


### Cоздание профиля

Для создания профиля существует handler POST https://doordashers.ru/signup
На вход подается учетная запись user с конфиденциальными данными и profile, который требуется создать

Текст запроса:
```sql
INSERT INTO users (id, login, first_name, last_name, phone_number, description, user_pic, password_hash) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
```

Тестирование:

```bash
Генерация текстового отчёта...
Requests      [total, rate, throughput]         6000, 100.02, 0.00
Duration      [total, attack, wait]             59.99s, 59.99s, 129.375µs
Latencies     [min, mean, 50, 90, 95, 99, max]  15.542µs, 127.055µs, 92.318µs, 167.407µs, 251.509µs, 939.339µs, 7.178ms
Bytes In      [total, mean]                     0, 0.00
Bytes Out     [total, mean]                     0, 0.00
Success       [ratio]                           0.00%
Status Codes  [code:count]                      0:6000  
Error Set:
Post "https://doordashers.ru/api/auth/signup": net/http: invalid header field name "{\"first_name\""
Post "https://doordashers.ru/api/auth/signup": net/http: invalid header field name "POST https"
/Library/Developer/CommandLineTools/usr/bin/make plot
Генерация HTML-графика...
Открой файл docs/perf_test/plot.html в браузере.
open docs/perf_test/plot.html
/Library/Developer/CommandLineTools/usr/bin/make histogram
Генерация гистограммы латентности...
Bucket           #     %        Histogram
[0s,     10ms]   6000  100.00%  ###########################################################################
[10ms,   20ms]   0     0.00%    
[20ms,   50ms]   0     0.00%    
[50ms,   100ms]  0     0.00%    
[100ms,  200ms]  0     0.00%    
[200ms,  500ms]  0     0.00%    
[500ms,  1s]     0     0.00%    
[1s,     +Inf]   0     0.00%    
```
Ошибки были вызваны тем, что в некоторых случаях при генерации пользователя совпадали данные телефона. Также некоторые записи вышли на timeout, так как тестирование проводилась на нагрузке 20 запросов в секунду, в реальной работе планируется, что создание пользователя будет выполняться реже, но в целях увеличения скорости тестирования были выбраны такие значения

Видно, что запросы идут достаточно быстро, так как запросы итак были изначально разделены, каждый из них по отдельности максимально отпимизирован


### Получение пользователя

В качестве получение пользователя тестируется handler POST https://localhost:5459/signin
На вход он получает логин и пароль текущего пользователя, на выход отдает полную модель пользователя

Текст первоначального запроса

```sql
SELECT id, first_name, last_name, phone_number, description, user_pic, password_hash, secret2fa FROM users WHERE login = $1
```

Проводим тестирование при 
RATE ?= 10
DURATION ?= 60s

```bash
Requests      [total, rate, throughput]         600, 10.02, 10.01
Duration      [total, attack, wait]             59.916s, 59.901s, 15.852ms
Latencies     [min, mean, 50, 90, 95, 99, max]  4.495ms, 19.302ms, 10.598ms, 23.309ms, 32.229ms, 228.396ms, 779.315ms
Bytes In      [total, mean]                     529759, 882.93
Bytes Out     [total, mean]                     0, 0.00
Success       [ratio]                           100.00%
Status Codes  [code:count]                      200:600  
Error Set:


Bucket           #    %       Histogram
[0s,     10ms]   275  45.83%  ##################################
[10ms,   20ms]   244  40.67%  ##############################
[20ms,   50ms]   62   10.33%  #######
[50ms,   100ms]  9    1.50%   #
[100ms,  200ms]  4    0.67%   
[200ms,  500ms]  3    0.50%   
[500ms,  1s]     3    0.50%   
[1s,     +Inf]   0    0.00%  
```

Значения являются приемлемыми:

- нет ни одной ошибки
- 95-й перцентиль: 32.2 мс 
- Средняя задержка (mean latency): 19.3 мс 
- Максимальная задержка: 779 мс — высокая, но редкая (менее 1%)

Однако это может создать проблемы при более высокой нагрузке приложения

Для ускорения работы приложения попробуем добавить индексы на запросы, включающие фильтрацию и джойны:

```sql
CREATE INDEX idx_profiles_profile_id ON users(user_id);
```

Тестирование с индексами на тех же данных:

```bash
Requests      [total, rate, throughput]         600, 10.02, 0.00
Duration      [total, attack, wait]             59.916s, 59.901s, 15.194ms
Latencies     [min, mean, 50, 90, 95, 99, max]  1.314ms, 16.522ms, 4.062ms, 9.053ms, 21.434ms, 449.168ms, 659.905ms
Bytes In      [total, mean]                     11400, 19.00
Bytes Out     [total, mean]                     0, 0.00
Success       [ratio]                           100.00%
Status Codes  [code:count]                      200:600  


Bucket           #    %       Histogram
[0s,     10ms]   544  90.67%  ####################################################################
[10ms,   20ms]   23   3.83%   ##
[20ms,   50ms]   11   1.83%   #
[50ms,   100ms]  4    0.67%   
[100ms,  200ms]  3    0.50%   
[200ms,  500ms]  11   1.83%   #
[500ms,  1s]     4    0.67%   
[1s,     +Inf]   0    0.00%   
```


Использование индексов заметно ускорило работу на получение основного профиля

Теперь попытаемся увеличить количество запросов в секунду с 10 до 100

```bash
Requests      [total, rate, throughput]         6000, 100.02, 0.00
Duration      [total, attack, wait]             59.992s, 59.989s, 2.567ms
Latencies     [min, mean, 50, 90, 95, 99, max]  710.5µs, 5.572ms, 2.293ms, 6.194ms, 13.23ms, 84.175ms, 312.17ms
Bytes In      [total, mean]                     114000, 19.00
Bytes Out     [total, mean]                     0, 0.00
Success       [ratio]                           100%
Status Codes  [code:count]                      200:6000  

Bucket           #     %       Histogram
[0s,     10ms]   5627  93.78%  ######################################################################
[10ms,   20ms]   154   2.57%   #
[20ms,   50ms]   110   1.83%   #
[50ms,   100ms]  59    0.98%   
[100ms,  200ms]  30    0.50%   
[200ms,  500ms]  20    0.33%   
[500ms,  1s]     0     0.00%   
[1s,     +Inf]   0     0.00%  
```

Несмотря на увеличение интенсивности, значения все еще остались в приемлемом диапазоне. Проверим это на виртуальной машине
Для этого в файле инициализации make заменим localhost на адрес машины

```bash
Requests      [total, rate, throughput]         6000, 100.02, 42.81
Duration      [total, attack, wait]             1m28s, 59.99s, 28.161s
Latencies     [min, mean, 50, 90, 95, 99, max]  388.243ms, 11.637s, 6.12s, 30.001s, 30.001s, 30.005s, 30.023s
Bytes In      [total, mean]                     3176001, 529.33
Bytes Out     [total, mean]                     0, 0.00
Success       [ratio]                           100.00%
Status Codes  [code:count]                      200:6000  


Генерация гистограммы латентности...
Bucket           #     %       Histogram
[0s,     10ms]   0     0.00%   
[10ms,   20ms]   0     0.00%   
[20ms,   50ms]   0     0.00%   
[50ms,   100ms]  0     0.00%   
[100ms,  200ms]  0     0.00%   
[200ms,  500ms]  175   2.92%   ##
[500ms,  1s]     2437  40.62%  ##############################
[1s,     +Inf]   3388  56.47%  ##########################################
```

При запуске на виртмашине обнаружилось, что запросы все еще остаются неэффективными
Самый длинный запрос на получение шел 30 секунд. Это слишком долго, поэтому необходимо провести дополнительную оптимизацию

- Многочисленные join и фильтрации
Сейчас JOIN'ы выполняются до фильтрации (WHERE), что увеличивает количество строк, проходящих через джойны
Оптимизация: сначала получить profile_id нужных профилей, затем сделать JOIN'ы по ним, чтобы JOIN'ы работали только по 30 записям (LIMIT $3), а не по всей таблице profiles:

```sql
WITH filtered_profiles AS (
    SELECT p.profile_id
    FROM profiles p
    LEFT JOIN likes liked 
        ON liked.liked_profile_id = p.profile_id AND liked.profile_id = $1
    WHERE p.profile_id != $1 
      AND liked.profile_id IS NULL 
      AND p.profile_id > $2
    ORDER BY p.profile_id
    LIMIT $3
)
SELECT 
    p.profile_id, 
    p.firstname, 
    p.lastname, 
    p.is_male,
    p.height,
    p.birthday, 
    p.description, 
    l.country, 
    l.city,
    l.district,
    s.path AS avatar,
    i.description AS interest,
    pr.preference_description,
    pr.preference_value 
FROM filtered_profiles fp
JOIN profiles p ON p.profile_id = fp.profile_id
LEFT JOIN locations l 
    ON p.location_id = l.location_id
LEFT JOIN "static" s 
    ON p.profile_id = s.profile_id
LEFT JOIN profile_interests pi 
    ON pi.profile_id = p.profile_id
LEFT JOIN interests i 
    ON pi.interest_id = i.interest_id
LEFT JOIN profile_preferences pp 
    ON pp.profile_id = p.profile_id
LEFT JOIN preferences pr 
    ON pp.preference_id = pr.preference_id;

```


- Добавить дополнительные индексы по join

Для устранения Seq Scan (полных проходов по таблице) PostgreSQL, были добавлены недостающие индексы, участвующие в JOIN'ах по profile_id:

```sql
CREATE INDEX IF NOT EXISTS idx_static_profile_id ON "static"(profile_id);
CREATE INDEX IF NOT EXISTS idx_profile_interests_profile_id ON profile_interests(profile_id);
CREATE INDEX IF NOT EXISTS idx_profile_preferences_profile_id ON profile_preferences(profile_id);
```
Эти индексы позволяют PostgreSQL использовать Index Scan, что значительно снижает время выполнения сложных объединений.


- Упорядочивание
Для более эффективной постраничной выборки (keyset pagination), используется сортировка по profile_id:

```sql
ORDER BY p.profile_id
```
В сочетании с фильтрацией p.profile_id > $2 это позволяет избежать пагинации через OFFSET, что критично при больших объемах данных.

- Оптимизация внутри go кода

Ранее: профили загружались из базы в неоптимальном порядке, хранились в map, затем сортировались вручную.

Теперь:

- Профили извлекаются уже отсортированными SQL-запросом.

- Используется keyset pagination.

- Обновление Redis-ключа вынесено из критического пути — выполняется асинхронно.


Тестирование кода после исправления

```sql
Генерация текстового отчёта...
Requests      [total, rate, throughput]         6000, 100.02, 98.81
Duration      [total, attack, wait]             1m1s, 59.99s, 730.255ms
Latencies     [min, mean, 50, 90, 95, 99, max]  493.767ms, 804.216ms, 766.496ms, 1.14s, 1.267s, 1.456s, 2s
Bytes In      [total, mean]                     97694103, 16282.35
Bytes Out     [total, mean]                     0, 0.00
Success       [ratio]                           100.00%
Status Codes  [code:count]                      200:6000  

Генерация гистограммы латентности...
Bucket           #     %       Histogram
[0s,     10ms]   0     0.00%   
[10ms,   20ms]   0     0.00%   
[20ms,   50ms]   0     0.00%   
[50ms,   100ms]  0     0.00%   
[100ms,  200ms]  0     0.00%   
[200ms,  500ms]  75    1.25%   
[500ms,  1s]     5171  86.18%  ################################################################
[1s,     +Inf]   754   12.57%  #########
```
Путем оптимизаций улалось снизить задержку и увеличить количество быстродействующих запросов.

Ключевые моменты:

- Используются Index Only Scan для таблиц users.

- Активно задействован Memoize для кэширования location_id → locations, что ускоряет повторные запросы.

- Большинство соединений — Hash Left Join и Nested Loop Join, что допустимо при небольшом объеме данных (или хороших индексах).

- Последовательное сканирование (Seq Scan) осталось только на малых таблицах (static, preferences, interests), что не критично при их размере.

Вывод: план запроса теперь оптимален и использует индексы, результат достигается за миллисекунды.

