COVERAGE_FILE=coverage.out
COVERAGE_HTML=coverage.html


test:
	go test ./...

cover:
	go test -coverprofile=$(COVERAGE_FILE) ./...
	go tool cover -html=$(COVERAGE_FILE) -o $(COVERAGE_HTML)
	go tool cover -func=$(COVERAGE_FILE)

view-coverage:
	open $(COVERAGE_HTML)

clean:
	rm -f $(COVERAGE_FILE) $(COVERAGE_HTML)