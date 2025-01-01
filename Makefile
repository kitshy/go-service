go-service:
	env GO111MODULE=on go build -v $(LDFLAGS)
.PHONY: go-service

clean:
	rm go-service

test:
	go test -v ./...
#
#protogo:
#	sh ./sh/go_compile.sh
#
#lint:
#	golangci-lint run ./...
#
#.PHONY: \
#	multichain-sync \
#	clean \
#	test \
#	protogo \
#	lint