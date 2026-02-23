test:
	go test .

bench:
	go test -bench=. -benchmem -count=5 -timeout=30m > results.txt

check:
	go mod tidy
	go test ./...
	go vet ./...
	gofmt -d .
