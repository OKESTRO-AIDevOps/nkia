#!/bin/bash

set -euxo pipefail

######

local_ip=$1

apt-get update

apt-get install -y haproxy

/bin/cp -Rf ./haproxy.cfg /etc/haproxy/haproxy.cfg

rand="$(openssl rand -hex 4)"

IP=$local_ip

RAND=$rand

sed -i "s/->/master-$RAND $IP:6443/" /etc/haproxy/haproxy.cfg

systemctl restart haproxy

systemctl enable haproxy
