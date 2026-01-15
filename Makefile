# tinyresp Makefile

GO      := go
PKG     := ./...
CMD     := ./cmd/tinyresp
BINARY  := tinyresp

.PHONY: run build test fmt lint clean

run:
	$(GO) run $(CMD)

build:
	$(GO) build -o $(BINARY) $(CMD)

test:
	$(GO) test $(PKG)

fmt:
	$(GO) fmt $(PKG)

lint:
	@command -v golangci-lint >/dev/null 2>&1 || \
		(echo "golangci-lint not installed"; exit 1)
	golangci-lint run

clean:
	rm -f $(BINARY)
	$(GO) clean