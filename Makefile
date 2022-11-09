GO_BIN = go

clean:
	@echo "Cleaning up"
	rm -f main
	
build:
	$(GO_BIN) build main.go

test:
	$(GO_BIN) test -v ./...
