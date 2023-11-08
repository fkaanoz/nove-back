tidy:
	go mod tidy && go mod vendor

run:
	go run ./app/shtil/main.go