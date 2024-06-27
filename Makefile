build:
	@go build -o ./bin/server ./cmd/server

test:
	@go test ./...

bench:
	@cd ./internal/services/calculator && go test -bench=.

HTTP_LISTEN_PORT=8080
start: build
	@./bin/server

tag=flights
docker:
	docker build -t ${tag} .
	docker run -e HTTP_LISTEN_PORT=8080 -p 8080:8080 --rm --name=flights ${tag}
