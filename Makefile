build:
	@go build -o ./bin/server ./cmd/server

test:
	@go test ./...

bench:
	@cd ./internal/services/calculator && go test -bench=.

start: build
	@./bin/server

tag=flights
docker:
	docker build -t ${tag} .
	docker run -p 8080:8080 --rm --name=flights ${tag}
