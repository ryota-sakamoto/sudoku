dev:
	go run *.go

less:
	make dev 2>&1 | less

test:
	go test -v ./...
