run-local:
	cp .env-local .env
	cp docker-compose-local.yml docker-compose.yml
	docker-compose down -v
	docker-compose up -d
	go build -o split-entrypoints cmd/main.go

run-docker:
	cp .env-docker .env
	cp docker-compose-prod.yml docker-compose.yml
	docker-compose down -v
	docker-compose up --build -d
