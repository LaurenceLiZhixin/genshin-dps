
imports:
	goimports -local github.com/laurencelizhixin/genshin-dps -w .

tidy: imports
	go mod tidy -compat=1.17

lint: tidy
	golangci-lint run

.PHONY:test
test: tidy
	go test ./... -cover -p 1

iocgen:
	iocli gen