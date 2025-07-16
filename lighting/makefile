deploy:
	cd ./vue/ && pnpm run build
	docker buildx build --platform linux/arm64 -t sattfatt/light-box:latest . --load
	docker push sattfatt/light-box:latest
