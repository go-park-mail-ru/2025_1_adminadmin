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

-	Поддержка HTTP-запросов с телом (необходима для POST /signup).
-	Гибкие параметры настройки (-rate, -duration, -connections).
-	Поддержка анализа метрик (latency, throughput, статус-коды).
-	Возможность построения гистограмм и графиков.

Команда для выполнения:

```bash
vegeta attack \
  -targets=signup-targets.txt \     # файл со списком HTTP-запросов (метод, URL, тело)
  -rate=10 \                 # частота запросов: 10 запросов в секунду
  -duration=60s \            # общее время атаки: 60 секунд
  -connections=5 \           # количество одновременных соединений (сокетов)
| tee /tmp/vegeta-test \     # сохраняем "сырые" результаты в файл для последующего анализа
| vegeta report > report.txt # формируем итоговый отчёт (latency, throughput, status codes и др.)

```

# Автоматизация тестирования через Makefile

Для удобства выполнения тестов и генерации отчетов был использован Makefile, обеспечивающий полный цикл нагрузочного тестирования: от подготовки данных до визуализации результатов.

```makefile
perf_tests_get_signup:
	clear
	go run build/perf_test/main.go
	@echo "Запуск нагрузки..."
	$(MAKE) clean
	$(MAKE) make_perf_test_get_signup
	$(MAKE) report
	$(MAKE) plot
	$(MAKE) histogram

perf_tests_get_signin:
	clear
	go run build/perf_test/main.go
	@echo "Запуск нагрузки..."
	$(MAKE) clean
	$(MAKE) make_perf_test_get_signin
	$(MAKE) report
	$(MAKE) plot
	$(MAKE) histogram

make_perf_test_get_signup:
	@echo "Запуск нагрузки на $(DURATION) с частотой $(RATE) запросов/сек..."
	vegeta attack -targets=$(TARGETS_FILE_SIGNUP) -rate=$(RATE) -duration=$(DURATION) | tee /tmp/vegeta-test | vegeta report > $(REPORT_FILE)

make_perf_test_get_signin:
	@echo "Запуск нагрузки на $(DURATION) с частотой $(RATE) запросов/сек..."
	vegeta attack -targets=$(TARGETS_FILE_SIGNIN) -rate=$(RATE) -duration=$(DURATION) | tee /tmp/vegeta-test | vegeta report > $(REPORT_FILE)


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

clear-target-files:
	rm docs/perf_test/signup-targets.txt docs/perf_test/auth-targets.txt
```

- go run docs/perf_test/main.go - Инициализация файлов запроса
- make perf_test_get_signin / perf_test_get_signup - запуск самого тестирования
- report - Генерация текстового отчёта
- plot - Генерация HTML-графика
- histogram - Генерация гистограммы латентности


## Процесс тестирования


### Cоздание профиля

Для создания профиля существует handler POST https://doordashers.ru/api/auth/signup
На вход подается учетная запись user с конфиденциальными данными и profile, который требуется создать

Текст запроса:
```sql
INSERT INTO users (id, login, first_name, last_name, phone_number, description, user_pic, password_hash) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
```

Тестирование:

```bash
Генерация текстового отчёта...
Requests      [total, rate, throughput]         100, 10.10, 10.01
Duration      [total, attack, wait]             9.989s, 9.9s, 89.042ms
Latencies     [min, mean, 50, 90, 95, 99, max]  84.993ms, 184.152ms, 105.697ms, 305.477ms, 620.595ms, 921.868ms, 928.086ms
Bytes In      [total, mean]                     22872, 228.72
Bytes Out     [total, mean]                     15872, 158.72
Success       [ratio]                           100.00%
Status Codes  [code:count]                      200:100  
Error Set:
/Library/Developer/CommandLineTools/usr/bin/make plot
Генерация HTML-графика...
Открой файл docs/perf_test/plot.html в браузере.
open docs/perf_test/plot.html
/Library/Developer/CommandLineTools/usr/bin/make histogram
Генерация гистограммы латентности...
Bucket           #   %       Histogram
[0s,     10ms]   0   0.00%   
[10ms,   20ms]   0   0.00%   
[20ms,   50ms]   0   0.00%   
[50ms,   100ms]  41  41.00%  ##############################
[100ms,  200ms]  28  28.00%  #####################
[200ms,  500ms]  25  25.00%  ##################
[500ms,  1s]     6   6.00%   ####
[1s,     +Inf]   0   0.00%   
```

Видно, что запросы идут достаточно быстро, так как запросы итак были изначально разделены, каждый из них по отдельности максимально отпимизирован


### Получение пользователя

В качестве получение пользователя тестируется handler POST https://doordashers.ru/api/auth/signin
На вход он получает логин и пароль текущего пользователя, на выход отдает полную модель пользователя

Текст первоначального запроса

```sql
SELECT id, first_name, last_name, phone_number, description, user_pic, password_hash, secret2fa FROM users WHERE login = $1
```

Проводим тестирование при 
RATE ?= 10
DURATION ?= 60s

```bash
 Генерация текстового отчёта...
Requests      [total, rate, throughput]         600, 10.02, 10.00
Duration      [total, attack, wait]             59.986s, 59.899s, 87.437ms
Latencies     [min, mean, 50, 90, 95, 99, max]  80.04ms, 117.768ms, 96.779ms, 168.915ms, 218.764ms, 397.627ms, 680.408ms
Bytes In      [total, mean]                     137772, 229.62
Bytes Out     [total, mean]                     37092, 61.82
Success       [ratio]                           100.00%
Status Codes  [code:count]                      200:600  
Error Set:
/Library/Developer/CommandLineTools/usr/bin/make plot
Генерация HTML-графика...
Открой файл docs/perf_test/plot.html в браузере.
open docs/perf_test/plot.html
/Library/Developer/CommandLineTools/usr/bin/make histogram
Генерация гистограммы латентности...
Bucket           #    %       Histogram
[0s,     10ms]   0    0.00%   
[10ms,   20ms]   0    0.00%   
[20ms,   50ms]   0    0.00%   
[50ms,   100ms]  346  57.67%  ###########################################
[100ms,  200ms]  221  36.83%  ###########################
[200ms,  500ms]  31   5.17%   ###
[500ms,  1s]     2    0.33%   
[1s,     +Inf]   0    0.00% 
```

Значения являются приемлемыми:

- нет ни одной ошибки
- 95-й перцентиль: 218.7 мс
- Средняя задержка (mean latency): 117 мс
- Максимальная задержка: 680.7 мс

Однако это может создать проблемы при более высокой нагрузке приложения

Для ускорения работы приложения попробуем добавить индексы на запросы, включающие фильтрацию и джойны:

```sql
CREATE UNIQUE INDEX idx_users_user_login ON users(login);
```

Тестирование с индексами на тех же данных:

```bash
Генерация текстового отчёта...
Requests      [total, rate, throughput]         600, 10.02, 10.00
Duration      [total, attack, wait]             59.985s, 59.9s, 85.138ms
Latencies     [min, mean, 50, 90, 95, 99, max]  80.045ms, 105.219ms, 88.832ms, 152.506ms, 178.409ms, 296.95ms, 406.705ms
Bytes In      [total, mean]                     137772, 229.62
Bytes Out     [total, mean]                     37092, 61.82
Success       [ratio]                           100.00%
Status Codes  [code:count]                      200:600  
Error Set:
/Library/Developer/CommandLineTools/usr/bin/make plot
Генерация HTML-графика...
Открой файл docs/perf_test/plot.html в браузере.
open docs/perf_test/plot.html
/Library/Developer/CommandLineTools/usr/bin/make histogram
Генерация гистограммы латентности...
Bucket           #    %       Histogram
[0s,     10ms]   0    0.00%   
[10ms,   20ms]   0    0.00%   
[20ms,   50ms]   0    0.00%   
[50ms,   100ms]  435  72.50%  ######################################################
[100ms,  200ms]  146  24.33%  ##################
[200ms,  500ms]  19   3.17%   ##
[500ms,  1s]     0    0.00%   
[1s,     +Inf]   0    0.00%  
```


Использование индексов немного ускорило работу на вход в профиль пользователя

Путем оптимизаций улалось снизить задержку и увеличить количество быстродействующих запросов.

Ключевые моменты:

- В результате применения индекса запрос начал использовать Index Only Scan, что ускоряет выполнение, особенно при больших объемах данных.

- Добавление уникального индекса на поле login в таблице users дало заметное снижение латентности — средняя задержка уменьшилась с 117 мс до 105 мс, 95-й перцентиль снизился на ~40 мс.


Вывод: план запроса теперь оптимален и использует индексы, результат достигается за меньшее время.

