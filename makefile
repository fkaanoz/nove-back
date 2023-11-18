main_go_path := ./app/shtil/main.go


tidy:
	go mod tidy && go mod vendor

run:
	go run $(main_go_path)

build:
	docker build -t fkaanoz/test:latest -f ./zarf/docker/Dockerfile .

hot-reload:
	air --build.cmd "go build -o bin/api $(main_go_path)" --build.bin "./bin/api"

rand-pass:
	 dd if=/dev/urandom bs=1 count=32 | shasum -a 256


