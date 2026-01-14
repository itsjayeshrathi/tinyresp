# tinyresp Makefile

GO      := go
PKG     := ./...
BINARY  := tinyresp

.PHONY: build test fmt lint clean

build:
	$(GO) build ./...

test:
	$(GO) test $(PKG)

fmt:
	$(GO) fmt $(PKG)

lint:
	@command -v golangci-lint >/dev/null 2>&1 || \
		(echo "golangci-lint not installed"; exit 1)
	golangci-lint run

clean:
	$(GO) clean
