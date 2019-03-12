#!/bin/sh

FROM=`pwd`

rm -rf ~/go/src/github.com/ipfs/ipfs-cluster
go get -u -v -fix github.com/ipfs/ipfs-cluster
systemctl --user stop ipfs-cluster

cd ~/go/src/github.com/ipfs/ipfs-cluster
make install
systemctl --user start ipfs-cluster

cd $FROM
