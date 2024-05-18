PACKAGE=github.com/tappoy/archive
WORKING_DIRS=tmp bin

SRC=$(shell find . -name "*.go")
BIN=bin/$(shell basename $(CURDIR))
DOC=Document.txt
OPENSTACK_DOC=openstack/Document.txt
AWS_DOC=aws/Document.txt
MOCK_DOC=mock/Document.txt
COVER=tmp/cover
COVER0=tmp/cover0


.PHONY: all clean fmt cover test lint build

all: $(WORKING_DIRS) fmt build test $(DOC) $(OPENSTACK_DOC) $(AWS_DOC) $(MOCK_DOC) lint

clean:
	rm -rf $(WORKING_DIRS) $(DOC) $(OPENSTACK_DOC) $(AWS_DOC) $(MOCK_DOC)

$(WORKING_DIRS):
	mkdir -p $(WORKING_DIRS)

fmt: $(SRC)
	go fmt ./...

go.sum: go.mod
	go mod tidy

build: $(SRC) go.sum
	go build -o $(BIN)

test: $(BIN)
	go test -v -tags=mock -vet=all -cover -coverprofile=$(COVER) ./...

$(DOC): *.go
	go doc -all > $(DOC)

$(OPENSTACK_DOC): openstack/*.go
	go doc -all openstack > $(OPENSTACK_DOC)

$(AWS_DOC): aws/*.go
	go doc -all aws > $(AWS_DOC)

$(MOCK_DOC): mock/*.go
	go doc -all mock > $(MOCK_DOC)

cover: $(COVER)
	grep "0$$" $(COVER) | sed 's!$(PACKAGE)!.!' | tee $(COVER0)

lint: $(BIN)
	go vet
