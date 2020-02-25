#!/bin/sh

wd=`pwd`

go get -u -v github.com/ipfs/ipfs-cluster

cd ~/go/src/github.com/ipfs/ipfs-cluster
git checkout -- .
git pull

systemctl --user stop ipfs-cluster
make install
systemctl --user start ipfs-cluster

cd $wd
