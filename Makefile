
all: godb

all: test vet fmt lint build run

run: godb
	./godb

build:
    go build -o bin/api ./cmd/api
    go build -o bin/worker ./cmd/worker

godb: *.go
	go build -ldflags="-s -w" .

test:
	go test -coverprofile=cover.out -v

coverage: test
	go tool cover -func=cover.out

coverage-html:
	go tool cover -html=cover.out

fmt:
    go list -f '{{.Dir}}' ./... | grep -v /vendor/ | xargs -L1 gofmt -l
    test -z $$(go list -f '{{.Dir}}' ./... | grep -v /vendor/ | xargs -L1 gofmt -l)

lint:
    go list ./... | grep -v /vendor/ | xargs -L1 golint -set_exit_status

vet:
    go vet ./...

install:
	go install .

clean:
	rm -rf test
	rm cover.out godb

.PHONY: all godb test coverage coverage-html
