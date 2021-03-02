docker-start:
	docker-compose up -d

containers-warmup:
	sleep 5

docker-stop:
	docker-compose down

go-test-integration:
	go test -tags=integration ./...

go-test-integration-update:
	go test -tags=integration ./... -update

integration-test: docker-start containers-warmup go-test-integration docker-stop

integration-test-update: docker-start containers-warmup go-test-integration-update docker-stop
