tidy:
	go mod tidy && go mod vendor

run:
	go run ./app/shtil/main.go


build:
	docker build -t fkaanoz/test:latest -f ./zarf/docker/Dockerfile .