.PHONEY: all

build-docker:
	@docker build -t app .

run-docker:
	@docker run -p 127.0.0.1:5000:5000 app

test:
	@go test -v ./...