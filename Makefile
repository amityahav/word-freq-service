.PHONEY: all


build-docker:
	@docker build -t app .

run-docker:
	@docker run -dp 127.0.0.1:5000:5000 app