#!/bin/sh
systemctl --user stop ipfs-cluster
systemctl --user stop ipfs

ipfs-cluster-service state cleanup
y

# ipfs pin ls -q --type recursive | xargs ipfs pin rm

systemctl --user start ipfs
systemctl --user start ipfs-cluster
