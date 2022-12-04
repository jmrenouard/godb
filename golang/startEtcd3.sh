#!/bin/bash

HostIP="$(ifconfig ${1:-"eth0"}| grep 'inet ' | awk '{print $2}')"

docker ps -a | grep etcd | awk '{print $1}'| xargs -n 1 docker rm -f

docker run -d -v /usr/share/ca-certificates/:/etc/ssl/certs -p 4001:4001 -p 2380:2380 -p 2379:2379 \
 --name etcd quay.io/coreos/etcd:latest \
 /usr/local/bin/etcd -listen-client-urls http://0.0.0.0:2379,http://0.0.0.0:4001 \
 -name etcd0 \
 -advertise-client-urls http://${HostIP}:2379,http://${HostIP}:4001 \
 -initial-advertise-peer-urls http://${HostIP}:2380 \
 -listen-peer-urls http://0.0.0.0:2380 \
 -initial-cluster-token etcd-cluster-1 \
 -initial-cluster etcd0=http://${HostIP}:2380 \
 -initial-cluster-state new

ETCDCTL_API=3 sudo etcdctl member list
ETCDCTL_API=3 sudo etcdctl set val1/val2/val test
ETCDCTL_API=3 sudo etcdctl ls -r
