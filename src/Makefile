build:
	go mod download
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o example-api main.go
test:
	go test -bench=. -cover -v example-api > test.out 
coverage:
	grep -o '[0-9]*\.[0-9]*\%' test.out > coverage.out
benchmark:
	grep -o '[0-9]*\.[0-9]* ns/op' test.out | sed s#\.[0-9]*.ns\/op##g > benchmark.out
ci: 
	make test && make coverage && make benchmark