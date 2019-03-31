#!/bin/sh

# Make sure you're not in $GOPATH
cd

# Go Get - IPFS
go get -u -v -fix github.com/ipfs/go-ipfs

# Go Get - IPFS Cluster
rm -rf ~/go/src/github.com/ipfs/ipfs-cluster
go get -u -v -fix github.com/ipfs/ipfs-cluster

# Dependencies
go get -u -v -fix github.com/apache/thrift/lib/go/thrift
go get -u -v -fix github.com/gorilla/mux
go get -u -v -fix github.com/hashicorp/raft
go get -u -v -fix github.com/hashicorp/raft-boltdb
go get -u -v -fix github.com/hsanjuan/go-libp2p-gostream
go get -u -v -fix github.com/hsanjuan/go-libp2p-http
go get -u -v -fix github.com/ipfs/go-fs-lock
go get -u -v -fix github.com/libp2p/go-libp2p-consensus
go get -u -v -fix github.com/libp2p/go-libp2p-pubsub
go get -u -v -fix github.com/libp2p/go-libp2p-raft
go get -u -v -fix github.com/prometheus/client_golang/prometheus
go get -u -v -fix github.com/prometheus/client_golang/prometheus/promhttp
go get -u -v -fix github.com/rs/cors
go get -u -v -fix github.com/zenground0/go-dot
go get -u -v -fix go.opencensus.io/vendor/google.golang.org/api/support/bundler
go get -u -v -fix google.golang.org/api/support/bundler
go get -u -v -fix google.golang.org/genproto/googleapis/rpc/status

# SystemCtl Stop
systemctl --user stop ipfs-cluster
systemctl --user stop ipfs

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
