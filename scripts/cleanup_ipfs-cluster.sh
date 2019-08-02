#!/bin/sh

# Stop
systemctl --user stop ipfs-cluster
systemctl --user stop ipfs

# Cleanup IPFS Cluster
ipfs-cluster-service state cleanup
y
rm -rf ~/.ipfs-cluster/raft.old.*

# Start
systemctl --user start ipfs
systemctl --user start ipfs-cluster
