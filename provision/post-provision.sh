IMAGE=sattfatt/light-box:latest

docker pull $IMAGE

docker stop light-box
docker rm light-box
docker run --privileged --detach --restart always --name light-box -p 80:80 -p 8080:8080 $IMAGE

docker stop watchtower
docker rm watchtower
docker run --detach --restart unless-stopped --name watchtower --volume /var/run/docker.sock:/var/run/docker.sock containrrr/watchtower --cleanup --interval 30 light-box
