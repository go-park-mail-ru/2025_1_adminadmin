COVERAGE_HTML=coverage.html
COVERPROFILE_TMP=coverprofile.tmp

test:
	go test -json ./... -coverprofile coverprofile_.tmp -coverpkg=./... ; \
	grep -v -e 'mocks.go' -e 'mock.go' -e 'docs.go' -e '_easyjson.go' -e 'gen_sql.go' -e '/gen/' coverprofile_.tmp > coverprofile.tmp ; \
    rm coverprofile_.tmp ; \
	go tool cover -html ${COVERPROFILE_TMP} -o  $(COVERAGE_HTML); \
    go tool cover -func ${COVERPROFILE_TMP}

view-coverage:
	open $(COVERAGE_HTML)

generate-mocks:
	mockgen -source=internal/pkg/restaurants/interfaces.go -destination=internal/pkg/restaurants/mocks/mocks.go -package=mocks
	mockgen -source=internal/pkg/cart/interfaces.go -destination=internal/pkg/cart/mocks/mocks.go -package=mocks
	mockgen -source=internal/pkg/auth/interfaces.go -destination=internal/pkg/auth/mocks/mocks.go -package=mocks

clean:
	rm -f $(COVERAGE_FILE) $(COVERAGE_HTML) ${COVERPROFILE_TMP} 