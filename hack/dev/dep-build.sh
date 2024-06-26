#!/bin/bash


U_PATH="/home/"

U_PATH="/home/$USER"

GO_VERSION="1.21.11"

sudo apt-get update

sudo apt-get install -y git make curl wget ca-certificates

curl -L https://golang.org/dl/go$GO_VERSION.linux-amd64.tar.gz -o go.tar.gz


sudo tar -C /usr/local -xvf go.tar.gz

rm go.tar.gz


sudo echo "export PATH=\$PATH:/usr/local/go/bin" | sudo tee "$U_PATH/.profile"

echo "successfully got build dependency"
echo "do: source ~/.profile"