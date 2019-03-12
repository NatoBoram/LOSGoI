#!/bin/sh
systemctl --user stop ipfs-cluster
systemctl --user stop ipfs

ipfs-cluster-service state cleanup
y

systemctl --user start ipfs
systemctl --user start ipfs-cluster
