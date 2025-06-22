docker buildx build --platform linux/arm64 -t sattfatt/light-box:latest . --load
docker push sattfatt/light-box:latest
