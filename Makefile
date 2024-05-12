PACKAGE=github.com/tappoy/cloud
WORKING_DIRS=tmp bin

SRC=$(shell find . -name "*.go")
BIN=bin/$(shell basename $(CURDIR))
DOC=Document.txt
COVER=tmp/cover
COVER0=tmp/cover0

.PHONY: all clean fmt cover test lint

all: $(WORKING_DIRS) fmt $(BIN) test $(DOC) lint

clean:
	rm -rf $(WORKING_DIRS)

$(WORKING_DIRS):
	mkdir -p $(WORKING_DIRS)

fmt: $(SRC)
	go fmt ./...

go.sum: go.mod
	go mod tidy

$(BIN): $(SRC) go.sum
	go build -o $(BIN)

test: $(BIN)
	go test -v -tags=mock -vet=all -cover -coverprofile=$(COVER) ./...

doc: $(SRC)
	go doc -all > Document.txt
	go doc -all conoha > conoha/Document.txt

cover: $(COVER)
	grep "0$$" $(COVER) | sed 's!$(PACKAGE)!.!' | tee $(COVER0)

lint: $(BIN)
	go vet
