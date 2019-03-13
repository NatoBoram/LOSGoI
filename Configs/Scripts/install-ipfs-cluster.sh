#!/bin/sh

FROM=`pwd`

rm -rf ~/go/src/github.com/ipfs/ipfs-cluster
go get -u -v -fix github.com/ipfs/ipfs-cluster

# Dependencies
go get -u -v -fix github.com/gorilla/mux
go get -u -v -fix github.com/hashicorp/raft
go get -u -v -fix github.com/hashicorp/raft-boltdb
go get -u -v -fix github.com/hsanjuan/go-libp2p-gostream
go get -u -v -fix github.com/hsanjuan/go-libp2p-http
go get -u -v -fix github.com/ipfs/go-fs-lock
go get -u -v -fix github.com/libp2p/go-libp2p-consensus
go get -u -v -fix github.com/libp2p/go-libp2p-pubsub
go get -u -v -fix github.com/libp2p/go-libp2p-raft
go get -u -v -fix github.com/rs/cors
go get -u -v -fix go.opencensus.io/vendor/google.golang.org/api/support/bundler
go get -u -v -fix google.golang.org/api/support/bundler
go get -u -v -fix google.golang.org/genproto/googleapis/rpc/status
go get -u -v -fix github.com/zenground0/go-dot
go get -u -v -fix github.com/prometheus/client_golang/prometheus
go get -u -v -fix github.com/prometheus/client_golang/prometheus/promhttp

# Make Install
systemctl --user stop ipfs-cluster
cd ~/go/src/github.com/ipfs/ipfs-cluster
make install
systemctl --user start ipfs-cluster

cd $FROM
