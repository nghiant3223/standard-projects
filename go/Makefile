.PHONY: docker-start containers-warmup docker-stop \
		unit-test unit-test-update go-test-integration go-test-integration-update \
		integration-test integration-test-update

docker-start:
	@docker-compose up -d

containers-warmup:
	@sleep 5

docker-stop:
	@docker-compose down

unit-test:
	@go test ./...

unit-test-update:
	@go test ./... -update

go-test-integration:
	@go test -tags=integration ./...

go-test-integration-update:
	@go test -tags=integration ./... -update

integration-test: docker-start containers-warmup go-test-integration docker-stop

integration-test-update: docker-start containers-warmup go-test-integration-update docker-stop
