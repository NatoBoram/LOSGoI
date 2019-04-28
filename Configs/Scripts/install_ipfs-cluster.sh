#!/bin/sh

# Make sure you're not in $GOPATH
cd

# Rm
rm -rf ~/go/src/github.com/ipfs/go-ipfs
rm -rf ~/go/src/github.com/ipfs/ipfs-cluster

# Go Get
go get -u -v github.com/ipfs/go-ipfs
go get -u -v github.com/ipfs/ipfs-cluster

# SystemCtl Stop
systemctl --user stop ipfs-cluster
systemctl --user stop ipfs

# Cleanup State
rm -rf ~/.ipfs-cluster/raft.old.*

# Make Install - IPFS
cd ~/go/src/github.com/ipfs/go-ipfs
GO111MODULE=on make install

# Make Install - IPFS Cluster
cd ~/go/src/github.com/ipfs/ipfs-cluster
GO111MODULE=on make install

# SystemCtl Start
systemctl --user start ipfs
systemctl --user start ipfs-cluster

cd
