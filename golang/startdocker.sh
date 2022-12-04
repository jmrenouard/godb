#!/bin/bash


DOCKER_DIR=/mnt/wsl/shared-docker
sudo mkdir -pm o=,ug=rwx "$DOCKER_DIR"
sudo chgrp docker "$DOCKER_DIR"
sudo service docker start 


sudo ln -sf //mnt/wsl/shared-docker/docker.sock /var/run/
docker run hello-world