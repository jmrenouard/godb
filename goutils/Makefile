UTILS :=  $(shell find $(SOURCEDIR) -maxdepth 1 -type f -name '*utils.go')
DAYDATE := $(shell date +%d%m%Y)
run: 
	go run GetTargetDrive.go $(UTILS)

build_exe:
	go.exe build GetTargetDrive.go $(UTILS)

build:
	go build GetTargetDrive.go $(UTILS)

zip: build build_exe
	zip GetTargetDrive_$(DAYDATE).zip GetTargetDrive GetTargetDrive.exe

clean:
	rm -rf GetTargetDrive.exe GetTargetDrive GetTargetDrive*.zip

all: clean zip

all: test vet fmt lint build

test:
    go test ./...

vet:
    go vet ./...

fmt:
    go list -f '{{.Dir}}' ./... | grep -v /vendor/ | xargs -L1 gofmt -l
    test -z $$(go list -f '{{.Dir}}' ./... | grep -v /vendor/ | xargs -L1 gofmt -l)

lint:
    go list ./... | grep -v /vendor/ | xargs -L1 golint -set_exit_status

build:
    go build -o bin/api ./cmd/api
    go build -o bin/worker ./cmd/worker