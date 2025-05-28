COVERAGE_HTML=coverage.html
COVERPROFILE_TMP=coverprofile.tmp
TARGETS_FILE ?= docs/perf_test/signup-targets.txt
RATE ?= 10
DURATION ?= 60s
REPORT_FILE ?= docs/perf_test/report.txt
PLOT_FILE ?= docs/perf_test/plot.html
HISTOGRAM_FILE ?= docs/perf_test/histogram.txt

test:
	go test -json ./... -coverprofile coverprofile_.tmp -coverpkg=./... ; \
	grep -v -e 'mocks.go' -e 'mock.go' -e 'docs.go' -e '_easyjson.go' -e 'gen_sql.go' -e '/redis/' -e '/gen/' -e '/metrics/'  -e '/cmd/' coverprofile_.tmp > coverprofile.tmp ; \
    rm coverprofile_.tmp ; \
	go tool cover -html ${COVERPROFILE_TMP} -o  $(COVERAGE_HTML); \
    go tool cover -func ${COVERPROFILE_TMP}

view-coverage:
	open $(COVERAGE_HTML)

generate-mocks:
	mockgen -source=internal/pkg/restaurants/interfaces.go -destination=internal/pkg/restaurants/mocks/mocks.go -package=mocks
	mockgen -source=internal/pkg/cart/interfaces.go -destination=internal/pkg/cart/mocks/mocks.go -package=mocks
	mockgen -source=internal/pkg/auth/interfaces.go -destination=internal/pkg/auth/mocks/mocks.go -package=mocks
	mockgen -source=internal/pkg/search/interfaces.go -destination=internal/pkg/search/mocks/mocks.go -package=mocks
	mockgen -source=internal/pkg/promocode/interfaces.go -destination=internal/pkg/promocode/mocks/mocks.go -package=mocks

easyjson:
	easyjson -all -pkg ./internal/models/

clean:
	rm -f $(COVERAGE_FILE) $(COVERAGE_HTML) ${COVERPROFILE_TMP} 

perf_tests_get:
	clear
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

clear-target-files:
	rm docs/perf_test/signup-targets.txt docs/perf_test/auth-targets.txt