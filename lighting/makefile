all: deploy

.PHONY: all deploy dev

deploy:
	cd ./vue/ && pnpm run build
	docker build --platform linux/arm64 -t sattfatt/light-box:latest . --load
	docker push sattfatt/light-box:latest

dev:
	cd ./vue/ && pnpm run build
	go run main.go
