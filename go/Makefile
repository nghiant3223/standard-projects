INTEGRATION_TEST_ENV=\
	DATABASE_URL=postgres://postgres:mysecretpassword@localhost:5432/gotodo?sslmode=disable

docker.start:
	docker-compose up -d

docker.stop:
	docker-compose down

go.test.integeration:
	$(INTEGRATION_TEST_ENV) \
	go test -tags=integration ./...

test.integration: docker.start go.test.integeration docker.stop
