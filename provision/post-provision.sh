IMAGE=sattfatt/light-box:latest

docker pull $IMAGE

docker run --detach --restart unless-stopped --name light-box $IMAGE

docker run \
  --detach \
  --restart unless-stopped \
  --name watchtower \
  --volume /var/run/docker.sock:/var/run/docker.sock \
  containrrr/watchtower \
  --cleanup \
  --interval 86400 \
  $IMAGE
