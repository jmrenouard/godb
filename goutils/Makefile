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