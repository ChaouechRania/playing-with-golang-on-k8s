all: serve

.PHONY: deploy

deploy:
	docker-compose -f ./deploy/local/docker-compose.yml up -d

teardown:
	docker-compose -f ./deploy/local/docker-compose.yml stop

serve:
	go run cmd/playing-with-golang-on-k8s.go --dbHost localhost --dbUser user  --dbPassword secret  --dbName productdb

test:
	go test ./...

.EXPORT_ALL_VARIABLES:
TEST_ES_HOST=127.0.0.1
TEST_ES_PORT=9200

build:
	go build cmd/playing-with-golang-on-k8s.go