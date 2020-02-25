#!/bin/sh

wd=`pwd`

go get -u -v github.com/ipfs/go-ipfs

cd ~/go/src/github.com/ipfs/go-ipfs
git checkout -- .
git pull

systemctl --user stop ipfs
make install
systemctl --user start ipfs

cd $wd
