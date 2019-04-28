#!/bin/sh

# Stop
systemctl --user stop ipfs-cluster
systemctl --user stop ipfs

# Cleanup IPFS Cluster
ipfs-cluster-service state cleanup
y
rm -rf ~/.ipfs-cluster/raft.old.*

# Cleanup IPFS
ipfs pin ls -q --type recursive | xargs ipfs pin rm
ipfs repo gc

# Start
systemctl --user start ipfs
systemctl --user start ipfs-cluster

exit
