#!/bin/sh

HERE=`pwd`
cd

# Discard changes - IPFS
cd ~/go/src/github.com/ipfs/go-ipfs
git checkout -- .

# Discard changes - IPFS Cluster
cd ~/go/src/github.com/ipfs/ipfs-cluster
git checkout -- .

cd

# Update dependencies
go get -u -v github.com/ipfs/go-ipfs
go get -u -v github.com/ipfs/ipfs-cluster

# Stop services
systemctl --user stop ipfs-cluster
systemctl --user stop ipfs

# Remove old states
rm -rf ~/.ipfs-cluster/raft.old.*

# Install IPFS
cd ~/go/src/github.com/ipfs/go-ipfs
GO111MODULE=on make install

# Install IPFS Cluster
cd ~/go/src/github.com/ipfs/ipfs-cluster
GO111MODULE=on make install

# Start services
systemctl --user start ipfs
systemctl --user start ipfs-cluster
ping lineageos-on-ipfs.com -c 15

cd $HERE

# Maintenance
ipfs-cluster-ctl sync
ipfs-cluster-ctl recover --local
ipfs repo gc

# Version
ipfs version --all
ipfs-cluster-service version

# Connection
ipfs-cluster-ctl peers ls
journalctl --user -fu ipfs-cluster --output cat
