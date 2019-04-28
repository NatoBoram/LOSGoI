#!/bin/sh

# Go Get
go get -u -v -fix github.com/ipfs/go-ipfs

# Make Install
systemctl --user stop ipfs
cd ~/go/src/github.com/ipfs/go-ipfs
GO111MODULE=on make install
systemctl --user start ipfs

cd