# Go variables
GO              := go
GOFMT           := $(GO)fmt
GOFLAGS         :=
TEST_TARGETS    := $(shell $(GO) list ./... | grep -v /vendor/)

.PHONY: all format test

all: format test

format:
	$(GOFMT) -w .

test:
	$(GO) test $(GOFLAGS) $(TEST_TARGETS)

