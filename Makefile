PACKAGE=github.com/tappoy/cloud
WORKING_DIRS=tmp bin

SRC=$(shell find . -name "*.go")
BIN=bin/$(shell basename $(CURDIR))
DOC=Document.txt
CONOHA_DOC=conoha/Document.txt
COVER=tmp/cover
COVER0=tmp/cover0


.PHONY: all clean fmt cover test lint

all: $(WORKING_DIRS) fmt $(BIN) test $(DOC) $(CONOHA_DOC) lint

clean:
	rm -rf $(WORKING_DIRS) $(DOC) $(CONOHA_DOC)

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

$(DOC): $(ls *.go)
	go doc -all > $(DOC)

$(CONOHA_DOC): $(ls conoha/*.go)
	go doc -all conoha > $(CONOHA_DOC)

cover: $(COVER)
	grep "0$$" $(COVER) | sed 's!$(PACKAGE)!.!' | tee $(COVER0)

lint: $(BIN)
	go vet
