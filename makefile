CGO_ENABLED=0

.PHONY: test
test:
	@go test ./... -v

.PHONY: bench
bench:
	@go test -bench . ./sort/... ./search/... ./checksum/... -benchmem -v