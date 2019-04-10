#!/bin/sh

# Make sure you're not in $GOPATH
cd

# Go Get - IPFS
GO111MODULE=on go get -u -v -fix github.com/ipfs/go-ipfs

# Go Get - IPFS Cluster
GO111MODULE=on go get -u -v -fix github.com/ipfs/ipfs-cluster

# SystemCtl Stop
systemctl --user stop ipfs-cluster
systemctl --user stop ipfs

# Cleanup State
#ipfs-cluster-service state cleanup
#y
rm -rf ~/.ipfs-cluster/raft.old.*

# Make Install - IPFS
cd ~/go/src/github.com/ipfs/go-ipfs
make install

# Make Install - IPFS Cluster
cd ~/go/src/github.com/ipfs/ipfs-cluster
make install

# SystemCtl Start
systemctl --user start ipfs
systemctl --user start ipfs-cluster

cd
